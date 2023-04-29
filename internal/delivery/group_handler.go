package delivery

import (
	"collector-telegram-bot/internal/usecase"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type GroupHandler interface {
	Great(c tele.Context) error
}

type GroupTgHandler struct {
	log     *logrus.Entry
	usecase usecase.GroupUsecase
}

func MakeGroupTgHandler(log *logrus.Entry, usecase usecase.GroupUsecase) GroupHandler {
	return &GroupTgHandler{log: log, usecase: usecase}
}

func (h *GroupTgHandler) Great(c tele.Context) error {
	h.log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return c.Send("Hello! Let's work together!")
}
