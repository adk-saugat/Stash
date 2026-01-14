package models

import (
	"encoding/json"
	"os"
	"time"

	"github.com/adk-saugat/stash/client/utils"
)

type Session struct {
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}

// SaveSession saves the session to .stash/session.json
func SaveSession(token, email string) error {
	// Ensure .stash directory exists
	err := utils.EnsureFolderExists(".stash")
	if err != nil && err != os.ErrExist {
		return err
	}

	session := Session{
		Token:     token,
		Email:     email,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	data, err := json.MarshalIndent(session, "", "    ")
	if err != nil {
		return err
	}

	return utils.WriteFileData(".stash/session.json", data)
}

// LoadSession loads the session from .stash/session.json
func LoadSession() (*Session, error) {
	data, err := os.ReadFile(".stash/session.json")
	if err != nil {
		return nil, err
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// IsSessionValid checks if session exists and is not expired
func IsSessionValid() bool {
	session, err := LoadSession()
	if err != nil {
		return false
	}
	return time.Now().Before(session.ExpiresAt)
}
