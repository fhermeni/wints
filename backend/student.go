package backend

import (
	"database/sql"
	"errors"
)

var (
	errUnknownStudent = errors.New("Unknown student")
	errStudentExists = errors.New("Student already exists")
)

type Student struct {
	P         User
	Promotion string
	Major string
}

func NewStudent(db *sql.DB, s Student) error {
	_, err := NewUser(db, s.P)
	if err != nil {
		return err
	}
	return SingleUpdate(db,
			errStudentExists,
			"insert into students(email, promotion, major) values($1,$2,$3)", s.P.Email, s.Promotion, "?");
}

func GetStudent(db *sql.DB, email string) (Student, error) {
	var fn, ln, tel, promo, major string
	err := db.QueryRow("select firstname, lastname, tel, promotion, major from students, users where students.email = users.email and users.email=$1", email).Scan(&fn, &ln, &tel, &promo, &major)
	if err != nil {
		return Student{}, errUnknownStudent
	}
	p := User{fn, ln, email, tel, ""}
	return Student{p, promo, major}, nil
}

func SetMajor(db *sql.DB, email string, m string) error {
	return SingleUpdate(db, errUnknownStudent, "update students set major=$2 where email=$1", email, m)
}
