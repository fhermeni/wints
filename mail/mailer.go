package mail

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

type Mailer interface {
	SendAdminInvitation(u internship.User, token []byte)
	SendStudentInvitation(u internship.User, token []byte)
	SendPasswordResetLink(u internship.User, token []byte)
	SendProfileUpdate(u internship.User)
	SendAccountRemoval(u internship.User)
	SendRoleUpdate(u internship.User)
	SendTutorUpdate(s internship.User, old internship.User, now internship.User)
	SendReportUploaded(s internship.User, t internship.User, kind string)
	SendGradeUploaded(s internship.User, t internship.User, kind string)
	SendReportDeadline(s internship.User, t internship.User, kind string, d time.Time)
}
