package models

import (
	"collector-telegram-bot/config"
	"database/sql"
	"fmt"
)

func NewPgSQLConnection(conn config.PostgresConnectionParams) *sql.DB {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conn.User, conn.Password,
		conn.Host, conn.Port, conn.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	return db
}
