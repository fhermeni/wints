package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	insertStudentDefense          = "insert into defenses(id, room, date, student, public, local) values ($1,$2,$3,$4,$5,$6)"
	updateStudentDefense          = "update defenses set date=$2, public=$3, local=$4 where student=$1"
	deleteStudentDefense          = "delete from defenses where student=$1"
	selectDefense                 = "select id, grade, public, local, date,room from defenses where student=$1 order by date"
	selectDefenses                = "select student, date, room, grade, public, local from defenses order by date"
	selectStudentInDefenseSession = "select student, grade, public, local, date from defenses where room=$1 and id=$2 order by date"
	updateDefenseGrade            = "update defenses set grade=$2 where student=$1"
	newDefenseSession             = "insert into defensesessions(room, id) values($1, $2)"
	defenseSessions               = "select room, id from defensesessions order by id"
	defenseSession                = "select room, id from defensesessions where room=$1 and id=$2"
	deleteDefenseSession          = "delete from defensesessions where room=$1 and id=$2"
	insertJuryToDefense           = "insert into defensejuries(room, id, jury) values($1,$2,$3)"
	deleteJuryToDefense           = "delete from defensejuries where room=$1 and id=$2 and jury=$3"
	selectJury                    = "select us.firstname, us.lastname, us.tel, us.email, us.role, us.lastVisit from defensejuries inner join users as us on (us.email = jury) where room=$1 and id=$2"
)

//SetDefenseGrade set the defense grade
func (s *Store) SetDefenseGrade(student string, g int) error {
	if g < 0 || g > 20 {
		return schema.ErrInvalidGrade
	}
	return s.singleUpdate(updateDefenseGrade, schema.ErrUnknownDefense, student, g)
}

func (s *Store) SetStudentDefense(session, room, student string, t time.Time, public, local bool) error {
	return s.singleUpdate(insertStudentDefense, schema.ErrUnknownStudent, session, room, t.Truncate(time.Minute).UTC(), student, public, local)
}

func (s *Store) UpdateStudentDefense(student string, t time.Time, public, local bool) error {
	return s.singleUpdate(updateStudentDefense, schema.ErrUnknownStudent, student, t.Truncate(time.Minute).UTC(), public, local)
}

func (s *Store) RmStudentDefense(student string) error {
	return s.singleUpdate(deleteStudentDefense, schema.ErrUnknownStudent, student)
}

//Defense get the defense of a student
func (s *Store) Defense(student string) (schema.Defense, error) {
	d := schema.Defense{}
	st := s.stmt(selectDefense)
	var grade sql.NullInt64
	err := st.QueryRow(student).Scan(
		&d.SessionId,
		&grade,
		&d.Public,
		&d.Local,
		&d.Time,
		&d.Room,
	)
	if err == sql.ErrNoRows {
		return d, schema.ErrUnknownDefense
	}
	if err != nil {
		return d, err
	}
	d.Grade = nullableInt(grade, -1)

	ints, err := s.Internship(student)
	if err != nil {
		return d, err
	}
	d.Student = ints.Convention.Student
	return d, err
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
		var grade sql.NullInt64
		err = rows.Scan(
			&s,
			&d.Time,
			&d.SessionId,
			&grade,
			&d.Public,
			&d.Local,
		)
		d.Grade = nullableInt(grade, -1)
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
	st := s.stmt(defenseSessions)
	rows, err := st.Query()
	if err != nil {
		return ss, err
	}
	defer rows.Close()
	for rows.Next() {
		var room string
		var id string
		err := rows.Scan(&room, &id)
		if err != nil {
			return ss, err
		}
		session, err := s.DefenseSession(room, id)

		if err != nil {
			return ss, err
		}
		ss = append(ss, session)
	}
	return ss, nil
}

func scanDefenseSession(row *sql.Rows) (schema.DefenseSession, error) {
	ss := schema.DefenseSession{
		Defenses: []schema.Defense{},
		Juries:   []schema.User{},
	}
	err := row.Scan(&ss.Room, &ss.Id)
	return ss, err
}
func (s *Store) DefenseSession(room, id string) (schema.DefenseSession, error) {
	st := s.stmt(defenseSession)
	rows, err := st.Query(room, id)
	if err != nil {
		return schema.DefenseSession{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return schema.DefenseSession{}, schema.ErrUnknownDefense
	}
	session, err := scanDefenseSession(rows)
	if err != nil {
		return session, err
	}
	//jury
	st = s.stmt(selectJury)
	rows2, err := st.Query(room, id)
	if err != nil {
		return session, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var last pq.NullTime
		u := schema.User{Person: schema.Person{}}
		var role string
		err := rows2.Scan(&u.Person.Firstname,
			&u.Person.Lastname,
			&u.Person.Tel,
			&u.Person.Email,
			&role,
			&last)
		u.Role = schema.Role(role)
		u.LastVisit = nullableTime(last)
		if err != nil {
			return session, err
		}
		session.Juries = append(session.Juries, u)
	}
	//students
	st = s.stmt(selectStudentInDefenseSession)
	rows3, err := st.Query(room, id)
	if err != nil {
		return session, err
	}
	defer rows3.Close()
	for rows3.Next() {

		var student string
		var grade sql.NullInt64
		def := schema.Defense{}
		err = rows3.Scan(&student, &grade, &def.Public, &def.Local, &def.Time)

		if err != nil {
			return session, err
		}
		ints, err := s.Internship(student)
		if err != nil {
			return session, err
		}
		if grade.Valid {
			def.Grade = int(grade.Int64)
		}
		def.Student = ints.Convention.Student
		def.Company = ints.Convention.Company
		session.Defenses = append(session.Defenses, def)
	}
	return session, err
}

func (s *Store) AddJuryToDefenseSession(room, id, jury string) error {
	return s.singleUpdate(insertJuryToDefense, schema.ErrUnknownDefense, room, id, jury)
}

func (s *Store) DelJuryToDefenseSession(room, id, jury string) error {
	return s.singleUpdate(deleteJuryToDefense, schema.ErrUnknownDefense, room, id, jury)
}

func (s *Store) NewDefenseSession(room string, id string) (schema.DefenseSession, error) {
	err := s.singleUpdate(newDefenseSession, schema.ErrDefenseSessionConflit, room, id)
	if err != nil {
		return schema.DefenseSession{}, err
	}
	return schema.DefenseSession{
		Room:     room,
		Id:       id,
		Juries:   []schema.User{},
		Defenses: []schema.Defense{},
	}, nil
}

func (s *Store) RmDefenseSession(room, id string) error {
	return s.singleUpdate(deleteDefenseSession, schema.ErrUnknownDefense, room, id)
}
