package models

type Session struct {
	UUID        int64
	CreatorID   int64
	ChatID      int64
	SessionName string
	StartedAt   string
	State       string
}

func NewSession(creatorID, chatID int64, sessionName string) *Session {
	return &Session{
		CreatorID:   creatorID,
		ChatID:      chatID,
		SessionName: sessionName,
		State:       "active",
	}
}
