package internship

import (
	"errors"
	"time"
)

var (
	ErrExists  = errors.New("Internship already exists")
	ErrUnknown = errors.New("Unknown internship")
)

type Company struct {
	Name string
	WWW  string
}

type Tutor struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
}

type Internship struct {
	Stu   string
	Sup   string
	Tut   Tutor
	Cpy   Company
	Begin time.Time
	End   time.Time
	Title string
}

type InternshipService interface {
	//	New(stu string, sup string, tut Tutor, from, to time.Time, title string) (Internship, error)
	Get(stu string) (Internship, error)
	SetTutor(stu string, t Tutor) error
	SetCompany(stu string, c Company) error
	List() ([]Internship, error)
}
