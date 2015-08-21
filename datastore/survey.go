package datastore

import (
	"encoding/json"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	SelectSurveyFromToken = "select student, kind from surveys where token=$1"
	SelectSurvey          = "select kind, deadline, timestamp, answers, token from surveys where student=$1 and kind=$2"
	UpdateSurveyContent   = "update surveys set answers=$1, timestamp=$2 where token=$3"
)

func (s *Service) SurveyToken(token string) (string, string, error) {
	student := ""
	kind := ""
	st, err := s.stmt(SelectSurveyFromToken)
	if err != nil {
		return student, kind, err
	}
	err = st.QueryRow(token).Scan(&student, &kind)
	err = mapCstrToError(err)
	return student, kind, err
}

func (s *Service) Survey(student, kind string) (internship.Survey, error) {
	var delivery pq.NullTime
	var buf []byte
	survey := internship.Survey{}
	st, err := s.stmt(SelectSurvey)
	if err != nil {
		return survey, err
	}

	err = st.QueryRow(student, kind).Scan(
		&survey.Kind,
		&survey.Deadline,
		delivery,
		&buf,
		&survey.Token)
	err = mapCstrToError(err)
	if err != nil {
		return survey, err
	}
	if delivery.Valid {
		survey.Delivery = &delivery.Time
		err = json.Unmarshal(buf, &survey.Answers)
	}
	return survey, nil
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
	return s.singleUpdate(UpdateSurveyContent, internship.ErrUnknownSurvey, buf, now, token)
}

/*

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
*/
