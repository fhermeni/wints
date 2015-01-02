package filter

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

//ownByStudent checks if the current user is a student and the target of the operation
func (v *Service) ownByStudent(email string) bool {
	return v.my.Role == internship.STUDENT && v.my.Email == email
}

//isTutoring checks if the current user tutors the student in parameters
func (v *Service) isTutoring(stu string) bool {
	i, err := v.srv.Internship(stu)
	if err != nil {
		return false
	}
	return i.Tutor.Email == v.my.Email
}

//NewService wraps a current unfiltered service to restrict the operation
//with regards to a given user
func NewService(srv internship.Service, u internship.User) (*Service, error) {
	return &Service{my: u, srv: srv}, nil
}

func (v *Service) NewInternship(c internship.Convention) error {
	if v.my.Role >= internship.ADMIN {
		return v.srv.NewInternship(c)
	}
	return ErrPermission
}

func (v *Service) NewConvention(c internship.Convention) error {
	if v.my.Role >= internship.ADMIN {
		return v.srv.NewConvention(c)
	}
	return ErrPermission
}

func (v *Service) Conventions() ([]internship.Convention, error) {
	if v.my.Role >= internship.ADMIN {
		return v.srv.Conventions()
	}
	return []internship.Convention{}, ErrPermission
}

func (v *Service) Internship(stu string) (internship.Internship, error) {
	if v.ownByStudent(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.Internship(stu)
	}
	i, err := v.srv.Internship(stu)
	if err != nil {
		return internship.Internship{}, err
	}
	if i.Tutor.Email != v.my.Email {
		return internship.Internship{}, ErrPermission
	}
	return i, nil
}

func (v *Service) SetSupervisor(stu string, sup internship.Person) error {
	if v.ownByStudent(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetSupervisor(stu, sup)
	}
	return ErrPermission
}

func (v *Service) SetMajor(stu, m string) error {
	if v.ownByStudent(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetMajor(stu, m)
	}
	return ErrPermission
}

func (v *Service) SetPromotion(stu, p string) error {
	if v.ownByStudent(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetPromotion(stu, p)
	}
	return ErrPermission
}

func (v *Service) SetCompany(stu string, c internship.Company) error {
	if v.ownByStudent(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetCompany(stu, c)
	}
	return ErrPermission
}

func (v *Service) Internships() ([]internship.Internship, error) {
	//Student get his own, others get everything
	if v.my.Role == internship.STUDENT {
		i, err := v.srv.Internship(v.my.Email)
		if err != nil {
			return []internship.Internship{}, err
		}
		return []internship.Internship{i}, err
	} else if v.my.Role >= internship.ADMIN {
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

func (v *Service) Registered(email string, password []byte) (internship.User, error) {
	return v.srv.Registered(email, password)
}

func (v *Service) NewUser(p internship.User) ([]byte, error) {
	if v.my.Role == internship.ROOT {
		return v.srv.NewUser(p)
	}
	return []byte{}, ErrPermission
}

func (v *Service) RmUser(email string) error {
	if v.my.Role == internship.ROOT && email != v.my.Email {
		return v.srv.RmUser(email)
	}
	return ErrPermission
}

func (v *Service) User(email string) (internship.User, error) {
	if v.ownByStudent(email) || v.my.Role >= internship.ADMIN {
		return v.srv.User(email)
	}
	return internship.User{}, ErrPermission
}

func (v *Service) Users() ([]internship.User, error) {
	if v.my.Role >= internship.ADMIN {
		return v.srv.Users()
	}
	return []internship.User{}, ErrPermission
}

//Change the user password
func (v *Service) SetUserPassword(email string, oldP, newP []byte) error {
	if v.my.Email == email {
		return v.srv.SetUserPassword(email, oldP, newP)
	}
	return ErrPermission
}

//Change user profile
func (v *Service) SetUserProfile(email, fn, ln, tel string) error {
	if v.my.Email == email {
		return v.srv.SetUserProfile(email, fn, ln, tel)
	}
	return ErrPermission
}

//Change user role
func (v *Service) SetUserRole(email string, priv internship.Privilege) error {
	if v.my.Role >= internship.ADMIN && email != v.my.Email {
		return v.srv.SetUserRole(email, priv)
	}
	return ErrPermission
}

func (v *Service) ResetPassword(email string) ([]byte, error) {
	return v.srv.ResetPassword(email)
}

func (v *Service) NewPassword(token, newP []byte) (string, error) {
	return v.srv.NewPassword(token, newP)
}

func (v *Service) PlanReport(student string, r internship.ReportHeader) error {
	if v.my.Role == internship.ROOT {
		return v.srv.PlanReport(student, r)
	}
	return ErrPermission
}

func (v *Service) ReportDefs() []internship.ReportDef {
	return v.srv.ReportDefs()
}

func (v *Service) Report(kind, email string) (internship.ReportHeader, error) {
	if v.ownByStudent(email) || v.my.Role >= internship.ADMIN {
		return v.srv.Report(kind, email)
	}
	return internship.ReportHeader{}, ErrPermission
}

func (v *Service) ReportContent(kind, email string) ([]byte, error) {
	if v.ownByStudent(email) || v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.ReportContent(kind, email)
	}
	return []byte{}, ErrPermission
}

func (v *Service) SetReportContent(kind, email string, cnt []byte) error {
	if v.ownByStudent(email) {
		return v.srv.SetReportContent(kind, email, cnt)
	}
	return ErrPermission
}

func (v *Service) SetReportGrade(kind, email string, r int) error {
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.SetReportGrade(kind, email, r)
	}
	return ErrPermission
}

func (v *Service) SetReportDeadline(kind, email string, t time.Time) error {
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.SetReportDeadline(kind, email, t)
	}
	return ErrPermission
}

func (v *Service) ResetRootAccount() error {
	if v.my.Role == internship.ROOT {
		return v.srv.ResetRootAccount()
	}
	return ErrPermission
}
