package db

import (
	"database/sql"
	"fmt"

	"github.com/cihanozhan/config"
)

const dbPrefix = "USERS_"

type (
	DB struct {
		db *sql.DB
	}

	UsersDB interface {
		Ping() error
		// GetUsers() ([]model.User, error)
	}
)

func NewDB(db *sql.DB) *DB {
	postgreSQLDB := DB{
		db: db,
	}
	return &postgreSQLDB
}

func CreateUsersDbConnection() (*sql.DB, error) {
	Global := config.SetConfig(dbPrefix)

	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		Global.DBHost,
		Global.DBPort,
		Global.DBName,
		Global.DBUser,
		Global.DBPassword)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Ping() error {
	return db.db.Ping()
}
