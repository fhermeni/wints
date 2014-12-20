//Define the internship service
package internship

import "time"

type Service interface {

	//Internship management

	Internship(stu string) (Internship, error)
	SetSupervisor(stu string, sup Supervisor) error
	SetCompany(stu string, c Company) error
	Internships() ([]Internship, error)

	//User management

	Login(email, password string) (User, error)
	NewUser(p User) error
	RmUser(email string) error
	User(email string) (User, error)
	Users() ([]User, error)
	SetUserPassword(email string, oldP, newP []byte) error
	SetUserProfile(email, fn, ln, tel string) error
	SetUserRole(email, priv string) error
	NewPassword(token string, newP []byte) (string, error)
	ResetPassword(email string) (string, error)

	//Reports management

	PlanReport(kind, email string, date time.Time) (ReportHeader, error)
	Report(kind, email string) (ReportHeader, error)
	ReportContent(kind, email string) ([]byte, error)
	SetReportContent(kind, email string, cnt []byte) error
	SetReportGrade(kind, email string, r int) error
	SetReportDeadline(kind, email string, t time.Time) error
}
