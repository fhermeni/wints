package datastore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
)

func (srv *Service) SetSupervisor(stu string, t internship.Supervisor) error {
	sql := "update internships set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	return SingleUpdate(srv.Db, internship.ErrUnknownInternship, sql, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

func (srv *Service) SetCompany(stu string, c internship.Company) error {
	sql := "update internships set companyWWW=$1, company=$2 where student=$3"
	return SingleUpdate(srv.Db, internship.ErrUnknownInternship, sql, c.WWW, c.Name, stu)
}

func (srv *Service) Internship(stu string) (internship.Internship, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, tut.firstname, tut.lastname, tut.email, tut.tel, tut.role, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, startTime, endTime, company, companyWWW from internships, users as stu, users as tut where internships.tutor=tut.email and internships.student = stu.email and stu.email = $1"
	rows, err := srv.Db.Query(sql, stu)
	if err != nil {
		return internship.Internship{}, internship.ErrUnknownInternship
	}
	defer rows.Close()
	return scanInternship(rows), nil
}

func scanInternship(r *sql.Rows) internship.Internship {
	var stuFn, stuLn, stuEmail, stuTel string
	var tutFn, tutLn, tutEmail, tutTel string
	var tutRole internship.Privilege
	var supFn, supLn, supEmail, supTel string
	var company, companyWWW string
	var title string
	var start, end time.Time
	r.Scan(&stuFn, &stuLn, &stuEmail, &stuTel, &tutFn, &tutLn, &tutEmail, &tutTel, &supFn, &supLn, &supEmail, &supTel, &title, &start, &end, &company, &companyWWW)
	stu := internship.User{Firstname: stuFn, Lastname: stuLn, Email: stuEmail, Tel: stuTel, Role: internship.STUDENT}
	tut := internship.User{Firstname: tutFn, Lastname: tutLn, Email: tutEmail, Tel: tutTel, Role: tutRole}
	c := internship.Company{Name: company, WWW: companyWWW}
	sup := internship.Supervisor{Firstname: supFn, Lastname: supLn, Email: supEmail, Tel: supTel}
	return internship.Internship{Student: stu, Sup: sup, Tutor: tut, Cpy: c, Begin: start, End: end, Title: title}
}

func (srv *Service) Internships() ([]internship.Internship, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, tut.firstname, tut.lastname, tut.email, tut.tel, tut.role, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, startTime, endTime, company, companyWWW from internships, users as stu, users as tut where internships.tutor=tut.email and internships.student = stu.email"
	internships := make([]internship.Internship, 0, 0)
	rows, err := srv.Db.Query(sql)
	if err != nil {
		return internships, err
	}
	defer rows.Close()
	for rows.Next() {
		i := scanInternship(rows)
		internships = append(internships, i)
	}
	return internships, nil
}
