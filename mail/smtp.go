package mail

import (
	"bytes"
	"crypto/tls"
	"log"
	"net"
	"net/smtp"
	"text/template"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
)

type SMTP struct {
	server string
	sender string
	path   string
	www    string
	auth   smtp.Auth
	tls    tls.Config
	fake   bool
	j      journal.Journal
}

func NewSMTP(srv, login, passwd, emitter, www, path string, j journal.Journal) (*SMTP, error) {
	hostname, _, err := net.SplitHostPort(srv)
	if err != nil {
		return &SMTP{}, err
	}
	m := SMTP{
		server: srv,
		sender: emitter,
		path:   path,
		www:    www,
		auth:   smtp.PlainAuth("", login, passwd, hostname),
		tls:    tls.Config{ServerName: hostname, InsecureSkipVerify: true},
	}
	return &m, err
}

func fill(path string, data interface{}) ([]byte, error) {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		return []byte{}, err
	}
	var b bytes.Buffer
	if tpl.Execute(&b, data); err != nil {
		return []byte{}, err
	}
	return b.Bytes(), err
}

func (m *SMTP) mail(to []string, cc []string, input string, data interface{}) error {
	body, err := fill(input, data)
	if err != nil {
		return err
	}
	if m.fake {
		log.Printf("To: %s\n", to)
		for _, m := range cc {
			log.Printf("Cc: %s\n", m)
		}
		log.Println(string(body))
		return nil
	} else {
		return m.sendMail(m.server, m.sender, to, cc, body)
	}

}

func (m *SMTP) sendMail(addr string, from string, to []string, cc []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if ok, _ := c.Extension("STARTTLS"); ok {
		if err = c.StartTLS(&m.tls); err != nil {
			return err
		}
	}
	if err = c.Auth(m.auth); err != nil {
		return err
	}

	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	for _, em := range cc {
		w.Write([]byte("Cc: "))
		w.Write([]byte(em))
		w.Write([]byte("\n"))
	}
	if _, err = w.Write(msg); err != nil {
		return err
	}
	if err = w.Close(); err != nil {
		return err
	}
	return c.Quit()
}

type InvitationData struct {
	WWW   string
	Token string
	Login string
}

func join(mails ...string) []string {
	res := make([]string, 0, 0)
	for _, m := range mails {
		res = append(res, m)
	}
	return res
}

func (m *SMTP) log(msg string, err error) {
	m.j.Log("mailer", msg, err)
}
func (m *SMTP) SendAdminInvitation(u internship.User, token []byte) {
	go func() {
		d := InvitationData{WWW: m.www, Token: string(token), Login: u.Email}
		err := m.mail(join(u.Email), join(), m.path+"/tutor_welcome.txt", d)
		m.log("Sending invitation to '"+u.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendStudentInvitation(u internship.User, token []byte) {
	go func() {
		d := InvitationData{WWW: m.www, Token: string(token), Login: u.Email}
		err := m.mail(join(u.Email), join(), m.path+"/student_welcome.txt", d)
		m.log("Sending invitation to '"+u.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendPasswordResetLink(u internship.User, token []byte) {
	go func() {
		d := InvitationData{WWW: m.www, Token: string(token)}
		err := m.mail(join(u.Email), join(), m.path+"/reset.txt", d)
		m.log("Sending the account reset link to '"+u.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendAccountRemoval(u internship.User) {
	go func() {
		empty := struct{}{}
		err := m.mail(join(u.Email), join(), m.path+"/delete.txt", empty)
		m.log("Sending the account removal mail to '"+u.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendRoleUpdate(u internship.User) {
	go func() {
		d := struct {
			WWW  string
			Role string
		}{
			m.www,
			u.Role.String(),
		}
		err := m.mail(join(u.Email), join(), m.path+"/privileges.txt", d)
		m.log("Sending the role update mail to '"+u.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendTutorNotification(s internship.Person, t internship.Person) {
	go func() {
		d := struct {
			Student internship.Person
			Tutor   internship.Person
			WWW     string
		}{
			s,
			t,
			m.www,
		}
		err := m.mail(join(t.Email), join(), m.path+"/tutor.txt", d)
		m.log("Notifying the new student '"+s.Fullname()+"' to '"+t.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendTutorUpdate(s internship.User, old internship.User, now internship.User) {
	go func() {
		d := struct {
			WWW string
			Old internship.User
			New internship.User
		}{
			m.www,
			old,
			now,
		}
		err := m.mail(join(s.Email, now.Email), join(old.Email), m.path+"/tutor_switch.txt", d)
		m.log("Sending the tutor switch of '"+s.Fullname()+"' to '"+now.Fullname()+"'(new) and '"+old.Fullname()+"' (old)", err)
	}()
}

func (m *SMTP) SendReportUploaded(s internship.User, t internship.User, kind string) {
	go func() {
		d := struct {
			WWW     string
			Student internship.User
			Tutor   internship.User
			Kind    string
		}{
			m.www,
			s,
			t,
			kind,
		}
		err := m.mail(join(t.Email), join(s.Email), m.path+"/report_uploaded.txt", d)
		m.log("Sending the '"+kind+"' report uploaded notification to '"+t.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendReportDeadline(s internship.User, t internship.User, kind string, deadline time.Time) {
	go func() {
		d := struct {
			WWW     string
			Student internship.User
			Tutor   internship.User
			Kind    string
			Date    time.Time
		}{
			m.www,
			s,
			t,
			kind,
			deadline,
		}
		err := m.mail(join(s.Email), join(t.Email), m.path+"/report_deadline.txt", d)
		m.log("Send the '"+kind+"' report deadline notification to '"+s.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendGradeUploaded(s internship.User, t internship.User, kind string) {
	go func() {
		d := struct {
			WWW     string
			Student internship.User
			Tutor   internship.User
			Kind    string
		}{
			m.www,
			s,
			t,
			kind,
		}
		err := m.mail(join(s.Email), join(t.Email), m.path+"/report_graded.txt", d)
		m.log("Send the '"+kind+"' report graded notification to '"+s.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendReportPrivate(s internship.User, t internship.User, kind string, b bool) {
	go func() {
		d := struct {
			WWW     string
			Student internship.User
			Tutor   internship.User
			Kind    string
			Status  string
		}{Student: s, Tutor: t, Kind: kind, Status: "public"}
		if b {
			d.Status = "private"
		}
		err := m.mail(join(t.Email), join(), m.path+"/report_status.txt", d)
		m.log("Sending the '"+kind+"' report status update to '"+t.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendSurveyUploaded(tutor internship.User, student internship.User, kind string) {
	go func() {
		d := struct {
			WWW     string
			Student internship.User
			Tutor   internship.User
			Kind    string
		}{
			m.www,
			student,
			tutor,
			kind,
		}
		err := m.mail(join(tutor.Email), join(), m.path+"/survey_uploaded.txt", d)
		m.log("Sending the '"+kind+"' survey uploaded status update to '"+tutor.Fullname()+"'", err)
	}()
}

func (m *SMTP) SendTest(s string) error {
	return m.mail(join(s), join(s), m.path+"/test.txt", struct{}{})
}

func (m *SMTP) Fake(b bool) {
	m.fake = b
}
