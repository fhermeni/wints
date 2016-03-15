package sqlstore

import (
	"database/sql"
	"strings"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	allStudents   = "select male, firstname, lastname, users.email, tel, role, lastVisit, promotion, major, nextPosition, nextFrance,nextPermanent,nextSameCompany, nextContact, skip from students inner join users on (students.email=users.email)"
	selectStudent = "select male, firstname, lastname, users.email, tel, role, lastVisit, promotion, major, nextPosition, nextFrance,nextPermanent,nextSameCompany, nextContact, skip from students inner join users on (students.email=users.email) where students.email=$1"
	insertStudent = "insert into students(email, male, major, promotion, skip) values ($1,$2,$3,$4,$5)"
	skipStudent   = "update students set skip=$2 where email=$1"
	setPromotion  = "update students set promotion=$2 where email=$1"
	setMajor      = "update students set major=$2 where email=$1"
	updateMale    = "update students set male=$2 where email=$1"
	updateAlumni  = "update students set nextPosition=$1, nextFrance=$2, nextPermanent=$3, nextSameCompany=$4, nextContact=$5 where email=$6"
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
	var nextPos sql.NullString
	var nextFrance, nextPermanent, nextSameCompany sql.NullBool
	var nextContact sql.NullString
	var role string
	s := schema.Student{
		User: schema.User{
			Person: schema.Person{},
		},
		Alumni: &schema.Alumni{},
	}
	err := rows.Scan(
		&s.Male,
		&s.User.Person.Firstname,
		&s.User.Person.Lastname,
		&s.User.Person.Email,
		&s.User.Person.Tel,
		&role,
		&lastVisit,
		&s.Promotion,
		&s.Major,
		&nextPos,
		&nextFrance,
		&nextPermanent,
		&nextSameCompany,
		&nextContact,
		&s.Skip)
	s.User.Role = schema.Role(role)
	s.User.LastVisit = nullableTime(lastVisit)
	if nextPos.Valid {
		s.Alumni.Contact = nullableString(nextContact)
		s.Alumni.Position = nullableString(nextPos)
		s.Alumni.France = nullableBool(nextFrance, false)
		s.Alumni.Permanent = nullableBool(nextPermanent, false)
		s.Alumni.SameCompany = nullableBool(nextSameCompany, false)
	} else {
		s.Alumni = nil
	}
	return s, err
}

//Students list all the registered students
func (s *Store) Students() (schema.Students, error) {
	var students []schema.Student
	rows, err := s.db.Query(allStudents)
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		s, e := scanStudent(rows)
		if e != nil {
			return students, e
		}
		students = append(students, s)
	}
	return students, err
}

//NewStudent create a new student using a given person, major, promotion and gender
//The major and the promotion must be valid against the supported values in config.Internships
//The underlying user account is created
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

//SetAlumni update a student alumni status
func (s *Store) SetAlumni(student string, a schema.Alumni) error {
	if !strings.Contains(a.Contact, "@") {
		return schema.ErrInvalidEmail
	}
	return s.singleUpdate(updateAlumni, schema.ErrUnknownUser, a.Position, a.France, a.Permanent, a.SameCompany, a.Contact, student)
}
