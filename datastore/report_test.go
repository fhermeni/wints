package datastore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestReportWorkflow(t *testing.T) {
	db := getDB(t)
	defer db.Close()
	s, err := NewService(db)
	assert.NoError(t, err)
	assert.NoError(t, s.Clean())
	assert.NoError(t, s.Install())

	_, err = s.Report("dow", "")
	assert.Equal(t, internship.ErrUnknownReport, err)
	assert.Equal(t, internship.ErrUnknownReport, s.SetReportContent("dow", "", []byte{}))
	assert.Equal(t, internship.ErrUnknownReport, s.SetReportGrade("dow", "", 1))
	assert.Equal(t, internship.ErrUnknownReport, s.SetReportDeadline("dow", "", time.Now()))
	_, err = s.ReportContent("dow", "")
	assert.Equal(t, internship.ErrUnknownReport, err)

	u := internship.User{Firstname: "stu", Lastname: "dent", Email: "stu@dent.com", Role: internship.STUDENT}
	_, err = s.NewUser(u)
	assert.NoError(t, err)
	r := internship.ReportHeader{Kind: "dow", Deadline: time.Now(), Grade: -1}
	assert.NoError(t, s.PlanReport(u.Email, r))
	r2, err := s.Report("dow", u.Email)
	assert.NoError(t, err)
	assert.Equal(t, r, r2)
	_, err = s.Report("di", u.Email)
	assert.Equal(t, internship.ErrUnknownReport, err)

	ti := time.Now()
	assert.NoError(t, s.SetReportDeadline("dow", u.Email, ti))
	r2, _ = s.Report("dow", u.Email)
	assert.Equal(t, ti, r2.Deadline)

	assert.Equal(t, internship.ErrInvalidGrade, s.SetReportGrade("dow", u.Email, 24))
	r2, _ = s.Report("dow", u.Email)
	assert.False(t, r2.Graded())

	assert.NoError(t, s.SetReportGrade("dow", u.Email, 3))
	r2, _ = s.Report("dow", u.Email)
	assert.Equal(t, 3, r2.Grade)
	assert.True(t, r2.Graded())

}
