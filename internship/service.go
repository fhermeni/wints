//Package internship defines the interface specifying common methods to manipulate student internships
package internship

import "time"

//Service specifies the core methods to manage student internships
type Service interface {

	//New internship
	NewInternship(c Convention) ([]byte, error)

	//Get the internship associated to a given student
	Internship(stu string) (Internship, error)

	//Set the supervisor for the internship of a given student
	SetSupervisor(stu string, sup Person) error

	//Set the tutor for the internship of a given student
	SetTutor(stu string, t string) error

	//Set the company for the internship of a given student
	SetCompany(stu string, c Company) error

	//Set the internship title
	SetTitle(stu string, title string) error

	//Set the student major
	SetMajor(stu string, m string) error

	//Get the possible majors
	Majors() []string

	//Set the student promotion
	SetPromotion(stu string, p string) error

	//Get all the internships
	Internships() ([]Internship, error)

	//Test if the credentials match a user, return a session token
	Registered(email string, password []byte) ([]byte, error)

	//Check if a session is opened for a given user and token
	OpenedSession(email, token string) error

	//Destroy a session
	Logout(email, token string) error

	Sessions() (map[string]time.Time, error)

	//Create a new user
	//Returns the resulting password
	NewTutor(p User) ([]byte, error)

	//Delete the user if it is not tutoring anyone
	RmUser(email string) error

	//Get the user
	User(email string) (User, error)

	//List the users
	Users() ([]User, error)

	//Change the user password
	SetUserPassword(email string, oldP, newP []byte) error

	//Change user profile
	SetUserProfile(email, fn, ln, tel string) error

	//Change user role
	SetUserRole(email string, priv Privilege) error

	//Ask for a password reset.
	//Return a token used to declare the new password (see NewPassword)
	ResetPassword(email string) ([]byte, error)

	//Declare the new password using an authentication token
	NewPassword(token, newP []byte) (string, error)

	//Reports management

	PlanReport(student string, r ReportHeader) error
	ReportDefs() []ReportDef
	Report(kind, email string) (ReportHeader, error)
	ReportContent(kind, email string) ([]byte, error)
	SetReportContent(kind, email string, cnt []byte) error
	SetReportGrade(kind, email string, r int, comment string) error
	SetReportDeadline(kind, email string, t time.Time) error
	SetReportPrivate(kind, email string, p bool) error

	//Survey management
	SurveyToken(kind string) (string, string, error)
	Survey(student, kind string) (Survey, error)
	SetSurveyContent(token string, cnt map[string]string) error
	SurveyDefs() []SurveyDef

	NewConvention(c Convention) error
	Conventions() ([]Convention, error)
	SkipConvention(student string, skip bool) error
	DeleteConvention(student string) error
	SetAlumni(student string, a Alumni) error

	//public statistics
	Statistics() ([]Stat, error)

	//Students that have to make an internship
	Students() ([]Student, error)
	InsertStudents(csv string) error
	AddStudent(s Student) error
	AlignWithInternship(student string, internship string) error
	HideStudent(em string, st bool) error

	//DefenseSessions returns all the registered defense sessions
	DefenseSessions() ([]DefenseSession, error)

	//SetDefenseSessions saves all the defense sessions
	SetDefenseSessions(defs []DefenseSession) error

	PublicDefenseSessions() ([]PublicDefenseSession, error)
	DefenseSession(student string) (DefenseSession, error)
	//SetDefenseGrade Set the grade for a given defense
	SetDefenseGrade(student string, g int) error

	RequestSurvey(stu, kind string) error
}
