package jobs

import (
	"time"

	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
)

//MissingReports scan the internships for reports or reviews that are missing
func MissingReports(provider schema.Internshipser, not notifier.NewsLetter) {
	ints, err := provider.Internships()
	logger.Log("event", "cron", "Scanning the internships for the tutor newsletters", err)
	if err != nil {
		return
	}
	byTutor := make(map[string][]schema.StudentReports)
	tutors := make(map[string]schema.User)
	for _, i := range ints {
		if i.Convention.Skip {
			continue
		}
		student := i.Convention.Student
		em := i.Convention.Tutor.Person.Email
		reports, ok := byTutor[em]
		if !ok {
			reports = make([]schema.StudentReports, 0)
			tutors[em] = i.Convention.Tutor
		}
		lates := lateReports(i.Reports)
		if len(lates) > 0 {
			x := schema.StudentReports{
				Student: student,
				Reports: lates,
			}
			reports = append(reports, x)
		}

		byTutor[em] = reports
	}

	for em, lates := range byTutor {
		if len(lates) > 0 {
			not.TutorNewsLetter(tutors[em], lates)
		}
	}
}

func lateReports(reports []schema.ReportHeader) []schema.ReportHeader {
	var res []schema.ReportHeader
	now := time.Now()
	for _, r := range reports {
		if r.Deadline.Before(now) && r.Reviewed == nil {
			//The deadline passed and it is not reviewed
			//So if not uploaded -> missing report
			//if uploaded -> missing review
			res = append(res, r)
		}
	}
	return res
}
