package db

// CachedMsgRow holds the raw columns from antidelete_cache.
type CachedMsgRow struct {
	ChatJID, SenderJID, SenderAlt string
	IsFromMe                      int
	MsgTS                         int64
	Blob                          []byte
}

// InsertCachedMessage stores a raw serialised message in antidelete_cache.
func InsertCachedMessage(msgID, chatJID, senderJID, senderAlt string, isFromMe int, msgTS, cachedAt int64, blob []byte) {
	if settingsDB == nil {
		return
	}
	settingsDB.Exec(
		`INSERT OR REPLACE INTO antidelete_cache (msg_id, chat_jid, sender_jid, sender_alt, is_from_me, msg_ts, message_blob, cached_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		msgID, chatJID, senderJID, senderAlt, isFromMe, msgTS, blob, cachedAt,
	)
}

// PopCachedMessage fetches and immediately deletes the entry for msgID.
func PopCachedMessage(msgID string) (CachedMsgRow, bool) {
	if settingsDB == nil {
		return CachedMsgRow{}, false
	}
	var row CachedMsgRow
	err := settingsDB.QueryRow(
		`SELECT chat_jid, sender_jid, sender_alt, is_from_me, msg_ts, message_blob FROM antidelete_cache WHERE msg_id = ?`, msgID,
	).Scan(&row.ChatJID, &row.SenderJID, &row.SenderAlt, &row.IsFromMe, &row.MsgTS, &row.Blob)
	if err != nil {
		return CachedMsgRow{}, false
	}
	settingsDB.Exec(`DELETE FROM antidelete_cache WHERE msg_id = ?`, msgID)
	return row, true
}

// PruneCache deletes all cache entries older than cutoff (unix timestamp).
func PruneCache(cutoff int64) {
	if settingsDB == nil {
		return
	}
	settingsDB.Exec(`DELETE FROM antidelete_cache WHERE cached_at < ?`, cutoff)
}
