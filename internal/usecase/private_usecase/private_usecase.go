package usecase

import (
	"collector-telegram-bot/internal/repository"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type PrivateUsecase interface {
	Sessions(c tele.Context)
}

type AppPrivateUsecase struct {
	log  *logrus.Entry
	repo repo.Repository
}

func NewPrivateUsecase(log *logrus.Entry, repo repo.Repository) PrivateUsecase {
	return &AppPrivateUsecase{log: log, repo: repo}
}

func (uc *AppPrivateUsecase) Sessions(c tele.Context) {}
