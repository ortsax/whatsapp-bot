package plugins

import (
	"context"
	"strings"

	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

// LIDResolver is the subset of the lid_map store needed by the plugin system.
type LIDResolver interface {
	GetLIDForPN(ctx context.Context, pn types.JID) (types.JID, error)
	GetPNForLID(ctx context.Context, lid types.JID) (types.JID, error)
	PutLIDMapping(ctx context.Context, lid, pn types.JID) error
}

var lidResolver LIDResolver
var ownerPhone string // bare phone number of the bot owner (e.g. "2348062795602")

// InitLIDStore wires the LID↔PN resolver and records the owner's phone number.
// ownerPN should be the bare phone number extracted from client.Store.ID.
func InitLIDStore(ls LIDResolver, ownerPN string) {
	lidResolver = ls
	ownerPhone = ownerPN
}

// GetAltID returns the alternate identifier for the given user string.
// Input may be:
//   - a bare phone number ("2348062795602") → returns the LID user part
//   - a full JID string ("2348062795602@s.whatsapp.net") → returns LID user part
//   - a LID user part or full LID JID ("270613692313713@lid") → returns phone number
//
// Returns an empty string when no mapping is found.
func GetAltID(id string) string {
	if lidResolver == nil {
		return ""
	}
	ctx := context.Background()

	var jid types.JID
	if strings.Contains(id, "@") {
		parsed, err := types.ParseJID(id)
		if err != nil {
			return ""
		}
		jid = parsed
	} else {
		// Bare string: treat as phone number.
		jid = types.NewJID(id, types.DefaultUserServer)
	}

	switch jid.Server {
	case types.DefaultUserServer:
		lid, err := lidResolver.GetLIDForPN(ctx, jid)
		if err != nil || lid.User == "" {
			return ""
		}
		return lid.User
	case types.HiddenUserServer:
		pn, err := lidResolver.GetPNForLID(ctx, jid)
		if err != nil || pn.User == "" {
			return ""
		}
		return pn.User
	}
	return ""
}

// SaveUser persists a LID↔PN mapping for the message sender when both can be
// inferred from the event. It is a no-op if the mapping already exists.
//
// Cases handled:
//  1. Incoming DM: sender is LID, chat JID is the sender's PN.
//  2. Our own outgoing message: sender is our LID, ownerPhone is our PN.
func SaveUser(evt *events.Message) {
	if lidResolver == nil {
		return
	}

	sender := evt.Info.Sender
	chat := evt.Info.Chat

	var lidUser, pnUser string

	switch {
	case sender.Server == types.HiddenUserServer &&
		!evt.Info.IsGroup &&
		chat.Server == types.DefaultUserServer &&
		!evt.Info.IsFromMe:
		// Incoming DM: sender is the other person's LID, chat is their PN.
		lidUser = sender.User
		pnUser = chat.User

	case sender.Server == types.HiddenUserServer &&
		evt.Info.IsFromMe &&
		ownerPhone != "":
		// Our own message: sender is our LID, ownerPhone is our PN.
		lidUser = sender.User
		pnUser = ownerPhone
	}

	if lidUser == "" || pnUser == "" {
		return
	}

	lidJID := types.NewJID(lidUser, types.HiddenUserServer)
	pnJID := types.NewJID(pnUser, types.DefaultUserServer)
	_ = lidResolver.PutLIDMapping(context.Background(), lidJID, pnJID)
}

// BootstrapOwnerSudoers adds the owner's phone (and LID if already resolvable)
// to the sudoers list and persists the change. Safe to call multiple times.
func BootstrapOwnerSudoers() {
	if ownerPhone == "" {
		return
	}
	changed := false

	if !BotSettings.IsSudo(ownerPhone) {
		BotSettings.AddSudo(ownerPhone)
		changed = true
	}

	if lid := GetAltID(ownerPhone); lid != "" && !BotSettings.IsSudo(lid) {
		BotSettings.AddSudo(lid)
		changed = true
	}

	if changed {
		_ = SaveSettings()
	}
}
