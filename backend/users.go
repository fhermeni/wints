package backend

import (
	_ "github.com/lib/pq"
	"database/sql"
	"errors"
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"log"
	"net/http"
)

const (
	selectUser     = "select email,password from users where email=$1"
	changePassword = "update users set password=$2 where uid=$1"
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
	Role	  string
}

func (p User) String() string {
	return p.Firstname + " " + p.Lastname + " (" + p.Email + ")";
}


func Register(db *sql.DB, c Credential) (User, error) {
	var fn, ln, tel, p, r string
	err := db.QueryRow("select firstname, lastname, tel, password, role from users where email=$1", c.Email).Scan(&fn, &ln, &tel, &p, &r)
	if err != nil {
		log.Printf("Unknown user %s: ‰s\n", c.Email, err)
		return User{}, errors.New("Unknown user or incorrect password")
	}
	if (bcrypt.CompareHashAndPassword([]byte(p), []byte(c.Password)) != nil) {
		log.Printf("Bad password for %s\n", c.Email)
		return User{}, errors.New("Unknown user or incorrect password")
	}

	return User{fn, ln, c.Email, tel, r}, nil
}

func NewUser(db *sql.DB, p User) error {
	newPassword := rand_str(8)
	log.Printf("New password for user '%s %s': %s\n", p.Firstname, p.Lastname, newPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	res, err := db.Exec("insert into users(email, firstname, lastname, tel, password, role) values ($1,$2,$3,$4,$5,$6)", p.Email, p.Firstname, p.Lastname, p.Tel, hashedPassword, p.Role)
	if err != nil {
		return err;
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb == 1 {
		return nil
	}
	return &BackendError{http.StatusConflict, "User already exists"}
}

func RmUser(db *sql.DB, email string) error {
	_,err := db.Exec("DELETE FROM users where email=$1", email)
 	return err
}

func AvailableTutors(db *sql.DB) ([]User, error) {
	sql := "select firstname, lastname, email, tel, role from users where email not in (select email from students)"
	rows, err := db.Query(sql)
	tutors := make([]User, 0, 0)
	if err != nil {
		return tutors, nil
	}
	defer rows.Close()
	var fn, ln, email, tel, role string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		tutors = append(tutors, User{fn, ln, email, tel, role})
	}
	return tutors, nil
}

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz$_-!@&,;:/"
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
	return User{fn, ln, email, tel, role}, err
}


func Admins(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select firstname, lastname, users.email, tel, role from users where users.email not in (select email from students)")
	admins := make([]User, 0, 0)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return admins, err
	}
	var fn, ln, email, tel, role string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		//Get the role if exists
		admins = append(admins, User{fn, ln, email, tel, role})
	}
	return admins, nil
}

func NewPassword(db *sql.DB, email string, oldPassword, newPassword []byte) error {
	//Get the password
	var p []byte
	err := db.QueryRow("select password from users where email=$1", email).Scan(&p)
	if err != nil {
		return err
	}
	if (bcrypt.CompareHashAndPassword(p, oldPassword) != nil) {
		return &BackendError{http.StatusConflict, "incorrect old password"}
	}

	hash, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = db.Exec(changePassword, hash)
	if err != nil {
		return err
	}
	return nil
}

func SetProfile(db *sql.DB, email, fn, ln, tel string) error {
	_, err := db.Exec("update users set firstname=$1, lastname=$2, tel=$3 where email=$4", fn, ln, tel, email)
	return err
}

func GrantPrivilege(db *sql.DB, email, priv string) error {
	_, err := db.Exec("update users set role=$2 where email=$1", email, priv)
	return err
}
