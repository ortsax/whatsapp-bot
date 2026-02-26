package plugins

import (
	"context"
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

// Command defines a bot command.
type Command struct {
	Pattern  string
	Aliases  []string
	IsSudo   bool
	IsGroup  bool
	Category string
	Func     func(ctx *Context) error
}

// Context carries per-invocation state passed to command handlers.
type Context struct {
	Client  *whatsmeow.Client
	Event   *events.Message
	Args    []string // whitespace-split words after the command name
	Text    string   // everything after the command name (unsplit)
	Prefix  string   // the matched prefix character(s)
	Matched string   // the matched command name (lowercased)
}

// Reply sends a plain-text message back to the originating chat.
func (c *Context) Reply(text string) (whatsmeow.SendResponse, error) {
	return c.Client.SendMessage(context.Background(), c.Event.Info.Chat, &waProto.Message{
		Conversation: proto.String(text),
	})
}

var registry []*Command

// Register adds a command to the global registry.
func Register(cmd *Command) {
	registry = append(registry, cmd)
}

// parseCommand tries to extract (prefix, commandName, rest) from text using the
// configured prefixes. Matching is case-insensitive. A space between the prefix
// and the command name is allowed (e.g. ". ping" == ".ping").
func parseCommand(text string, prefixes []string) (prefix, name, rest string, ok bool) {
	lower := strings.ToLower(text)
	for _, p := range prefixes {
		var after string
		if p == "" {
			// No prefix – command is the first word.
			after = lower
		} else {
			lp := strings.ToLower(p)
			if !strings.HasPrefix(lower, lp) {
				continue
			}
			after = lower[len(lp):]
		}
		after = strings.TrimLeft(after, " ")
		if after == "" {
			continue
		}
		parts := strings.SplitN(after, " ", 2)
		name = parts[0]
		if len(parts) > 1 {
			rest = strings.TrimSpace(parts[1])
		}
		return p, name, rest, true
	}
	return "", "", "", false
}

// findCommand looks up a command by its pattern or any alias (case-insensitive).
func findCommand(name string) *Command {
	for _, cmd := range registry {
		if strings.ToLower(cmd.Pattern) == name {
			return cmd
		}
		for _, alias := range cmd.Aliases {
			if strings.ToLower(alias) == name {
				return cmd
			}
		}
	}
	return nil
}

// extractText returns the human-readable text from a message, checking plain
// conversation, extended text, and image/video captions.
func extractText(evt *events.Message) string {
	if t := evt.Message.GetConversation(); t != "" {
		return t
	}
	if t := evt.Message.GetExtendedTextMessage().GetText(); t != "" {
		return t
	}
	if t := evt.Message.GetImageMessage().GetCaption(); t != "" {
		return t
	}
	if t := evt.Message.GetVideoMessage().GetCaption(); t != "" {
		return t
	}
	return ""
}

// Dispatch parses a message event, resolves the matching command, enforces all
// access checks, and calls the command handler.
func Dispatch(client *whatsmeow.Client, evt *events.Message) {
	text := extractText(evt)
	if text == "" {
		return
	}

	senderPhone := evt.Info.Sender.User
	isGroup := evt.Info.Chat.Server == types.GroupServer

	prefix, name, rest, ok := parseCommand(text, BotSettings.GetPrefixes())
	if !ok {
		return
	}

	cmd := findCommand(name)
	if cmd == nil {
		return
	}

	ctx := &Context{
		Client:  client,
		Event:   evt,
		Args:    strings.Fields(rest),
		Text:    rest,
		Prefix:  prefix,
		Matched: name,
	}

	isSudo := BotSettings.IsSudo(senderPhone)
	mode := BotSettings.GetMode()

	// Private mode: silently ignore non-sudo users – no response at all.
	if mode == ModePrivate && !isSudo {
		return
	}

	if cmd.IsGroup && !isGroup {
		ctx.Reply("❌ This command can only be used in groups.")
		return
	}

	if cmd.IsSudo && !isSudo {
		ctx.Reply("🔒 This command is for sudo users only.")
		return
	}

	_ = cmd.Func(ctx)
}
