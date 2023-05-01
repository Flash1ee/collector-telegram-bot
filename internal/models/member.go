package models

import "collector-telegram-bot/internal"

type Member struct {
	ID          uint64
	SessionUUID internal.UUID
	UserID      uint64
}

func NewEmptyMember() *Member {
	return &Member{}
}

func NewMember(sessionUUID internal.UUID, userID uint64) *Member {
	return &Member{
		ID:          0,
		SessionUUID: sessionUUID,
		UserID:      userID,
	}
}
