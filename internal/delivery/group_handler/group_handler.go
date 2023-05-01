package group_handler

import (
	"collector-telegram-bot/internal"
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/usecase"
	"collector-telegram-bot/internal/usecase/group_usecase"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

type GroupTgHandler struct {
	log     internal.Logger
	usecase group_usecase.GroupUsecase
}

func New(log internal.Logger, usecase group_usecase.GroupUsecase) GroupHandler {
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
		return c.Send("Please, add session name after command!")
	}

	sessionName := c.Args()[0]
	info := dto.CreateSessionDTO{
		UserID:      userID,
		ChatID:      chatID,
		Username:    username,
		SessionName: sessionName,
	}

	err = h.usecase.CreateSession(info)
	switch {
	case err == usecase.SessionExistsErr:
		return c.Send("Session is exists now, you can't create new session :(")
	case err != nil:
		h.log.Warnf("Create session err: %v", err)
		return c.Send("Sorry, internal problems")
	default:
		responseText = fmt.Sprintf("Session %s successfully created!", sessionName)
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
