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
	AddExpenseToSession(info dto.AddExpenseDTO) error
}

type AppGroupUsecase struct {
	log  *logrus.Entry
	repo repo.Repository
}

func NewGroupUsecase(log *logrus.Entry, repo repo.Repository) GroupUsecase {
	return &AppGroupUsecase{log: log, repo: repo}
}

func (uc *AppGroupUsecase) CreateSession(info dto.CreateSessionDTO) error {
	userID, err := uc.upsertUser(info.UserID, info.Username)

	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	// Check, that aren't any active sessions in chat
	curSession, err := uc.repo.GetActiveSessionByChatID(info.ChatID)
	switch {
	case err != nil:
		return err
	case curSession.UUID != 0:
		return SessionExistsErr
	default:
	}

	session := models.NewSession(userID, info.ChatID, info.SessionName)
	err = uc.repo.CreateNewSession(session)
	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	// Check that session is created and get uuid
	curSession, err = uc.repo.GetActiveSessionByChatID(info.ChatID)
	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}
	// Add creator to members
	return uc.repo.AddMemberToSession(curSession.UUID, userID)
}

func (uc *AppGroupUsecase) upsertUser(userID int64, username string) (int64, error) {
	user, err := uc.repo.GetUser(userID)

	if err != nil {
		return 0, fmt.Errorf("usecase: %v", err)
	}

	// If user not exists, you should create user
	if user.ID == 0 {
		user.Username = username
		user.TgID = userID
		err = uc.repo.CreateUser(user)
		if err != nil {
			return 0, fmt.Errorf("usecase: %v", err.Error())
		}
	}

	// Check that user was created and get id from Database
	user, err = uc.repo.GetUser(userID)
	if err != nil {
		return 0, fmt.Errorf("usecase: %v", err.Error())
	}

	return user.ID, nil
}

func (uc *AppGroupUsecase) AddExpenseToSession(info dto.AddExpenseDTO) error {
	// Get session by chat id
	session, err := uc.repo.GetActiveSessionByChatID(info.ChatID)
	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	// If session not exist -- return error
	if session.UUID == 0 {
		return SessionNotExistsErr
	}

	// Check user is exists in db
	userID, upsertErr := uc.upsertUser(info.UserID, info.Username)
	if upsertErr != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	// Check, that user is member of session
	member, err := uc.repo.GetMemberBySession(session.UUID, userID)
	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	var memberID = member.ID

	// If user not in session, add as member
	if memberID == 0 {
		err = uc.repo.AddMemberToSession(session.UUID, userID)
		if err != nil {
			return fmt.Errorf("usecase: %v", err.Error())
		}
		// Get UUID
		member, _ = uc.repo.GetMemberBySession(session.UUID, info.UserID)
		memberID = member.ID
	}

	// Add user costs
	return uc.repo.AddUserCosts(memberID, info.Cost, info.Product)
}
