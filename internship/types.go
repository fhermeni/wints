package internship

import (
	"errors"
	"strings"
	"time"
)

//Privilege aims at giving the possible level of privilege for a user.
type Privilege int

//The different level of privileges
const (
	NONE Privilege = iota
	TUTOR
	MAJOR
	ADMIN
	ROOT
	//DateLayout indicates the date format
	DateLayout = "02/01/2006 15:04"
)

//User denotes a person that can authenticate.
//Students, major administrator, administrator or super user can authenticate
type User struct {
	//User firstname
	Firstname string
	//User lastname
	Lastname string
	//User email
	Email string
	//User telephon
	Tel string
	//User role
	Role Privilege
}

//Privileges denotes the string value associated to each level of privilege
var Privileges = [...]string{"none", "tutor", "major", "admin", "root"}

func (p Privilege) String() string {
	return Privileges[p]
}

//Returns the user fullname plus the email in parenthesis
func (p User) String() string {
	return p.Fullname() + " (" + p.Email + ")"
}

//Fullname provides the user fullname, starting with its firstname
func (p User) Fullname() string {
	return strings.Title(p.Firstname) + " " + strings.Title(p.Lastname)
}

//ReportHeader provides the metadata associated to a student report
type ReportHeader struct {
	//The report identifier
	Kind string
	//The report deadline
	Deadline time.Time
	//The grade, between 0 and 20 if graded. <-2 if their is no report, -1 if it is waiting for a review. >=0 once graded
	Grade int
	//The tutor comment
	Comment string
	Private bool
	ToGrade bool
}

//ReportsDef allows to define internship reports
type ReportDef struct {
	Name     string
	Deadline string
	Value    string
	ToGrade  bool
}

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) Graded() bool {
	return m.Grade >= 0
}

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) In() bool {
	return m.Grade != -2
}

//ReportHeader allows to instantiate a report definition to a header
func (def *ReportDef) Instantiate(from time.Time) (ReportHeader, error) {
	switch def.Deadline {
	case "relative":
		shift, err := time.ParseDuration(def.Value)
		if err != nil {
			return ReportHeader{}, err
		}
		return ReportHeader{Kind: def.Name, Deadline: from.Add(shift), ToGrade: def.ToGrade}, nil
	case "absolute":
		at, err := time.Parse(DateLayout, def.Value)
		if err != nil {
			return ReportHeader{}, err
		}
		return ReportHeader{Kind: def.Name, Deadline: at, ToGrade: def.ToGrade}, nil
	}
	return ReportHeader{}, errors.New("Unsupported time definition '" + def.Deadline + "'")
}

//Company is just a composite to store meaningful information for a company hosting a student.
type Company struct {
	//The company name
	Name string
	//The company website
	WWW string
}

//Person just gather informations to contact a company supervisor
type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
}

//Fullname provides the user fullname, starting with its firstname
func (p Person) Fullname() string {
	return strings.Title(p.Firstname) + " " + strings.Title(p.Lastname)
}

//Internship is the core type to specify required data related to an internship
type Internship struct {
	//The student
	Student User
	//The academic tutor
	Tutor User
	//the company supervisor
	Sup Person
	//The company
	Cpy Company
	//First day of the internship
	Begin time.Time
	//Last day of the internship
	End time.Time
	//Internship subject
	Title string
	//The headers for each
	Reports []ReportHeader
	//The student promotion
	Promotion string
	//The student major
	Major string
	//Student future
	Future Alumni
}

type Convention struct {
	Student    Person
	Tutor      Person
	Supervisor Person
	Cpy        Company
	Begin      time.Time
	End        time.Time
	Title      string
	Promotion  string
	//Ignore this convention
	Skip bool
}

type Alumni struct {
	Contact  string
	Position int
}

//Predefined errors
var (
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
)
