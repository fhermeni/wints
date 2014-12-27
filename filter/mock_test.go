package filter

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

type MockBackend struct{}

func (m *MockBackend) NewInternship(student, tutor string, from, to time.Time, c internship.Company, sup internship.Supervisor, Title string) error {
	return nil
}
func (m *MockBackend) Internship(stu string) (internship.Internship, error) {
	return internship.Internship{}, nil
}
func (m *MockBackend) SetSupervisor(stu string, sup internship.Supervisor) error {
	return nil
}
func (m *MockBackend) SetCompany(stu string, c internship.Company) error {
	return nil
}
func (m *MockBackend) Internships() ([]internship.Internship, error) {
	return []internship.Internship{}, nil
}
func (m *MockBackend) Registered(email string, password []byte) (internship.User, error) {
	return internship.User{}, nil
}
func (m *MockBackend) NewUser(p internship.User) ([]byte, error) {
	return []byte{}, nil
}
func (m *MockBackend) RmUser(email string) error {
	return nil
}
func (m *MockBackend) User(email string) (internship.User, error) {
	return internship.User{}, nil
}
func (m *MockBackend) Users() ([]internship.User, error) {
	return []internship.User{
		internship.User{Email: "stu", Role: internship.STUDENT},
		internship.User{Email: "adm", Role: internship.ADMIN},
		internship.User{Email: "stu2", Role: internship.STUDENT},
	}, nil
}
func (m *MockBackend) SetUserPassword(email string, oldP, newP []byte) error {
	return nil
}
func (m *MockBackend) SetUserProfile(email, fn, ln, tel string) error {
	return nil
}
func (m *MockBackend) SetUserRole(email string, priv internship.Privilege) error {
	return nil
}
func (m *MockBackend) ResetPassword(email string) (string, error) {
	return "", nil
}
func (m *MockBackend) NewPassword(token, newP []byte) (string, error) {
	return "", nil
}
func (m *MockBackend) PlanReport(student string, r internship.ReportHeader) error {
	return nil
}
func (m *MockBackend) Report(kind, email string) (internship.ReportHeader, error) {
	return internship.ReportHeader{}, nil
}
func (m *MockBackend) ReportContent(kind, email string) ([]byte, error) {
	return []byte{}, nil
}
func (m *MockBackend) SetReportContent(kind, email string, cnt []byte) error {
	return nil
}
func (m *MockBackend) SetReportGrade(kind, email string, r int) error {
	return nil
}
func (m *MockBackend) SetReportDeadline(kind, email string, t time.Time) error {
	return nil
}
func (m *MockBackend) ResetRootAccount() error {
	return nil
}
