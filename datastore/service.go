package datastore

import (
	"database/sql"
	"errors"

	"code.google.com/p/go.crypto/bcrypt"
)

const (
	DEFAULT_LOGIN    = "root@localhost.com"
	DEFAULT_PASSWORD = "wints"
)

type Service struct {
	Db *sql.DB
}

func NewService(db *sql.DB) (*Service, error) {
	s := Service{Db: db}
	//Test the connection just in case of lazyness
	_, err := s.Internships()
	if err != nil {
		return &Service{}, err
	}
	return &s, err
}

func Reset(db *sql.DB) error {
	_, err := db.Exec("delete from users where email=$1", DEFAULT_LOGIN)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(DEFAULT_PASSWORD), bcrypt.MinCost)
	if err != nil {
		return err
	}
	err = SingleUpdate(db, errors.New("Unable to set the default password"), "insert into users (email, password) values ($1,$2)", DEFAULT_LOGIN, hashedPassword)
	return err
}
