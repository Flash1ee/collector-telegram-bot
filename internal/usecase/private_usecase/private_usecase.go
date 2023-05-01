package private_usecase

import (
	"collector-telegram-bot/internal"
	"collector-telegram-bot/internal/repository"
	tele "gopkg.in/telebot.v3"
)

type AppPrivateUsecase struct {
	log  internal.Logger
	repo repo.Repository
}

func New(log internal.Logger, repo repo.Repository) PrivateUsecase {
	return &AppPrivateUsecase{log: log, repo: repo}
}

func (uc *AppPrivateUsecase) Sessions(c tele.Context) {}
