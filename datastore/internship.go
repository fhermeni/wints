package datastore

import (
	"database/sql"
	"log"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

func rollback(e error, tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return e
}
func (srv *Service) NewInternship(c internship.Convention) error {
	if c.Begin.After(c.End) {
		return internship.ErrInvalidPeriod
	}
	tx, err := srv.DB.Begin()
	if err != nil {
		return err
	}
	//Create the student account
	sql := "insert into users(email, firstname, lastname, tel, password, role) values ($1,$2,$3,$4,$5,$6)"
	_, err = tx.Exec(sql, c.Student.Email, c.Student.Firstname, c.Student.Lastname, c.Student.Tel, randomBytes(64), internship.STUDENT)
	if err != nil {
		return rollback(internship.ErrUserExists, tx)
	}
	srv.ResetPassword(c.Student.Email)
	//Create the internship
	sql = "insert into internships(student, promotion, tutor, startTime, endTime, title, supervisorFn, supervisorLn, supervisorTel, supervisorEmail, company, companyWWW) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_, err = tx.Exec(sql, c.Student.Email, c.Promotion, c.Tutor.Email, c.Begin, c.End, c.Title, c.Supervisor.Firstname, c.Supervisor.Lastname, c.Supervisor.Tel, c.Supervisor.Email, c.Cpy.Name, c.Cpy.WWW)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		if e, ok := err.(*pq.Error); ok {
			if e.Constraint == "internships_tutor_fkey" {
				return internship.ErrUnknownUser
			}
			if e.Constraint == "internships_pkey" {
				return internship.ErrInternshipExists
			}
		}
		return err
	}
	//Plan the reports
	for _, def := range srv.ReportDefs() {
		hdr, err := def.Instantiate(c.Begin)
		if err != nil {
			return rollback(err, tx)
		}
		sql = "insert into reports(student, kind, deadline) values($1,$2,$3)"
		_, err = tx.Exec(sql, c.Student.Email, hdr.Kind, hdr.Deadline)
		if err != nil {
			return rollback(err, tx)
		}
	}
	//Delete the old convetion if needed
	_, err = tx.Exec("delete from conventions where studentEmail=$1", c.Student.Email)
	if err != nil {
		return rollback(err, tx)
	}
	return tx.Commit()
}

func (srv *Service) SetSupervisor(stu string, t internship.Person) error {
	sql := "update internships set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

func (srv *Service) SetTutor(stu string, t string) error {
	sql := "update internships set tutor=$1 where student=$2"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, t, stu)
}

func (srv *Service) SetCompany(stu string, c internship.Company) error {
	sql := "update internships set companyWWW=$1, company=$2 where student=$3"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, c.WWW, c.Name, stu)
}

func (srv *Service) SetMajor(stu, m string) error {
	for _, x := range srv.majors {
		if x == m {
			sql := "update internships set major=$2 where student=$1"
			return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, stu, m)
		}
	}
	return internship.ErrInvalidMajor
}

func (srv *Service) Majors() []string {
	return srv.majors
}

func (srv *Service) SetPromotion(stu, p string) error {
	sql := "update internships set promotion=$2 where student=$1"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, stu, p)
}

func (srv *Service) Internship(stu string) (internship.Internship, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, promotion, major, tut.firstname, tut.lastname, tut.email, tut.tel, tut.role, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, startTime, endTime, company, companyWWW from internships, users as stu, users as tut where internships.tutor=tut.email and internships.student = stu.email and stu.email = $1"
	rows, err := srv.DB.Query(sql, stu)
	if err != nil || !rows.Next() {
		return internship.Internship{}, internship.ErrUnknownInternship
	}
	defer rows.Close()
	i, err := scanInternship(rows)
	if err != nil {
		return internship.Internship{}, err
	}
	err = srv.appendReports(&i)
	return i, err
}

func scanInternship(r *sql.Rows) (internship.Internship, error) {
	var stuFn, stuLn, stuEmail, stuTel, prom string
	var major sql.NullString
	var tutFn, tutLn, tutEmail, tutTel string
	var tutRole internship.Privilege
	var supFn, supLn, supEmail, supTel string
	var company, companyWWW string
	var title string
	var start, end time.Time
	err := r.Scan(&stuFn, &stuLn, &stuEmail, &stuTel, &prom, &major, &tutFn, &tutLn, &tutEmail, &tutTel, &tutRole, &supFn, &supLn, &supEmail, &supTel, &title, &start, &end, &company, &companyWWW)
	if err != nil {
		log.Println(err.Error())
		return internship.Internship{}, err
	}
	stu := internship.User{Firstname: stuFn, Lastname: stuLn, Email: stuEmail, Tel: stuTel, Role: internship.STUDENT}
	tut := internship.User{Firstname: tutFn, Lastname: tutLn, Email: tutEmail, Tel: tutTel, Role: tutRole}
	c := internship.Company{Name: company, WWW: companyWWW}
	sup := internship.Person{Firstname: supFn, Lastname: supLn, Email: supEmail, Tel: supTel}
	i := internship.Internship{Student: stu, Promotion: prom, Sup: sup, Tutor: tut, Cpy: c, Begin: start, End: end, Title: title}
	if major.Valid {
		i.Major = major.String
	}
	return i, err
}

func (srv *Service) appendReports(i *internship.Internship) error {
	s := "select kind, grade, deadline, comment from reports where student=$1 order by deadline"
	rows, err := srv.DB.Query(s, i.Student.Email)
	if err != nil {
		return err
	}
	i.Reports = make([]internship.ReportHeader, 0, 0)
	for rows.Next() {
		var kind string
		var comment sql.NullString
		var grade sql.NullInt64
		var deadline time.Time
		rows.Scan(&kind, &grade, &deadline, &comment)
		hdr := internship.ReportHeader{Kind: kind, Deadline: deadline, Grade: -1}
		if comment.Valid {
			hdr.Comment = comment.String
		}
		if grade.Valid {
			hdr.Grade = int(grade.Int64)
		}
		i.Reports = append(i.Reports, hdr)
	}
	defer rows.Close()
	return err
}

func (srv *Service) Internships() ([]internship.Internship, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, promotion, major, tut.firstname, tut.lastname, tut.email, tut.tel, tut.role, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, startTime, endTime, company, companyWWW from internships, users as stu, users as tut where internships.tutor=tut.email and internships.student = stu.email"
	internships := make([]internship.Internship, 0, 0)
	rows, err := srv.DB.Query(sql)
	if err != nil {
		return internships, err
	}
	defer rows.Close()
	for rows.Next() {
		i, err := scanInternship(rows)
		if err != nil {
			return internships, err
		}
		err = srv.appendReports(&i)
		if err != nil {
			return internships, err
		}
		internships = append(internships, i)
	}
	return internships, err
}
