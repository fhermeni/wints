package sqlstore

import "github.com/fhermeni/wints/schema"

var (
	updateDefensePrivacy  = "update defenses set private=$1 where student=$2"
	updateDefenseLocality = "update defenses set local=$1 where student=$2"
	selectDefense         = "select date, room, grade, private, local from defenses where student=$1"
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
	return d, err
}

//DefenseSessions get all the defense sessions
func (s *Store) DefenseSessions() ([]schema.DefenseSession, error) {
	ss := make([]schema.DefenseSession, 0, 0)
	return ss, nil
}
