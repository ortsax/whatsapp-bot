package plugins

import (
	"context"
	"fmt"
	"time"

	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"
)

func init() {
	Register(&Command{
		Pattern:  "ping",
		Category: "utility",
		Func: func(ctx *Context) error {
			start := time.Now()

			resp, err := ctx.Reply(T().Pong)
			if err != nil {
				return err
			}

			elapsed := time.Since(start)

			edit := ctx.Client.BuildEdit(ctx.Event.Info.Chat, resp.ID, &waProto.Message{
				Conversation: proto.String(fmt.Sprintf(T().PongLatency, elapsed.Milliseconds())),
			})
			_, err = ctx.Client.SendMessage(context.Background(), ctx.Event.Info.Chat, edit)
			return err
		},
	})
}
