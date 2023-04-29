package usecase

import (
	repo "collector-telegram-bot/internal/repository"
	"github.com/sirupsen/logrus"
)

type GroupUsecase interface {
}

type AppGroupUsecase struct {
	log  *logrus.Entry
	repo repo.Repository
}

func NewGroupUsecase(log *logrus.Entry, repo repo.Repository) GroupUsecase {
	return &AppGroupUsecase{log: log, repo: repo}
}

func (uc *AppGroupUsecase) Sessions() {}
