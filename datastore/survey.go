package datastore

import (
	"encoding/json"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

func (s *Service) SurveyToken(token string) (string, string, error) {
	var student, kind string
	err := s.DB.QueryRow("select student, kind from surveys where token=$1", token).Scan(&student, &kind)
	if err != nil {
		return "", "", internship.ErrUnknownSurvey
	}
	return student, kind, err
}

func (s *Service) Survey(student, kind string) (internship.Survey, error) {
	var deadline time.Time
	var timeStamp pq.NullTime
	var buf []byte
	var token string
	r := s.DB.QueryRow("select kind, deadline, timestamp, answers, token from surveys where student=$1 and kind=$2", student, kind)
	err := r.Scan(&kind, &deadline, &timeStamp, &buf, &token)
	if err != nil {
		return internship.Survey{}, err
	}
	hdr := internship.Survey{Kind: kind, Deadline: deadline, Answers: make(map[string]string), Token: token}
	if timeStamp.Valid {
		hdr.Timestamp = timeStamp.Time
		err = json.Unmarshal(buf, &hdr.Answers)
		if err != nil {
			return internship.Survey{}, err
		}
	}
	return hdr, nil
}

func (s *Service) SetSurveyContent(token string, cnt map[string]string) error {
	now := time.Now()
	if len(cnt) > 1000 {
		return internship.ErrInvalidSurvey
	}
	buf, err := json.Marshal(cnt)
	if err != nil {
		return internship.ErrInvalidSurvey
	}
	return SingleUpdate(s.DB, internship.ErrUnknownSurvey, "update surveys set answers=$1, timestamp=$2 where token=$3", buf, now, token)
}

func (s *Service) SurveyDefs() []internship.SurveyDef {
	return s.surveyDefs
}

func (s *Service) RequestSurvey(stu, kind string) error {
	if i, err := s.Internship(stu); err != nil {
		return err
	} else {
		s.mailer.SendSurveyRequest(i, kind)
	}
	return nil
}
