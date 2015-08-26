//Package internship defines the interface specifying common methods to manipulate student internships
package internship

import (
	"errors"
	"time"

	"github.com/fhermeni/wints/config"
)

//Service specifies the core methods to manage student internships
type Service interface {
	RmUser(email string) error

	Users() ([]User, error)

	User(email string) (User, error)

	Visit(u string) error
	SetPassword(email string, oldP, newP []byte) error
	SetUserPerson(p Person) error

	SetUserRole(email string, priv Privilege) error

	NewSession(email string, password []byte, expire time.Duration) (Session, error)
	RmSession(token []byte) error
	Session(token string) (Session, error)

	ResetPassword(email string, d time.Duration) ([]byte, error)

	NewPassword(token, newP []byte) (string, error)

	ReplaceUserWith(src, dst string) error

	SetStudentSkippable(em string, st bool) error
	Students() ([]Student, error)
	Student(email string) (Student, error)
	SetNextContact(student string, em string) error
	SetNextPosition(student string, pos int) error
	SetMajor(stu string, m string) error
	SetPromotion(stu string, p string) error
	SetMale(stu string, male bool) error

	Convention(student string) (Convention, error)
	Conventions() ([]Convention, error)
	SetSupervisor(stu string, sup Person) error
	SetTutor(stu string, t string) error
	SetCompany(stu string, c Company) error
	SetTitle(stu string, title string) error
	SetConventionSkippable(student string, skip bool) error

	ValidateConvention(stu string, cfg config.Config) error

	Internships() ([]Internship, error)
	Internship(stu string) (Internship, error)

	Reports(email string) (map[string]ReportHeader, error)
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
	//SetDefenseSessions(defs []DefenseSession) error
	Defense(student string) (Defense, error)
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
	NewUser(p Person) error

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
	ErrUnknownReport      = errors.New("Unknown report")
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
