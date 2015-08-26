package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	selectSurveyFromToken = "select student, kind from surveys where token=$1"
	selectSurvey          = "select kind, deadline, timestamp, cnt, token from surveys where student=$1 and kind=$2"
	selectSurveys         = "select kind, deadline, timestamp, cnt, token from surveys where student=$1"
	updateSurveyContent   = "update surveys set cnt=$1, timestamp=$2 where token=$3"
)

//SurveyToken returns a given survey identifier
func (s *Store) SurveyToken(token string) (string, string, error) {
	student := ""
	kind := ""
	st := s.stmt(selectSurveyFromToken)
	err := st.QueryRow(token).Scan(&student, &kind)
	return student, kind, mapCstrToError(err)
}

func scanSurvey(rows *sql.Rows) (internship.SurveyHeader, error) {
	survey := internship.SurveyHeader{}
	var delivery pq.NullTime

	err := rows.Scan(
		&survey.Kind,
		&survey.Deadline,
		delivery,
		&survey.Cnt,
		&survey.Token)
	err = mapCstrToError(err)
	if err != nil {
		return survey, err
	}
	survey.Delivery = nullableTime(delivery)
	return survey, err
}

//Surveys returns all the survey related to a student
func (s *Store) Surveys(student string) (map[string]internship.SurveyHeader, error) {
	res := make(map[string]internship.SurveyHeader)
	st := s.stmt(selectSurveys)
	rows, err := st.Query(student)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		s, err := scanSurvey(rows)
		if err != nil {
			return res, nil
		}
		res[s.Kind] = s
	}
	return res, nil
}

//Survey returns the required survey
func (s *Store) Survey(student, kind string) (internship.SurveyHeader, error) {
	st := s.stmt(selectSurvey)
	rows, err := st.Query(student, kind)
	if err != nil {
		return internship.SurveyHeader{}, err
	}
	rows.Next()
	defer rows.Close()
	return scanSurvey(rows)
}

//SetSurveyContent stores the survey answers
func (s *Store) SetSurveyContent(token string, cnt []byte) error {
	return s.singleUpdate(updateSurveyContent, internship.ErrUnknownSurvey, cnt, time.Now(), token)
}
