package sqlstore

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	selectSurveyFromToken  = "select student, kind, invitation, lastInvitation, deadline, delivery, cnt, token from surveys where token=$1"
	selectSurvey           = "select student, kind, invitation, lastInvitation, deadline, delivery, cnt, token from surveys where student=$1 and kind=$2"
	selectSurveys          = "select student, kind, invitation, lastInvitation, deadline, delivery, cnt, token from surveys where student=$1 order by deadline asc"
	selectAllSurveys       = "select student, kind, invitation, lastInvitation, deadline, delivery, cnt, token from surveys order by deadline asc"
	updateSurveyContent    = "update surveys set cnt=$1, delivery=$2 where token=$3"
	resetSurveyContent     = "update surveys set cnt=null, delivery=null where student=$1 and kind=$2"
	updateSurveyInvitation = "update surveys set lastInvitation=$3 where student=$1 and kind=$2"
)

//SurveyFromToken returns a given survey identifier
func (s *Store) SurveyFromToken(token string) (string, schema.SurveyHeader, error) {
	st := s.stmt(selectSurveyFromToken)
	rows, err := st.Query(token)
	if err != nil {
		return "", schema.SurveyHeader{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return "", schema.SurveyHeader{}, schema.ErrUnknownSurvey
	}
	return scanSurvey(rows)
}

func scanSurvey(rows *sql.Rows) (string, schema.SurveyHeader, error) {
	survey := schema.SurveyHeader{}
	var delivery pq.NullTime
	var student string
	var buf sql.NullString
	err := rows.Scan(
		&student,
		&survey.Kind,
		&survey.Invitation,
		&survey.LastInvitation,
		&survey.Deadline,
		&delivery,
		&buf,
		&survey.Token)
	err = mapCstrToError(err)
	if err != nil {
		return student, survey, err
	}
	if buf.Valid {
		if e := json.Unmarshal([]byte(buf.String), &survey.Cnt); e != nil {
			return student, survey, e
		}
	}
	survey.Deadline = survey.Deadline.Truncate(time.Minute).UTC()
	survey.Invitation = survey.Invitation.Truncate(time.Minute).UTC()
	survey.LastInvitation = survey.LastInvitation.Truncate(time.Minute).UTC()
	survey.Delivery = nullableTime(delivery)
	return student, survey, err
}

func (s *Store) allSurveys() (map[string][]schema.SurveyHeader, error) {
	res := make(map[string][]schema.SurveyHeader)
	st := s.stmt(selectAllSurveys)
	rows, err := st.Query()
	if err != nil {
		return res, nil
	}
	defer rows.Close()
	for rows.Next() {
		stu, s, err := scanSurvey(rows)
		if err != nil {
			return res, err
		}
		surveys, ok := res[stu]
		if !ok {
			surveys = make([]schema.SurveyHeader, 0, 0)
		}
		surveys = append(surveys, s)
		res[stu] = surveys
	}
	return res, nil
}

//Surveys returns all the survey related to a student
func (s *Store) Surveys(student string) ([]schema.SurveyHeader, error) {
	var res []schema.SurveyHeader
	st := s.stmt(selectSurveys)
	rows, err := st.Query(student)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		_, s, e := scanSurvey(rows)
		if e != nil {
			return res, e
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
	defer rows.Close()
	if !rows.Next() {
		return schema.SurveyHeader{}, schema.ErrUnknownSurvey
	}
	_, srv, err := scanSurvey(rows)
	return srv, err
}

//SetSurveyInvitation stores the invitation date for a survey to fullfil
func (s *Store) SetSurveyInvitation(student, kind string) (time.Time, error) {
	t := time.Now()
	return t, s.singleUpdate(updateSurveyInvitation, schema.ErrUnknownSurvey, student, kind, t)
}

//SetSurveyContent stores the survey answers
func (s *Store) SetSurveyContent(token string, cnt interface{}) (time.Time, error) {
	buf, err := json.Marshal(cnt)
	if err != nil {
		return time.Now(), err
	}
	now := time.Now().Truncate(time.Minute).UTC()
	_, sr, err := s.SurveyFromToken(token)
	if err != nil {
		return now, err
	}
	if sr.Delivery != nil {
		return now, schema.ErrSurveyUploaded
	}
	return now, s.singleUpdate(updateSurveyContent, schema.ErrUnknownSurvey, buf, now, token)
}

//ResetSurveyContent delete an uploaded survey
func (s *Store) ResetSurveyContent(student, kind string) error {
	return s.singleUpdate(resetSurveyContent, schema.ErrUnknownSurvey, student, kind)
}
