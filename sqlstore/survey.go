package sqlstore

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	selectSurveyFromToken = "select student, kind from surveys where token=$1"
	selectSurvey          = "select kind, deadline, timestamp, answers, token from surveys where student=$1 and kind=$2"
	selectSurveys         = "select kind, deadline, timestamp, answers, token from surveys where student=$1"
	updateSurveyContent   = "update surveys set answers=$1, timestamp=$2 where token=$3"
)

func (s *Store) SurveyToken(token string) (string, string, error) {
	student := ""
	kind := ""
	st, err := s.stmt(selectSurveyFromToken)
	if err != nil {
		return student, kind, err
	}
	err = st.QueryRow(token).Scan(&student, &kind)
	return student, kind, mapCstrToError(err)
}

func scanSurvey(rows *sql.Rows) (internship.SurveyHeader, error) {
	survey := internship.SurveyHeader{}
	var delivery pq.NullTime
	var buf []byte

	err := rows.Scan(
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
	return survey, err
}

func (s *Store) Surveys(student string) (map[string]internship.SurveyHeader, error) {
	res := make(map[string]internship.SurveyHeader)
	st, err := s.stmt(selectSurveys)
	if err != nil {
		return res, err
	}
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

func (s *Store) Survey(student, kind string) (internship.SurveyHeader, error) {
	st, err := s.stmt(selectSurvey)
	if err != nil {
		return internship.SurveyHeader{}, err
	}
	rows, err := st.Query(student, kind)
	if err != nil {
		return internship.SurveyHeader{}, err
	}
	defer rows.Close()
	return scanSurvey(rows)
}

func (s *Store) SetSurveyContent(token string, cnt map[string]string) error {
	now := time.Now()
	if len(cnt) > 1000 {
		return internship.ErrInvalidSurvey
	}
	buf, err := json.Marshal(cnt)
	if err != nil {
		return internship.ErrInvalidSurvey
	}
	return s.singleUpdate(updateSurveyContent, internship.ErrUnknownSurvey, buf, now, token)
}
