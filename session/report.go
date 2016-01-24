package session

import (
	"time"

	"github.com/fhermeni/wints/schema"
)

//SetReportContent store a report.
//If the emitter is the current student, it is allowed if the deadline is not passed
//or if no report has already been uploaded
func (s *Session) SetReportContent(kind, email string, cnt []byte) (time.Time, error) {
	h, err := s.store.Report(kind, email)
	if err != nil {
		return time.Now(), err
	}
	if s.Myself(email) && (time.Now().Before(h.Deadline) || !h.In()) {
		return s.store.SetReportContent(kind, email, cnt)
	}
	return time.Now(), ErrPermission
}

//SetReportGrade set the report grade if the emitter is the student tutor or an admin at minimum
func (s *Session) SetReportGrade(kind, email string, g int, comment string) (time.Time, error) {
	if s.Tutoring(email) || s.Role().Level() >= schema.AdminLevel {
		return s.store.SetReportGrade(kind, email, g, comment)
	}
	return time.Time{}, ErrPermission
}

//SetReportDeadline changes the report deadline if the emitter tutors the student or is an admin at minimum
func (s *Session) SetReportDeadline(kind, email string, t time.Time) error {
	if s.Tutoring(email) || s.Role().Level() >= schema.AdminLevel {
		return s.store.SetReportDeadline(kind, email, t)
	}
	return ErrPermission
}

//SetReportPrivacy changes the report privacy status if the emitter tutors the student or is an admin at minimum
func (s *Session) SetReportPrivacy(kind, email string, p bool) error {
	if s.Tutoring(email) || s.Role().Level() >= schema.AdminLevel {
		return s.store.SetReportPrivacy(kind, email, p)
	}
	return ErrPermission
}

//Report returns a report header if the emitter if the student or if he is watching the student
func (s *Session) Report(kind, email string) (schema.ReportHeader, error) {
	if s.Myself(email) || s.Watching(email) {
		return s.store.Report(kind, email)
	}
	return schema.ReportHeader{}, ErrPermission
}

//ReportContent returns the report content if the emitter is the report owner, the tutor or a root.
//if the report is private, no one else can access the report. Otherwise, the major leader and the admins can access the reports
func (s *Session) ReportContent(kind, email string) ([]byte, error) {
	h, err := s.store.Report(kind, email)
	if err != nil {
		return []byte{}, err
	}
	if s.Myself(email) ||
		s.Tutoring(email) ||
		s.Role().Level() >= schema.RootLevel {
		return s.store.ReportContent(kind, email)
	}
	if s.Role().Level() >= schema.MajorLevel && !h.Private {
		return s.store.ReportContent(kind, email)
	}
	return []byte{}, ErrPermission
}
