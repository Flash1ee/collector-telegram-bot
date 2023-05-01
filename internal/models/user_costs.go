package models

import "sort"

type UserCost struct {
	Money       int
	Description string
}

type AllUserCosts struct {
	Sum   int
	Costs []UserCost
}

func (c *AllUserCosts) SortByCost() {
	sort.Slice(c.Costs, func(i, j int) bool {
		return c.Costs[i].Money > c.Costs[j].Money
	})
}
