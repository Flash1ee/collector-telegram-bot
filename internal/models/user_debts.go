package models

import "sort"

type UserDebt struct {
	DebtorName string
	Money      int
}

type AllUserDebts struct {
	Debts []UserDebt
}

func (d *AllUserDebts) SortByDebt() {
	sort.Slice(d.Debts, func(i, j int) bool {
		return d.Debts[i].Money > d.Debts[j].Money
	})
}
