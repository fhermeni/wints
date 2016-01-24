package session

import (
	"time"

	"github.com/fhermeni/wints/schema"
)

//Survey get the given survey for a student if the emitter is the student tutor or a major admin at least
func (s *Session) Survey(student, kind string) (schema.SurveyHeader, error) {
	if s.Tutoring(student) || s.Role().Level() >= schema.MajorLevel {
		return s.store.Survey(student, kind)
	}
	return schema.SurveyHeader{}, ErrPermission
}

//ResetSurvey if the emitter is at least an administrator
func (s *Session) ResetSurvey(student, kind string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.ResetSurveyContent(student, kind)
	}
	return ErrPermission
}

//SetSurveyInvitation is ok if the emitter is at least an administrator
func (s *Session) SetSurveyInvitation(student, kind string) (time.Time, error) {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.SetSurveyInvitation(student, kind)
	}
	return time.Time{}, ErrPermission
}
