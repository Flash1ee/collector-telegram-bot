package delivery

import (
	"collector-telegram-bot/internal/useCase"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type PrivateHandler interface {
	Info(c tele.Context) error
	Start(c tele.Context) error
	Sessions(c tele.Context) error
}

type PrivateTgHandler struct {
	log     *logrus.Entry
	useCase useCase.PrivateUseCase
}

func MakePrivateTgHandler(log *logrus.Entry, useCase useCase.PrivateUseCase) PrivateHandler {
	return &PrivateTgHandler{log: log, useCase: useCase}
}

func (h *PrivateTgHandler) Info(c tele.Context) error {
	h.log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())

	return c.Send("Hello! You can work with this bot!")
}

func (h *PrivateTgHandler) Start(c tele.Context) error {
	h.log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return c.Send("Hello! Let's work together!")
}

func (h *PrivateTgHandler) Sessions(c tele.Context) error {
	h.log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return nil
}
