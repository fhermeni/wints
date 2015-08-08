package filter

import (
	"strconv"

	"github.com/fhermeni/wints/internship"
)

func (v *Service) NewInternship(c internship.Convention) ([]byte, error) {
	err := ErrPermission
	res := []byte{}
	if v.my.Role >= internship.ADMIN {
		res, err = v.srv.NewInternship(c)
	}
	v.UserLog("import convention of '"+c.Student.Email+"'", err)
	if err == nil {
		s := internship.User{Firstname: c.Student.Firstname, Lastname: c.Student.Lastname, Email: c.Student.Email}
		v.mailer.SendStudentInvitation(s, res)
		v.mailer.SendTutorNotification(c.Student, c.Tutor)
	}
	return res, err
}

func (v *Service) NewConvention(c internship.Convention) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN {
		err = v.srv.NewConvention(c)
	}
	v.UserLog("add convention '"+c.Student.Email+"'", err)
	return err
}

func (v *Service) Conventions() ([]internship.Convention, error) {
	if v.my.Role >= internship.ADMIN {
		return v.srv.Conventions()
	}
	return []internship.Convention{}, ErrPermission
}

func (v *Service) SkipConvention(student string, skip bool) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN {
		err = v.srv.SkipConvention(student, skip)
	}
	v.UserLog("skip convention of '"+student+"'", err)
	return err
}

func (v *Service) DeleteConvention(student string) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN {
		err = v.srv.DeleteConvention(student)
	}
	v.UserLog("delete convention '"+student+"'", err)
	return err
}

func (v *Service) Internship(stu string) (internship.Internship, error) {
	err := ErrPermission
	i := internship.Internship{}
	if v.mine(stu) || v.my.Role >= internship.MAJOR {
		i, err = v.srv.Internship(stu)
	} else {
		i, err = v.srv.Internship(stu)
		if err == nil && i.Tutor.Email != v.my.Email {
			i = internship.Internship{}
			err = ErrPermission
		}
	}
	if !v.mine(stu) {
		v.UserLog("wants internship '"+stu+"'", err)
	}
	return i, nil
}

func (v *Service) SetSupervisor(stu string, sup internship.Person) error {
	err := ErrPermission
	if v.mine(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		err = v.srv.SetSupervisor(stu, sup)
	}
	v.UserLog("set '"+sup.Email+"' as supervisor for '"+stu+"'", ErrPermission)
	return err
}

func (v *Service) SetTutor(stu string, t string) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN {
		err = v.srv.SetTutor(stu, t)
	}
	v.UserLog("set '"+t+"' as tutor for internship '"+stu+"'", err)
	if err == nil {
		i, err := v.srv.Internship(stu)
		if err == nil {
			if u, err := v.srv.User(t); err == nil {
				v.mailer.SendTutorUpdate(i.Student, i.Tutor, u)
			}
		}
	}

	return err
}

func (v *Service) SetMajor(stu, m string) error {
	err := ErrPermission
	if v.mine(stu) || v.my.Role >= internship.MAJOR {
		err = v.srv.SetMajor(stu, m)
	}
	v.UserLog("set '"+m+"' as major for internship '"+stu+"'", err)
	return err
}

func (v *Service) Majors() []string {
	return v.srv.Majors()
}
func (v *Service) SetPromotion(stu, p string) error {
	err := ErrPermission
	if v.mine(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		err = v.srv.SetPromotion(stu, p)
	}
	v.UserLog("set '"+p+"' as promotion for internship '"+stu+"'", err)
	return err
}

func (v *Service) SetCompany(stu string, c internship.Company) error {
	err := ErrPermission
	if v.mine(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		err = v.srv.SetCompany(stu, c)
	}
	v.UserLog("set '"+c.Name+"' ("+c.WWW+") as company for internship '"+stu+"'", err)
	return err
}

func (v *Service) SetTitle(stu string, title string) error {
	err := ErrPermission
	if v.mine(stu) || v.my.Role >= internship.ADMIN {
		err = v.srv.SetTitle(stu, title)
	}
	v.UserLog("set '"+title+"' as title for internship '"+stu+"'", err)
	return err
}

func (v *Service) Internships() ([]internship.Internship, error) {
	//Student get his own, others get everything
	if v.my.Role == internship.NONE {
		i, err := v.srv.Internship(v.my.Email)
		if err != nil {
			return []internship.Internship{}, err
		}
		return []internship.Internship{i}, err
	} else if v.my.Role >= internship.MAJOR {
		return v.srv.Internships()
	}
	//filter out internship I don't handle
	mine := make([]internship.Internship, 0, 0)
	is, err := v.srv.Internships()
	if err != nil {
		return is, err
	}
	for _, i := range is {
		if i.Tutor.Email == v.my.Email {
			mine = append(mine, i)
		}
	}
	return mine, nil
}

func (v *Service) SetAlumni(student string, a internship.Alumni) error {
	err := ErrPermission
	if v.my.Email == student || v.my.Role >= internship.ADMIN || v.isTutoring(student) {
		err = v.srv.SetAlumni(student, a)
	}
	v.UserLog("set alumni of '"+student+"' to '"+a.Contact+"','"+strconv.Itoa(a.Position)+"'", err)
	return err
}

func (v *Service) SetNextPosition(student string, p int) error {
	def, err := v.DefenseSession(student)
	if err != nil {
		return err
	}
	if !def.InJury(v.my.Email) {
		return ErrPermission
	}
	err = v.SetNextPosition(student, p)
	v.UserLog("set next position of '"+student+"' to '"+strconv.Itoa(p)+"'", err)
	return err
}
