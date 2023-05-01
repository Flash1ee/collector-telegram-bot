package delivery

import (
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/usecase"
	"fmt"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

type GroupHandler interface {
	Great(c tele.Context) error
	StartSession(c tele.Context) error
	AddExpense(c tele.Context) error
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
	username := c.Message().Sender.Username

	if len(c.Args()) == 0 {
		responseText = fmt.Sprintf("Please, add session name after command!")
	} else {
		sessionName := c.Args()[0]
		info := dto.CreateSessionDTO{
			UserID:      userID,
			ChatID:      chatID,
			Username:    username,
			SessionName: sessionName,
		}

		err = h.usecase.CreateSession(info)
		if err == nil {
			responseText = fmt.Sprintf("Session %s successfully created!", sessionName)
		} else {
			h.log.Warnf("Create session err: %v", err)
			responseText = fmt.Sprintf("Sorry, internal problems")
		}
	}
	return c.Send(responseText)
}

func (h *GroupTgHandler) AddExpense(c tele.Context) error {
	var (
		err          error
		responseText string
	)
	h.log.Infof("Recieved message from %s, text = %s", c.Message().Sender.Username, c.Text())

	if len(c.Args()) != 2 {
		return c.Send("Please, create command with params /add <SessionName> <Cost>!")
	}

	productName := c.Args()[0]
	cost, costErr := strconv.Atoi(c.Args()[1])
	if costErr != nil {
		return c.Send("Cost must be integer!")
	}
	info := dto.AddExpenseDTO{
		ChatID:   c.Chat().ID,
		Product:  productName,
		Cost:     cost,
		UserID:   c.Message().Sender.ID,
		Username: c.Message().Sender.Username,
	}

	err = h.usecase.AddExpenseToSession(info)
	switch err {
	case usecase.SessionNotExistsErr:
		responseText = fmt.Sprintf("You should start session!")
	case nil:
		responseText = fmt.Sprintf("Expense is added!")
	default:
		h.log.Warnf("Create session err: %v", err)
		responseText = fmt.Sprintf("Sorry, internal problems")
	}
	return c.Send(responseText)
}
