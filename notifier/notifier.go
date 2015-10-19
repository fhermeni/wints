package notifier

import (
	"log"

	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/mail"
	"github.com/fhermeni/wints/schema"
)

type Notifier struct {
	mailer mail.Mailer
	Log    journal.Journal
}

func New(m mail.Mailer, l journal.Journal) *Notifier {
	return &Notifier{mailer: m, Log: l}
}

func (n *Notifier) Println(s string, err error) {
	n.Log.Log("wintsd", s, err)
	if err == nil {
		log.Println(s + ": OK")
	} else {
		log.Println(s + ": " + err.Error())
	}
}

func (n *Notifier) Fatalln(s string, err error) {
	n.Log.Log("wintsd", s, err)
	if err == nil {
		log.Println(s + ": OK")
	} else {
		log.Fatalln(s + ": " + err.Error())
	}
}

/*func (n *Notifier) NewInternship(i schema.Internship, token []byte, err error) {
	//
	ed.notifier.InviteStudent(ex.s.Me(), i, string(token), err)

	//invite the student
	//Invitation mail
	//notify the tutor

	//log
}*/
func (n *Notifier) AccountReseted(em string, token []byte, err error) error {
	//mail with token
	//log
	n.Log.Log("wintsd", "start password reset for "+em, err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "reset.txt", string(token))
}

func (n *Notifier) ReportUploaded() {
	//mail tutor, cc student
}

func (n *Notifier) ReportReviewed() {
	//mail student, cc tutor
}

func (n *Notifier) ReportDeadlineChanged() {
	//mail student, cc tutor
}

func (n *Notifier) Login(s schema.Session, err error) {
	n.Log.Log(s.Email, "login", err)
}

func (n *Notifier) Logout(s schema.User, err error) {
	n.Log.UserLog(s, "logout", err)
}

func (n *Notifier) PrivilegeUpdated(from schema.User, em string, p schema.Privilege, err error) error {
	n.Log.UserLog(from, p.String()+" privileges for "+em, err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "privileges.txt", p, from.Person)
}

func (n *Notifier) SurveyUploaded() {
	//mail tutor, cc supervisor
}

func (n *Notifier) SurveyReseted() {
	//mail supervisor, cc tutor
}

func (n *Notifier) RmAccount(from schema.User, em string, err error) error {
	n.Log.UserLog(from, "delete account '"+em+"'", err)
	if err == nil {
		return n.mailer.Send(schema.Person{Email: em}, "delete.txt", nil, from.Person)
	}
	return err
}

func (n *Notifier) InviteStudent(from schema.User, i schema.Internship, token string, err error) error {
	err = n.invite(from, i.Convention.Student.User.Person, token, "student_welcome.txt", err)
	if err == nil {
		err = n.mailer.Send(i.Convention.Tutor.Person, "tutor.txt", i.Convention.Student)
		n.Log.UserLog(from, "Notify tutor ", err)
	}
	return err
}

func (n *Notifier) invite(from schema.User, p schema.Person, token string, tpl string, err error) error {
	n.Log.UserLog(from, p.Email+" invited", err)
	if err != nil {
		return err
	}
	data := struct {
		Login string
		Token string
	}{Login: p.Email, Token: string(token)}
	err = n.mailer.Send(p, tpl, data)
	n.Log.UserLog(from, "invitation sent", err)
	return err
}

func (n *Notifier) PasswordChanged(em string, err error) error {
	n.Log.Log(em, "password changed", err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "password-changed.txt", nil)
}

func (n *Notifier) ProfileEdited(from schema.User, p schema.Person, err error) {
	n.Log.UserLog(from, "profile updated for "+p.Email+" ("+p.Fullname()+")", err)
}

func (n *Notifier) InviteTeacher(from schema.User, p schema.Person, token string, err error) error {
	return n.invite(from, p, token, "tutor_welcome.txt", err)
}

func (n *Notifier) InviteRoot(from schema.User, p schema.Person, token string, err error) error {
	return n.invite(from, p, token, "root_welcome.txt", err)
}

func (n *Notifier) NewStudent(from schema.User, st schema.Student, err error) {
	n.Log.UserLog(from, "new student "+st.User.Fullname()+"("+st.User.Person.Email+") "+st.Major+"/"+st.Promotion, err)
}
