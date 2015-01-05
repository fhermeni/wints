package mail

import (
	"log"
	"time"

	"github.com/fhermeni/wints/internship"
)

type Mock struct{}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) SendAdminInvitation(u internship.User, token []byte) {
	log.Printf("Invitation for '%s'. Token is %s\n", u.Fullname(), string(token))
}

func (m *Mock) SendStudentInvitation(u internship.User, token []byte) {
	log.Printf("Invitation student '%s'. Token: %s\n", u.Fullname(), string(token))
}

func (m *Mock) SendPasswordResetLink(u internship.User, token []byte) {
	log.Printf("Reset token for  %s: %s\n", u.Fullname(), string(token))
}

func (m *Mock) SendAccountRemoval(u internship.User) {
	log.Printf("Account for %s deleted\n", u.Fullname())
}

func (m *Mock) SendRoleUpdate(u internship.User) {
	log.Printf("New privileges for %s: %s\n", u.Fullname(), u.Role.String())
}

func (m *Mock) SendTutorUpdate(s internship.User, old internship.User, now internship.User) {
	log.Printf("New tutor for %s: Was %s, now %s\n", s.Fullname(), old.Fullname(), now.Fullname())
}

func (m *Mock) SendReportUploaded(s internship.User, t internship.User, kind string) {
	log.Printf("Hey %s, your student %s uploaded the report '%s'\n", t.Fullname(), s.Fullname(), kind)
}

func (m *Mock) SendReportDeadline(s internship.User, t internship.User, kind string, d time.Time) {
	log.Printf("Hey %s, your tutor %s changed the deadline for report '%s' to %s\n", s.Fullname(), t.Fullname(), kind, d.Format("02/01/2006"))
}

func (m *Mock) SendGradeUploaded(s internship.User, t internship.User, kind string) {
	log.Printf("Hey %s, your tutor %s uploaded the grade for the '%s' report\n", s.Fullname(), t.Fullname(), kind)
}

func (m *Mock) SendReportPrivate(s internship.User, t internship.User, kind string, b bool) {
	if b {
		log.Printf("Hey %s, your report '%s' is now private", s.Fullname(), kind)
	} else {
		log.Printf("Hey %s, your report '%s' is now public", s.Fullname(), kind)
	}
}

func (m *Mock) SendTest(e string) {
	log.Println("Test mail to %s\n", e)
}
