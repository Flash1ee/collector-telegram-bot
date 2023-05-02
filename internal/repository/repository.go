package repo

import (
	"collector-telegram-bot/internal"
	"collector-telegram-bot/internal/models"
	"database/sql"
	"fmt"
)

const (
	UserTable     = "users"
	SessionTable  = "sessions"
	MembersTable  = "members"
	CostsTable    = "costs"
	ClosedSession = "closed"
)

type Repository interface {
	GetUserSessions()
	CreateUser(user *models.User) (uint64, error)
	GetUser(tgID int64) (*models.User, error)
	CreateNewSession(session *models.Session) error
	GetActiveSessionByChatID(chatID int64) (*models.Session, error)
	AddMemberToSession(sessionUUID internal.UUID, userID uint64) (uint64, error)
	GetMemberBySession(sessionUUID internal.UUID, userID uint64) (*models.Member, error)
	AddUserCosts(memberID uint64, money int, description string) error
	GetUsersCosts(sessionUUID internal.UUID) ([]*models.Expanse, error)
	GetAllCosts(sessionUUID internal.UUID) ([]*models.Cost, error)
	GetAllUsers(sessionUUID internal.UUID) ([]*models.User, error)
	GetUserById(ID uint64) (*models.User, error)
	FinishSession(sessionUUID internal.UUID) error
}

type PgRepository struct {
	log  internal.Logger
	Conn *sql.DB
}

func NewPgRepository(log internal.Logger, conn *sql.DB) Repository {
	return &PgRepository{log: log, Conn: conn}
}

func (r *PgRepository) GetUserSessions() {

}

func (r *PgRepository) GetUser(tgID int64) (*models.User, error) {
	var (
		user = models.NewUser()
		err  error
	)
	queryString := fmt.Sprintf(`SELECT 
	id, 
	tg_id, 
	username, 
	created_at, 
	requisites
	FROM`+" %s "+`WHERE tg_id = $1;`, UserTable)

	rows, err := r.Conn.Query(queryString, tgID)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.TgID, &user.Username, &user.CreatedAt, &user.Requisites)
		}
	}
	return user, err
}

func (r *PgRepository) GetUserById(ID uint64) (*models.User, error) {
	var (
		user = models.NewUser()
		err  error
	)
	queryString := fmt.Sprintf(`SELECT 
	id, 
	tg_id, 
	username, 
	created_at, 
	requisites
	FROM`+" %s "+`WHERE id = $1;`, UserTable)

	rows, err := r.Conn.Query(queryString, ID)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.TgID, &user.Username, &user.CreatedAt, &user.Requisites)
		}
	}
	return user, err
}

func (r *PgRepository) CreateUser(user *models.User) (uint64, error) {
	var id uint64
	queryString := fmt.Sprintf(`INSERT INTO`+" %s "+`
		(tg_id, username, created_at, requisites) VALUES 
		($1, $2, current_timestamp, $3) returning id;`, UserTable)

	row := r.Conn.QueryRow(queryString, user.TgID, user.Username, user.Requisites)
	err := row.Scan(&id)
	return id, err
}

func (r *PgRepository) CreateNewSession(session *models.Session) error {
	queryString := fmt.Sprintf(`INSERT INTO`+" %s "+`
		(uuid, creator_id, chat_id, session_name, started_at, state) VALUES 
		($1, $2, $3, $4, current_timestamp, $5);`, SessionTable)

	_, err := r.Conn.Exec(queryString, session.UUID, session.CreatorID, session.ChatID,
		session.SessionName, session.State)
	return err
}

func (r *PgRepository) GetActiveSessionByChatID(chatID int64) (*models.Session, error) {
	var (
		session = models.NewEmptySession()
		err     error
	)
	queryString := fmt.Sprintf(`SELECT 
	uuid, 
	creator_id, 
	chat_id, 
	session_name,
	state
	FROM`+" %s "+`WHERE chat_id = $1 AND state='active';`, SessionTable)

	rows, err := r.Conn.Query(queryString, chatID)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&session.UUID, &session.CreatorID, &session.ChatID, &session.SessionName,
				&session.State)
		}
	}
	return session, err
}

func (r *PgRepository) AddMemberToSession(sessionUUID internal.UUID, userID uint64) (uint64, error) {
	var id uint64
	queryString := fmt.Sprintf(`INSERT INTO`+" %s "+`
		(session_id, user_id) VALUES 
		($1, $2) returning id;`, MembersTable)

	row := r.Conn.QueryRow(queryString, sessionUUID, userID)
	err := row.Scan(&id)
	return id, err
}

func (r *PgRepository) GetMemberBySession(sessionUUID internal.UUID, userID uint64) (*models.Member, error) {
	var (
		member = models.NewEmptyMember()
		err    error
	)
	queryString := fmt.Sprintf(`SELECT 
	id, 
	session_id, 
	user_id
	FROM`+" %s "+`WHERE session_id = $1 AND user_id = $2;`, MembersTable)

	rows, err := r.Conn.Query(queryString, sessionUUID, userID)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&member.ID, &member.SessionUUID, &member.UserID)
		}
	}
	return member, err
}

func (r *PgRepository) AddUserCosts(memberID uint64, money int, description string) error {
	queryString := fmt.Sprintf(`INSERT INTO`+" %s "+`
		(member_id, money, description, created_at) VALUES 
		($1, $2, $3, current_timestamp);`, CostsTable)

	_, err := r.Conn.Exec(queryString, memberID, money, description)
	return err
}

func (r *PgRepository) GetUsersCosts(sessionUUID internal.UUID) ([]*models.Expanse, error) {
	result := make([]*models.Expanse, 0)

	queryString := fmt.Sprintf(`SELECT U.username, C.money, C.description 
	FROM`+" %s "+`as M JOIN `+" %s "+`as C on M.id = C.member_id
		JOIN`+" %s "+`as U on M.user_id = U.id
	WHERE M.session_id = $1`, MembersTable, CostsTable, UserTable)

	rows, err := r.Conn.Query(queryString, sessionUUID)
	if err == nil {
		for rows.Next() {
			var tmpExpenses = &models.Expanse{}
			err = rows.Scan(&tmpExpenses.Username, &tmpExpenses.Cost, &tmpExpenses.Description)
			if err == nil {
				result = append(result, tmpExpenses)
			}
		}
	}

	return result, err
}

func (r *PgRepository) GetAllCosts(sessionUUID internal.UUID) ([]*models.Cost, error) {
	result := make([]*models.Cost, 0)

	queryString := fmt.Sprintf(`SELECT M.user_id, C.money
	FROM`+" %s "+`as M JOIN `+" %s "+`as C on M.id = C.member_id
	WHERE M.session_id = $1`, MembersTable, CostsTable)

	rows, err := r.Conn.Query(queryString, sessionUUID)
	if err == nil {
		for rows.Next() {
			var tmpCosts = &models.Cost{}
			err = rows.Scan(&tmpCosts.UserID, &tmpCosts.Money)
			if err == nil {
				result = append(result, tmpCosts)
			}
		}
	}

	return result, err
}

func (r *PgRepository) FinishSession(sessionUUID internal.UUID) error {
	queryString := fmt.Sprintf(`UPDATE`+" %s "+`SET state = $1 WHERE uuid = $2`, SessionTable)

	_, err := r.Conn.Exec(queryString, ClosedSession, sessionUUID)
	return err
}

func (r *PgRepository) GetAllUsers(sessionUUID internal.UUID) ([]*models.User, error) {
	result := make([]*models.User, 0)

	queryString := fmt.Sprintf(`select U.id, U.tg_id, U.username, U.created_at, U.requisites 
	from`+" %s "+`as U join`+" %s "+`as M on U.id = M.user_id where M.session_id = $1`, UserTable, MembersTable)

	rows, err := r.Conn.Query(queryString, sessionUUID)
	if err == nil {
		for rows.Next() {
			var tmpUser = &models.User{}
			err = rows.Scan(&tmpUser.ID, &tmpUser.TgID, &tmpUser.Username, &tmpUser.CreatedAt, &tmpUser.Requisites)
			if err == nil {
				result = append(result, tmpUser)
			}
		}
	}
	return result, err
}
