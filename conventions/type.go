package conventions

import (
	"errors"
	"time"
)

var (
	ErrExists  = errors.New("Convention already exists")
	ErrUnknown = errors.New("Unknown convention")
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

type Convention struct {
	Stu   string
	Sup   string
	Tut   Tutor
	Cpy   Company
	Begin time.Time
	End   time.Time
	Title string
}

func (m *MetaData) isGraded() bool {
	return m.Grade >= 0
}

type ConventionService interface {
	New(stu string, sup string, tut Tutor, from, to time.Time, title string) (Convention, error)
	Get(stu string) (Convention, error)
	SetTutor(stu email, t Tutor) error
	SetCompany(stu email, c Company) error
	List() ([]Convention, error)
}
