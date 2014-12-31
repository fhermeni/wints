package mail

import (
	"github.com/fhermeni/wints/internship"
)

type Mailer interface {
	SendAdminInvitation(u internship.User) error
	SendStudentInvitation(u internship.User) error
	SendPasswordResetLink(u internship.User) error
	SendProfileUpdate(u internship.User) error
	SendAccountRemoval(u internship.User) error
}
