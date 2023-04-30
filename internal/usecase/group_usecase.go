package usecase

import (
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/models"
	repo "collector-telegram-bot/internal/repository"
	"fmt"
	"github.com/sirupsen/logrus"
)

type GroupUsecase interface {
	CreateSession(info dto.CreateSessionDTO) error
}

type AppGroupUsecase struct {
	log  *logrus.Entry
	repo repo.Repository
}

func NewGroupUsecase(log *logrus.Entry, repo repo.Repository) GroupUsecase {
	return &AppGroupUsecase{log: log, repo: repo}
}

func (uc *AppGroupUsecase) CreateSession(info dto.CreateSessionDTO) error {
	user, err := uc.repo.GetUser(info.UserID)

	if err != nil {
		return fmt.Errorf("usecase: %v", err)
	}

	// If user not exists, you should create user
	if user.ID == 0 {
		user.Username = info.Username
		user.TgID = info.UserID
		err = uc.repo.CreateUser(user)
		if err != nil {
			return fmt.Errorf("usecase: %v", err)
		}
	}

	// Check that user was created and get id from Database
	user, err = uc.repo.GetUser(info.UserID)
	if err != nil {
		return fmt.Errorf("usecase: %v", err)
	}

	session := models.NewSession(user.ID, info.ChatID, info.SessionName)
	return uc.repo.CreateNewSession(session)
}
