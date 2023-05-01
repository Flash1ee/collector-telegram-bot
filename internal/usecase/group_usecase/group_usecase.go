package group_usecase

import (
	"collector-telegram-bot/internal"
	"collector-telegram-bot/internal/dto"
	"collector-telegram-bot/internal/models"
	repo "collector-telegram-bot/internal/repository"
	"collector-telegram-bot/internal/usecase"
	"fmt"
	"github.com/google/uuid"
)

const (
	ActiveSession = "active"
	EmptyString   = ""
)

type AppGroupUsecase struct {
	log  internal.Logger
	repo repo.Repository
}

func New(log internal.Logger, repo repo.Repository) GroupUsecase {
	return &AppGroupUsecase{log: log, repo: repo}
}

func (uc *AppGroupUsecase) CreateSession(info dto.CreateSessionDTO) error {
	sessionUUID := uuid.New()

	userID, err := uc.upsertUser(info.UserID, info.Username)

	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	// Check, that aren't any active sessions in chat
	curSession, err := uc.repo.GetActiveSessionByChatID(info.ChatID)
	switch {
	case err != nil:
		return err
	case curSession.State == ActiveSession:
		return usecase.SessionExistsErr
	default:
	}

	session := models.NewSession(sessionUUID, userID, info.ChatID, info.SessionName)
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
	_, err = uc.repo.AddMemberToSession(curSession.UUID, userID)
	return err
}

func (uc *AppGroupUsecase) upsertUser(userID int64, username string) (uint64, error) {
	user, err := uc.repo.GetUser(userID)

	if err != nil {
		return 0, fmt.Errorf("usecase: %v", err)
	}

	// If user not exists, you should create user
	if user.ID == 0 {
		user.Username = username
		user.TgID = userID
		user.ID, err = uc.repo.CreateUser(user)
		if err != nil {
			return 0, fmt.Errorf("usecase: %v", err.Error())
		}
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
	if session.State != ActiveSession {
		return usecase.SessionNotExistsErr
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
		memberID, err = uc.repo.AddMemberToSession(session.UUID, userID)
		if err != nil {
			return fmt.Errorf("usecase: %v", err.Error())
		}
	}

	// Add user costs
	return uc.repo.AddUserCosts(memberID, info.Cost, info.Product)
}
