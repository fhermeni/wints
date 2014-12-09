package user;

import (	
	"database/sql"	
	"github.com/fhermeni/wints/db"
	"code.google.com/p/go.crypto/bcrypt"	
	"time"
)

type DbService struct {
	db *sql.DB
}

func (s *DbService) Register(email, password string) (User, error) {
	var fn, ln, tel, p, r string
	if err := s.db.QueryRow("select firstname, lastname, tel, password, role from users where email=$1", email).Scan(&fn, &ln, &tel, &p, &r); err != nil {
		return User{}, ErrCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password)); err != nil {
		return User{}, ErrCredentials
	}
	return User{fn, ln, email, tel, r}, nil
}

func (s *DbService) New(p User) error {
	newPassword := rand_str(64) //Hard to guess
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("insert into users(email, firstname, lastname, tel, password, role) values ($1,$2,$3,$4,$5,$6)", p.Email, p.Firstname, p.Lastname, p.Tel, hashedPassword, p.Role)
	if err != nil {
		return ErrExists
	}
	return nil
}

func (s *DbService) Get(email string) (User, error) {
	var fn, ln, tel, role string
	err := s.db.QueryRow("select firstname, lastname, tel, role from users where email=$1", email).Scan(&fn, &ln, &tel, &role)
	if err != nil {
		return User{}, ErrNotFound
	}
	return User{fn, ln, email, tel, role}, nil	
}

func (s *DbService) GetByRole(p string) ([]User, error) {
	rows, err := s.db.Query("select firstname, lastname, users.email, tel from users where role=$1", p)
	users := make([]User, 0, 0)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	var fn, ln, email, tel, role string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		//Get the role if exists
		users = append(users, User{fn, ln, email, tel, role})
	}
	return users, nil
}

func (s *DbService) SetPassword(email string, oldP, newP []byte) error {
	//Get the password
	var p []byte
	if err := s.db.QueryRow("select password from users where email=$1", email).Scan(&p); err != nil {
		return ErrCredentials
	}
	if (bcrypt.CompareHashAndPassword(p, oldP) != nil) {
		return ErrCredentials
	}

	//Delete possible renew requests
	s.db.QueryRow("delete from password_renewal where user=$1", email)

	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return err
	}
	return db.SingleUpdate(s.db, ErrNotFound, "update users set password=$2 where email=$1", email, hash)
}

func (s *DbService) SetProfile(email, fn, ln, tel string) error {
	return db.SingleUpdate(s.db, ErrNotFound, "update users set firstname=$1, lastname=$2, tel=$3 where email=$4", fn, ln, tel, email)
}

func (s *DbService) SetRole(email, priv string) error {
	return db.SingleUpdate(s.db, ErrNotFound, "update users set role=$2 where email=$1", email, priv)
}

func (s *DbService) NewPassword(token string, newP []byte) (string, error) {
	//Delete possible renew requests
	var email string
	err := s.db.QueryRow("select email from password_renewal where token=$1", token).Scan(&email)
	if err != nil {
		return "", ErrNoPendingRequests
	}
	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	if err := db.SingleUpdate(s.db, ErrNotFound, "update users set password=$2 where email=$1", email, hash); err != nil {
		return "", err
	}
	_, err = s.db.Exec("delete from password_renewal where token=$1", token)
	return email, err
}

func (s *DbService) Rm(email string) error {	
	return SingleUpdate(db, ErrUserNotFound, "DELETE FROM users where email=$1", email)
}

func (s *DbService) PasswordRenewalRequest(email string) (string, error) {
	//In case a request already exists
	s.db.QueryRow("delete from password_renewal where email=$1", email)
	token := rand_str(32)
	d, _ := time.ParseDuration("48h")
	_, err := s.db.Exec("insert into password_renewal(email,token,deadline) values($1,$2,$3)", email, token, time.Now().Add(d))
	if err != nil {
		//Not very sure, might be a SQL error or a connexion error as well.
		return "", ErrNotFound
	}
 	return token, err
}