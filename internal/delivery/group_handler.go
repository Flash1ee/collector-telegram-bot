package delivery

import (
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type GroupHandler interface {
	Great(c tele.Context) error
}

type GroupTgHandler struct {
	Log *logrus.Entry
}

func MakeGroupTgHandler(log *logrus.Entry) GroupHandler {
	return &GroupTgHandler{Log: log}
}

func (h *GroupTgHandler) Great(c tele.Context) error {
	h.Log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return c.Send("Hello! Let's work together!")
}
