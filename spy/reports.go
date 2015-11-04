package spy

import (
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
)

//ReminderReporter signals reports that are late (not uploaded)
//or late in terms of review
type ReminderReporter func(t TutorReminder) error

//ReportReminder signals a deadline reminder for a given report
type ReportReminder struct {
	//Kind is the report kind
	Kind string
	//Deadline to remind
	Deadline time.Time
	//Status is the reminder message
	Status string
}

//StudentReminder gathers all the reminders related to a same student
type StudentReminder struct {
	Student   schema.Student
	Reminders []ReportReminder
}

//TutorReminder gathers all the StudentReminder for a same tutor
type TutorReminder struct {
	Tutor    schema.User
	Students []StudentReminder
}

//ReportsScanner scans the internships to report TutorReminders
func ReportsScanner(cfg config.Internships, s schema.Internshipser, reporter ReminderReporter) error {
	reviews := make(map[string]time.Duration)
	tutors := make(map[string]TutorReminder)

	for _, r := range cfg.Reports {
		reviews[r.Kind] = r.Review.Duration
	}

	ints, err := s.Internships()
	if err != nil {
		return err
	}
	for _, i := range ints {
		em := i.Convention.Tutor.Person.Email
		tut, ok := tutors[em]
		if !ok {
			tut = TutorReminder{
				Tutor:    i.Convention.Tutor,
				Students: make([]StudentReminder, 0),
			}
		}
		rem := StudentReminder{
			Student:   i.Convention.Student,
			Reminders: make([]ReportReminder, 0),
		}
		rem.Reminders = append(rem.Reminders, missingReports(i.Reports, reviews)...)
		tut.Students = append(tut.Students, rem)
		tutors[em] = tut
	}
	for _, tut := range tutors {
		reporter(tut)
	}
	return nil
}

//Scan the reminders to report for a given set of reports
func missingReports(reports []schema.ReportHeader, reviews map[string]time.Duration) []ReportReminder {
	reminders := make([]ReportReminder, 0)
	now := time.Now()
	for _, r := range reports {
		if r.Delivery == nil && r.Deadline.Before(now) {
			//Late upload
			rem := ReportReminder{
				Kind:     r.Kind,
				Status:   "no delivery",
				Deadline: r.Deadline,
			}
			reminders = append(reminders, rem)
		} else if r.Delivery != nil && r.Reviewed == nil {
			//Missing review
			readyForReview := r.Deadline
			if r.Delivery.After(r.Deadline) {
				readyForReview = *r.Delivery
			}
			ps, _ := reviews[r.Kind]
			rem := ReportReminder{
				Kind:     r.Kind,
				Status:   "no review",
				Deadline: readyForReview.Add(ps),
			}
			reminders = append(reminders, rem)
		}
	}
	return reminders
}
