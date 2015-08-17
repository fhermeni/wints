package internship

import "strings"

//The different level of privileges
const (
	NONE Privilege = iota
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
var Privileges = [...]string{"none", "tutor", "major", "admin", "root"}

func (p Privilege) String() string {
	if p < 0 {
		return "tutor"
	}
	return Privileges[p]
}

//User just gather informations to contact someone. He may have an account on the platform
type User struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
	Role      Privilege
}

//Alumni denotes the basic information for a student future
type Alumni struct {
	Contact  string
	Position int
}

type Student struct {
	U         User
	Promotion string
	Major     string
	Alumni    Alumni
	Skip      bool
}

//Returns the user fullname plus the email in parenthesis
func (p User) String() string {
	return p.Fullname() + " (" + p.Email + ")"
}

//Fullname provides the user fullname, starting with its firstname
func (p User) Fullname() string {
	return strings.Title(p.Firstname) + " " + strings.Title(p.Lastname)
}
