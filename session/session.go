//Package session provides the material to implement the security layer.
//A session restrict the operation that are authorized by the session owner, depending on its identity or his role
package session

import (
	"errors"

	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/sqlstore"
)

var (
	//ErrPermission indicates an operation that is not permitted
	ErrPermission = errors.New("Permission denied")
	//ErrConfidentialReport indicates the report cannot be access due to its confidential status
	ErrConfidentialReport = errors.New("Confidential report")
)

//Session restricts the operation that can be executed by the current user with regards
//to its role and or relationships
type Session struct {
	my    schema.User
	store sqlstore.Store
	//log    journal.Journal
	//mailer mail.Mailer
}

//Role returns the current user role
func (s *Session) Role() schema.Privilege {
	return s.my.Role
}

//Myself checks if the given email matches the session one
func (s *Session) Myself(email string) bool {
	return s.my.Person.Email == email
}

//Tutoring checks if the session owner is a tutor of the given student
func (s *Session) Tutoring(student string) bool {
	c, err := s.store.Convention(student)
	if err != nil {
		return false
	}
	return c.Tutor.Person.Email == s.my.Person.Email
}

//JuryOf checks if I attend to the student defense
func (s *Session) JuryOf(student string) bool {
	_, err := s.store.Defense(student)
	if err != nil {
		return false
	}
	return false
}
