package private_handler

import (
	"collector-telegram-bot/internal"
	"collector-telegram-bot/internal/usecase/private_usecase"
	tele "gopkg.in/telebot.v3"
)

type PrivateTgHandler struct {
	log     internal.Logger
	usecase private_usecase.PrivateUsecase
}

func New(log internal.Logger, usecase private_usecase.PrivateUsecase) PrivateHandler {
	return &PrivateTgHandler{log: log, usecase: usecase}
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
