package schema

import (
	"database/sql/driver"
	"strings"
	"time"
)

//The different level of Roles
const (
	STUDENT_LEVEL int = iota
	TUTOR_LEVEL
	MAJOR_LEVEL
	HEAD_LEVEL
	ADMIN_LEVEL
	ROOT_LEVEL

	STUDENT Role = "student"
	TUTOR   Role = "tutor"
	ROOT    Role = "root"
	//DateLayout indicates the date format
	DateLayout = "02/01/2006 15:04"
)

//Role aims at giving the possible level of Role for a user.
type Role string

func (p Role) Level() int {
	if p == "student" {
		return STUDENT_LEVEL
	} else if p == "tutor" {
		return TUTOR_LEVEL
	} else if strings.Index(string(p), "major") == 0 {
		return MAJOR_LEVEL
	} else if p == "head" {
		return HEAD_LEVEL
	} else if p == "admin" {
		return ADMIN_LEVEL
	} else if p == "root" {
		return ROOT_LEVEL
	}
	return -1
}

func (p Role) SubRole() string {
	from := strings.LastIndex(string(p), "-") + 1
	return string(p)[from:len(p)]
}
func (p Role) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p Role) String() string {
	return string(p)
}

//Person just gathers contact information for someone.
type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
}

//Session denotes a user session
type Session struct {
	Email  string
	Token  []byte
	Expire time.Time
}

//User is a person with an account
type User struct {
	Person    Person
	Role      Role
	LastVisit *time.Time `,json:"omitempty"`
}

//Alumni denotes the basic information for a student future
type Alumni struct {
	Contact     string
	Position    string
	France      bool
	Permanent   bool
	SameCompany bool
}

//Student denotes a student that as a promotion, a major.
type Student struct {
	User      User
	Promotion string
	Major     string
	Alumni    *Alumni `,json:"omitempty"`
	//Skip indicates we don't care about this student. Might left the school for example
	Skip bool
	Male bool
}

//Fullname provides the user fullname, starting with its firstname
func (p Person) Fullname() string {
	return strings.Title(p.Firstname) + " " + strings.Title(p.Lastname)
}

//String returns the person email
func (p Person) String() string {
	return p.Email
}

//String delegates to Person.String()
func (u User) String() string {
	return u.Person.String()
}

//Fullname delegates to Person.String()
func (u User) Fullname() string {
	return u.Person.Fullname()
}
