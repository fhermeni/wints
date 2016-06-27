package jobs

import (
	"time"

	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
)

//NeverLogged scan students that never logged in since
//one week before the beginning of the internship
func NeverLogged(provider schema.Internshipser, not *notifier.Notifier) {
	ints, err := provider.Internships()
	logger.Log("event", "cron", "Scanning the internships for student that never connected", err)
	if err != nil {
		return
	}
	for _, i := range ints {
		last := i.Convention.Student.User.LastVisit
		if last == nil && i.Convention.Begin.Before(time.Now()) {
			//The internship began
			not.ReportIdleAccount(i.Convention.Student.User)
		}
	}
}
