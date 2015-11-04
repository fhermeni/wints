package spy

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/testutil"
	"github.com/stretchr/testify/assert"
)

var tut1 = testutil.Tutor()
var tut2 = testutil.Tutor()
var i1 = testutil.Internship(tut1)
var i2 = testutil.Internship(tut1)
var i3 = testutil.Internship(tut2)

var now = time.Now()
var p1 = now.AddDate(0, 0, -1)
var p2 = now.AddDate(0, 0, -2)
var n1 = now.Add(time.Hour)

//No upload but ok
var noUp = testutil.ReportHeader("noUp", n1)

//upload and ok
var upOk = testutil.ReportHeader("upOk", now, p1)

//no upload and late
var upKo = testutil.ReportHeader("upKo", p1)

var rKo = testutil.ReportHeader("rKo", p2, p1)
var rOk = testutil.ReportHeader("rOk", p2, p1, now)

var expUpKo = ReportReminder{
	Kind:     upKo.Kind,
	Status:   "missing upload",
	Deadline: p1,
}

var expRKo = ReportReminder{
	Kind:   rKo.Kind,
	Status: "missing review",
}

func makeReporter(t *testing.T) ReminderReporter {
	return func(rem TutorReminder) error {
		if rem.Tutor == tut1 {
			assert.Equal(t, 2, len(rem.Students))
			for _, x := range rem.Students {
				assert.True(t, x.Student == i1.Convention.Student || x.Student == i2.Convention.Student)
				//assert.Contains(t, x.Reminders, expRKo, expUpKo)
			}
		} else if rem.Tutor == tut2 {
			assert.Equal(t, 1, len(rem.Students))
			assert.Equal(t, i3.Convention.Student, rem.Students[0].Student)
			//assert.Contains(t, rem.Students[0].Reminders, expRKo, expUpKo)
		}
		return nil
	}
}

func TestTutorSpy(t *testing.T) {

	reports := []schema.ReportHeader{noUp, upOk, upKo, rKo, rOk}

	cfg := config.Internships{
		Reports: make([]config.Report, 0),
	}

	for _, r := range reports {
		c := config.Report{
			Delivery: config.AbsoluteDeadline(now),
			Review:   config.Duration{Duration: time.Duration(48) * time.Hour}, //2 days for reviewing
			Kind:     r.Kind,
		}
		cfg.Reports = append(cfg.Reports, c)
	}
	i1.Reports = reports
	i2.Reports = reports
	i3.Reports = reports

	f := &testutil.FakeInternshipser{
		Ints: []schema.Internship{i1, i2, i3},
	}

	ReportsScanner(cfg, f, makeReporter(t))
}
