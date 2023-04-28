package useCase

import (
	repo "collector-telegram-bot/internal/repository"
	"github.com/sirupsen/logrus"
)

type GroupUseCase interface {
}

type AppGroupUseCase struct {
	log  *logrus.Entry
	repo repo.Repository
}

func MakeGroupUseCase(log *logrus.Entry, repo repo.Repository) GroupUseCase {
	return &AppGroupUseCase{log: log, repo: repo}
}

func (uc *AppGroupUseCase) Sessions() {}
