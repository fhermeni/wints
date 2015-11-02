//+build integration

package sqlstore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

func TestReports(t *testing.T) {

	stu := newStudent(t)
	em := stu.User.Person.Email
	unknown := string(randomBytes(12))
	tut := newTutor(t)
	ints := newInternship(t, stu, tut)

	//Get the report
	r, err := store.Report("foo", em)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ints.Reports))
	assert.Equal(t, ints.Reports[0], r)

	//Bad reports
	_, err = store.Report("", em)
	assert.Equal(t, schema.ErrUnknownReport, err)
	_, err = store.Report("foo", unknown)
	assert.Equal(t, schema.ErrUnknownReport, err)

	//Setters
	assert.Nil(t, store.SetReportPrivacy(r.Kind, em, true))

	now := time.Now()
	assert.Nil(t, store.SetReportDeadline(r.Kind, em, now))

	r.Private = true
	r.Deadline = now.Truncate(time.Minute).UTC()

	r2, err := store.Report(r.Kind, em)
	assert.Nil(t, err)
	assert.Equal(t, r, r2)

	//Grade
	cnt := randomBytes(12)
	at, err := store.SetReportContent(r.Kind, em, cnt)
	assert.Nil(t, err)
	at, err = store.SetReportContent(r.Kind, em, cnt)
	assert.Nil(t, err)

	//Place the deadline in the past so no more upload possible
	deadline := at.AddDate(0, -1, 0)
	assert.Nil(t, store.SetReportDeadline(r.Kind, em, deadline))
	_, err = store.SetReportContent(r.Kind, em, cnt)
	assert.Equal(t, schema.ErrDeadlinePassed, err)

	_, err = store.SetReportGrade(r.Kind, em, 22, "fooooo")
	assert.Equal(t, schema.ErrInvalidGrade, err)
	_, err = store.SetReportGrade(r.Kind, em, -2, "fooooo")
	assert.Equal(t, schema.ErrInvalidGrade, err)
	_, err = store.SetReportGrade("bar", em, 3, "fooooo")
	assert.Equal(t, schema.ErrUnknownReport, err)
	_, err = store.SetReportGrade(r.Kind, "", 3, "fooooo")
	assert.Equal(t, schema.ErrUnknownReport, err)

	reviewed, err := store.SetReportGrade(r.Kind, em, 3, "fooooo")
	assert.Nil(t, err)
	r.Deadline = deadline
	r.Delivery = &at
	r.Reviewed = &reviewed
	r.Grade = 3
	r.Comment = "fooooo"
	r2, err = store.Report(r.Kind, em)
	assert.Nil(t, err)
	assert.Equal(t, r, r2)

	//Cannot resubmit once graded
	_, err = store.SetReportContent(r.Kind, em, cnt)
	assert.Equal(t, schema.ErrGradedReport, err)
	r2, err = store.Report(r.Kind, em)
	assert.Nil(t, err)
	assert.Equal(t, r, r2)
}
