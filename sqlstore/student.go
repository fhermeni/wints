package sqlstore

import (
	"database/sql"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	allStudents     = "select male, firstname, lastname, users.email, tel, role, lastVisit, promotion, major, nextPosition, nextContact, skip from students inner join users on (students.email=users.email)"
	selectStudent   = "select male, firstname, lastname, users.email, tel, role, lastVisit, promotion, major, nextPosition, nextContact, skip from students inner join users on (students.email=users.email) where students.email=$1"
	insertStudent   = "insert into students(email, male, major, promotion, skip) values ($1,$2,$3,$4,$5)"
	skipStudent     = "update students set skip=$2 where email=$1"
	setPromotion    = "update students set promotion=$2 where email=$1"
	setMajor        = "update students set major=$2 where email=$1"
	updateMale      = "update students set male=$2 where email=$1"
	setNextPosition = "update students set nextPosition=$1 where email=$2"
	setNextContact  = "update students set nextContact=$1 where email=$2"
)

//Student returns a given student
func (s *Store) Student(email string) (schema.Student, error) {
	st := s.stmt(selectStudent)
	rows, err := st.Query(email)
	if err != nil {
		return schema.Student{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return schema.Student{}, schema.ErrUnknownStudent
	}
	return scanStudent(rows)
}

func scanStudent(rows *sql.Rows) (schema.Student, error) {
	var lastVisit pq.NullTime
	var nextPos sql.NullInt64
	var nextContact sql.NullString
	s := schema.Student{
		User: schema.User{
			Person: schema.Person{},
		},
		Alumni: schema.Alumni{},
	}
	err := rows.Scan(
		&s.Male,
		&s.User.Person.Firstname,
		&s.User.Person.Lastname,
		&s.User.Person.Email,
		&s.User.Person.Tel,
		&s.User.Role,
		&lastVisit,
		&s.Promotion,
		&s.Major,
		&nextPos,
		&nextContact,
		&s.Skip)
	s.User.LastVisit = nullableTime(lastVisit)
	s.Alumni.Contact = nullableString(nextContact)
	s.Alumni.Position = nullableInt(nextPos, 0)
	return s, err
}

//Students list all the registered students
func (s *Store) Students() ([]schema.Student, error) {
	rows, err := s.db.Query(allStudents)

	students := make([]schema.Student, 0, 0)
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		s, err := scanStudent(rows)
		if err != nil {
			return students, err
		}
		students = append(students, s)
	}
	return students, err
}

func (s *Store) NewStudent(p schema.Person, major, promotion string, male bool) error {
	if !s.config.ValidMajor(major) {
		return schema.ErrInvalidMajor
	}
	if !s.config.ValidPromotion(promotion) {
		return schema.ErrInvalidPromotion
	}
	tx := newTxErr(s.db)
	u := schema.User{
		Person: p,
		Role:   schema.STUDENT,
	}
	s.addUser(&tx, u)
	tx.Update(insertStudent, p.Email, male, major, promotion, false)
	return tx.Done()
}

//SetStudentSkippable indicates if it is not required for the student to get an internship (an abandom typically)
func (s *Store) SetStudentSkippable(em string, st bool) error {
	return s.singleUpdate(skipStudent, schema.ErrUnknownUser, em, st)
}

//SetPromotion updates the student promotion
func (s *Store) SetPromotion(stu, p string) error {
	if !s.config.ValidPromotion(p) {
		return schema.ErrInvalidPromotion
	}
	return s.singleUpdate(setPromotion, schema.ErrUnknownStudent, stu, p)
}

//SetMajor updates the student major
func (s *Store) SetMajor(stu, m string) error {
	if !s.config.ValidMajor(m) {
		return schema.ErrInvalidMajor
	}
	return s.singleUpdate(setMajor, schema.ErrUnknownStudent, stu, m)
}

//SetMale stores the student gender
func (s *Store) SetMale(stu string, m bool) error {
	return s.singleUpdate(updateMale, schema.ErrUnknownStudent, stu, m)
}

//SetNextPosition updates the student next position
func (s *Store) SetNextPosition(student string, pos int) error {
	return s.singleUpdate(setNextPosition, schema.ErrUnknownUser, pos, student)
}

//SetNextContact updates the student next contact email.
func (s *Store) SetNextContact(student string, em string) error {
	return s.singleUpdate(setNextContact, schema.ErrUnknownUser, em, student)
}
