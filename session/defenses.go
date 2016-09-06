package session

import (
	"time"

	"github.com/fhermeni/wints/schema"
)

func (s *Session) NewDefenseSession(room string, id string) (schema.DefenseSession, error) {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.NewDefenseSession(room, id)
	}
	return schema.DefenseSession{}, ErrPermission
}

func (s *Session) RmDefenseSession(room string, id string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.RmDefenseSession(room, id)
	}
	return ErrPermission
}

func (s *Session) SetStudentDefense(session, room, student string, t time.Time, public, local bool) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.SetStudentDefense(session, room, student, t, public, local)
	}
	return ErrPermission
}

func (s *Session) UpdateStudentDefense(student string, t time.Time, public, local bool) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.UpdateStudentDefense(student, t, public, local)
	}
	return ErrPermission
}

func (s *Session) RmDefense(student string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.RmStudentDefense(student)
	}
	return ErrPermission
}
func (s *Session) AddJuryToDefenseSession(room, id, jury string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.AddJuryToDefenseSession(room, id, jury)
	}
	return ErrPermission
}

func (s *Session) DelJuryToDefenseSession(room, id, jury string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.DelJuryToDefenseSession(room, id, jury)
	}
	return ErrPermission
}

//DefenseSessions returns all the sessions
//Unfiltered for an admin and more. Otherwise, filter out the sessions I am not a jury member of
func (s *Session) DefenseSessions() ([]schema.DefenseSession, error) {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.DefenseSessions()
	}
	//filter out the sessions I am not a jury member of
	sessions, err := s.store.DefenseSessions()
	if err != nil {
		return []schema.DefenseSession{}, err
	}
	mine := []schema.DefenseSession{}
	for _, session := range sessions {
		if session.InJury(s.Me().Person.Email) {
			mine = append(mine, session)
		}
	}
	return mine, nil
}

func (s *Session) DefenseProgram() ([]schema.DefenseSession, error) {
	ss, err := s.store.DefenseSessions()
	if err != nil {
		return []schema.DefenseSession{}, err
	}
	/*for idx, x := range ss {
		x.Anonymise()
		ss[idx] = x
	}*/
	return ss, nil
}
func (s *Session) DefenseSession(room, id string) (schema.DefenseSession, error) {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.DefenseSession(room, id)
	}
	//Return the whole session if jury
	mySession, err := s.store.DefenseSession(room, id)
	if err != nil {
		return schema.DefenseSession{}, err
	}
	if mySession.InJury(s.Me().Person.Email) {
		return mySession, nil
	}

	ss2 := schema.DefenseSession{}
	ss2.Id = mySession.Id
	ss2.Room = mySession.Room
	ss2.Juries = mySession.Juries
	ss2.Defenses = []schema.Defense{}
	for _, def := range mySession.Defenses {
		if def.Student.User.Person.Email == s.Me().Person.Email {
			ss2.Defenses = append(ss2.Defenses, def)
		}
	}
	return ss2, nil
}

func (s *Session) Defense(stu string) (schema.Defense, error) {
	if s.Myself(stu) || s.JuryOf(stu) {
		return s.store.Defense(stu)
	}
	return schema.Defense{}, ErrPermission
}

func (s *Session) SetDefenseGrade(stu string, grade int) error {
	if s.JuryOf(stu) {
		return s.store.SetDefenseGrade(stu, grade)
	}
	return ErrPermission
}
