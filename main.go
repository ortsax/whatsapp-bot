package main

import (
	"context"
	"flag"
	"fmt"
	"orstax/store/sqlstore"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"

	_ "modernc.org/sqlite"
)

var client *whatsmeow.Client
var metaJID = types.NewMetaAIJID
var pendingReplies = make(map[string]types.JID)

var lastProcessedResponse = make(map[string]string)
var sentMessageIDs = make(map[string]types.MessageID)

func chatMetaAi(v *events.Message) {
	var responseText string
	resID := v.Message.GetMessageContextInfo().GetBotMetadata().GetBotResponseID()

	if v.Message.Conversation != nil {
		responseText = v.Message.GetConversation()
	} else if v.Message.ExtendedTextMessage != nil {
		responseText = v.Message.GetExtendedTextMessage().GetText()
	} else if v.Message.ProtocolMessage != nil && v.Message.ProtocolMessage.GetType() == waProto.ProtocolMessage_MESSAGE_EDIT {
		edit := v.Message.ProtocolMessage.EditedMessage
		if edit != nil {
			if edit.Conversation != nil {
				responseText = edit.GetConversation()
			} else if edit.ExtendedTextMessage != nil {
				responseText = edit.ExtendedTextMessage.GetText()
			}
		}
	}

	if responseText != "" && resID != "" {
		lastText, seen := lastProcessedResponse[resID]
		if !seen || len(responseText) > len(lastText) {
			if targetJID, ok := pendingReplies[v.Info.Sender.String()]; ok {
				if msgID, exists := sentMessageIDs[resID]; exists {
					editMsg := client.BuildEdit(targetJID, msgID, &waProto.Message{
						Conversation: proto.String(responseText),
					})
					_, err := client.SendMessage(context.Background(), targetJID, editMsg)
					if err == nil {
						lastProcessedResponse[resID] = responseText
					}
				} else {
					resp, err := client.SendMessage(context.Background(), targetJID, &waProto.Message{
						Conversation: proto.String(responseText),
					})
					if err == nil {
						sentMessageIDs[resID] = resp.ID
						lastProcessedResponse[resID] = responseText
					}
				}
			}
		}
	}
}

func eventHandler(evt any) {
	switch v := evt.(type) {
	case *events.Message:

		text := v.Message.GetConversation()
		if text == "" {
			text = v.Message.GetExtendedTextMessage().GetText()
		}

		if strings.HasPrefix(strings.ToLower(text), "meta ") {
			query := strings.TrimPrefix(text, "meta ")
			_, err := client.SendMessage(context.Background(), metaJID, &waProto.Message{
				Conversation: proto.String(query),
			})
			if err != nil {
				return
			}
			pendingReplies[metaJID.String()] = v.Info.Chat
			return
		}

		if v.Info.Sender.User == metaJID.User {
			chatMetaAi(v)
		}
	}
}

func main() {
	phoneArg := flag.String("phone-number", "", "The phone number to pair with (international format)")
	flag.Parse()

	dbLog := waLog.Stdout("Database", "ERROR", true)
	ctx := context.Background()

	dbAddr := "file:database.db?" +
		"_pragma=foreign_keys(1)&" +
		"_pragma=journal_mode(WAL)&" +
		"_pragma=synchronous(NORMAL)&" +
		"_pragma=busy_timeout(10000)&" +
		"_pragma=cache_size(-64000)&" +
		"_pragma=mmap_size(2147483648)&" +
		"_pragma=temp_store(MEMORY)"

	container, err := sqlstore.New(ctx, "sqlite", dbAddr, dbLog)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "ERROR", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	err = client.Connect()
	if err != nil {
		panic(err)
	}

	if client.Store.ID == nil {
		if *phoneArg == "" {
			fmt.Println("No session found. Please provide a phone number using --phone-number")
			return
		}

		fmt.Println("Waiting 10 seconds before generating pairing code...")
		time.Sleep(10 * time.Second)

		code, err := client.PairPhone(ctx, *phoneArg, true, whatsmeow.PairClientChrome, "Chrome (Linux)")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Your pairing code is: %s\n", code)
	} else {
		fmt.Println("Already logged in.")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
