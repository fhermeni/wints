package notifier

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fhermeni/wints/config"
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

func (n *Notifier) ReportPrivacyUpdated(from schema.User, stu, kind string, priv bool, err error) error {
	buf := fmt.Sprintf("set privacy status for '%s' report of '%s' to '%t'", kind, stu, priv)
	n.Log.UserLog(from, buf, err)
	return err
}

func (n *Notifier) ReportDeadlineUpdated(from schema.User, stu, kind string, d time.Time, err error) error {
	s := d.Format(config.DateTimeLayout)
	buf := fmt.Sprintf("set deadline for '%s' report of '%s' to '%s'", kind, stu, s)
	n.Log.UserLog(from, buf, err)
	return err
}

func (n *Notifier) Fatalln(s string, err error) {
	n.Log.Log("wintsd", s, err)
	if err == nil {
		log.Println(s + ": OK")
	} else {
		log.Fatalln(s + ": " + err.Error())
	}
}

func (n *Notifier) AccountReseted(em string, token []byte, err error) error {
	//mail with token
	//log
	n.Log.Log("wintsd", "start password reset for "+em, err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "reset.txt", string(token))
}

func (n *Notifier) ReportUploaded(student schema.User, tutor schema.User, kind string, err error) error {
	n.Log.UserLog(student, "Report "+kind+" uploaded", err)
	data := struct {
		Kind    string
		Student schema.Person
	}{Kind: kind, Student: student.Person}

	return n.mailer.Send(student.Person, "report_uploaded.txt", data, tutor.Person)
}

func (n *Notifier) ReportReviewed(from, student, tutor schema.User, kind string, err error) error {
	//mail student, cc tutor
	//Report XX has been reviewed. Log to access the comments
	data := struct {
		Kind  string
		Tutor schema.User
	}{Kind: kind, Tutor: tutor}
	n.Log.UserLog(from, "report '"+kind+"' reviewed for '"+student.Person.Email+"'", err)
	if err != nil {
		return err
	}
	return n.mailer.Send(student.Person, "report_reviewed.txt", data, tutor.Person)
}

func (n *Notifier) Login(s schema.Session, err error) {
	n.Log.Log(s.Email, "login", err)
}

func (n *Notifier) Logout(s schema.User, err error) {
	n.Log.UserLog(s, "logout", err)
}

func (n *Notifier) PrivilegeUpdated(from schema.User, em string, p schema.Role, err error) error {
	n.Log.UserLog(from, "'"+string(p)+"' privileges for "+em, err)
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

func (n *Notifier) CompanyUpdated(from schema.User, c schema.Company, err error) {
	buf := fmt.Sprintf("company data updated '%+v'", c)
	n.Log.UserLog(from, buf, err)
}

func (n *Notifier) AlumniUpdated(from schema.User, a schema.Alumni, err error) {
	buf := fmt.Sprintf("alumni data updated '%+v'", a)
	n.Log.UserLog(from, buf, err)
}
func (n *Notifier) SupervisorUpdated(from schema.User, sup schema.Person, err error) {
	buf := fmt.Sprintf("supervisor data updated '%+v'", sup)
	n.Log.UserLog(from, buf, err)
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

func (n *Notifier) SkipStudent(from schema.User, em string, skip bool, err error) {
	n.Log.UserLog(from, "set student "+em+" skip status to "+strconv.FormatBool(skip), err)
}
func (n *Notifier) NewTutor(from schema.User, stu string, old, now schema.User, err error) error {
	n.Log.UserLog(from, stu+" was tutored by "+old.Person.Email+", now "+now.Person.Email, err)
	if err != nil {
		return err
	}
	dta := struct {
		Old schema.User
		Now schema.User
	}{Old: old, Now: now}
	//Mail the student, cc the tutors
	return n.mailer.Send(schema.Person{Email: stu}, "tutor_switch.txt", dta, old.Person, now.Person, from.Person)
}
