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

type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
}

type User struct {
	P     Person
	Privs []string
}

func (u User) isAdmin() bool {
	for _, p := range u.Privs {
		if (p == "admin" || p == "root") {
			return true;
		}
	}
	return false;
}

func (u User) isTutor() bool {
	for _, p := range u.Privs {
		if (p == "tutor") {
			return true;
		}
	}
	return false;
}


func (p Person) String() string {
	return p.Firstname + " " + p.Lastname + " (" + p.Email + ")";
}



func Register(db *sql.DB, c Credential) (User, error) {
	var fn, ln, tel, p string
	err := db.QueryRow("select firstname, lastname, tel, password from users where email=$1", c.Email).Scan(&fn, &ln, &tel, &p)
	if err != nil {
		log.Printf("Unknown user %s: ‰s\n", c.Email, err)
		return User{}, errors.New("Unknown user or incorrect password")
	}
	if (bcrypt.CompareHashAndPassword([]byte(p), []byte(c.Password)) != nil) {
		log.Printf("Bas paddword for %s: ‰s\n", c.Email, err)
		return User{}, errors.New("Unknown user or incorrect password")
	}
	rows, err := db.Query("select role from roles where email=$1", c.Email)
	if err != nil {
		log.Printf("Error while getting the roles for %s:%s\n", c.Email, err)
		return User{}, err
	}
	roles := make([]string, 0, 0)
	var r string
	for rows.Next() {
		rows.Scan(&r)
		roles = append(roles, r)
	}
	return User{Person{fn, ln, c.Email, tel}, roles}, nil
}

func NewUser(db *sql.DB, p Person) error {
	newPassword := rand_str(8)
	log.Printf("New password for user '%s %s': %s\n", p.Firstname, p.Lastname, newPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	res, err := db.Exec("insert into users(email, firstname, lastname, tel, password) values ($1,$2,$3,$4,$5)", p.Email, p.Firstname, p.Lastname, p.Tel, hashedPassword)
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



func NewTutor(db *sql.DB, p Person) error {
	err := NewUser(db, p)
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into roles(email,role) values($1,'tutor')", p.Email)
	if err != nil {
		return err
	}

	return err
}



func Tutors(db *sql.DB) ([]Person, error) {
	sql := "select firstname, lastname, roles.email, tel from users,roles where roles.email = users.email and role='tutor'"
	rows, err := db.Query(sql)
	tutors := make([]Person, 0, 0)
	if err != nil {
		return tutors, nil
	}
	defer rows.Close()
	var fn, ln, email, tel string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel)
		tutors = append(tutors, Person{fn, ln, email, tel})
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

func GetPerson(db *sql.DB, email string) (Person, error) {
	var fn, ln, tel string
	err := db.QueryRow("select firstname, lastname, tel from users where email=$1", email).Scan(&fn, &ln, &tel)
	return Person{fn, ln, email, tel}, err
}


func Admins(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select firstname, lastname, users.email, tel, role from users,roles where roles.email=users.email and role != 'student'")
	admins := make([]User, 0, 0)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return admins, err
	}
	var fn, ln, email, tel, role string
	managers := make(map[string]User)
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		m, ok := managers[email]
		if !ok {
			p := Person{fn, ln, email, tel}
			m = User{p, make([]string, 0, 0)}
			managers[email] = m
		}
		m.Privs = append(m.Privs, role)
		managers[email] = m
	}
	for _, v := range managers {
		admins = append(admins, v)
	}
	return admins, nil
}

func NewPassword(db *sql.DB, uid int, oldPassword, newPassword []byte) error {
	//Get the password
	var p []byte
	err := db.QueryRow("select password from users where uid=$1", uid).Scan(&p)
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
	_, err := db.Exec("insert into roles(email, role) values ($1, $2)", email, priv)
	return err
}
