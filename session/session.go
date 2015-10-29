//Package session provides the material to implement the security layer.
//A session restrict the operation that are authorized by the session owner, depending on its identity or his role
package session

import (
	"errors"

	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/sqlstore"
)

var (
	//ErrPermission indicates an operation that is not permitted
	ErrPermission = errors.New("Permission denied")
)

//Session restricts the operation that can be executed by the current user with regards
//to its role and or relationships
type Session struct {
	my          schema.User
	store       *sqlstore.Store
	conventions feeder.Conventions
	//log    journal.Journal
	//mailer mail.Mailer
}

//NewSession creates a new session
func NewSession(u schema.User, store *sqlstore.Store, conventions feeder.Conventions) Session {
	return Session{my: u, store: store, conventions: conventions}
}

func (s *Session) RmSession(em string) error {
	if s.Myself(em) || s.Role().Level() >= schema.ADMIN_LEVEL {
		return s.store.RmSession(em)
	}
	return ErrPermission
}
func (s *Session) Me() schema.User {
	return s.my
}

//Role returns the current user role
func (s *Session) Role() schema.Role {
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
