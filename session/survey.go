package session

import (
	"github.com/fhermeni/wints/schema"
)

//Survey get the given survey for a student if the emitter is the student tutor or a major admin at least
func (s *Session) Survey(student, kind string) (schema.SurveyHeader, error) {
	if s.Tutoring(student) || s.Role().Level() >= schema.MAJOR_LEVEL {
		return s.store.Survey(student, kind)
	}
	return schema.SurveyHeader{}, ErrPermission
}
