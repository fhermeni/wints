package jobs

import (
	"time"

	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
)

//MissingSurveys scans the missing surveys
//One week before the deadine, it send a first mail, then a reminder every week until the survey is uploaded
func MissingSurveys(ints []schema.Internship, not *notifier.Notifier) {
	now := time.Now()
	for _, i := range ints {
		for _, s := range i.Surveys {
			if s.Deadline.Before(now) && s.Delivery == nil {
				if sameDay(s.Deadline, now) || diffDays(s.Deadline, now)%7 == 0 {
					not.SurveyRequest(i.Convention.Supervisor, i.Convention.Tutor, i.Convention.Student, s, nil)
				}
			}
		}
	}
}

func sameDay(d1, d2 time.Time) bool {
	return d1.Month() == d2.Month() && d1.Day() == d2.Day() && d1.Year() == d2.Year()
}

func diffDays(d1, d2 time.Time) int {
	delta := d2.Sub(d1)
	return int(delta.Hours() / 24)
}
