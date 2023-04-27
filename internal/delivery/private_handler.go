package delivery

import (
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type PrivateHandler interface {
	Info(c tele.Context) error
	Start(c tele.Context) error
}

type PrivateTgHandler struct {
	Log *logrus.Entry
}

func MakePrivateTgHandler(log *logrus.Entry) PrivateHandler {
	return &PrivateTgHandler{Log: log}
}

func (h *PrivateTgHandler) Info(c tele.Context) error {
	h.Log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())

	return c.Send("Hello! You can work with this bot!")
}

func (h *PrivateTgHandler) Start(c tele.Context) error {
	h.Log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return c.Send("Hello! Let's work together!")
}
