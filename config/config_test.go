package config

import (
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRelativeReportHeader(t *testing.T) {
	d, _ := time.ParseDuration("2h")
	def := ReportsDef{Name: "DoW", Deadline: "relative", Value: "2h"}
	now := time.Now()
	r, err := ReportHeader(now, def)
	assert.NoError(t, err)
	at := now.Add(d)
	assert.Equal(t, def.Name, r.Kind)
	assert.Equal(t, at, r.Deadline)
	assert.Equal(t, -1, r.Grade)
	assert.False(t, r.Graded())
}

func TestAbsoluteReportHeader(t *testing.T) {
	d, _ := time.Parse(DateLayout, "01/01/1990 12:00")
	def := ReportsDef{Name: "DoW", Deadline: "absolute", Value: "01/01/1990 12:00"}
	now := time.Now()
	r, err := ReportHeader(now, def)
	assert.NoError(t, err)
	assert.Equal(t, def.Name, r.Kind)
	assert.Equal(t, d, r.Deadline)
	assert.Equal(t, -1, r.Grade)
	assert.False(t, r.Graded())
}
