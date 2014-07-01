package backend

import (
	"database/sql"
	"log"
	"errors"
)

type Student struct {
	P         Person
	Promotion string
	Major string
}

func NewStudent(db *sql.DB, s Student) error {
	log.Printf("Creating student %s\n", s)
	err := NewUser(db, s.P)
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into roles(email,role) values($1,'student')", s.P.Email)
	if err != nil {
		return err
	}
	res, err := db.Exec("insert into students(email, promotion) values($1,$2)", s.P.Email, s.Promotion)
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

func Students(db *sql.DB) ([]Student, error) {
	sql := "select firstname, lastname, email, tel, promotion, major from users,students where role='tutor' and users.emails = student.email"
	rows, err := db.Query(sql)
	students := make([]Student, 0, 0)
	if err != nil {
		return students, nil
	}
	defer rows.Close()
	var fn, ln, email, tel, p, m string
	for rows.Next() {
		rows.Scan(&fn, &ln, &email, &tel, &p, &m)
		students = append(students, Student{Person{fn, ln, email, tel}, p, m})
	}
	return students, nil
}

func GetStudent(db *sql.DB, email string) (Student, error) {
	var fn, ln, tel, promo, major string
	err := db.QueryRow("select firstname, lastname, tel, promotion, major from students, users where students.email = users.email and users.email=$1", email).Scan(&fn, &ln, &tel, &promo, &major)
	if err != nil {
		return Student{}, err
	}
	p := Person{fn, ln, email, tel}
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
