package session

import "github.com/fhermeni/wints/schema"

//Conventions lists the conventions if the emitter is an admin at minimum
func (s *Session) Conventions() ([]schema.Convention, error) {
	if s.Role().Level() >= schema.ADMIN_LEVEL {
		return s.conventions.Import()
	}
	return []schema.Convention{}, ErrPermission
}

//Convention returns the convention of a given student if the emitter is the student or at least an admin
func (s *Session) Convention(stu string) (schema.Convention, error) {
	if s.Myself(stu) || s.Role().Level() >= schema.ADMIN_LEVEL {
		return s.store.Convention(stu)
	}
	return schema.Convention{}, ErrPermission
}

//SetSupervisor changes the supervisor if the emitter is the student or an admin at minimum
func (s *Session) SetSupervisor(stu string, sup schema.Person) error {
	if s.Myself(stu) || s.Role().Level() >= schema.ADMIN_LEVEL {
		return s.store.SetSupervisor(stu, sup)
	}
	return ErrPermission
}

//SetTutor changes the tutor if the emitter is an admin at minimum
func (s *Session) SetTutor(stu string, t string) error {
	if s.Role().Level() > schema.ADMIN_LEVEL {
		return s.store.SetTutor(stu, t)
	}
	return ErrPermission
}

//SetCompany changes the company if the emitter is the student or an admin at minimum
func (s *Session) SetCompany(stu string, c schema.Company) error {
	if s.Myself(stu) || s.Role().Level() >= schema.ADMIN_LEVEL {
		return s.store.SetCompany(stu, c)
	}
	return ErrPermission
}

//NewInternship validates the convention if the emitter is an admin at minimum
func (s *Session) NewInternship(c schema.Convention) (schema.Internship, []byte, error) {
	if s.Role().Level() >= schema.ADMIN_LEVEL {
		return s.store.NewInternship(c)
	}
	return schema.Internship{}, []byte{}, ErrPermission
}

//Internships list the internships if the emitter is at least a major leader.
//Otherwise, all the internships now tutored by the emitter are removed
func (s *Session) Internships() (schema.Internships, error) {
	is, err := s.store.Internships()
	if s.Role().Level() >= schema.HEAD_LEVEL {
		return is, err
	} else if s.Role().Level() == schema.MAJOR_LEVEL {
		return is.Filter(schema.InMajor(s.my.Role.SubRole())), err
	} else if s.Role().Level() == schema.TUTOR_LEVEL {
		return is.Filter(schema.Tutoring(s.my.Person.Email)), err
	}
	return schema.Internships{}, ErrPermission

}

//Internship returns the internship of the emitter or if the emitter is at least a major leader
func (s *Session) Internship(stu string) (schema.Internship, error) {
	if s.Myself(stu) || s.Role().Level() >= schema.MAJOR_LEVEL {
		return s.store.Internship(stu)
	}
	return schema.Internship{}, ErrPermission
}
