package sqlstore

import (
	"database/sql"

	"github.com/fhermeni/wints/schema"
)

var (
	updateDefensePrivacy  = "update defenses set private=$1 where student=$2"
	updateDefenseLocality = "update defenses set local=$1 where student=$2"
	selectDefense         = "select date, room, grade, private, local from defenses where student=$1"
	selectDefenses        = "select student, date, room, grade, private, local from defenses"
	updateDefenseGrade    = "update defenses set grade=$2 where student=$1"
)

//SetDefensePrivacy indicates if the defense is private or public
func (s *Store) SetDefensePrivacy(student string, private bool) error {
	return s.singleUpdate(updateDefensePrivacy, schema.ErrUnknownStudent, private, student)
}

//SetDefenseLocality indicates if the defense is done in live or by visio
func (s *Store) SetDefenseLocality(student string, local bool) error {
	return s.singleUpdate(updateDefenseLocality, schema.ErrUnknownStudent, local, student)
}

//SetDefenseGrade set the defense grade
func (s *Store) SetDefenseGrade(student string, g int) error {
	if g < 0 || g > 20 {
		return schema.ErrInvalidGrade
	}
	return s.singleUpdate(updateDefenseGrade, schema.ErrUnknownDefense, student, g)
}

//Defense get the defense of a student
func (s *Store) Defense(student string) (schema.Defense, error) {
	d := schema.Defense{}
	st := s.stmt(selectDefense)
	err := st.QueryRow(student).Scan(
		&d.Time,
		&d.Room,
		&d.Grade,
		&d.Private,
		&d.Local,
	)
	if err != sql.ErrNoRows {
		return d, err
	}
	return d, nil
}

func (s *Store) defenses() (map[string]schema.Defense, error) {
	defs := make(map[string]schema.Defense)
	st := s.stmt(selectDefenses)
	rows, err := st.Query()
	if err != nil {
		return defs, err
	}
	defer rows.Close()
	for rows.Next() {
		d := schema.Defense{}
		var s string
		err = rows.Scan(
			&s,
			&d.Time,
			&d.Room,
			&d.Grade,
			&d.Private,
			&d.Local,
		)
		if err != nil {
			return defs, err
		}
		defs[s] = d
	}
	return defs, nil
}

//DefenseSessions get all the defense sessions
func (s *Store) DefenseSessions() ([]schema.DefenseSession, error) {
	var ss []schema.DefenseSession
	return ss, nil
}
