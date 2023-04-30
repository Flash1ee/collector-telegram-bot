package models

type Session struct {
	UUID        uint64
	CreatorID   uint64
	ChatID      int64
	SessionName string
	StartedAt   string
	State       string
}

func NewSession(creatorID uint64, chatID int64, sessionName string) *Session {
	return &Session{
		CreatorID:   creatorID,
		ChatID:      chatID,
		SessionName: sessionName,
		State:       "active",
	}
}

func NewEmptySession() *Session {
	return &Session{}
}
