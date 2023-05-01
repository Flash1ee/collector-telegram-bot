package group_usecase

import (
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/models"
)

type GroupUsecase interface {
	CreateSession(info dto.CreateSessionDTO) error
	AddExpenseToSession(info dto.AddExpenseDTO) error
	GetAllExpenses(info dto.GetCostsDTO) (map[string]models.AllUserCosts, error)
	FinishSession(info dto.FinishSessionDTO) error
}
