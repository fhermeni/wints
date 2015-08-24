//Package internship defines the interface specifying common methods to manipulate student internships
package internship

import (
	"errors"
	"time"

	"github.com/fhermeni/wints/config"
)

//Service specifies the core methods to manage student internships
type Service interface {

	//Delete the user if it is not tutoring anyone
	RmUser(email string) error

	//List the users
	Users() ([]User, error)

	Visit(u string) error
	//Change the user password
	SetPassword(email string, oldP, newP []byte) error
	//Change user profile
	SetUserProfile(email, fn, ln, tel string) error

	//Change user role
	SetUserRole(email string, priv Privilege) error

	//Test if the credentials match a user, return a session token
	Login(email string, password []byte) ([]byte, error)
	//Destroy a session
	Logout(email, token string) error

	//Check if a session is opened for a given user and token
	OpenedSession(email, token string) error

	//Ask for a password reset.
	//Return a token used to declare the new password (see NewPassword)
	ResetPassword(email string, d time.Duration) ([]byte, error)

	//Declare the new password using an authentication token
	NewPassword(token, newP []byte) (string, error)

	ReplaceUserWith(src, dst string) error

	//Student that left ?
	SkipStudent(em string, st bool) error
	Students() ([]Student, error)
	SetNextContact(student string, em string) error
	SetNextPosition(student string, pos int) error
	SetMajor(stu string, m string) error
	SetPromotion(stu string, p string) error
	//SetMale states a student gender
	SetMale(stu string, male bool) error
	Conventions() ([]Convention, error)
	SetSupervisor(stu string, sup User) error
	SetTutor(stu string, t string) error
	SetCompany(stu string, c Company) error
	SetTitle(stu string, title string) error

	//I don't manage it
	SetConventionSkippable(student string, skip bool) error

	//Turns the convention to an internship
	//align student & tutor. Tutor must exists. Send student credentials, instantiate reports, surveys
	ValidateConvention(stu string, cfg config.Config) error

	//Internship management
	Internships() ([]Internship, error)
	Internship(stu string) (Internship, error)

	//Reports management

	//PlanReport(student string, r ReportHeader) error
	//ReportDefs() []ReportDef
	Report(kind, email string) (ReportHeader, error)
	ReportContent(kind, email string) ([]byte, error)
	SetReportContent(kind, email string, cnt []byte) error
	SetReportGrade(kind, email string, r int, comment string) error
	SetReportDeadline(kind, email string, t time.Time) error
	SetReportPrivacy(kind, email string, p bool) error

	//Survey management
	SurveyToken(kind string) (string, string, error)
	Survey(student, kind string) (SurveyHeader, error)
	SetSurveyContent(token string, cnt map[string]string) error
	//SurveyDefs() []SurveyDef
	//RequestSurvey(stu, kind string) error

	//DefenseSessions returns all the registered defense sessions
	DefenseSessions() ([]DefenseSession, error)
	//SetDefenseSessions saves all the defense sessions
	SetDefenseSessions(defs []DefenseSession) error
	DefenseSession(student string) (DefenseSession, error)
	//SetDefenseGrade Set the grade for a given defense
	SetDefenseGrade(student string, g int) error
	SetDefensePrivacy(student string, private bool) error
	SetDefenseLocality(student string, local bool) error
}

//Adder abstracts the definition of new users or conventions.
//Things that are usually imported from external sourfces
type Adder interface {

	//NewUser creates a new user
	//By default, the user has no privileges
	NewUser(fn, ln, tel, email string) error

	//ToStudent casts the given user to a student
	ToStudent(email, major, promotion string, male bool) error

	//NewConvention adds a convention
	//The student and the tutor must have an account despite it is not required to be a valid one
	NewConvention(student string, startTime, endTime time.Time, tutor string, cpy Company, sup Person, title string, creation time.Time, foreignCountry, lab bool, gratification int) (bool, error)

	//SetMale states a student gender
	SetMale(stu string, male bool) error
}

var (
	ErrUnknownStudent     = errors.New("Unknown student")
	ErrUnknownConvention  = errors.New("No convention associated to this student")
	ErrStudentExists      = errors.New("Student already exists")
	ErrReportExists       = errors.New("Report already exists")
	ErrUnknownReport      = errors.New("Unknown report or user")
	ErrInvalidGrade       = errors.New("The grade must be between 0 and 20 (inclusive)")
	ErrReportConflict     = errors.New("The report has not been uploaded")
	ErrInternshipExists   = errors.New("Internship already exists")
	ErrUnknownInternship  = errors.New("Unknown internship")
	ErrUserExists         = errors.New("User already exists")
	ErrUnknownUser        = errors.New("User not found")
	ErrUserTutoring       = errors.New("The user is tutoring students")
	ErrCredentials        = errors.New("Incorrect credentials")
	ErrNoPendingRequests  = errors.New("No password renewable request pending")
	ErrInvalidPeriod      = errors.New("invalid internship period")
	ErrConventionExists   = errors.New("convention already scanned")
	ErrInvalidMajor       = errors.New("Invalid major")
	ErrDeadlinePassed     = errors.New("Deadline passed")
	ErrGradedReport       = errors.New("Report already graded")
	ErrSessionExpired     = errors.New("Session expired")
	ErrInvalidToken       = errors.New("Invalid session")
	ErrUnknownSurvey      = errors.New("Unknown survey")
	ErrInvalidSurvey      = errors.New("Invalid answers")
	ErrUnknownAlumni      = errors.New("No informations for future alumni")
	ErrInvalidAlumniEmail = errors.New("Invalid email. It must not be served by polytech' or unice")
	ErrUnknownDefense     = errors.New("Unknown defense")
)
