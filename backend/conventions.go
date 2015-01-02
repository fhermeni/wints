package backend

import (
	"database/sql"
	"errors"
	"time"
)

type Convention struct {
	Stu           Student
	Sup           User
	Tutor         User
	Company       string
	CompanyWWW    string
	Begin         time.Time
	End           time.Time
	MidtermReport ReportMetaData
	FinalReport   ReportMetaData
	SupReport     ReportMetaData
	Title         string
}

const (
	TWO_MONTHS = 2 * time.Hour * 24 * 30
)

var (
	ErrUnknownConvention           = errors.New("Unknown convention")
	ErrConventionExists            = errors.New("Convention already exists")
	ErrInvalidTutor                = errors.New("The new tutor is a student")
	allMyConventions     *sql.Stmt = nil
	allConventions       *sql.Stmt = nil
)

func IsTutoring(db *sql.DB, email string) (bool, error) {
	rows, err := db.Query("select internships.student,pending_internships.student from internships,pending_internships where tutor=$1", email)
	if err != nil {
		return false, err
	}
	return rows.Next(), nil
}

func scanConvention(db *sql.DB, rows *sql.Rows) (Convention, error) {
	var stuFn, stuLn, stuMail, stuTel, promo, major, title string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW string
	var start, end, midDeadline time.Time

	err := rows.Scan(&stuFn, &stuLn, &stuMail, &stuTel, &promo, &major, &start, &end,
		&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
		&company, &companyWWW, &supFn, &supLn, &supEmail, &supTel, &title)
	if err != nil {
		return Convention{}, err
	}
	stu := Student{User{stuFn, stuLn, stuMail, stuTel, ""}, promo, major}
	tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel, ""}
	sup := User{supFn, supLn, supEmail, supTel, ""}
	mid, err := GetReportMetaData(db, stuMail, "midterm")
	if err != nil {
		return Convention{}, err
	}
	final, err := GetReportMetaData(db, stuMail, "final")
	if err != nil {
		return Convention{}, err
	}
	supReport, err := GetReportMetaData(db, stuMail, "supReport")
	if err != nil {
		return Convention{}, err
	}

	return Convention{stu, sup, tutor, company, companyWWW, start, end, mid, final, supReport, title}, nil
}

func GetConvention(db *sql.DB, email string) (Convention, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel," +
		"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title " +
		" from internships, users as stu, users as tut, students where students.email = $1 and stu.email = $1 and internships.student = $1 and tut.email = internships.tutor"

	rows, err := db.Query(sql, email)
	if err != nil {
		return Convention{}, err
	}
	if !rows.Next() {
		return Convention{}, ErrUnknownConvention
	}
	return scanConvention(db, rows)
}

func GetConventions2(db *sql.DB, emitter string) ([]Convention, error) {
	conventions := make([]Convention, 0, 0)
	u, err := GetUser(db, emitter)
	if err != nil {
		return conventions, err
	}
	var q string
	var rows *sql.Rows
	if len(u.Role) == 0 {
		q = "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel, midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title " +
			" from internships" +
			" join users as stu on stu.email = internships.student" +
			" join users as tut on tut.email = internships.tutor" +
			" join students on students.email = stu.email" +
			" and tutor=$1"
		if allMyConventions == nil {
			allMyConventions, err = db.Prepare(q)
		}
		if err != nil {
			return conventions, err
		}
		rows, err = allMyConventions.Query(emitter)
	} else {
		q = "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel, midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title " +
			" from internships" +
			" join users as stu on stu.email = internships.student" +
			" join users as tut on tut.email = internships.tutor" +
			" join students on students.email = stu.email"
		if allConventions == nil {
			allConventions, err = db.Prepare(q)
		}
		if err != nil {
			return conventions, err
		}
		rows, err = allConventions.Query()
	}

	if err != nil {
		return conventions, err
	}
	defer rows.Close()
	for rows.Next() {
		c, err := scanConvention(db, rows)
		if err != nil {
			return conventions, err
		}
		conventions = append(conventions, c)
	}
	return conventions, nil
}

func UpdateTutor(db *sql.DB, student string, newTutor string) error {
	//new tutor is not a student
	_, err := GetConvention(db, newTutor)
	if err == nil {
		return ErrInvalidTutor
	}
	err = SingleUpdate(db, ErrUserNotFound, "update internships set tutor=$2 where student=$1", student, newTutor)
	//userNotFound might be the student or the tutor
	return err
}

func UpdateCompany(db *sql.DB, stu, name, www, title string) error {
	return SingleUpdate(db, ErrUnknownConvention,
		"update internships set company=$2, companywww=$3, title=$4 where student=$1", stu, name, www, title)
}

func UpdateSupervisor(db *sql.DB, stu string, u User) error {
	return SingleUpdate(db, ErrUnknownConvention,
		"update internships set supervisorfn=$2, supervisorln=$3, supervisoremail=$4, supervisortel=$5 where student=$1", stu, u.Firstname, u.Lastname, u.Email, u.Tel)
}
