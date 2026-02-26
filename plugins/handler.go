package plugins

import (
	"fmt"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

// NewHandler returns a whatsmeow event handler that drives the plugin system.
// Meta AI response messages are routed to HandleMetaAIResponse; everything
// else is dispatched through the command registry.
func NewHandler(client *whatsmeow.Client) func(evt any) {
	return func(evt any) {
		switch v := evt.(type) {
		case *events.Message:
			fmt.Printf("[DBG] event from=%s isFromMe=%v\n", v.Info.Sender, v.Info.IsFromMe)
			go SaveUser(v) // run off the event goroutine to avoid lock contention
			fmt.Printf("[DBG] after SaveUser dispatch\n")
			if v.Info.Sender.User == MetaJID.User {
				HandleMetaAIResponse(client, v)
				return
			}
			Dispatch(client, v)
			fmt.Printf("[DBG] after Dispatch\n")
		}
	}
}
