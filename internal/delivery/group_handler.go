package delivery

import (
	"collector-telegram-bot/internal/useCase"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type GroupHandler interface {
	Great(c tele.Context) error
}

type GroupTgHandler struct {
	log     *logrus.Entry
	useCase useCase.GroupUseCase
}

func MakeGroupTgHandler(log *logrus.Entry, useCase useCase.GroupUseCase) GroupHandler {
	return &GroupTgHandler{log: log, useCase: useCase}
}

func (h *GroupTgHandler) Great(c tele.Context) error {
	h.log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return c.Send("Hello! Let's work together!")
}
