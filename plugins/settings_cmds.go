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
			if len(ctx.Args) < 1 {
				ctx.Reply(T().SetSudoUsage)
				return nil
			}
			action := strings.ToLower(ctx.Args[0])

			// Target: reply (no arg), @mention, bare phone, or LID.
			targetArg := ""
			if len(ctx.Args) >= 2 {
				targetArg = ctx.Args[1]
			}
			phone, lid := ResolveTarget(ctx, targetArg)
			if phone == "" && lid == "" {
				ctx.Reply(T().SetSudoUsage)
				return nil
			}

			display := phone
			if display == "" {
				display = lid
			}

			switch action {
			case "add":
				if phone != "" {
					BotSettings.AddSudo(phone)
				}
				if lid != "" {
					BotSettings.AddSudo(lid)
				}
				if err := SaveSettings(); err != nil {
					ctx.Reply(fmt.Sprintf(T().SaveFailed, err.Error()))
					return err
				}
				ctx.Reply(fmt.Sprintf(T().SudoAdded, display))
			case "remove":
				removed := false
				if phone != "" && BotSettings.RemoveSudo(phone) {
					removed = true
				}
				if lid != "" && BotSettings.RemoveSudo(lid) {
					removed = true
				}
				if removed {
					_ = SaveSettings()
					ctx.Reply(fmt.Sprintf(T().SudoRemoved, display))
				} else {
					ctx.Reply(fmt.Sprintf(T().SudoNotFound, display))
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
