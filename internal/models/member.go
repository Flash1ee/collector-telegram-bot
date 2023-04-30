package models

type Member struct {
	ID          uint64
	SessionUUID uint64
	UserID      uint64
}

func NewEmptyMember() *Member {
	return &Member{}
}

func NewMember(sessionUUID, userID uint64) *Member {
	return &Member{
		ID:          0,
		SessionUUID: sessionUUID,
		UserID:      userID,
	}
}
