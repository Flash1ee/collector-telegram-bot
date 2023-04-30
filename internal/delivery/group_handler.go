package delivery

import (
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/usecase"
	"fmt"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type GroupHandler interface {
	Great(c tele.Context) error
	StartSession(c tele.Context) error
}

type GroupTgHandler struct {
	log     *logrus.Entry
	usecase usecase.GroupUsecase
}

func NewGroupTgHandler(log *logrus.Entry, usecase usecase.GroupUsecase) GroupHandler {
	return &GroupTgHandler{log: log, usecase: usecase}
}

func (h *GroupTgHandler) Great(c tele.Context) error {
	h.log.Infof("Recieved message from %s, text = %s", c.Chat().Username, c.Text())
	return c.Send("Hello! Let's work together!")
}

func (h *GroupTgHandler) StartSession(c tele.Context) error {
	var (
		err          error
		responseText string
	)
	h.log.Infof("Recieved message from %s, text = %s", c.Message().Sender.Username, c.Text())

	chatID := c.Chat().ID
	userID := c.Message().Sender.ID
	userName := c.Message().Sender.Username

	h.log.Infof("Params: %d %s %d", userID, userName, chatID)
	if len(c.Args()) == 0 {
		responseText = fmt.Sprintf("Please, add session name after command!")
	} else {
		sessionName := c.Args()[0]
		dto := dto.CreateSessionDTO{
			UserID:      0,
			ChatID:      chatID,
			Username:    userName,
			SessionName: sessionName,
		}

		err = h.usecase.CreateSession(dto)
		if err == nil {
			responseText = fmt.Sprintf("Session %s successfully created!", sessionName)
		} else {
			h.log.Warnf("Create session err: %v", err)
			responseText = fmt.Sprintf("Sorry, internal problems")
		}
	}
	return c.Send(responseText)
}
