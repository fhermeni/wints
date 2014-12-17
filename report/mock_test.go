package report

import (
	"time"
	"log"
)

type MockService struct {
	Reports  []MetaData
	Contents map[string][]byte
}

func NewMock() *MockService {
	return &MockService{make([]MetaData, 0, 0), make(map[string][]byte)}
}

func (s *MockService) New(kind, email string, date time.Time) (MetaData, error) {
	///Check existance
	for _, m := range s.Reports {
		if m.Kind == kind && m.Email == email {
			return MetaData{}, ErrExists
		}
	}
	m := MetaData{kind, email, date, -1}
	s.Reports = append(s.Reports, m)
	return m, nil
}

func (s *MockService) Get(kind, email string) (MetaData, error) {
	for _, m := range s.Reports {
		if m.Kind == kind && m.Email == email {
			return m, nil
		}
	}
	return MetaData{}, ErrUnknown
}

func (s *MockService) Content(kind, email string) ([]byte, error) {
	for k, cnt := range s.Contents {
		if kind+"|"+email == k {
			return cnt, nil
		}
	}
	return []byte{}, ErrUnknown
}

func (s *MockService) SetContent(kind, email string, cnt []byte) error {
	for _, m := range s.Reports {
		if m.Kind == kind && m.Email == email {
			s.Contents[kind+"|"+email] = cnt
			return nil
		}
	}
	return ErrUnknown
}

func (s *MockService) SetGrade(kind, email string, g int) error {
	if g < 0 || g > 20 {
		log.Println("oops")
		return ErrInvalidGrade
	}
	for i, m := range s.Reports {
		if m.Kind == kind && m.Email == email {
			if _, ok := s.Contents[kind+"|"+email]; !ok {
				log.Println("not uploaded")
				return ErrConflict
			}
			m.Grade = g
			s.Reports[i] = m
			return nil
		}
	}
	return ErrUnknown
}

func (s *MockService) SetDeadline(kind, email string, t time.Time) error {
	for i, m := range s.Reports {
		if m.Kind == kind && m.Email == email {
			m.Deadline = t
			s.Reports[i] = m
			return nil
		}
	}
	return ErrUnknown
}

func (s *MockService) List(email string) ([]MetaData, error) {
	reports := make([]MetaData, 0, 0)
	for _, m := range s.Reports {
		if m.Email == email {
			reports = append(reports, m)
		}
	}
	if len(reports) == 0 {
		return reports, ErrUnknown
	}
	return reports, nil
}
