package models

type Member struct {
	ID          int64
	SessionUUID int64
	UserID      int64
}

func NewEmptyMember() *Member {
	return &Member{}
}

func NewMember(sessionUUID, userID int64) *Member {
	return &Member{
		ID:          0,
		SessionUUID: sessionUUID,
		UserID:      userID,
	}
}
