package backend

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"time"
)

var (
	ErrUserExists        = errors.New("User already exists")
	ErrUserNotFound      = errors.New("User not found")
	ErrUserTutoring      = errors.New("The user is tutoring students")
	ErrCredentials       = errors.New("Incorrect credentials")
	ErrNoPendingRequests = errors.New("No password renewable request pending")
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
	Role      string
}

func (p User) String() string {
	return p.Firstname + " " + p.Lastname + " (" + p.Email + ")"
}

func (p User) CompatibleRole(r string) bool {
	if p.Role == "root" {
		return true
	}
	if p.Role == "admin" && r != "root" {
		return true
	}
	if p.Role == "major" && r == "major" {
		return true
	}
	return false
}

func Register(db *sql.DB, email, password string) (User, error) {
	var fn, ln, tel, p, r string
	if err := db.QueryRow("select firstname, lastname, tel, password, role from users where email=$1", email).Scan(&fn, &ln, &tel, &p, &r); err != nil {
		return User{}, ErrCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password)); err != nil {
		return User{}, ErrCredentials
	}
	return User{fn, ln, email, tel, r}, nil
}

func NewUser(db *sql.DB, p User) (string, error) {
	newPassword := rand_str(64) //Hard to guess
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("insert into users(email, firstname, lastname, tel, password, role) values ($1,$2,$3,$4,$5,$6)", p.Email, p.Firstname, p.Lastname, p.Tel, hashedPassword, p.Role)
	if err != nil {
		return "", ErrUserExists
	}
	token, err := PasswordRenewalRequest(db, p.Email)
	return token, err
}

func RmUser(db *sql.DB, email string) error {
	ok, err := IsTutoring(db, email)
	if err != nil {
		return err
	}
	if ok {
		return ErrUserTutoring
	}
	return SingleUpdate(db, ErrUserNotFound, "DELETE FROM users where email=$1", email)
}

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func GetUser(db *sql.DB, email string) (User, error) {
	var fn, ln, tel, role string
	err := db.QueryRow("select firstname, lastname, tel, role from users where email=$1", email).Scan(&fn, &ln, &tel, &role)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	return User{fn, ln, email, tel, role}, nil
}

func Admins(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select firstname, lastname, users.email, tel, role from users where users.email not in (select email from students)")
	admins := make([]User, 0, 0)
	if err != nil {
		return admins, err
	}
	defer rows.Close()
	var fn, ln, email, tel, role string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		//Get the role if exists
		admins = append(admins, User{fn, ln, email, tel, role})
	}
	return admins, nil
}

func ChangePassword(db *sql.DB, email string, oldPassword, newPassword []byte) error {
	//Get the password
	var p []byte
	if err := db.QueryRow("select password from users where email=$1", email).Scan(&p); err != nil {
		return ErrCredentials
	}
	if bcrypt.CompareHashAndPassword(p, oldPassword) != nil {
		return ErrCredentials
	}

	//Delete possible renew requests
	db.QueryRow("delete from password_renewal where user=$1", email)

	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	return SingleUpdate(db, ErrUserNotFound, "update users set password=$2 where email=$1", email, hash)
}

func SetProfile(db *sql.DB, email, fn, ln, tel string) error {
	return SingleUpdate(db, ErrUserNotFound, "update users set firstname=$1, lastname=$2, tel=$3 where email=$4", fn, ln, tel, email)
}

func GrantPrivilege(db *sql.DB, email, priv string) error {
	return SingleUpdate(db, ErrUserNotFound, "update users set role=$2 where email=$1", email, priv)
}

func PasswordRenewalRequest(db *sql.DB, email string) (string, error) {
	//In case a request already exists
	db.QueryRow("delete from password_renewal where email=$1", email)
	token := rand_str(32)
	d, _ := time.ParseDuration("48h")
	_, err := db.Exec("insert into password_renewal(email,token,deadline) values($1,$2,$3)", email, token, time.Now().Add(d))
	if err != nil {
		//Not very sure, might be a SQL error or a connexion error as well.
		return "", ErrUserNotFound
	}
	return token, err
}

func NewPassword(db *sql.DB, token string, newPassword []byte) (string, error) {
	//Delete possible renew requests
	var email string
	err := db.QueryRow("select email from password_renewal where token=$1", token).Scan(&email)
	if err != nil {
		return "", ErrNoPendingRequests
	}
	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	if err := SingleUpdate(db, ErrUserNotFound, "update users set password=$2 where email=$1", email, hash); err != nil {
		return "", err
	}
	_, err = db.Exec("delete from password_renewal where token=$1", token)
	return email, err
}
