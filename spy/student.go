package spy

import (
	"log"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/sqlstore"
)

type Students struct {
	st        *sqlstore.Store
	reminders map[string]config.Report
}

/*
 Reminder about the deadlines
*/

func (stu *Students) Run() {
	log.Println("Running Tutored spy")

	ints, err := stu.st.Internships()
	if err != nil {
		return
	}
	for _, i := range ints {
		stu.spy(i.Convention.Student, i.Reports)
	}
}

func (stu *Students) spy(s schema.Student, reports map[string]schema.ReportHeader) {
	notify := false
	now := time.Now()
	for _, r := range reports {
		if r.Delivery == nil {
			d := stu.reminders[r.Kind].Reminder
			if r.Delivery.Add(d.Duration).After(now) {
				notify = true
				break
			}
		}
	}
	if notify {
		//show the deadline. Explicit late
	}
}
