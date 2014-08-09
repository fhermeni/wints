package backend

import (
	"time"
	"database/sql"
	"log"
	"errors"
	"strconv"
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

func IsTutoring(db *sql.DB, email string) (bool, error) {
	rows, err := db.Query("select internships.student,pending_internships.student from internships,pending_internships where tutor=$1", email)
	if err != nil {
		return false, err;
	}
	return rows.Next(), nil
}

func RegisterInternship(db *sql.DB, c Convention, move bool) error {
	supervisor := c.Sup
	_, err := db.Exec("insert into internships (student, startTime, endTime, tutor, midtermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11, $12)",
		c.Stu.P.Email, c.Begin, c.End, c.Tutor.Email, c.Begin.Add(TWO_MONTHS), c.Company, c.CompanyWWW, supervisor.Firstname, supervisor.Lastname, supervisor.Email, supervisor.Tel, c.Title)
	if err != nil {
		return err
	}
	if move {
		res, err := db.Exec("delete from pending_internships where student=$1", c.Stu.P.Email)
		if err != nil {
			return err
		}
		nb, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if nb != 1 {
			return errors.New("No pending internship deleted with email '" + c.Stu.P.Email + "'")
		}
	}
	//Create the reports
	_, err = NewReport(db, c.Stu.P.Email, "midterm", c.Begin.Add(TWO_MONTHS))
	if err != nil {
		log.Printf("midterm: %s\n", err)
	}
	t, _ := time.Parse("02/01/2006", "01/09/" + strconv.Itoa(time.Now().Year()))
	_, err = NewReport(db, c.Stu.P.Email, "final", t)
	if err != nil {
		log.Printf("final: %s\n", err)
	}

	_, err = NewReport(db, c.Stu.P.Email, "supReport", t)
	if err != nil {
		log.Printf("supervisor report: %s\n", err)
	}

	log.Printf("Convention of %s validated\n", c.Stu)
	return err
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
		log.Printf("Here")
		return Convention{}, err
	}
	stu := Student{User{stuFn, stuLn, stuMail, stuTel,""}, promo, major}
	tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel, ""}
	sup := User{supFn, supLn, supEmail, supTel, ""}
	mid, err := GetReportMetaData(db, stuMail, "midterm")
	if err != nil {
		log.Printf("Hop")
		return Convention{}, err
	}
	final, err := GetReportMetaData(db, stuMail, "final")
	if err != nil {
		log.Printf("Hip")
		return Convention{}, err
	}
	supReport, err := GetReportMetaData(db, stuMail, "supReport")
	if err != nil {
		log.Printf("Poo")
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
		return Convention{}, errors.New("No pending convention for " + email)
	}
	return scanConvention(db, rows)
}

func GetConventions(db *sql.DB) ([]Convention, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel,"+
			"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title "+
			" from internships, users as stu, users as tut, students where students.email = stu.email and internships.student = stu.email and tut.email = internships.tutor";
	rows, err := db.Query(sql)
	conventions := make([]Convention, 0, 0)
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


func GetConventions2(db *sql.DB, emitter string) ([]Convention, error) {
	conventions := make([]Convention, 0, 0)
	u, err := GetUser(db, emitter)
	if err != nil {
		return conventions, err
	}
	var q string
	var rows *sql.Rows
	if len(u.Role) == 0 {
		q = "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel," +
				"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title " +
				" from internships, users as stu, users as tut, students where students.email = stu.email and internships.student = stu.email and tut.email = internships.tutor" +
				" and tutor=$1";
		rows, err = db.Query(q, emitter)
	} else {
		q = "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel," +
				"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title " +
				" from internships, users as stu, users as tut, students where students.email = stu.email and internships.student = stu.email and tut.email = internships.tutor";
		rows, err = db.Query(q)
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
