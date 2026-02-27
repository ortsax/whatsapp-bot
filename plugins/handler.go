package plugins

import (
	"encoding/json"
	"fmt"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

// NewHandler returns a whatsmeow event handler that drives the plugin system.
// Each message is handled in its own goroutine so the whatsmeow event loop is
// never blocked by command processing or network I/O.
func NewHandler(client *whatsmeow.Client) func(evt any) {
	return func(evt any) {
		switch v := evt.(type) {
		case *events.Message:
			if b, err := json.MarshalIndent(v, "", "  "); err == nil {
				fmt.Printf("[DEBUG] Message:\n%s\n", b)
			} else {
				fmt.Printf("[DEBUG] Message (raw): %+v\n", v)
			}
			go SaveUser(v)
			if v.Info.Sender.User == MetaJID.User {
				go HandleMetaAIResponse(client, v)
				return
			}
			go Dispatch(client, v)
		}
	}
}
