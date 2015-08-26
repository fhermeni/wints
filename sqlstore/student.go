package sqlstore

import (
	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	allStudents     = "select male, firstname, lastname, email, tel, role, lastVisit, promotion, major, nextPosition, nextContact, skip from students inner join users on (students.email=users.email)"
	insertStudent   = "insert into students(email, male, major, promotion, skip) values ($1,$2,$3,$4,$5)"
	skipStudent     = "update students set skip=$2 where email=$1"
	setPromotion    = "update students set promotion=$2 where email=$1"
	setMajor        = "update students set major=$2 where email=$1"
	updateMale      = "update students set male=$2 where email=$1"
	setNextPosition = "update students set nextPosition=$1 where email=$2"
	setNextContact  = "update students set nextContact=$1 where email=$2"
)

//Students list all the registered students
func (s *Store) Students() ([]internship.Student, error) {
	rows, err := s.db.Query(allStudents)

	students := make([]internship.Student, 0, 0)
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		var lastVisit pq.NullTime
		s := internship.Student{
			User: internship.User{
				Person: internship.Person{},
			},
			Alumni: internship.Alumni{},
		}
		rows.Scan(
			&s.Male,
			&s.User.Person.Firstname,
			&s.User.Person.Lastname,
			&s.User.Person.Email,
			&s.User.Person.Tel,
			&s.User.Role,
			&lastVisit,
			&s.Promotion,
			&s.Major,
			&s.Alumni.Position,
			&s.Alumni.Contact,
			&s.Skip)
		s.User.LastVisit = nullableTime(lastVisit)
		students = append(students, s)
	}
	return students, err
}

//ToStudent convert a user account to a student account
func (s *Store) ToStudent(email, major, promotion string, male bool) error {
	tx := newTxErr(s.db)
	tx.Update(insertStudent, email, male, major, promotion, false)
	tx.Update(updateUserRole, email, internship.STUDENT)
	return tx.Done()
}

//SkipStudent indicates if it is not required for the student to get an internship (an abandom typically)
func (s *Store) SkipStudent(em string, st bool) error {
	return s.singleUpdate(skipStudent, internship.ErrUnknownUser, em, st)
}

//SetPromotion updates the student promotion
func (s *Store) SetPromotion(stu, p string) error {
	return s.singleUpdate(setPromotion, internship.ErrUnknownStudent, stu, p)
}

//SetMajor updates the student major
func (s *Store) SetMajor(stu, m string) error {
	return s.singleUpdate(setMajor, internship.ErrUnknownStudent, stu, m)
}

//SetMale stores the student gender
func (s *Store) SetMale(stu string, m bool) error {
	return s.singleUpdate(updateMale, internship.ErrUnknownStudent, stu, m)
}

//SetNextPosition updates the student next position
func (s *Store) SetNextPosition(student string, pos int) error {
	return s.singleUpdate(setNextPosition, internship.ErrUnknownUser, pos, student)
}

//SetNextContact updates the student next contact email.
func (s *Store) SetNextContact(student string, em string) error {
	return s.singleUpdate(setNextContact, internship.ErrUnknownUser, em, student)
}
