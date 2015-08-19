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
	InsertStudent   = "insert into students(major, promotion, skip, nextContact, nextPosition) values ($1,$2,$3,$4,$5)"
	SkipStudent     = "update students set skip=$2 where email=$1"
	SetPromotion    = "update students set promotion=$2 where email=$1"
	SetMajor        = "update students set major=$2 where email=$1"
	SetNextPosition = "update students set nextPosition=$1 where email=$2"
	SetNextContact  = "update students set nextContact=$1 where email=$2"
)

//Students list all the registered students
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

//AddStudent add a student to the database.
//The underlying user is declared in prior with a random password
/*func (s *Service) AddStudent(st internship.Student) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	if err := s.addUser(tx, st.User); err != nil {
		return rollback(err, tx)
	}
	_, err = tx.Exec(InsertStudent, st.Major, st.Promotion, st.Skip, st.Alumni.Contact, st.Alumni.Position)
	if violated(err, "pk_students_email") {
		return rollback(internship.ErrStudentExists, tx)
	}
	if violated(err, "fk_students_email") {
		return rollback(internship.ErrUnknownUser, tx)
	}
	if err != nil {
		return rollback(err, tx)
	}
	return tx.Commit()
}*/

func (s *Service) AddStudent(st internship.Student) error {
	rb := newRollbackable(s.DB)
	rb.err = s.addUser(rb.tx, st.User)
	rb.Exec(InsertStudent, st.Major, st.Promotion, st.Skip, st.Alumni.Contact, st.Alumni.Position)
	rb.err = violationAsErr(rb.err, "pk_students_email", internship.ErrStudentExists)
	rb.err = violationAsErr(rb.err, "fk_students_email", internship.ErrUnknownUser)
	return rb.Done()
}

//SkipStudent indicates if it is not required for the student to get an internship (an abandom typically)
func (s *Service) SkipStudent(em string, st bool) error {
	return s.singleUpdate(SkipStudent, internship.ErrUnknownUser, em, st)
}

//InsertStudents inserts all the students provided in the given CSV file.
//Expected format: firstname;lastname;email;tel
//The other fields are set to the default values
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
			User: internship.User{
				Firstname: fn,
				Lastname:  ln,
				Email:     email,
				Tel:       "not available",
				Role:      internship.STUDENT,
			},
			Promotion: p,
			Major:     major,
			Skip:      false,
			Alumni: internship.Alumni{
				Contact:  "", //Not available
				Position: 0,  //Not available
			},
		}
		e2 = s.AddStudent(st)
	}
	return nil
}

//SetPromotion updates the student promotion
func (s *Service) SetPromotion(stu, p string) error {
	return s.singleUpdate(SetPromotion, internship.ErrUnknownUser, stu, p)
}

//SetMajor updates the student major
func (s *Service) SetMajor(stu, m string) error {
	return s.singleUpdate(SetMajor, internship.ErrUnknownUser, stu, m)
}

//SetNextPosition updates the student next position
func (s *Service) SetNextPosition(student string, pos int) error {
	return s.singleUpdate(SetNextPosition, internship.ErrUnknownUser, pos, student)
}

//SetNextContact updates the student next contact email.
//The email must contains one '@' and must not contains '@unice.fr' or 'polytech' to prevent from incorrect emails
func (s *Service) SetNextContact(student string, em string) error {
	if !strings.Contains(em, "@") || strings.Contains(em, "@unice.fr") || strings.Contains(em, "polytech") {
		return internship.ErrInvalidAlumniEmail
	}
	return s.singleUpdate(SetNextContact, internship.ErrUnknownUser, em, student)
}
