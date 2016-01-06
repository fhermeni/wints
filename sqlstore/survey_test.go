// +build integration

package sqlstore

import (
	"testing"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

func TestSurveyMgnt(t *testing.T) {
	student := newStudent(t)
	em := student.User.Person.Email
	tutor := newTutor(t)
	newInternship(t, student, tutor)

	i, err := store.Internship(em)
	assert.Nil(t, err)
	s := i.Surveys[0]
	//Getter
	s2, err := store.Survey(em, s.Kind)
	assert.Nil(t, err)
	assert.Equal(t, s, s2)
	_, err = store.Survey(student.Major, s.Kind)
	assert.Equal(t, schema.ErrUnknownSurvey, err)
	_, err = store.Survey(em, em)
	assert.Equal(t, schema.ErrUnknownSurvey, err)

	ss, err := store.Surveys(em)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ss))
	assert.Contains(t, ss, s)
	_, s2, err = store.SurveyFromToken(s.Token)
	assert.Nil(t, err)
	assert.Equal(t, s, s2)

	_, _, err = store.SurveyFromToken(em)
	assert.Equal(t, schema.ErrUnknownSurvey, err)

	cnt := map[string]string{
		"rsc": "3711",
		"r":   "2138",
		"gri": "1908",
		"adg": "912",
	}
	d, err := store.SetSurveyContent(s.Token, cnt)
	assert.Nil(t, err)

	_, err = store.SetSurveyContent(em, cnt)
	assert.Equal(t, schema.ErrUnknownSurvey, err)

	//Reupload disallowed
	_, err = store.SetSurveyContent(s.Token, cnt)
	assert.Equal(t, schema.ErrSurveyUploaded, err)

	//Reset & re-upload
	err = store.ResetSurveyContent(em, s.Kind)
	assert.Nil(t, err)

	d, err = store.SetSurveyContent(s.Token, cnt)
	assert.Nil(t, err)

	s2, err = store.Survey(em, s.Kind)
	assert.Nil(t, err)
	s.Delivery = &d
	s.Cnt = cnt
	assert.Equal(t, s, s2)
}
