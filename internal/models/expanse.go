package models

type Expanse struct {
	Username    string
	Cost        int
	Description string
}

func NewEmptyExpanse() *Expanse {
	return &Expanse{}
}

func NewExpanse(username string, cost int, description string) *Expanse {
	return &Expanse{
		Username:    username,
		Cost:        cost,
		Description: description,
	}
}
