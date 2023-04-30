package repo

import (
	"collector-telegram-bot/internal/models"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	UserTable    = "users"
	SessionTable = "sessions"
)

type Repository interface {
	GetUserSessions()
	CreateUser(user *models.User) error
	GetUser(tgID int64) (*models.User, error)
	CreateNewSession(session *models.Session) error
}

type PgRepository struct {
	log  *logrus.Entry
	Conn *sql.DB
}

func NewPgRepository(log *logrus.Entry, conn *sql.DB) Repository {
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
	FROM`+" %s "+`WHERE tg_id = $1`, UserTable)

	rows, err := r.Conn.Query(queryString, tgID)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.TgID, &user.Username, &user.CreatedAt, &user.Requisites)
		}
	}
	return user, err
}

func (r *PgRepository) CreateUser(user *models.User) error {
	queryString := fmt.Sprintf(`INSERT INTO`+" %s "+`
		(tg_id, username, created_at, requisites) VALUES 
		($1, $2, current_date, $3);`, UserTable)

	_, err := r.Conn.Exec(queryString, user.TgID, user.Username, user.Requisites)
	return err
}

func (r *PgRepository) CreateNewSession(session *models.Session) error {
	queryString := fmt.Sprintf(`INSERT INTO`+" %s "+`
		(creator_id, chat_id, session_name, started_at, state) VALUES 
		($1, $2, $3, current_date, $4);`, SessionTable)

	_, err := r.Conn.Exec(queryString, session.CreatorID, session.ChatID, session.SessionName, session.State)
	return err
}
