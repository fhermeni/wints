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
}

//AnonSession creates a session that is not attached to a particular user
func AnonSession(store *sqlstore.Store) Session {
	return Session{store: store}
}

//NewSession creates a new session
func NewSession(u schema.User, store *sqlstore.Store, conventions feeder.Conventions) Session {
	return Session{my: u, store: store, conventions: conventions}
}

//RmSession delete the session if the emitter is the session owner or at least an admin
func (s *Session) RmSession(em string) error {
	if s.Myself(em) || s.Role().Level() >= schema.AdminLevel {
		return s.store.RmSession(em)
	}
	return ErrPermission
}

//Me returns the session emitter
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

//Watching checks if the student is in the major I am the leader of, an head or not or
//the student tutor
func (s *Session) Watching(student string) bool {
	//watching
	if s.my.Role.Level() >= schema.HeadLevel {
		return true
	}
	c, err := s.store.Convention(student)
	if err != nil {
		return false
	}
	//wathing
	if s.my.Role.Level() == schema.MajorLevel && s.my.Role.SubRole() == c.Student.Major {
		return true
	}
	//tutoring
	return s.my.Person.Email == c.Tutor.Person.Email
}

//InMyMajor checks if the student is in the major I am the leader of
func (s *Session) InMyMajor(student string) bool {
	stu, err := s.store.Student(student)
	if err != nil {
		return false
	}
	return s.my.Role.SubRole() == stu.Major
}

//JuryOf checks if I am in a jury for a defense.
//That if indeed I am in the jury, or an admin
func (s *Session) JuryOf(student string) bool {
	if s.my.Role.Level() >= schema.AdminLevel {
		return true
	}
	def, err := s.store.Defense(student)
	if err != nil {
		return false
	}
	mySession, err := s.store.DefenseSession(def.Room, def.SessionId)
	if err != nil {
		return false
	}
	return mySession.InJury(s.my.Person.Email)
}
