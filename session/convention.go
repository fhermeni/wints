package session

import (
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
)

//Conventions lists the conventions if the emitter is an admin at minimum
func (s *Session) Conventions() ([]schema.Convention, error) {
	if s.Role() >= schema.ADMIN {
		return s.store.Conventions()
	}
	return []schema.Convention{}, ErrPermission
}

//SetSupervisor changes the supervisor if the emitter is the student or an admin at minimum
func (s *Session) SetSupervisor(stu string, sup schema.Person) error {
	if s.Myself(stu) || s.Role() >= schema.ADMIN {
		return s.store.SetSupervisor(stu, sup)
	}
	return ErrPermission
}

//SetTutor changes the tutor if the emitter is an admin at minimum
func (s *Session) SetTutor(stu string, t string) error {
	if s.Role() > schema.ADMIN {
		return s.store.SetTutor(stu, t)
	}
	return ErrPermission
}

//SetCompany changes the company if the emitter is the student or an admin at minimum
func (s *Session) SetCompany(stu string, c schema.Company) error {
	if s.Myself(stu) || s.Role() >= schema.ADMIN {
		return s.store.SetCompany(stu, c)
	}
	return ErrPermission
}

//SetTitle changes the title if the emitter is the student or an admin at minimum
func (s *Session) SetTitle(stu string, title string) error {
	if s.Myself(stu) || s.Role() >= schema.ADMIN {
		return s.store.SetTitle(stu, title)
	}
	return ErrPermission
}

//SetConventionSkippable changes the skippable status if the emitter is an admin at minimum
func (s *Session) SetConventionSkippable(student string, skip bool) error {
	if s.Role() >= schema.ADMIN {
		return s.store.SetConventionSkippable(student, skip)
	}
	return ErrPermission

}

//ValidateConvention validates the convention if the emitter is an admin at minimum
func (s *Session) ValidateConvention(stu string, cfg config.Internships) error {
	if s.Role() >= schema.ADMIN {
		return s.store.ValidateConvention(stu, cfg)
	}
	return ErrPermission
}

//Internships list the internships if the emitter is at least a major leader.
//Otherwise, all the internships now tutored by the emitter are removed
func (s *Session) Internships() (schema.Internships, error) {
	if s.Role() >= schema.MAJOR {
		return s.store.Internships()
	}
	is, err := s.store.Internships()
	return is.Filter(schema.Tutoring(s.my.Person.Email)), err
}

//Internship returns the internship of the emitter
func (s *Session) Internship(stu string) (schema.Internship, error) {
	if s.Myself(stu) {
		return s.store.Internship(stu)
	}
	return schema.Internship{}, ErrPermission
}
