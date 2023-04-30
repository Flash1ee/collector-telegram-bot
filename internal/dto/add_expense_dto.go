package dto

type AddExpenseDTO struct {
	Product  string
	ChatID   int64
	UserID   int64
	Username string
	Cost     int
}
