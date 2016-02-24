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

//Notifier just embed notification mechanisms
type Notifier struct {
	mailer         mail.Mailer
	Log            journal.Journal
	reviewDuration map[string]time.Duration
}

//New creates a new notifier
func New(m mail.Mailer, l journal.Journal, cfg config.Config) *Notifier {
	reports := make(map[string]time.Duration)
	for _, r := range cfg.Internships.Reports {
		reports[r.Kind] = r.Review.Duration
	}
	return &Notifier{mailer: m, Log: l, reviewDuration: reports}
}

//Println logs a message with an error if needed
func (n *Notifier) Println(s string, err error) {
	n.Log.Log("wintsd", s, err)
	if err == nil {
		log.Println(s + ": OK")
	} else {
		log.Println(s + ": " + err.Error())
	}
}

//ReportPrivacyUpdated logs a privacy status change for a given report
func (n *Notifier) ReportPrivacyUpdated(from schema.User, stu, kind string, priv bool, err error) error {
	buf := fmt.Sprintf("set privacy status for '%s' report of '%s' to '%t'", kind, stu, priv)
	n.Log.UserLog(from, buf, err)
	return err
}

//ReportDeadlineUpdated logs a new deadline for a report
func (n *Notifier) ReportDeadlineUpdated(from schema.User, stu, kind string, d time.Time, err error) error {
	s := d.Format(config.DateTimeLayout)
	buf := fmt.Sprintf("set deadline for '%s' report of '%s' to '%s'", kind, stu, s)
	n.Log.UserLog(from, buf, err)
	return err
}

//Fatalln logs a fatal error and quit
func (n *Notifier) Fatalln(s string, err error) {
	n.Log.Log("wintsd", s, err)
	if err == nil {
		log.Println(s + ": OK")
	} else {
		log.Fatalln(s + ": " + err.Error())
	}
}

//AccountReseted logs an account reset and send the mail to the targeted user
func (n *Notifier) AccountReseted(em string, token []byte, err error) error {
	//mail with token
	//log
	n.Log.Log("wintsd", "start password reset for "+em, err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "reset.txt", string(token))
}

//ReportUploaded logs the upload and notify the student with the tutor in cc.
func (n *Notifier) ReportUploaded(student schema.User, tutor schema.User, kind string, err error) error {
	n.Log.UserLog(student, "Report "+kind+" uploaded", err)
	data := struct {
		Kind    string
		Student schema.Person
	}{Kind: kind, Student: student.Person}

	return n.mailer.Send(student.Person, "report_uploaded.txt", data, tutor.Person)
}

//ReportReviewed logs the action and notifies the student (tutor in cc.) by mail
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

//Login logs the login
func (n *Notifier) Login(s schema.Session, err error) {
	n.Log.Log(s.Email, "login", err)
}

//Logout logs the logout
func (n *Notifier) Logout(s schema.User, err error) {
	n.Log.UserLog(s, "logout", err)
}

//PrivilegeUpdated logs the action and notify the target
func (n *Notifier) PrivilegeUpdated(from schema.User, em string, p schema.Role, err error) error {
	n.Log.UserLog(from, "'"+string(p)+"' privileges for "+em, err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "privileges.txt", p, from.Person)
}

//SurveyUploaded logs the action and notify the tutor
func (n *Notifier) SurveyUploaded(student, tutor schema.User, supervisor schema.Person, kind string, err error) error {
	data := struct {
		Student schema.Person
		Kind    string
	}{Student: student.Person, Kind: kind}
	n.Log.UserLog(schema.User{Person: supervisor}, "'"+kind+"' survey uploaded for student '"+student.Person.Email+"'", err)
	return n.mailer.Send(tutor.Person, "survey_uploaded.txt", data)
}

//SurveyReseted logs the action
func (n *Notifier) SurveyReseted(from schema.User, student, kind string, err error) error {
	n.Log.UserLog(from, "'"+kind+"' survey content reseted for '"+student+"'", err)
	return err
}

//CompanyUpdated logs the change
func (n *Notifier) CompanyUpdated(from schema.User, c schema.Company, err error) {
	buf := fmt.Sprintf("company data updated '%+v'", c)
	n.Log.UserLog(from, buf, err)
}

//AlumniUpdated logs the change
func (n *Notifier) AlumniUpdated(from schema.User, a schema.Alumni, err error) {
	buf := fmt.Sprintf("alumni data updated '%+v'", a)
	n.Log.UserLog(from, buf, err)
}

//SupervisorUpdated logs the change
func (n *Notifier) SupervisorUpdated(from schema.User, sup schema.Person, err error) {
	buf := fmt.Sprintf("supervisor data updated '%+v'", sup)
	n.Log.UserLog(from, buf, err)
}

//RmAccount logs the target and notify it by mail
func (n *Notifier) RmAccount(from schema.User, em string, err error) error {
	n.Log.UserLog(from, "delete account '"+em+"'", err)
	if err == nil {
		return n.mailer.Send(schema.Person{Email: em}, "delete.txt", nil, from.Person)
	}
	return err
}

//InviteStudent invite() and mail the tutor in case of success
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

//PasswordChanged logs the event and notify the target
func (n *Notifier) PasswordChanged(em string, err error) error {
	n.Log.Log(em, "password changed", err)
	if err != nil {
		return err
	}
	return n.mailer.Send(schema.Person{Email: em}, "password-changed.txt", nil)
}

//ProfileEdited logs the change
func (n *Notifier) ProfileEdited(from schema.User, p schema.Person, err error) {
	n.Log.UserLog(from, "new profile for "+p.Email+" ("+p.Fullname()+")", err)
}

//InviteTeacher calls invite()
func (n *Notifier) InviteTeacher(from schema.User, p schema.Person, token string, err error) error {
	return n.invite(from, p, token, "tutor_welcome.txt", err)
}

//InviteRoot calls invite
func (n *Notifier) InviteRoot(from schema.User, p schema.Person, token string, err error) error {
	return n.invite(from, p, token, "root_welcome.txt", err)
}

//NewStudent logs the student addition
func (n *Notifier) NewStudent(from schema.User, st schema.Student, err error) {
	n.Log.UserLog(from, "new student "+st.User.Fullname()+" ("+st.User.Person.Email+") "+st.Major+"/"+st.Promotion, err)
}

//SkipStudent logs the change
func (n *Notifier) SkipStudent(from schema.User, em string, skip bool, err error) {
	n.Log.UserLog(from, "set student "+em+" skip status to "+strconv.FormatBool(skip), err)
}

//NewTutor logs the change and mail the old, the new tutor and the emitter
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

//SurveyRequest sends the survey request to the supervisor
func (n *Notifier) SurveyRequest(sup schema.Person, tutor schema.User, student schema.Student, survey schema.SurveyHeader, err error) error {
	n.Log.Log("cron", "send invitation for survey '"+student.User.Person.Email+"/"+survey.Kind+"' to '"+sup.Email+"'", err)
	frMonths := []string{"", "Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Decembre"}
	if err == nil {
		dta := struct {
			Student schema.User
			Survey  schema.SurveyHeader
			FrDate  string
			EnDate  string
		}{Student: student.User,
			Survey: survey,
			FrDate: fmt.Sprintf("%d %s %d", survey.Deadline.Day(), frMonths[survey.Deadline.Month()], survey.Deadline.Year()),
			EnDate: survey.Deadline.Format("Mon, 02 Jan 2006"),
		}
		return n.mailer.Send(sup, "survey_request.txt", dta, tutor.Person)
	}
	return err
}

type studentStatus struct {
	Student schema.User
	Status  []string
}

//TutorNewsLetter send to a tutor a mail about the missing reports or reviews
func (n *Notifier) TutorNewsLetter(tut schema.User, lates []schema.StudentReports) {
	now := time.Now()
	var statuses []studentStatus

	for _, stu := range lates {
		s := studentStatus{Student: stu.Student.User, Status: make([]string, 0)}
		for _, r := range stu.Reports {
			//missing deadlines
			if r.Delivery == nil {
				s.Status = append(s.Status, r.Kind+" deadline passed; was "+r.Deadline.Format(config.DateLayout))
			}
			//missing reviews
			if r.Deadline.Before(now) && r.Reviewed == nil {
				deadline := r.Deadline.Add(n.reviewDuration[r.Kind])
				s.Status = append(s.Status, r.Kind+" waiting for review (deadline: "+deadline.Format(config.DateLayout)+")")
			}
		}
		statuses = append(statuses, s)
	}
	buf := fmt.Sprintf("send the weekly report to '%s' with %d warning(s)\n", tut.Fullname(), len(statuses))
	err := n.mailer.Send(tut.Person, "tutor_newsletter.txt", statuses)
	n.Log.Log("cron", buf, err)
}
