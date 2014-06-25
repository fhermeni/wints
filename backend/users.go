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
	selectUser = "select email,password from users where email=$1"
	changePassword = "update users set password=$2 where uid=$1"
)

type Person struct {
	Firstname string
	Lastname string
	Email string
	Tel string
}

type Admin struct {
	P Person
	Privs []string
}

func (p Person) String() string {
	return p.Firstname + " " + p.Lastname + " (" + p.Email + ")";
}

type Student struct {
 	P Person
	Promotion string
}

func Register(db *sql.DB, c Credential) error {
	var uid int
	var p []byte
	var username,r string
	err := db.QueryRow(selectUser, c.Email).Scan(&uid, &username, &p,&r)
	if err != nil {
		return err
	}
	if (bcrypt.CompareHashAndPassword(p, []byte(c.Password)) != nil) {
		return errors.New("Unknown user or incorrect password")
	}
	return nil
}

func newUser(db *sql.DB, p Person) error {
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

func NewStudent(db *sql.DB, s Student) error {
	log.Printf("Creating student %s\n", s)
	err := newUser(db, s.P)
	if err != nil {
		return err
	}
	_,err = db.Exec("insert into roles(email,role) values($1,'student')", s.P.Email)
	if err != nil {
		return err
	}
	res, err := db.Exec("insert into students(email, promotion) values($1,$2)",s.P.Email, s.Promotion)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errors.New("Unable to declare the student")
	}
	return nil
}

func NewTutor(db *sql.DB, p Person) error {
	err := newUser(db, p)
	if err != nil {
		return err
	}
	_,err = db.Exec("insert into roles(email,role) values($1,'tutor')", p.Email)
	if err != nil {
		return err
	}

	return err
}

func Students(db *sql.DB) ([]Student, error) {
	sql := "select firstname, lastname, email, tel, promotion from users,students where role='tutor' and users.emails = student.email"
	rows, err := db.Query(sql)
	students := make([]Student, 0, 0)
	if err != nil {
		return students, nil
	}
	defer rows.Close()
	var fn, ln, email, tel, p string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &p)
		students = append(students, Student{Person{fn, ln, email, tel}, p})
	}
	return students, nil
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

func GetStudent(db *sql.DB, email string) (Student, error) {
	var fn, ln, tel, promo string
	err := db.QueryRow("select firstname, lastname, tel, promotion from students, users where students.email = users.email and users.email=$1",email).Scan(&fn, &ln, &tel, &promo)
	if err != nil {
		return Student{}, err
	}
	p := Person{fn, ln, email, tel}
	return Student{p, promo}, nil
}

func Admins(db *sql.DB) ([]Admin, error) {
	rows, err := db.Query("select firstname, lastname, users.email, tel, role from users,roles where roles.email=users.email and role != 'student'")
	admins := make([]Admin, 0, 0)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return admins, err
	}
	var fn, ln, email, tel, role string
	managers := make(map[string]Admin)
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &role)
		m, ok := managers[email]
		if !ok {
			p := Person{fn, ln, email, tel}
			m = Admin{p, make([]string, 0, 0)}
			managers[email] = m
			log.Printf("%s\n", m)
		} else {
			log.Printf("else: %s\n", m)
		}
		m.Privs = append(m.Privs, role)
		managers[email] = m
	}
	for _, v := range managers {
		admins = append(admins, v)
	}
	return admins, nil
}
