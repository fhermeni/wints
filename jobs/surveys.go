package jobs

import (
	"time"

	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
)

//MissingSurveys scans the missing surveys
//One week before the deadine, it send a first mail, then a reminder every week until the survey is uploaded
func MissingSurveys(provider schema.Internshipser, not *notifier.Notifier) {
	ints, err := provider.Internships()
	logger.Log("event", "cron", "Scanning the internships for surveys to request", err)
	if err != nil {
		return
	}

	now := time.Now()
	for _, i := range ints {
		if i.Convention.Skip {
			continue
		}
		for _, s := range i.Surveys {
			if s.Delivery == nil && s.Invitation.Before(now) {
				//Its time to notify the supervisor
				not.SurveyRequest(i.Convention.Supervisor, i.Convention.Tutor, i.Convention.Student, s, nil)
			}
		}
	}
}
