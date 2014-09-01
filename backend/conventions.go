package backend

import (
	"time"
	"database/sql"
	"errors"
	"strconv"
	"log"
)

type Convention struct {
	Stu        Student
	Sup        User
	Tutor      User
	Company    string
	CompanyWWW string
	Begin      time.Time
	End        time.Time
	MidtermReport ReportMetaData
	FinalReport ReportMetaData
	SupReport ReportMetaData
	Title	string
}


const (
	TWO_MONTHS = 2*time.Hour*24*30
)

var (
	ErrUnknownConvention = errors.New("Unknown convention")
	ErrConventionExists = errors.New("Convention already exists")
	ErrInvalidTutor = errors.New("The new tutor is a student")
	allMyConventions *sql.Stmt = nil
	allConventions *sql.Stmt = nil
)

func IsTutoring(db *sql.DB, email string) (bool, error) {
	rows, err := db.Query("select internships.student,pending_internships.student from internships,pending_internships where tutor=$1", email)
	if err != nil {
		return false, err;
	}
	return rows.Next(), nil
}

func RegisterInternship(db *sql.DB, c Convention, move bool) error {
	supervisor := c.Sup
	err := SingleUpdate(db, ErrConventionExists, "insert into internships (student, startTime, endTime, tutor, midtermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11, $12)",
		c.Stu.P.Email, c.Begin, c.End, c.Tutor.Email, c.Begin.Add(TWO_MONTHS), c.Company, c.CompanyWWW, supervisor.Firstname, supervisor.Lastname, supervisor.Email, supervisor.Tel, c.Title)
	if move {
		_, err := db.Exec("delete from pending_internships where student=$1", c.Stu.P.Email)
		if err != nil {
			return err
		}
	}
	//Create the reports
	_, err = NewReport(db, c.Stu.P.Email, "midterm", c.Begin.Add(TWO_MONTHS))
	if err != nil {
		return err
	}
	t, _ := time.Parse("02/01/2006", "01/09/" + strconv.Itoa(time.Now().Year()))
	_, err = NewReport(db, c.Stu.P.Email, "final", t)
	if err != nil {
		return err
	}
	_, err = NewReport(db, c.Stu.P.Email, "supReport", t)
	if err != nil {
		return err
	}
	return nil
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
	stu := Student{User{stuFn, stuLn, stuMail, stuTel,""}, promo, major}
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
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel,"+
	"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title "+
	" from internships, users as stu, users as tut, students where students.email = $1 and stu.email = $1 and internships.student = $1 and tut.email = internships.tutor";

	rows, err := db.Query(sql, email)
	if err != nil {
		return Convention{}, err
	}
	if (!rows.Next()) {
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

func UpdateTutor(db *sql.DB, student string, newTutor string)  error {
	//new tutor is not a student
	_, err := GetConvention(db, newTutor)
	if err == nil {
		return ErrInvalidTutor
	}
	log.Printf("student=%s new=%s\n", student, newTutor)
	err = SingleUpdate(db, ErrUserNotFound, "update internships set tutor=$2 where student=$1", student, newTutor)
	//userNotFound might be the student or the tutor
	return err
}
