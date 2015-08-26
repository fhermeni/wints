package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"

	"crypto/rand"
)

//Store allows to communicate with a database
type Store struct {
	db *sql.DB
	//mailer     mail.Mailer
	stmts map[string]*stmtErr
}

//NewStore initiate the storage servive
func NewStore(d *sql.DB) (*Store, error) {
	s := Store{db: d,
		stmts: make(map[string]*stmtErr),
	}
	return &s, nil
}

//Install the tables on the database
func (s *Store) Install() error {
	_, err := s.db.Exec(create)
	return err
}

func (s *Store) stmt(q string) *stmtErr {
	st, ok := s.stmts[q]
	if !ok {
		x, err := s.db.Prepare(q)
		st = &stmtErr{err: err, st: x}
		s.stmts[q] = st
	}
	return st
}

//SingleUpdate executes a query that aims are affecting only 1 row.
func (s *Store) singleUpdate(q string, errNoUpdate error, args ...interface{}) error {
	st := s.stmt(q)
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

func randomBytes(s int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, s)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func nullableTime(t pq.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func nullableInt(t sql.NullInt64, def int) int {
	if t.Valid {
		return int(t.Int64)
	}
	return def
}

func noRowsTo(got error, with error) error {
	if got == sql.ErrNoRows {
		return with
	}
	return got
}

func mapCstrToError(err error) error {
	if err == nil {
		return err
	}
	if e, ok := err.(*pq.Error); ok {
		switch e.Constraint {
		case "pk_email":
			return internship.ErrUserExists
		case "fk_sessions_email", "fk_student_email", "fk_password_renewal_email":
			return internship.ErrUnknownUser
		case "pk_students_email":
			return internship.ErrStudentExists
		case "pk_conventions_student":
			return internship.ErrConventionExists
		case "fk_conventions_student", "fk_reports_student", "fk_surveys_student", "fk_defenses_student":
			return internship.ErrUnknownStudent
		case "fk_conventions_tutor":
			return internship.ErrUserTutoring
		}
	}
	return err
}
