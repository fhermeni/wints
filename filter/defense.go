package filter

import (
	"strconv"

	"github.com/fhermeni/wints/internship"
)

//Defenses
func (s *Service) DefenseSessions() ([]internship.DefenseSession, error) {
	err := ErrPermission
	defs := []internship.DefenseSession{}

	if s.my.Role >= internship.ADMIN {
		defs, err = s.srv.DefenseSessions()
	} else if s.my.Role != internship.NONE {
		var sessions []internship.DefenseSession
		sessions, err = s.srv.DefenseSessions()
		if err == nil {
			defs = make([]internship.DefenseSession, 0, 0)
			for _, session := range sessions {
				if session.InJury(s.my.Email) {
					defs = append(defs, session)
				}
			}
		}
	}
	return defs, err
}

func (s *Service) DefenseSession(student string) (internship.DefenseSession, error) {
	err := ErrPermission
	def := internship.DefenseSession{}
	if s.mine(student) || s.my.Role >= internship.MAJOR {
		def, err = s.srv.DefenseSession(student)
	}
	s.UserLog("get defense of '"+student+"'", err)
	return def, err
}

func (s *Service) SetDefenseGrade(student string, g int) error {
	err := ErrPermission
	if s.my.Role >= internship.ADMIN {
		err = s.srv.SetDefenseGrade(student, g)
	}
	def, err := s.DefenseSession(student)
	if err == nil && def.InJury(s.my.Email) {
		err = s.srv.SetDefenseGrade(student, g)
	}
	s.UserLog("grades defense of '"+student+"' to "+strconv.Itoa(g), err)
	return err
}

func (s *Service) SetDefenseSessions(defs []internship.DefenseSession) error {
	err := ErrPermission
	if s.my.Role >= internship.ADMIN {
		err = s.srv.SetDefenseSessions(defs)
	}
	s.UserLog("new defense program", err)
	return err
}

func (s *Service) PublicDefenseSessions() ([]internship.PublicDefenseSession, error) {
	return s.srv.PublicDefenseSessions()
}
