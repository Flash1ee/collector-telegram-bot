package useCase

import (
	"collector-telegram-bot/internal/repository"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type PrivateUseCase interface {
	Sessions(c tele.Context)
}

type AppPrivateUseCase struct {
	log  *logrus.Entry
	repo repo.Repository
}

func MakePrivateUseCase(log *logrus.Entry, repo repo.Repository) PrivateUseCase {
	return &AppPrivateUseCase{log: log, repo: repo}
}

func (uc *AppPrivateUseCase) Sessions(c tele.Context) {}
