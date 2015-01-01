package datastore

import (
	"database/sql"
	"errors"

	"github.com/fhermeni/wints/internship"

	"code.google.com/p/go.crypto/bcrypt"
)

const (
	//DefaultLogin is the login for the initial root account
	DefaultLogin = "root@localhost.com"
	//DefaultPassword is the password for the initial root account
	DefaultPassword = "wints"
)

//Service allows to communicate with a database
type Service struct {
	DB *sql.DB
}

//NewService initiate the storage servive
func NewService(db *sql.DB) (*Service, error) {
	s := Service{DB: db}
	//Test the connection just in case of lazyness
	/*_, err := s.Internships()
	if err != nil {
		return &Service{}, err
	}*/
	return &s, nil
}

func hash(buf []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(buf, bcrypt.MinCost)
}

//ResetRootAccount create the root account using the default credentials
func (s *Service) ResetRootAccount() error {
	_, err := s.DB.Exec("delete from users where email=$1", DefaultLogin)
	hashedPassword, err := hash([]byte(DefaultPassword))
	if err != nil {
		return err
	}
	err = SingleUpdate(s.DB, errors.New("Unable to set the default password"), "insert into users (email, password, firstname, lastname, tel, role) values ($1,$2,'The','Root','118218',$3)", DefaultLogin, hashedPassword, internship.ROOT)
	return err
}

//Install the tables on the database
func (s *Service) Install() error {
	_, err := s.DB.Exec(create)
	return err
}
