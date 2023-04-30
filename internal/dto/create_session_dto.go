package dto

type CreateSessionDTO struct {
	UserID      int64
	ChatID      int64
	Username    string
	SessionName string
}
