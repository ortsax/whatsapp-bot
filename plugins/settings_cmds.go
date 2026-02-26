package plugins

import (
	"fmt"
	"strings"
)

func init() {
	Register(&Command{
		Pattern:  "setprefix",
		IsSudo:   true,
		Category: "settings",
		Func: func(ctx *Context) error {
			if ctx.Text == "" {
				ctx.Reply("Usage: .setprefix <p1> <p2> ...\nUse the token *empty* for a no-prefix entry.\nExample: .setprefix . / #")
				return nil
			}
			BotSettings.SetPrefixes(ctx.Text)
			if err := SaveSettings(); err != nil {
				ctx.Reply("❌ Failed to save settings: " + err.Error())
				return err
			}
			display := strings.Join(BotSettings.GetPrefixes(), "  ")
			if display == "" {
				display = "(empty)"
			}
			ctx.Reply(fmt.Sprintf("✅ Prefix updated: %s", display))
			return nil
		},
	})

	Register(&Command{
		Pattern:  "setsudo",
		IsSudo:   true,
		Category: "settings",
		Func: func(ctx *Context) error {
			if len(ctx.Args) < 2 {
				ctx.Reply("Usage: .setsudo add|remove <phone>\nExample: .setsudo add 1234567890")
				return nil
			}
			action := strings.ToLower(ctx.Args[0])
			phone := ctx.Args[1]
			switch action {
			case "add":
				BotSettings.AddSudo(phone)
				if err := SaveSettings(); err != nil {
					ctx.Reply("❌ Failed to save settings: " + err.Error())
					return err
				}
				ctx.Reply(fmt.Sprintf("✅ %s added as sudo user.", phone))
			case "remove":
				if BotSettings.RemoveSudo(phone) {
					_ = SaveSettings()
					ctx.Reply(fmt.Sprintf("✅ %s removed from sudo users.", phone))
				} else {
					ctx.Reply(fmt.Sprintf("❌ %s is not a sudo user.", phone))
				}
			default:
				ctx.Reply("❌ Unknown action. Use: add or remove")
			}
			return nil
		},
	})

	Register(&Command{
		Pattern:  "setmode",
		IsSudo:   true,
		Category: "settings",
		Func: func(ctx *Context) error {
			switch strings.ToLower(ctx.Text) {
			case "public":
				BotSettings.SetMode(ModePublic)
				_ = SaveSettings()
				ctx.Reply("✅ Mode set to: *public* – anyone can use commands.")
			case "private":
				BotSettings.SetMode(ModePrivate)
				_ = SaveSettings()
				ctx.Reply("✅ Mode set to: *private* – sudo users only.")
			default:
				ctx.Reply("Usage: .setmode public|private")
			}
			return nil
		},
	})
}
