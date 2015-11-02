//+build integration

package sqlstore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

func TestReports(t *testing.T) {

	//p, passwd := preparePerson(t)

	//The error status
	fake := string(randomBytes(10))
	_, err := store.Report("", fake)
	assert.Equal(t, schema.ErrUnknownStudent, err)
	_, err = store.ReportContent("", fake)
	assert.Equal(t, schema.ErrUnknownReport, err)
	_, err = store.SetReportContent("", fake, []byte{})
	assert.Equal(t, schema.ErrUnknownReport, err)
	assert.Equal(t, schema.ErrUnknownReport, store.SetReportGrade("", fake, 0, ""))
	assert.Equal(t, schema.ErrUnknownReport, store.SetReportDeadline("", fake, time.Now()))
	assert.Equal(t, schema.ErrUnknownReport, store.SetReportPrivacy("", fake, true))
}
