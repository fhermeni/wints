package filter

import (
	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/mail"
)

//ownByStudent checks if the current user is a student and the target of the operation
func (v *Service) mine(email string) bool {
	return v.my.Email == email
}

//isTutoring checks if the current user tutors the student in parameters
func (v *Service) isTutoring(stu string) bool {
	i, err := v.srv.Internship(stu)
	if err != nil {
		return false
	}
	return i.Tutor.Email == v.my.Email
}

func (v *Service) UserLog(msg string, err error) {
	v.log.UserLog(v.my, msg, err)
}

//NewService wraps a current unfiltered service to restrict the operation
//with regards to a given user
func NewService(l journal.Journal, srv internship.Service, u internship.User, m mail.Mailer) (*Service, error) {
	return &Service{my: u, srv: srv, log: l, mailer: m}, nil
}
