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

func (uc *AppGroupUsecase) GetAllExpenses(info dto.GetCostsDTO) (map[string]models.AllUserCosts, error) {
	// Get session by chat id
	session, err := uc.repo.GetActiveSessionByChatID(info.ChatID)
	if err != nil {
		return nil, fmt.Errorf("usecase: %v", err.Error())
	}

	// If session not exist -- return error
	if session.State != ActiveSession {
		return nil, usecase.SessionNotExistsErr
	}

	costs, err := uc.repo.GetUsersCosts(session.UUID)
	if err != nil {
		return nil, fmt.Errorf("usecase: %v", err.Error())
	}

	var UsersCosts = map[string]models.AllUserCosts{}
	for _, curCost := range costs {
		username := curCost.Username
		curRec := UsersCosts[username]
		curRec.Sum += curCost.Cost

		newUserCost := models.UserCost{
			Money:       curCost.Cost,
			Description: curCost.Description,
		}

		curRec.Costs = append(curRec.Costs, newUserCost)
		UsersCosts[username] = curRec
	}
	return UsersCosts, nil
}

func (uc *AppGroupUsecase) FinishSession(info dto.FinishSessionDTO) error {
	// Get session by chat id
	session, err := uc.repo.GetActiveSessionByChatID(info.ChatID)
	if err != nil {
		return fmt.Errorf("usecase: %v", err.Error())
	}

	// If session not exist -- return error
	if session.State != ActiveSession {
		return usecase.SessionNotExistsErr
	}

	return uc.repo.FinishSession(session.UUID)
}

type DebtsMtr map[uint64]map[uint64]int

func (uc *AppGroupUsecase) formDebtMtr(sessionUUID uuid.UUID) (DebtsMtr, error) {
	allUsers, err := uc.repo.GetAllUsers(sessionUUID)
	if err != nil {
		return nil, fmt.Errorf("usecase: %v", err.Error())
	}

	var debtsMtr = make(DebtsMtr)
	for _, currUser := range allUsers {
		currDebtor := make(map[uint64]int)
		for _, tmpUser := range allUsers {
			currDebtor[tmpUser.ID] = 0
		}
		debtsMtr[currUser.ID] = currDebtor
	}

	allCosts, err := uc.repo.GetAllCosts(sessionUUID)
	if err != nil {
		return nil, fmt.Errorf("usecase: %v", err.Error())
	}

	for _, curCost := range allCosts {
		userID := curCost.UserID
		debt := curCost.Money / len(allUsers)
		curDebtors := debtsMtr[userID]
		for curDebtor := range curDebtors {
			if curDebtor != userID {
				debtsMtr[userID][curDebtor] += debt
			}
		}
	}

	for curUser, curDebtors := range debtsMtr {
		for curDebtor := range curDebtors {
			if debtsMtr[curUser][curDebtor] != 0 && debtsMtr[curDebtor][curUser] != 0 {
				if debtsMtr[curUser][curDebtor] > debtsMtr[curDebtor][curUser] {
					debtsMtr[curUser][curDebtor] -= debtsMtr[curDebtor][curUser]
					debtsMtr[curDebtor][curUser] = 0
				} else {
					debtsMtr[curDebtor][curUser] -= debtsMtr[curUser][curDebtor]
					debtsMtr[curUser][curDebtor] = 0
				}
			}
		}
	}
	return debtsMtr, nil
}

func (uc *AppGroupUsecase) GetAllDebts(info dto.GetDebtsDTO) (map[string]models.AllUserDebts, error) {
	// Get session by chat id
	session, err := uc.repo.GetActiveSessionByChatID(info.ChatID)
	if err != nil {
		return nil, fmt.Errorf("usecase: %v", err.Error())
	}

	// If session not exist -- return error
	if session.State != ActiveSession {
		return nil, usecase.SessionNotExistsErr
	}

	debtsMtr, err := uc.formDebtMtr(session.UUID)
	if err != nil {
		return nil, err
	}

	UserDebts := map[string]models.AllUserDebts{}
	for curUser, curDebtors := range debtsMtr {
		for curDebtor := range curDebtors {
			if debtsMtr[curUser][curDebtor] != 0 {
				creditor, _ := uc.repo.GetUserById(curUser)
				debtor, _ := uc.repo.GetUserById(curDebtor)
				debt := debtsMtr[curUser][curDebtor]

				curUserDebts := UserDebts[creditor.Username]
				newUserDebt := models.UserDebt{
					DebtorName: debtor.Username,
					Money:      debt,
				}
				curUserDebts.Debts = append(curUserDebts.Debts, newUserDebt)
				UserDebts[creditor.Username] = curUserDebts
			}
		}
	}

	return UserDebts, nil
}
