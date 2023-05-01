package group_handler

import tele "gopkg.in/telebot.v3"

type GroupHandler interface {
	Great(c tele.Context) error
	StartSession(c tele.Context) error
	AddExpense(c tele.Context) error
	GetCosts(c tele.Context) error
	FinishSession(c tele.Context) error
}
