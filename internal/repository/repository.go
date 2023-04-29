package repo

import (
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetUserSessions()
}

type PgRepository struct {
	log *logrus.Entry
}

func MakePgRepository(log *logrus.Entry) Repository {
	return &PgRepository{log: log}
}

func (r *PgRepository) GetUserSessions() {

}
