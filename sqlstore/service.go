package sqlstore

import (
	"database/sql"

	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"

	"code.google.com/p/go.crypto/bcrypt"
)

//Service allows to communicate with a database
type Service struct {
	DB         *sql.DB
	reportDefs []internship.ReportDef
	surveyDefs []internship.SurveyDef
	majors     []string
	j          journal.Journal
	//mailer     mail.Mailer
	stmts map[string]*sql.Stmt
}

//NewService initiate the storage servive
func NewService(db *sql.DB, reportDefs []internship.ReportDef, surveyDefs []internship.SurveyDef, majors []string, j journal.Journal /*, m mail.Mailer*/) (*Service, error) {
	s := Service{DB: db, reportDefs: reportDefs, surveyDefs: surveyDefs, majors: majors, j: j /*, mailer: m, */, stmts: make(map[string]*sql.Stmt)}
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

func (s *Service) stmt(q string) (*sql.Stmt, error) {
	st, ok := s.stmts[q]
	if !ok {
		st, err := s.DB.Prepare(q) //Bad, loss prepare failure
		if err != nil {
			return nil, err
		}
		s.stmts[q] = st
	}
	return st, nil
}

//SingleUpdate executes a query that aims are affecting only 1 row.
func (s *Service) singleUpdate(q string, errNoUpdate error, args ...interface{}) error {
	st, err := s.stmt(q)
	if err != nil {
		return mapCstrToError(err)
	}
	res, err := st.Exec(args...)
	if err != nil {
		return mapCstrToError(err)
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errNoUpdate
	}
	return nil
}
