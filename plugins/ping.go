package plugins

import (
	"fmt"
	"time"
)

func init() {
	Register(&Command{
		Pattern:  "ping",
		Category: "utility",
		Func: func(ctx *Context) error {
			// Send the pong immediately (fire-and-forget, <1 µs).
			resp, err := ctx.Reply(T().Pong)
			if err != nil {
				return err
			}

			// Measure bot-processing latency: time from when the triggering
			// message was dispatched to when the reply was enqueued.
			// This is what Baileys measures — socket-write latency, not
			// server-ACK round-trip time (~100–500 ms).
			elapsed := time.Since(ctx.ReceivedAt)
			ms := float64(elapsed.Microseconds()) / 1000

			// Queue the edit; sendWorker sends it after the pong is delivered.
			ctx.QueueEdit(resp.ID, fmt.Sprintf(T().PongLatency, ms))
			return nil
		},
	})
}
