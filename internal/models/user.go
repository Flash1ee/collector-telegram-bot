package models

type User struct {
	ID         int64
	Username   string
	CreatedAt  string
	Requisites string
	TgID       int64
}

func NewUser() *User {
	return &User{}
}
