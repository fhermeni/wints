package datastore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

func (srv *Service) NewInternship(student, tutor string, from, to time.Time, c internship.Company, sup internship.Supervisor, title string) error {
	if from.After(to) {
		return internship.ErrInvalidPeriod
	}
	sql := "insert into internships(student, tutor, startTime, endTime, title, supervisorFn, supervisorLn, supervisorTel, supervisorEmail, company, companyWWW) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	res, err := srv.DB.Exec(sql, student, tutor, from, to, title, sup.Firstname, sup.Lastname, sup.Tel, sup.Email, c.Name, c.WWW)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Constraint == "internships_student_fkey" || e.Constraint == "internships_tutor_fkey" {
				return internship.ErrUnknownUser
			}
			if e.Constraint == "internships_pkey" {
				return internship.ErrInternshipExists
			}
		}
		return err

	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return internship.ErrInternshipExists
	}
	return err
}

func (srv *Service) SetSupervisor(stu string, t internship.Supervisor) error {
	sql := "update internships set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

func (srv *Service) SetCompany(stu string, c internship.Company) error {
	sql := "update internships set companyWWW=$1, company=$2 where student=$3"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, c.WWW, c.Name, stu)
}

func (srv *Service) SetMajor(stu, m string) error {
	sql := "update internships set major=$2 where student=$1"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, stu, m)
}

func (srv *Service) SetPromotion(stu, p string) error {
	sql := "update internships set promotion=$2 where student=$1"
	return SingleUpdate(srv.DB, internship.ErrUnknownInternship, sql, stu, p)
}

func (srv *Service) Internship(stu string) (internship.Internship, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, tut.firstname, tut.lastname, tut.email, tut.tel, tut.role, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, startTime, endTime, company, companyWWW from internships, users as stu, users as tut where internships.tutor=tut.email and internships.student = stu.email and stu.email = $1"
	rows, err := srv.DB.Query(sql, stu)
	if err != nil || !rows.Next() {
		return internship.Internship{}, internship.ErrUnknownInternship
	}

	defer rows.Close()
	return scanInternship(rows)
}

func scanInternship(r *sql.Rows) (internship.Internship, error) {
	var stuFn, stuLn, stuEmail, stuTel string
	var tutFn, tutLn, tutEmail, tutTel string
	var tutRole internship.Privilege
	var supFn, supLn, supEmail, supTel string
	var company, companyWWW string
	var title string
	var start, end time.Time
	err := r.Scan(&stuFn, &stuLn, &stuEmail, &stuTel, &tutFn, &tutLn, &tutEmail, &tutTel, &tutRole, &supFn, &supLn, &supEmail, &supTel, &title, &start, &end, &company, &companyWWW)
	if err != nil {
		return internship.Internship{}, err
	}
	stu := internship.User{Firstname: stuFn, Lastname: stuLn, Email: stuEmail, Tel: stuTel, Role: internship.STUDENT}
	tut := internship.User{Firstname: tutFn, Lastname: tutLn, Email: tutEmail, Tel: tutTel, Role: tutRole}
	c := internship.Company{Name: company, WWW: companyWWW}
	sup := internship.Supervisor{Firstname: supFn, Lastname: supLn, Email: supEmail, Tel: supTel}
	return internship.Internship{Student: stu, Sup: sup, Tutor: tut, Cpy: c, Begin: start, End: end, Title: title}, nil
}

func (srv *Service) Internships() ([]internship.Internship, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, tut.firstname, tut.lastname, tut.email, tut.tel, tut.role, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, startTime, endTime, company, companyWWW from internships, users as stu, users as tut where internships.tutor=tut.email and internships.student = stu.email"
	internships := make([]internship.Internship, 0, 0)
	rows, err := srv.DB.Query(sql)
	if err != nil {
		return internships, err
	}
	defer rows.Close()
	for rows.Next() {
		i, _ := scanInternship(rows)
		internships = append(internships, i)
	}
	return internships, nil
}
