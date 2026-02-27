package plugins

import (
	"database/sql"
	"encoding/json"
	"strings"
	"sync"
)

// Mode controls who can use bot commands.
type Mode string

const (
	ModePublic  Mode = "public"
	ModePrivate Mode = "private"
)

// Settings holds the in-memory bot configuration.
type Settings struct {
	mu       sync.RWMutex
	Prefixes []string
	Sudoers  []string
	Mode     Mode
	Language string
}

// BotSettings is the global settings instance, seeded with defaults.
var BotSettings = &Settings{
	Prefixes: []string{"."},
	Sudoers:  []string{},
	Mode:     ModePublic,
	Language: "en",
}

var settingsDB   *sql.DB
var settingsUser string // bare phone number of the bot owner

// InitDB creates the bot_settings table if it doesn't exist.
// Call this as soon as the database is available (before Connect).
func InitDB(db *sql.DB) error {
	settingsDB = db
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS bot_settings (
		user  TEXT NOT NULL,
		key   TEXT NOT NULL,
		value TEXT NOT NULL,
		PRIMARY KEY (user, key)
	)`)
	return err
}

// InitSettings sets the active user and loads their settings from the database.
// Call this after Connect() once the owner phone is known.
// If no rows exist for this user the in-memory defaults are kept and a first
// save is written so the row appears in the table.
func InitSettings(user string) error {
	settingsUser = user
	if err := LoadSettings(); err != nil {
		return err
	}
	// Write defaults for this user if nothing was stored yet.
	return SaveSettings()
}

// LoadSettings reads all rows for the current user from bot_settings.
func LoadSettings() error {
	if settingsUser == "" {
		return nil
	}
	rows, err := settingsDB.Query(
		`SELECT key, value FROM bot_settings WHERE user = ?`, settingsUser)
	if err != nil {
		return err
	}
	defer rows.Close()

	BotSettings.mu.Lock()
	defer BotSettings.mu.Unlock()

	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return err
		}
		switch key {
		case "prefixes":
			var p []string
			if json.Unmarshal([]byte(value), &p) == nil {
				BotSettings.Prefixes = p
			}
		case "sudoers":
			var s []string
			if json.Unmarshal([]byte(value), &s) == nil {
				BotSettings.Sudoers = s
			}
		case "mode":
			BotSettings.Mode = Mode(value)
		case "language":
			BotSettings.Language = value
		}
	}
	return rows.Err()
}

// SaveSettings persists the entire current state to the database for the active user.
func SaveSettings() error {
	if settingsUser == "" {
		return nil
	}

	BotSettings.mu.RLock()
	prefixes := BotSettings.Prefixes
	sudoers := BotSettings.Sudoers
	mode := BotSettings.Mode
	language := BotSettings.Language
	BotSettings.mu.RUnlock()

	pData, _ := json.Marshal(prefixes)
	sData, _ := json.Marshal(sudoers)

	upsert := `INSERT INTO bot_settings (user, key, value) VALUES (?, ?, ?)
		ON CONFLICT(user, key) DO UPDATE SET value = excluded.value`

	tx, err := settingsDB.Begin()
	if err != nil {
		return err
	}
	for _, row := range [][2]string{
		{"prefixes", string(pData)},
		{"sudoers", string(sData)},
		{"mode", string(mode)},
		{"language", language},
	} {
		if _, err = tx.Exec(upsert, settingsUser, row[0], row[1]); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (s *Settings) IsSudo(phone string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.Sudoers {
		if p == phone {
			return true
		}
	}
	return false
}

func (s *Settings) GetPrefixes() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]string, len(s.Prefixes))
	copy(result, s.Prefixes)
	return result
}

func (s *Settings) GetMode() Mode {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Mode
}

// SetPrefixes parses a space-separated list of prefixes.
// Use the token "empty" to include an empty (no-prefix) entry.
func (s *Settings) SetPrefixes(raw string) {
	parts := strings.Fields(raw)
	for i, p := range parts {
		if strings.ToLower(p) == "empty" {
			parts[i] = ""
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Prefixes = parts
}

func (s *Settings) AddSudo(phone string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.Sudoers {
		if p == phone {
			return
		}
	}
	s.Sudoers = append(s.Sudoers, phone)
}

// RemoveSudo removes a phone from sudoers and returns true if it was present.
func (s *Settings) RemoveSudo(phone string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.Sudoers {
		if p == phone {
			s.Sudoers = append(s.Sudoers[:i], s.Sudoers[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Settings) SetMode(m Mode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Mode = m
}

func (s *Settings) GetLanguage() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.Language == "" {
		return "en"
	}
	return s.Language
}

func (s *Settings) SetLanguage(lang string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Language = lang
}
