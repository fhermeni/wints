package filter

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

//ownByStudent checks if the current user is a student and the target of the operation
func (v *Service) mine(email string) bool {
	return v.my.Email == email
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

func (v *Service) NewInternship(c internship.Convention) ([]byte, error) {
	if v.my.Role >= internship.ADMIN {
		return v.srv.NewInternship(c)
	}
	return []byte{}, ErrPermission
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

func (v *Service) SkipConvention(student string, skip bool) error {
	if v.my.Role >= internship.ADMIN {
		return v.srv.SkipConvention(student, skip)
	}
	return ErrPermission
}

func (v *Service) DeleteConvention(student string) error {
	if v.my.Role >= internship.ADMIN {
		return v.srv.DeleteConvention(student)
	}
	return ErrPermission
}

func (v *Service) Internship(stu string) (internship.Internship, error) {
	if v.mine(stu) || v.my.Role >= internship.MAJOR {
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
	if v.mine(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetSupervisor(stu, sup)
	}
	return ErrPermission
}

func (v *Service) SetTutor(stu string, t string) error {
	if v.my.Role >= internship.ADMIN {
		return v.srv.SetTutor(stu, t)
	}
	return ErrPermission
}

func (v *Service) SetMajor(stu, m string) error {
	if v.mine(stu) || v.my.Role >= internship.MAJOR {
		return v.srv.SetMajor(stu, m)
	}
	return ErrPermission
}

func (v *Service) Majors() []string {
	return v.srv.Majors()
}
func (v *Service) SetPromotion(stu, p string) error {
	if v.mine(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetPromotion(stu, p)
	}
	return ErrPermission
}

func (v *Service) SetCompany(stu string, c internship.Company) error {
	if v.mine(stu) || v.isTutoring(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetCompany(stu, c)
	}
	return ErrPermission
}

func (v *Service) SetTitle(stu string, title string) error {
	if v.mine(stu) || v.my.Role >= internship.ADMIN {
		return v.srv.SetTitle(stu, title)
	}
	return ErrPermission
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

func (v *Service) Registered(email string, password []byte) ([]byte, error) {
	return v.srv.Registered(email, password)
}

func (v *Service) OpenedSession(email, token string) error {
	return v.srv.OpenedSession(email, token)
}

func (v *Service) Logout(email, token string) error {
	return v.srv.Logout(email, token)
}

func (v *Service) NewTutor(p internship.User) ([]byte, error) {
	if v.my.Role == internship.ROOT {
		return v.srv.NewTutor(p)
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
	if v.mine(email) || v.my.Role >= internship.MAJOR {
		return v.srv.User(email)
	}
	return internship.User{}, ErrPermission
}

func (v *Service) Users() ([]internship.User, error) {
	if v.my.Role >= internship.MAJOR {
		return v.srv.Users()
	}
	return []internship.User{v.my}, nil
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
	if v.mine(email) || v.isTutoring(email) || v.my.Role >= internship.MAJOR {
		return v.srv.Report(kind, email)
	}
	return internship.ReportHeader{}, ErrPermission
}

func (v *Service) ReportContent(kind, email string) ([]byte, error) {
	if v.mine(email) || v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.ReportContent(kind, email)
	}
	return []byte{}, ErrPermission
}

func (v *Service) SetReportContent(kind, email string, cnt []byte) error {
	if v.mine(email) {
		return v.srv.SetReportContent(kind, email, cnt)
	}
	return ErrPermission
}

func (v *Service) SetReportGrade(kind, email string, r int, comment string) error {
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.SetReportGrade(kind, email, r, comment)
	}
	return ErrPermission
}

func (v *Service) SetReportDeadline(kind, email string, t time.Time) error {
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.SetReportDeadline(kind, email, t)
	}
	return ErrPermission
}

func (v *Service) SetReportPrivate(kind, email string, p bool) error {
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		return v.srv.SetReportPrivate(kind, email, p)
	}
	return ErrPermission
}

func (v *Service) SetAlumni(student string, a internship.Alumni) error {
	if v.my.Email == student || v.my.Role >= internship.ADMIN || v.isTutoring(student) {
		return v.srv.SetAlumni(student, a)
	}
	return ErrPermission
}

func (v *Service) SurveyToken(kind string) (string, string, error) {
	return v.srv.SurveyToken(kind)
}

func (v *Service) Survey(student, kind string) (internship.Survey, error) {
	if v.my.Role >= internship.MAJOR || v.isTutoring(student) {
		return v.srv.Survey(student, kind)
	}
	return internship.Survey{}, ErrPermission
}

func (v *Service) SetSurveyContent(token string, cnt map[string]string) error {
	return v.SetSurveyContent(token, cnt)
}

func (v *Service) SurveyDefs() []internship.SurveyDef {
	return v.srv.SurveyDefs()
}

func (v *Service) Statistics() ([]internship.Stat, error) {
	return v.srv.Statistics()
}
