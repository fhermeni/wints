package report;

import (
	"time"	
	"errors"	
)

var (
	ErrExists = errors.New("Report already exists")
	ErrUnknown = errors.New("Unknown report")
	ErrInvalidGrade = errors.New("The grade must be between 0 and 20 (inclusive)")
)


type MetaData struct {
	Kind string
	Email string
	Deadline time.Time
	Grade int	
}

func (m *MetaData) isGraded() bool {
	return m.Grade >= 0
}

type ReportService interface {
	New(kind, email string, date time.Time) (MetaData, error)
	Get(kind, email string) (MetaData, error)
	Content(kind, email string) ([]byte, error)
	SetContent(kind, email string, cnt []byte) error
	SetGrade(kind, email string, r int) error
	SetDeadline(kind, email string, t time.Time) error	
}