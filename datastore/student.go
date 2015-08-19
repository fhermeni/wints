package datastore

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	AllStudents     = "select firstname, lastname, email, tel, role, lastVisit, promotion, major, nextPosition, nextContact, skip from students inner join users on (students.email=users.email)"
	InsertStudent   = "insert into students(major, promotion, nextContact, nextPosition, skip) values ($1,$2,$3,$4,$5)"
	SkipStudent     = "update student set skip=$2 where email=$1"
	SetPromotion    = "update students set promotion=$2 where email=$1"
	SetMajor        = "update internships set major=$2 where email=$1"
	SetNextPosition = "update students set nextPosition=$1 where email=$2"
	SetNextContact  = "update students set nextContact=$1 where email=$2"
)

func (s *Service) Students() ([]internship.Student, error) {
	rows, err := s.DB.Query(AllStudents)

	students := make([]internship.Student, 0, 0)
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		var lastVisit pq.NullTime
		s := internship.Student{User: internship.User{}, Alumni: internship.Alumni{}}
		rows.Scan(
			&s.User.Firstname,
			&s.User.Lastname,
			&s.User.Email,
			&s.User.Tel,
			&s.User.Role,
			&lastVisit,
			&s.Promotion,
			&s.Major,
			&s.Alumni.Position,
			&s.Alumni.Contact,
			&s.Skip)

		if lastVisit.Valid {
			s.User.LastVisit = &lastVisit.Time
		}
		students = append(students, s)
	}
	return students, err
}

func (s *Service) AddStudent(st internship.Student) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return rollback(err, tx)
	}
	if _, err := s.addUser(tx, st.User); err != nil {
		return rollback(err, tx)
	}
	res, err := tx.Exec(InsertStudent, st.Major, st.Promotion, st.Alumni.Contact, st.Alumni.Position, false)
	if err != nil {
		return rollback(err, tx)
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return rollback(err, tx)
	}
	if nb != 1 {
		return internship.ErrUnknownUser
	}
}

func (s *Service) SkipStudent(em string, st bool) error {
	return s.singleUpdate(SkipStudent, internship.ErrUnknownUser, em, st)
}

func (s *Service) InsertStudents(file string) error {
	in := csv.NewReader(strings.NewReader(file))
	in.Comma = ';'
	//Get rid of the header
	in.Read()
	for {
		record, e2 := in.Read()
		if e2 == io.EOF {
			break
		}
		fn := record[1]
		ln := record[0]
		email := record[2]
		p := record[3]
		major := record[4]
		st := internship.Student{
			Firstname:  fn,
			Lastname:   ln,
			Email:      email,
			Promotion:  p,
			Major:      major,
			Internship: "",
		}
		e2 = s.AddStudent(st)
		// end-of-file is fitted into err
	}
	return nil
}

func (s *Service) SetPromotion(stu, p string) error {
	return s.singleUpdate(SetPromotion, internship.ErrUnknownUser, stu, p)
}

func (s *Service) SetMajor(stu, m string) error {
	return s.singleUpdate(SetMajor, internship.ErrUnknownUser, stu, m)
}

func (s *Service) SetNextPosition(student string, pos int) error {
	return s.singleUpdate(SetNextPosition, internship.ErrUnknownUser, pos, student)
}

func (s *Service) SetNextContact(student string, em string) error {
	if !strings.Contains(em, "@") || strings.Contains(em, "@unice.fr") || strings.Contains(em, "polytech") {
		return internship.ErrInvalidAlumniEmail
	}
	return s.singleUpdate(SetNextContact, internship.ErrUnknownUser, em, student)
}
