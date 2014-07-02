package backend

import (
	"database/sql"
	"log"
	"errors"
)

type Student struct {
	P         User
	Promotion string
	Major string
}

func NewStudent(db *sql.DB, s Student) error {
	log.Printf("Creating student %s\n", s)
	err := NewUser(db, s.P)
	if err != nil {
		return err
	}
	res, err := db.Exec("insert into students(email, promotion, major) values($1,$2,$3)", s.P.Email, s.Promotion, "?")
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

func GetStudent(db *sql.DB, email string) (Student, error) {
	var fn, ln, tel, promo, major string
	err := db.QueryRow("select firstname, lastname, tel, promotion, major from students, users where students.email = users.email and users.email=$1", email).Scan(&fn, &ln, &tel, &promo, &major)
	if err != nil {
		return Student{}, err
	}
	p := User{fn, ln, email, tel, ""}
	return Student{p, promo, major}, nil
}

func SetPromotion(db *sql.DB, email string, p string) error {
	res, err := db.Exec("update students set promotion=$2 where email=$1", email, p)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errors.New("Unknown student '" + email + "'\n")
	}
	return nil
}

func SetMajor(db *sql.DB, email string, m string) error {
	res, err := db.Exec("update students set major=$2 where email=$1", email, m)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errors.New("Unknown student '" + email + "'\n")
	}
	return nil
}
