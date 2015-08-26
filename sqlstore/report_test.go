//+build integration

package sqlstore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestReports(t *testing.T) {

	//p, passwd := preparePerson(t)

	//The error status
	fake := string(randomBytes(10))
	_, err := store.Report("", fake)
	assert.Equal(t, internship.ErrUnknownStudent, err)
	_, err = store.ReportContent("", fake)
	assert.Equal(t, internship.ErrUnknownReport, err)
	assert.Equal(t, internship.ErrUnknownReport, store.SetReportContent("", fake, []byte{}))
	assert.Equal(t, internship.ErrUnknownReport, store.SetReportGrade("", fake, 0, ""))
	assert.Equal(t, internship.ErrUnknownReport, store.SetReportDeadline("", fake, time.Now()))
	assert.Equal(t, internship.ErrUnknownReport, store.SetReportPrivacy("", fake, true))
}
