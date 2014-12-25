package internship

import (
	"errors"
	"time"
)

//Privilege aims at giving the possible level of privilege for a user.
type Privilege int

//The different level of privileges
const (
	STUDENT Privilege = iota
	MAJOR
	ADMIN
	ROOT
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
var Privileges = [...]string{"student", "major", "admin", "root"}

func (p Privilege) String() string {
	return Privileges[p]
}

//Returns the user fullname plus the email in parenthesis
func (p User) String() string {
	return p.Fullname() + " (" + p.Email + ")"
}

//Fullname provides the user fullname, starting with its firstname
func (p User) Fullname() string {
	return p.Firstname + " " + p.Lastname
}

//ReportHeader provides the metadata associated to a student report
type ReportHeader struct {
	//The report identifier
	Kind string
	//The report deadline
	Deadline time.Time
	//The grade, between 0 and 20 if graded. <0 if it is waiting for a grade
	Grade int
}

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) Graded() bool {
	return m.Grade >= 0
}

//Company is just a composite to store meaningful information for a company hosting a student.
type Company struct {
	//The company name
	Name string
	//The company website
	WWW string
}

//Supervisor just gather informations to contact a company supervisor
type Supervisor struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
}

//Internship is the core type to specify required data related to an internship
type Internship struct {
	//The student
	Student User
	//The academic tutor
	Tutor User
	//the company supervisor
	Sup Supervisor
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
}

//Predefined errors
var (
	ErrReportExists      = errors.New("Report already exists")
	ErrUnknownReport     = errors.New("Unknown report or user")
	ErrInvalidGrade      = errors.New("The grade must be between 0 and 20 (inclusive)")
	ErrReportConflict    = errors.New("The report has not been uploaded")
	ErrInternshipExists  = errors.New("Internship already exists")
	ErrUnknownInternship = errors.New("Unknown internship")
	ErrUserExists        = errors.New("User already exists")
	ErrUnknownUser       = errors.New("User not found")
	ErrUserTutoring      = errors.New("The user is tutoring students")
	ErrCredentials       = errors.New("Incorrect credentials")
	ErrNoPendingRequests = errors.New("No password renewable request pending")
	ErrInvalidPeriod     = errors.New("invalid internship period")
)
