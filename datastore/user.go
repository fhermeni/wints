package datastore

import (
	"time"

	"github.com/fhermeni/wints/internship"

	"code.google.com/p/go.crypto/bcrypt"
)

//Prepared requests

func (s *Service) Login(email, password string) (internship.User, error) {
	var fn, ln, tel, p string
	var r internship.Privilege
	if err := s.Db.QueryRow("select firstname, lastname, tel, password, role from users where email=$1", email).Scan(&fn, &ln, &tel, &p, &r); err != nil {
		return internship.User{}, internship.ErrCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password)); err != nil {
		return internship.User{}, internship.ErrCredentials
	}
	return internship.User{Firstname: fn, Lastname: ln, Email: email, Tel: tel, Role: r}, nil
}

func (s *Service) NewUser(p internship.User) error {
	newPassword := rand_str(64) //Hard to guess
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = s.Db.Exec("insert into users(email, firstname, lastname, tel, password, role) values ($1,$2,$3,$4,$5,$6)", p.Email, p.Firstname, p.Lastname, p.Tel, hashedPassword, p.Role)
	if err != nil {
		return internship.ErrUserExists
	}
	return nil
}

func (s *Service) Get(email string) (internship.User, error) {
	var fn, ln, tel string
	var role internship.Privilege
	err := s.Db.QueryRow("select firstname, lastname, tel, role from users where email=$1", email).Scan(&fn, &ln, &tel, &role)
	if err != nil {
		return internship.User{}, internship.ErrUnknownUser
	}
	return internship.User{Firstname: fn, Lastname: ln, Email: email, Tel: tel, Role: role}, nil
}

func (s *Service) GetByRole(p string) ([]internship.User, error) {
	rows, err := s.Db.Query("select firstname, lastname, users.email, tel from users where role=$1", p)
	users := make([]internship.User, 0, 0)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	var fn, ln, email, tel string
	var role internship.Privilege
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		//Get the role if exists
		users = append(users, internship.User{Firstname: fn, Lastname: ln, Email: email, Tel: tel, Role: role})
	}
	return users, nil
}

func (s *Service) SetPassword(email string, oldP, newP []byte) error {
	//Get the password
	var p []byte
	if err := s.Db.QueryRow("select password from users where email=$1", email).Scan(&p); err != nil {
		return internship.ErrCredentials
	}
	if bcrypt.CompareHashAndPassword(p, oldP) != nil {
		return internship.ErrCredentials
	}

	//Delete possible renew requests
	s.Db.QueryRow("delete from password_renewal where user=$1", email)

	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return err
	}
	return SingleUpdate(s.Db, internship.ErrUnknownUser, "update users set password=$2 where email=$1", email, hash)
}

func (s *Service) SetProfile(email, fn, ln, tel string) error {
	return SingleUpdate(s.Db, internship.ErrUnknownUser, "update users set firstname=$1, lastname=$2, tel=$3 where email=$4", fn, ln, tel, email)
}

func (s *Service) SetRole(email, priv string) error {
	return SingleUpdate(s.Db, internship.ErrUnknownUser, "update users set role=$2 where email=$1", email, priv)
}

func (s *Service) NewPassword(token string, newP []byte) (string, error) {
	//Delete possible renew requests
	var email string
	err := s.Db.QueryRow("select email from password_renewal where token=$1", token).Scan(&email)
	if err != nil {
		return "", internship.ErrNoPendingRequests
	}
	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	if err := SingleUpdate(s.Db, internship.ErrUnknownUser, "update users set password=$2 where email=$1", email, hash); err != nil {
		return "", err
	}
	_, err = s.Db.Exec("delete from password_renewal where token=$1", token)
	return email, err
}

func (s *Service) Rm(email string) error {
	return SingleUpdate(s.Db, internship.ErrUnknownUser, "DELETE FROM users where email=$1", email)
}

func (s *Service) ResetPassword(email string) (string, error) {
	//In case a request already exists
	s.Db.QueryRow("delete from password_renewal where email=$1", email)
	token := rand_str(32)
	d, _ := time.ParseDuration("48h")
	_, err := s.Db.Exec("insert into password_renewal(email,token,deadline) values($1,$2,$3)", email, token, time.Now().Add(d))
	if err != nil {
		//Not very sure, might be a SQL error or a connexion error as well.
		return "", internship.ErrUnknownUser
	}
	return token, err
}
