package session

import "github.com/fhermeni/wints/schema"

//SetStudentSkippable change the skippable status if the emitter is a major leader at minimum
func (s *Session) SetStudentSkippable(em string, st bool) error {
	if s.Role() >= schema.MAJOR {
		return s.store.SetStudentSkippable(em, st)
	}
	return ErrPermission
}

//Students lists all the students if the emitter is an admin at least
func (s *Session) Students() ([]schema.Student, error) {
	if s.Role() >= schema.ADMIN {
		return s.store.Students()
	}
	return []schema.Student{}, ErrPermission
}

func (s *Session) Student(stu string) (schema.Student, error) {
	if s.Myself(stu) || s.Tutoring(stu) || s.Role() >= schema.MAJOR {
		return s.store.Student(stu)
	}
	return schema.Student{}, ErrPermission
}

//SetNextPosition changes the student next position if the emitter is the targetted student,
//the tutor, a member of his jury or a major leader at minimum
func (s *Session) SetAlumni(student string, a schema.Alumni) error {
	if s.Myself(student) || s.Role() >= schema.MAJOR || s.Tutoring(student) || s.JuryOf(student) {
		return s.store.SetAlumni(student, a)
	}
	return ErrPermission
}

//SetMajor changes the student major if the emitter is the student or at least a major leader
func (s *Session) SetMajor(student string, m string) error {
	if s.Myself(student) || s.Role() >= schema.MAJOR {
		return s.store.SetMajor(student, m)
	}
	return ErrPermission
}

//SetPromotion changes the student promotion if the emitter is the student himself or a major leader at least
func (s *Session) SetPromotion(student string, p string) error {
	if s.Myself(student) || s.Role() >= schema.MAJOR {
		return s.store.SetPromotion(student, p)
	}
	return ErrPermission
}

//SetMale changes the student gender if the emitter is the student itself, the tutor or an admin at minimum
func (s *Session) SetMale(student string, male bool) error {
	if s.Myself(student) || s.Tutoring(student) || s.Role() >= schema.MAJOR {
		return s.store.SetMale(student, male)
	}
	return ErrPermission
}
