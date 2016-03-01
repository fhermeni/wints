package notifier

import "github.com/fhermeni/wints/schema"

//NewsLetter describes the news to send to users
type NewsLetter interface {
	//TutorNewsLetter notifies a tutor about some late reports or reviews
	TutorNewsLetter(tut schema.User, reminders []schema.StudentReports)
}

type SupervisorLetter interface {
	//SurveyRequest sends the survey request to the supervisor
	SurveyRequest(sup schema.Person, tutor schema.User, student schema.Student, survey schema.SurveyHeader, err error)
}
