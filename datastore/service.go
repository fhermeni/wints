package datastore

import (
	"database/sql"
	"errors"

	"code.google.com/p/go.crypto/bcrypt"
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

func Reset(db *sql.DB, login, password string) error {
	_, err := db.Exec("delete from users where email=$1", login)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	err = SingleUpdate(db, errors.New("Unable to set the default password"), "insert into users (email, password) values ($1,$2)", login, hashedPassword)
	return err
}
