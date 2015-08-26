package sqlstore

import "github.com/fhermeni/wints/internship"

var (
	updateDefensePrivacy  = "update defenses set private=$1 where student=$2"
	updateDefenseLocality = "update defenses set local=$1 where student=$2"
	selectDefense         = "select date, room, grade, private, local from defenses where student=$1"
	updateDefenseGrade    = "update defenses set grade=$2 where student=$1"
)

//SetDefensePrivacy indicates if the defense is private or public
func (s *Store) SetDefensePrivacy(student string, private bool) error {
	return s.singleUpdate(updateDefensePrivacy, internship.ErrUnknownStudent, private, student)
}

//SetDefenseLocality indicates if the defense is done in live or by visio
func (s *Store) SetDefenseLocality(student string, local bool) error {
	return s.singleUpdate(updateDefenseLocality, internship.ErrUnknownStudent, local, student)
}

//SetDefenseGrade set the defense grade
func (s *Store) SetDefenseGrade(student string, g int) error {
	return s.singleUpdate(updateDefenseGrade, internship.ErrUnknownDefense, student, g)
}

//Defense get the defense of a student
func (s *Store) Defense(student string) (internship.Defense, error) {
	d := internship.Defense{}
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
func (s *Store) DefenseSessions() ([]internship.DefenseSession, error) {
	ss := make([]internship.DefenseSession, 0, 0)
	return ss, nil
}
