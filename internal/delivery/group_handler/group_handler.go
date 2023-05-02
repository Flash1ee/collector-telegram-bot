package group_handler

import (
	"collector-telegram-bot/internal"
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/models"
	"collector-telegram-bot/internal/usecase"
	"collector-telegram-bot/internal/usecase/group_usecase"
	"fmt"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

const (
	bigSeparateString   = "===========\n"
	smallSeparateString = "----------\n"
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
	return c.Send("Привет!")
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
		return c.Send("Пожалуйста, добавьте название сессии после команды!")
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
		return c.Send("Сессия уже существует – новую создать нельзя. :(")
	case err != nil:
		h.log.Warnf("Create session err: %v", err)
		return c.Send("Извини, технические проблемы")
	default:
		responseText = fmt.Sprintf("Сессия '%s' успешно создана!", sessionName)
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
		return c.Send("Пожалуйста, укажи так: /add <Название продукта> <Цена>!")
	}

	productName := c.Args()[0]
	cost, costErr := strconv.Atoi(c.Args()[1])
	if costErr != nil {
		return c.Send("Цена должна быть целым числом!")
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
		responseText = fmt.Sprintf("Сессия уже существует – новую создать нельзя!")
	case nil:
		responseText = fmt.Sprintf("Добавлена новая трата!")
	default:
		h.log.Warnf("Add expense err: %v", err)
		responseText = "Извини, технические проблемы :("
	}
	return c.Send(responseText)
}

func (h *GroupTgHandler) GetCosts(c tele.Context) error {
	var responseText string
	h.log.Infof("Recieved message from %s, text = %s", c.Message().Sender.Username, c.Text())

	info := dto.GetCostsDTO{
		ChatID: c.Chat().ID,
	}

	allCosts, err := h.usecase.GetAllExpenses(info)

	if err == usecase.SessionNotExistsErr {
		return c.Send("Для выполнения этой команды нужно начать сессию!")
	}

	if err != nil {
		h.log.Warnf("Get costs err: %v", err)
		return c.Send("Извини, техническая ошибка :(")
	}

	if len(allCosts) == 0 {
		return c.Send("Трат пока еще не было :(")
	}

	responseText += "Все траты на текущий момент\n" + bigSeparateString
	responseText += h.createOutput(allCosts)

	return c.Send(responseText)
}

func (h *GroupTgHandler) createOutput(allCosts map[string]models.AllUserCosts) string {
	var responseText string
	for username, allUserCosts := range allCosts {
		responseText += fmt.Sprintf("Пользователь @%s \n", username)
		responseText += fmt.Sprintf("Общая сумма: %d рублей\n"+smallSeparateString, allUserCosts.Sum)

		// Sorting for pretty output
		allUserCosts.SortByCost()

		for _, cost := range allUserCosts.Costs {
			responseText += fmt.Sprintf("%s - %d рублей \n", cost.Description, cost.Money)
		}

		responseText += bigSeparateString
	}
	return responseText
}

func (h *GroupTgHandler) FinishSession(c tele.Context) error {
	var responseText string
	h.log.Infof("Recieved message from %s, text = %s", c.Message().Sender.Username, c.Text())

	info := dto.GetCostsDTO{
		ChatID: c.Chat().ID,
	}

	allCosts, err := h.usecase.GetAllExpenses(info)

	if err == usecase.SessionNotExistsErr {
		return c.Send("Нельзя закончить сессию, если ее еще нет!")
	}

	if err != nil {
		h.log.Warnf("Finish session err: %v", err)
		return c.Send("Извини, технические проблемы")
	}

	err = h.usecase.FinishSession(dto.FinishSessionDTO{ChatID: c.Chat().ID})

	if err != nil {
		h.log.Warnf("Finish session err: %v", err)
		return c.Send("Извини, технические проблемы")
	}

	responseText += "Сессия завершена! Итоговые траты: \n" + bigSeparateString
	responseText += h.createOutput(allCosts)

	return c.Send(responseText)
}

func (h *GroupTgHandler) createOutputDebts(allDebts map[string]models.AllUserDebts) string {
	var responseText string
	for username, allUserDebts := range allDebts {
		responseText += fmt.Sprintf("Пользователю @%s \n", username)

		// Sorting for pretty output
		allUserDebts.SortByDebt()

		for _, cost := range allUserDebts.Debts {
			responseText += fmt.Sprintf("%s - %d рублей \n", cost.DebtorName, cost.Money)
		}

		responseText += bigSeparateString
	}
	return responseText
}

func (h *GroupTgHandler) GetDebts(c tele.Context) error {
	var responseText string
	h.log.Infof("Recieved message from %s, text = %s", c.Message().Sender.Username, c.Text())

	info := dto.GetDebtsDTO{
		ChatID: c.Chat().ID,
	}

	allDebts, err := h.usecase.GetAllDebts(info)

	if err == usecase.SessionNotExistsErr {
		return c.Send("Для выполнения этой команды нужно начать сессию!")
	}

	if err != nil {
		h.log.Warnf("Get debts err: %v", err)
		return c.Send("Извини, техническая ошибка :(")
	}

	if len(allDebts) == 0 {
		return c.Send("Трат пока еще не было :(")
	}

	responseText += "Все долги на текущий момент\n" + bigSeparateString
	responseText += h.createOutputDebts(allDebts)

	return c.Send(responseText)
}
