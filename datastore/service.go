package datastore

import (
	"database/sql"
	"errors"

	"github.com/fhermeni/wints/internship"

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

func hash(buf []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(buf, bcrypt.MinCost)
}

func (s *Service) ResetRootAccount() error {
	_, err := s.Db.Exec("delete from users where email=$1", DEFAULT_LOGIN)
	hashedPassword, err := hash([]byte(DEFAULT_PASSWORD))
	if err != nil {
		return err
	}
	err = SingleUpdate(s.Db, errors.New("Unable to set the default password"), "insert into users (email, password, firstname, lastname, tel, role) values ($1,$2,'The','Root','118218',$3)", DEFAULT_LOGIN, hashedPassword, internship.ROOT)
	return err
}
