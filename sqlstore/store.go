package sqlstore

import (
	"database/sql"
	"sync"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"

	"crypto/rand"
)

//Store allows to communicate with a database
type Store struct {
	db        *sql.DB
	stmts     map[string]*stmtErr
	config    config.Internships
	lock      *sync.Mutex
	cacheLock *sync.Mutex
	cache     *schema.Internships
}

//NewStore initiate the storage servive
func NewStore(d *sql.DB, config config.Internships) (*Store, error) {
	s := Store{
		db:        d,
		stmts:     make(map[string]*stmtErr),
		config:    config,
		lock:      &sync.Mutex{},
		cacheLock: &sync.Mutex{},
		cache:     nil,
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
		return mapCstrToError(err)
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
		x := t.Time.Truncate(time.Minute).UTC()
		return &x
	}
	return nil
}

func nullableString(t sql.NullString) string {
	if t.Valid {
		return t.String
	}
	return ""
}

func nullableBool(t sql.NullBool, def bool) bool {
	if t.Valid {
		return t.Bool
	}
	return def
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
			return schema.ErrUserExists
		case "fk_sessions_email", "fk_student_email", "fk_password_renewal_email":
			return schema.ErrUnknownUser
		case "pk_students_email":
			return schema.ErrStudentExists
		case "pk_conventions_student":
			return schema.ErrConventionExists
		case "fk_conventions_student", "fk_reports_student", "fk_surveys_student", "fk_defenses_student":
			return schema.ErrUnknownStudent
		case "fk_conventions_tutor":
			return schema.ErrUserTutoring
		case "pk_defensejuries", "pk_jury":
			return schema.ErrDefenseJuryConflict
		case "pk_defenses_student":
			return schema.ErrDefenseExists
		case "pk_defensesessions":
			return schema.ErrDefenseConflict
		}
	}
	return err
}

//flush the internship cache
/*func (s *Store) flush() {
	s.cacheLock.Lock()
	defer s.cacheLock.Unlock()
	s.cache = nil
}*/

func (s *Store) cached() *schema.Internships {
	s.cacheLock.Lock()
	defer s.cacheLock.Unlock()
	return s.cache
}

/*func (s *Store) encache(i schema.Internships) {
	s.cacheLock.Lock()
	defer s.cacheLock.Unlock()
	s.cache = &i
}*/
