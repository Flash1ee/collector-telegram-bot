package group_usecase

import "collector-telegram-bot/internal/dto"

type GroupUsecase interface {
	CreateSession(info dto.CreateSessionDTO) error
	AddExpenseToSession(info dto.AddExpenseDTO) error
}
