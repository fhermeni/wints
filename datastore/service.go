package datastore

import (
	"database/sql"

	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/mail"

	"code.google.com/p/go.crypto/bcrypt"
)

//Service allows to communicate with a database
type Service struct {
	DB         *sql.DB
	reportDefs []internship.ReportDef
	surveyDefs []internship.SurveyDef
	majors     []string
	j          journal.Journal
	mailer     mail.Mailer
}

//NewService initiate the storage servive
func NewService(db *sql.DB, reportDefs []internship.ReportDef, surveyDefs []internship.SurveyDef, majors []string, j journal.Journal, m mail.Mailer) (*Service, error) {
	s := Service{DB: db, reportDefs: reportDefs, surveyDefs: surveyDefs, majors: majors, j: j, mailer: m}
	return &s, nil
}

func hash(buf []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(buf, bcrypt.MinCost)
}

//Install the tables on the database
func (s *Service) Install() error {
	_, err := s.DB.Exec(create)
	return err
}
