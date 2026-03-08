package db

import "database/sql"

var settingsDB *sql.DB

// InitDB stores d in settingsDB and creates all required tables.
func InitDB(d *sql.DB) error {
	settingsDB = d
	return initTables()
}

// DB returns the active database handle.
func DB() *sql.DB {
	return settingsDB
}

func initTables() error {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS bot_settings (
			user  TEXT,
			key   TEXT,
			value TEXT,
			PRIMARY KEY (user, key)
		)`,
		`CREATE TABLE IF NOT EXISTS warns (
			chat_jid TEXT,
			user_id  TEXT,
			count    INTEGER DEFAULT 0,
			PRIMARY KEY (chat_jid, user_id)
		)`,
		`CREATE TABLE IF NOT EXISTS shh_users (
			chat_jid TEXT,
			user_id  TEXT,
			PRIMARY KEY (chat_jid, user_id)
		)`,
		`CREATE TABLE IF NOT EXISTS antilink_settings (
			chat_jid TEXT PRIMARY KEY,
			mode     TEXT DEFAULT 'off'
		)`,
		`CREATE TABLE IF NOT EXISTS antiword_settings (
			chat_jid TEXT,
			word     TEXT,
			PRIMARY KEY (chat_jid, word)
		)`,
		`CREATE TABLE IF NOT EXISTS antispam_settings (
			chat_jid TEXT PRIMARY KEY,
			mode     TEXT DEFAULT 'off'
		)`,
		`CREATE TABLE IF NOT EXISTS antispam_whitelist (
			chat_jid TEXT,
			user_id  TEXT,
			PRIMARY KEY (chat_jid, user_id)
		)`,
		`CREATE TABLE IF NOT EXISTS afk_status (
			user_id TEXT PRIMARY KEY,
			message TEXT,
			set_at  INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS filters (
			scope    TEXT,
			chat_jid TEXT,
			keyword  TEXT,
			response TEXT,
			PRIMARY KEY (scope, chat_jid, keyword)
		)`,
		`CREATE TABLE IF NOT EXISTS antistatus_settings (
			chat_jid TEXT PRIMARY KEY,
			enabled  INTEGER DEFAULT 0
		)`,
		`CREATE TABLE IF NOT EXISTS antidelete_cache (
			msg_id       TEXT PRIMARY KEY,
			chat_jid     TEXT NOT NULL,
			sender_jid   TEXT NOT NULL,
			sender_alt   TEXT NOT NULL DEFAULT '',
			is_from_me   INTEGER NOT NULL DEFAULT 0,
			msg_ts       INTEGER NOT NULL,
			message_blob BLOB NOT NULL,
			cached_at    INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS anticall_settings (
			key   TEXT PRIMARY KEY,
			value TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS meta_messages (
			msg_id        TEXT PRIMARY KEY,
			chat_jid      TEXT NOT NULL,
			response_text TEXT NOT NULL,
			created_at    INTEGER NOT NULL
		)`,
	}
	for _, q := range tables {
		if _, err := settingsDB.Exec(q); err != nil {
			return err
		}
	}
	return nil
}
