package filter

import "github.com/fhermeni/wints/internship"

func (s *Service) Students() ([]internship.Student, error) {
	err := ErrPermission
	stus := []internship.Student{}
	if s.my.Role >= internship.ADMIN {
		stus, err = s.srv.Students()
	}
	s.UserLog("wants to list students", err)
	return stus, err
}

func (s *Service) AddStudent(st internship.Student) error {
	err := ErrPermission
	if s.my.Role >= internship.ADMIN {
		err = s.srv.AddStudent(st)
	}
	s.UserLog("add student '"+st.Email+"'", err)
	return err
}

func (s *Service) AlignWithInternship(student string, intern string) error {
	err := ErrPermission
	if s.my.Role >= internship.ADMIN {
		err = s.srv.AlignWithInternship(student, intern)
	}
	s.UserLog("Student '"+student+"' aligned with internship '"+intern+"'", err)
	return err
}

func (s *Service) InsertStudents(file string) error {
	err := ErrPermission
	if s.my.Role >= internship.ADMIN {
		err = s.srv.InsertStudents(file)
	}
	s.UserLog("declares the students", err)
	return err
}

func (s *Service) HideStudent(em string, st bool) error {
	err := ErrPermission
	if s.my.Role >= internship.ADMIN {
		err = s.srv.HideStudent(em, st)
	}
	status := "hide"
	if !st {
		status = "visible"
	}
	s.UserLog("Status for student '"+em+"' set to '"+status+"'", err)
	return err
}
