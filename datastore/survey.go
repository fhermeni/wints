package datastore

import (
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

func (s *Service) SurveyToken(token string) (string, string, error) {
	var student, kind string
	err := s.DB.QueryRow("select student, kind from survey where token=$1", token).Scan(&student, &kind)
	if err != nil {
		return "", "", internship.ErrUnknownSurvey
	}
	return student, kind, err
}

func (s *Service) Survey(student, kind string) (internship.SurveyHeader, error) {
	var stuFn, stuLn, stuMail, stuTel string
	var tutFn, tutLn, tutMail, tutTel string
	var deadline time.Time
	var timeStamp pq.NullTime
	r := s.DB.QueryRow("select s.Firstname, s.Lastname, s.Email, s.Tel, t.Firstname, t.Lastname, t.Email, t.Tel, kind, deadline, timestamp from user as s, user as t, internships, survey where student=$1 and kind=$2 and internships.student=s.email and internships.tutor=t.email and student=internships.student", student, kind)
	err := r.Scan(&stuFn, &stuLn, &stuMail, &stuTel, &tutFn, &tutLn, &tutMail, &tutTel, &kind, &deadline, &timeStamp)
	if err != nil {
		return internship.SurveyHeader{}, err
	}
	stu := internship.Person{Firstname: stuFn, Lastname: stuLn, Email: stuMail, Tel: stuTel}
	tut := internship.Person{Firstname: tutFn, Lastname: tutLn, Email: tutMail, Tel: tutTel}
	hdr := internship.SurveyHeader{Student: stu, Tutor: tut, Kind: kind, Deadline: deadline}
	if timeStamp.Valid {
		hdr.Timestamp = timeStamp.Time
	}
	return hdr, nil
}

func (s *Service) SurveyContent(student, kind string) (map[string]string, error) {
	rows, err := s.DB.Query("select id, value from survey_questions where student=$1 and kind=$2", student, kind)
	if err != nil {
		return make(map[string]string), err
	}
	defer rows.Close()
	res := make(map[string]string)
	for rows.Next() {
		var id, value string
		rows.Scan(&id, &value)
		res[id] = value
	}
	return res, nil
}
