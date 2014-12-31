package datastore

import (
	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

func (srv *Service) NewConvention(c internship.Convention) error {
	sql := "insert into internships(studentEmail, studentFn, studentLn, studentTel, tutorEmail, tutorFn, tutorLn, tutorTel, supervisorEmail, supervisorFn, supervisorLn, supervisorTel, startTime, endTime, title, company, companyWWW) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	_, err := srv.DB.Exec(sql, c.Student.Email, c.Student.Firstname, c.Student.Lastname, c.Student.Tel, c.Tutor.Email, c.Tutor.Firstname, c.Tutor.Lastname, c.Tutor.Tel, c.Supervisor.Email, c.Supervisor.Firstname, c.Supervisor.Lastname, c.Supervisor.Tel, c.Begin, c.End, c.Title, c.Cpy.Name, c.Cpy.WWW)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Constraint == "conventions_pkey" {
				return internship.ErrConventionExists
			}
		}
		return err

	}
	return err
}
