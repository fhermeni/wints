package sqlstore

import (
	"strings"

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
func (s *Service) Students() ([]internship.Student, error) {
	rows, err := s.DB.Query(allStudents)

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

		if lastVisit.Valid {
			s.User.LastVisit = &lastVisit.Time
		}
		students = append(students, s)
	}
	return students, err
}

func (s *Service) ToStudent(email, major, promotion string, male bool) error {
	tx := newTxErr(s.DB)
	tx.Update(insertStudent, email, male, major, promotion, false)
	tx.Update(updateUserRole, email, internship.STUDENT)
	return tx.Done()
}

//SkipStudent indicates if it is not required for the student to get an internship (an abandom typically)
func (s *Service) SkipStudent(em string, st bool) error {
	return s.singleUpdate(skipStudent, internship.ErrUnknownUser, em, st)
}

//InsertStudents inserts all the students provided in the given CSV file.
//Expected format: firstname;lastname;email;tel
//The other fields are set to the default values
/*func (s *Service) InsertStudents(file string) error {
	in := csv.NewReader(strings.NewReader(file))
	in.Comma = ';'
	//Get rid of the header
	in.Read()
	for {
		record, e2 := in.Read()
		if e2 == io.EOF {
			break
		}
		male, err := strconv.ParseBool(record[0])
		if err != nil {
			male = true
		}
		ln := record[0]
		fn := record[1]
		email := record[2]
		p := record[3]
		major := record[4]
		st := internship.Student{
			User: internship.User{
				Person: internship.Person{
					Firstname: fn,
					Lastname:  ln,
					Email:     email,
					Tel:       "not available",
				},
				Role: internship.STUDENT,
			},
			Promotion: p,
			Major:     major,
			Skip:      false,
			Alumni: internship.Alumni{
				Contact:  "", //Not available
				Position: 0,  //Not available
			},
			Male: male,
		}
		e2 = s.AddStudent(st)
	}
	return nil
}*/

//SetPromotion updates the student promotion
func (s *Service) SetPromotion(stu, p string) error {
	return s.singleUpdate(setPromotion, internship.ErrUnknownStudent, stu, p)
}

//SetMajor updates the student major
func (s *Service) SetMajor(stu, m string) error {
	return s.singleUpdate(setMajor, internship.ErrUnknownStudent, stu, m)
}

func (s *Service) SetMale(stu, m bool) error {
	return s.singleUpdate(updateMale, internship.ErrUnknownStudent, stu, m)
}

//SetNextPosition updates the student next position
func (s *Service) SetNextPosition(student string, pos int) error {
	return s.singleUpdate(setNextPosition, internship.ErrUnknownUser, pos, student)
}

//SetNextContact updates the student next contact email.
//The email must contains one '@' and must not contains '@unice.fr' or 'polytech' to prevent from incorrect emails
func (s *Service) SetNextContact(student string, em string) error {
	if !strings.Contains(em, "@") || strings.Contains(em, "@unice.fr") || strings.Contains(em, "polytech") {
		return internship.ErrInvalidAlumniEmail
	}
	return s.singleUpdate(setNextContact, internship.ErrUnknownUser, em, student)
}
