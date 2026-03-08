package db

// ChatActivity holds a chat JID and its message count.
type ChatActivity struct {
	JID   string
	Count int
}

// GetTopChats returns the top 30 chats by message count from message_secrets.
func GetTopChats() ([]ChatActivity, error) {
	rows, err := settingsDB.Query(`SELECT chat_jid, COUNT(*) as cnt FROM message_secrets WHERE chat_jid != 'status@broadcast' GROUP BY chat_jid ORDER BY cnt DESC LIMIT 30`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []ChatActivity
	for rows.Next() {
		var a ChatActivity
		if rows.Scan(&a.JID, &a.Count) == nil {
			result = append(result, a)
		}
	}
	return result, nil
}

// GetContactName returns the push_name for a full JID string (e.g. "123@s.whatsapp.net").
func GetContactName(jidStr string) string {
	var name string
	settingsDB.QueryRow(`SELECT push_name FROM contacts WHERE their_jid = ?`, jidStr).Scan(&name)
	return name
}

// GetContactNameByLID returns the push_name for a LID user part (without @lid).
func GetContactNameByLID(lidUser string) string {
	var name string
	settingsDB.QueryRow(
		`SELECT c.push_name FROM lid_map l JOIN contacts c ON c.their_jid = l.pn || '@s.whatsapp.net' WHERE l.lid = ?`, lidUser,
	).Scan(&name)
	return name
}

// SenderCount holds a sender JID and message count.
type SenderCount struct {
	SenderJID string
	Count     int
}

// GetActiveSenders returns senders in a chat ordered by message count descending (limit 20).
func GetActiveSenders(chatJID string) ([]SenderCount, error) {
	rows, err := settingsDB.Query(`SELECT sender_jid, COUNT(*) as cnt FROM message_secrets WHERE chat_jid = ? GROUP BY sender_jid ORDER BY cnt DESC LIMIT 20`, chatJID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []SenderCount
	for rows.Next() {
		var s SenderCount
		if rows.Scan(&s.SenderJID, &s.Count) == nil {
			result = append(result, s)
		}
	}
	return result, nil
}

// GetAllSenderCounts returns all senders in a chat with their message counts.
func GetAllSenderCounts(chatJID string) ([]SenderCount, error) {
	rows, err := settingsDB.Query(`SELECT sender_jid, COUNT(*) as cnt FROM message_secrets WHERE chat_jid = ? GROUP BY sender_jid`, chatJID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []SenderCount
	for rows.Next() {
		var s SenderCount
		if rows.Scan(&s.SenderJID, &s.Count) == nil {
			result = append(result, s)
		}
	}
	return result, nil
}
