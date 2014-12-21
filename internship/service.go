//Define the internship service
package internship

import "time"

type Service interface {

	//Internship management

	//New internship
	NewInternship(student, tutor string, from, to time.Time, c Company, sup Supervisor, Title string) error

	//Get the internship associated to a given student
	Internship(stu string) (Internship, error)

	//Set the supervisor for the internship of a given student
	SetSupervisor(stu string, sup Supervisor) error

	//Set the company for the internship of a given student
	SetCompany(stu string, c Company) error

	//Get all the internships
	Internships() ([]Internship, error)

	//User management

	//Test if the credentials match a user
	Registered(email string, password []byte) (User, error)

	//Create a new user
	//Returns the resulting password
	NewUser(p User) ([]byte, error)
	RmUser(email string) error
	User(email string) (User, error)
	Users() ([]User, error)
	SetUserPassword(email string, oldP, newP []byte) error
	SetUserProfile(email, fn, ln, tel string) error
	SetUserRole(email, priv string) error
	NewPassword(token, newP []byte) (string, error)
	ResetPassword(email string) (string, error)

	//Reports management

	PlanReport(kind, email string, date time.Time) (ReportHeader, error)
	Report(kind, email string) (ReportHeader, error)
	ReportContent(kind, email string) ([]byte, error)
	SetReportContent(kind, email string, cnt []byte) error
	SetReportGrade(kind, email string, r int) error
	SetReportDeadline(kind, email string, t time.Time) error

	ResetRootAccount() error
}
