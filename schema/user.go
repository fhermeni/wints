package schema

import (
	"database/sql/driver"
	"strings"
	"time"
)

//The different level of Roles
const (
	StudentLevel int = iota
	TutorLevel
	MajorLevel
	HeadLevel
	AdminLevel
	RootLevel

	STUDENT Role = "student"
	TUTOR   Role = "tutor"
	ROOT    Role = "root"
	//DateLayout indicates the date format
	DateLayout = "02/01/2006 15:04"
)

//Role aims at giving the possible level of Role for a user.
type Role string

//Level returns the authentication level associated to a given role
func (p Role) Level() int {
	if p == "student" {
		return StudentLevel
	} else if p == "tutor" {
		return TutorLevel
	} else if strings.Index(string(p), "major") == 0 {
		return MajorLevel
	} else if p == "head" {
		return HeadLevel
	} else if p == "admin" {
		return AdminLevel
	} else if p == "root" {
		return RootLevel
	}
	return -1
}

//SubRole extract the second level role if exists.
func (p Role) SubRole() string {
	from := strings.LastIndex(string(p), "-") + 1
	return string(p)[from:len(p)]
}

//Value returns String()
func (p Role) Value() (driver.Value, error) {
	return p.String(), nil
}

//String() versions of the role
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

//Anonymise the person by removing any personal information
func (p *Person) Anonymise() {
	p.Firstname = ""
	p.Lastname = ""
	p.Email = ""
	p.Tel = ""
}

//Students aliases slices of student to exhibit filtering methods
type Students []Student

//Filter returns the students that pass the given filter
func (ss Students) Filter(filter func(Student) bool) Students {
	res := make(Students, 0)
	for _, s := range ss {
		if filter(s) {
			res = append(res, s)
		}
	}
	return res
}

//StudentInMajor is a filter that keep only the students in the given major
func StudentInMajor(major string) func(Student) bool {
	return func(s Student) bool {
		return s.Major == major
	}
}
