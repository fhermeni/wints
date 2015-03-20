package datastore

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/fhermeni/wints/internship"
)

func (s *Service) Students() ([]internship.Student, error) {
	req := "select firstname, lastname, email, promotion, major, internship, hidden from pending order by lastname"
	rows, err := s.DB.Query(req)

	students := make([]internship.Student, 0, 0)
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		var fn, ln, email, promotion, major string
		var intern sql.NullString
		var hidden bool
		rows.Scan(&fn, &ln, &email, &promotion, &major, &intern, &hidden)
		st := internship.Student{
			Firstname: fn,
			Lastname:  ln,
			Email:     email,
			Promotion: promotion,
			Major:     major,
			Hidden:    hidden,
		}
		if intern.Valid {
			st.Internship = intern.String
		}
		students = append(students, st)
	}
	return students, err
}

func (s *Service) AlignWithInternship(student string, intern string) error {
	req := "update pending set internship=$2 where email=$1"
	log.Println(student + " |" + intern + "|")
	return SingleUpdate(s.DB, internship.ErrUnknownUser, req, student, intern)
}

func (s *Service) AddStudent(st internship.Student) error {
	req := "insert into pending(firstname, lastname, email, promotion, major, internship,hidden) values ($1,$2,$3,$4,$5,$6,$7)"
	return SingleUpdate(s.DB, internship.ErrUserExists, req, st.Firstname, st.Lastname, st.Email, st.Promotion, st.Major, st.Internship, false)
}

func (s *Service) HideStudent(em string, st bool) error {
	req := "update pending set hidden=$2 where email=$1"
	return SingleUpdate(s.DB, internship.ErrUnknownUser, req, em, st)
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
