package models

type Cost struct {
	UserID uint64
	Money  int
}

func NewEmptyCost() *Cost {
	return &Cost{}
}

func NewCost(userID uint64, money int) *Cost {
	return &Cost{
		UserID: userID,
		Money:  money,
	}
}
