package private_handler

import tele "gopkg.in/telebot.v3"

type PrivateHandler interface {
	Info(c tele.Context) error
	Start(c tele.Context) error
	Sessions(c tele.Context) error
}
