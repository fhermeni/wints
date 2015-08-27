package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	selectSurveyFromToken = "select kind, deadline, timestamp, cnt, token from surveys where token=$1"
	selectSurvey          = "select kind, deadline, timestamp, cnt, token from surveys where student=$1 and kind=$2"
	selectSurveys         = "select kind, deadline, timestamp, cnt, token from surveys where student=$1"
	updateSurveyContent   = "update surveys set cnt=$1, timestamp=$2 where token=$3"
)

//SurveyFromToken returns a given survey identifier
func (s *Store) SurveyFromToken(token string) (schema.SurveyHeader, error) {
	st := s.stmt(selectSurveyFromToken)
	rows, err := st.Query(token)
	if err != nil {
		return schema.SurveyHeader{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return schema.SurveyHeader{}, schema.ErrUnknownSurvey
	}
	return scanSurvey(rows)
}

func scanSurvey(rows *sql.Rows) (schema.SurveyHeader, error) {
	survey := schema.SurveyHeader{}
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
func (s *Store) Surveys(student string) (map[string]schema.SurveyHeader, error) {
	res := make(map[string]schema.SurveyHeader)
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
func (s *Store) Survey(student, kind string) (schema.SurveyHeader, error) {
	st := s.stmt(selectSurvey)
	rows, err := st.Query(student, kind)
	if err != nil {
		return schema.SurveyHeader{}, err
	}
	rows.Next()
	defer rows.Close()
	return scanSurvey(rows)
}

//SetSurveyContent stores the survey answers
func (s *Store) SetSurveyContent(token string, cnt []byte) error {
	return s.singleUpdate(updateSurveyContent, schema.ErrUnknownSurvey, cnt, time.Now(), token)
}
