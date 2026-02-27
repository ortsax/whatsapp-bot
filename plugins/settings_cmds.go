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
				ctx.Reply(T().SetPrefixUsage)
				return nil
			}
			BotSettings.SetPrefixes(ctx.Text)
			if err := SaveSettings(); err != nil {
				ctx.Reply(fmt.Sprintf(T().SaveFailed, err.Error()))
				return err
			}
			display := strings.Join(BotSettings.GetPrefixes(), "  ")
			if display == "" {
				display = "(empty)"
			}
			ctx.Reply(fmt.Sprintf(T().SetPrefixUpdated, display))
			return nil
		},
	})

	Register(&Command{
		Pattern:  "setsudo",
		IsSudo:   true,
		Category: "settings",
		Func: func(ctx *Context) error {
			if len(ctx.Args) < 2 {
				ctx.Reply(T().SetSudoUsage)
				return nil
			}
			action := strings.ToLower(ctx.Args[0])
			phone := ctx.Args[1]
			switch action {
			case "add":
				BotSettings.AddSudo(phone)
				if err := SaveSettings(); err != nil {
					ctx.Reply(fmt.Sprintf(T().SaveFailed, err.Error()))
					return err
				}
				ctx.Reply(fmt.Sprintf(T().SudoAdded, phone))
			case "remove":
				if BotSettings.RemoveSudo(phone) {
					_ = SaveSettings()
					ctx.Reply(fmt.Sprintf(T().SudoRemoved, phone))
				} else {
					ctx.Reply(fmt.Sprintf(T().SudoNotFound, phone))
				}
			default:
				ctx.Reply(T().UnknownAction)
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
				ctx.Reply(T().ModePublicSet)
			case "private":
				BotSettings.SetMode(ModePrivate)
				_ = SaveSettings()
				ctx.Reply(T().ModePrivateSet)
			default:
				ctx.Reply(T().SetModeUsage)
			}
			return nil
		},
	})
}
