package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	selectSurveyFromToken = "select kind, deadline, delivery, cnt, token from surveys where token=$1"
	selectSurvey          = "select kind, deadline, delivery, cnt, token from surveys where student=$1 and kind=$2"
	selectSurveys         = "select kind, deadline, delivery, cnt, token from surveys where student=$1 order by deadline asc"
	updateSurveyContent   = "update surveys set cnt=$1, delivery=$2 where token=$3"
	resetSurveyContent    = "update surveys set cnt=null and delivery=null where student=$1 and kind=$2"
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
		&delivery,
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
func (s *Store) Surveys(student string) ([]schema.SurveyHeader, error) {
	res := make([]schema.SurveyHeader, 0, 0)
	st := s.stmt(selectSurveys)
	rows, err := st.Query(student)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		s, err := scanSurvey(rows)
		if err != nil {
			return res, err
		}
		res = append(res, s)
	}
	return res, err
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

func (s *Store) ResetSurveyContent(student, kind string) error {
	return s.singleUpdate(resetSurveyContent, schema.ErrUnknownSurvey, student, kind)
}
