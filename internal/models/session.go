package models

import "github.com/google/uuid"

const SessionActive = "active"

type Session struct {
	UUID        uuid.UUID
	CreatorID   uint64
	ChatID      int64
	SessionName string
	StartedAt   string
	State       string
}

func NewSession(UUID uuid.UUID, creatorID uint64, chatID int64, sessionName string) *Session {
	return &Session{
		UUID:        UUID,
		CreatorID:   creatorID,
		ChatID:      chatID,
		SessionName: sessionName,
		State:       SessionActive,
	}
}

func NewEmptySession() *Session {
	return &Session{}
}
