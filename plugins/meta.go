package plugins

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

// MetaJID is the WhatsApp JID for the Meta AI assistant.
var MetaJID = types.NewMetaAIJID

var metaMu sync.Mutex

// pendingReplies maps the Meta AI JID string to the chat JID that issued the query.
var pendingReplies = make(map[string]types.JID)

// lastProcessedResponse tracks the last text seen for a given Meta AI response ID
// so we only forward longer (more complete) streaming updates.
var lastProcessedResponse = make(map[string]string)

// sentMessageIDs maps a Meta AI response ID to the message ID we already sent,
// so we can edit it in place as the streamed response grows.
var sentMessageIDs = make(map[string]types.MessageID)

// buildMetaQuery constructs a context-aware query string for Meta AI.
// When pastContext is empty a simpler format without context is returned.
func buildMetaQuery(senderID, pushName, pastContext, userQuery string) string {
	if pastContext != "" {
		return fmt.Sprintf(
			"User ID: %s, Their Name: %s, Past Context — You Meta AI: %s, Their reply to your message: %s",
			senderID, pushName, pastContext, userQuery,
		)
	}
	return fmt.Sprintf("User ID: %s, Their Name: %s, Query: %s", senderID, pushName, userQuery)
}

// senderPushName returns the push name from the event, falling back to the
// sender's user ID when the name is not available.
func senderPushName(evt *events.Message) string {
	if evt.Info.PushName != "" {
		return evt.Info.PushName
	}
	return evt.Info.Sender.User
}

// HandleMetaAIResponse processes incoming messages from the Meta AI JID and
// forwards the response text (or edits the previous message) back to the
// original requester's chat.
func HandleMetaAIResponse(client *whatsmeow.Client, v *events.Message) {
	var responseText string
	resID := v.Message.GetMessageContextInfo().GetBotMetadata().GetBotResponseID()

	// Check for image response from Meta AI (e.g. image generation).
	if img := v.Message.GetImageMessage(); img != nil {
		metaMu.Lock()
		targetJID, ok := pendingReplies[v.Info.Sender.String()]
		metaMu.Unlock()
		if ok {
			client.SendMessage(context.Background(), targetJID, &waProto.Message{ImageMessage: img}) //nolint:errcheck
		}
		return
	}

	if v.Message.Conversation != nil {
		responseText = v.Message.GetConversation()
	} else if v.Message.ExtendedTextMessage != nil {
		responseText = v.Message.GetExtendedTextMessage().GetText()
	} else if v.Message.ProtocolMessage != nil &&
		v.Message.ProtocolMessage.GetType() == waProto.ProtocolMessage_MESSAGE_EDIT {
		edit := v.Message.ProtocolMessage.EditedMessage
		if edit != nil {
			if edit.Conversation != nil {
				responseText = edit.GetConversation()
			} else if edit.ExtendedTextMessage != nil {
				responseText = edit.ExtendedTextMessage.GetText()
			}
		}
	}

	if responseText == "" || resID == "" {
		return
	}

	metaMu.Lock()
	defer metaMu.Unlock()

	if lastText, seen := lastProcessedResponse[resID]; seen && len(responseText) <= len(lastText) {
		return
	}

	targetJID, ok := pendingReplies[v.Info.Sender.String()]
	if !ok {
		return
	}

	if msgID, exists := sentMessageIDs[resID]; exists {
		editMsg := client.BuildEdit(targetJID, msgID, &waProto.Message{
			Conversation: proto.String(responseText),
		})
		if _, err := client.SendMessage(context.Background(), targetJID, editMsg); err == nil {
			lastProcessedResponse[resID] = responseText
			updateMetaMessageText(string(msgID), responseText)
		}
	} else {
		if resp, err := client.SendMessage(context.Background(), targetJID, &waProto.Message{
			Conversation: proto.String(responseText),
		}); err == nil {
			sentMessageIDs[resID] = resp.ID
			lastProcessedResponse[resID] = responseText
			saveMetaMessage(string(resp.ID), targetJID.String(), responseText)
		}
	}
}

// handleMetaAutoReply is a moderation hook that detects when a user replies
// directly to a forwarded Meta AI message (without using the .meta command)
// and automatically continues the conversation with full context.
func handleMetaAutoReply(client *whatsmeow.Client, evt *events.Message) {
	if evt.Info.IsFromMe {
		return
	}

	ci := evt.Message.GetExtendedTextMessage().GetContextInfo()
	if ci == nil {
		return
	}
	stanzaID := ci.GetStanzaID()
	if stanzaID == "" {
		return
	}

	pastResponse, found := getMetaMessageText(stanzaID)
	if !found {
		return
	}

	replyText := evt.Message.GetExtendedTextMessage().GetText()
	if replyText == "" {
		return
	}

	// If the reply looks like a command, let the command handler deal with it.
	for _, p := range BotSettings.GetPrefixes() {
		if p != "" && strings.HasPrefix(strings.ToLower(replyText), strings.ToLower(p)) {
			return
		}
	}

	query := buildMetaQuery(evt.Info.Sender.User, senderPushName(evt), pastResponse, replyText)
	if _, err := client.SendMessage(context.Background(), MetaJID, &waProto.Message{
		Conversation: proto.String(query),
	}); err != nil {
		return
	}

	metaMu.Lock()
	pendingReplies[MetaJID.String()] = evt.Info.Chat
	metaMu.Unlock()
}

func init() {
	RegisterModerationHook(handleMetaAutoReply)

	Register(&Command{
		Pattern:  "meta",
		Category: "ai",
		Func: func(ctx *Context) error {
			query := ctx.Text
			senderID := ctx.Event.Info.Sender.User
			pushName := senderPushName(ctx.Event)

			var outMsg *waProto.Message
			nonTextQuoted := false

			// Inspect any quoted (replied-to) message.
			ci := ctx.Event.Message.GetExtendedTextMessage().GetContextInfo()
			if ci != nil {
				quoted := ci.GetQuotedMessage()
				stanzaID := ci.GetStanzaID()

				if quoted != nil {
					// Case A: quoting a forwarded Meta AI response — add full context.
					if stanzaID != "" {
						if pastResponse, found := getMetaMessageText(stanzaID); found {
							if query == "" {
								ctx.Reply(T().MetaUsage)
								return nil
							}
							q := buildMetaQuery(senderID, pushName, pastResponse, query)
							outMsg = &waProto.Message{Conversation: proto.String(q)}
						}
					}

					// Case B: quoting a plain text message — include it as context.
					if outMsg == nil {
						quotedText := quoted.GetConversation()
						if quotedText == "" {
							quotedText = quoted.GetExtendedTextMessage().GetText()
						}
						if quotedText != "" {
							if query == "" {
								ctx.Reply(T().MetaUsage)
								return nil
							}
							q := fmt.Sprintf(
								"User ID: %s, Their Name: %s, Quoted context: %s, User query: %s",
								senderID, pushName, quotedText, query,
							)
							outMsg = &waProto.Message{Conversation: proto.String(q)}
						} else {
							// Quoted message is non-text (image, video, audio, etc.).
							nonTextQuoted = true
						}
					}
				}
			}

			// Case C: the message itself carries media (caption = ".meta query").
			if outMsg == nil {
				if img := ctx.Event.Message.GetImageMessage(); img != nil {
					img.Caption = proto.String(query)
					outMsg = &waProto.Message{ImageMessage: img}
				} else if vid := ctx.Event.Message.GetVideoMessage(); vid != nil {
					vid.Caption = proto.String(query)
					outMsg = &waProto.Message{VideoMessage: vid}
				}
			}

			// Case D: plain text query with user context.
			if outMsg == nil {
				if query == "" {
					ctx.Reply(T().MetaUsage)
					return nil
				}
				q := buildMetaQuery(senderID, pushName, "", query)
				outMsg = &waProto.Message{Conversation: proto.String(q)}
			}

			// Warn the user that the quoted media was dropped before sending.
			if nonTextQuoted {
				ctx.Reply(T().MetaNonTextWarning)
			}

			resp, err := ctx.Client.SendMessage(context.Background(), MetaJID, outMsg)
			if err != nil {
				return err
			}
			_ = resp

			metaMu.Lock()
			pendingReplies[MetaJID.String()] = ctx.Event.Info.Chat
			metaMu.Unlock()
			return nil
		},
	})
}
