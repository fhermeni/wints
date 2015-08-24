package internship

import (
	"strings"
	"time"
)

//The different level of privileges
const (
	NONE Privilege = iota //Not validated. Cannot connect
	STUDENT
	TUTOR
	MAJOR
	ADMIN
	ROOT
	//DateLayout indicates the date format
	DateLayout = "02/01/2006 15:04"
)

//Privilege aims at giving the possible level of privilege for a user.
type Privilege int

//Privileges denotes the string value associated to each level of privilege
var Privileges = [...]string{"none", "student", "tutor", "major", "admin", "root"}

func (p Privilege) String() string {
	return Privileges[p]
}

//Person just gathers contact information for someone.
type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
}

//User is a person with an account
type User struct {
	Person    Person
	Role      Privilege
	LastVisit *time.Time `,json:"omitempty"`
}

//Alumni denotes the basic information for a student future
type Alumni struct {
	Contact  string
	Position int
}

//Student denotes a student that as a promotion, a major.
type Student struct {
	User      User
	Promotion string
	Major     string
	Alumni    Alumni
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
func (U User) Fullname() string {
	return U.Person.Fullname()
}
