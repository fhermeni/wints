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
)

type SMTP struct {
	server   string
	username string
	password string
	sender   string
	path     string
	www      string
	auth     smtp.Auth
	tls      tls.Config
}

func NewSMTP(srv, login, passwd, emitter, www, path string) (*SMTP, error) {
	hostname, _, err := net.SplitHostPort(srv)
	if err != nil {
		return &SMTP{}, err
	}
	m := SMTP{
		server:   srv,
		username: login,
		password: passwd,
		sender:   emitter,
		path:     path,
		www:      www,
		auth:     smtp.PlainAuth("", login, passwd, hostname),
		tls:      tls.Config{ServerName: hostname, InsecureSkipVerify: true},
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
	return m.sendMail(m.server, m.sender, to, cc, body)
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
}

func join(mails ...string) []string {
	res := make([]string, 0, 0)
	for _, m := range mails {
		res = append(res, m)
	}
	return res
}
func one(s string) []string {
	return []string{s}
}
func none() []string {
	return []string{}
}
func (m *SMTP) SendAdminInvitation(u internship.User, token []byte) {
	d := InvitationData{WWW: m.www, Token: string(token)}
	if err := m.mail(join(u.Email), join(), m.path+"/tutor_welcome.txt", d); err != nil {
		log.Printf("Unable to send invitation to '%s': %s\n", u.Fullname(), err.Error())
	}
}

func (m *SMTP) SendStudentInvitation(u internship.User, token []byte) {
	d := InvitationData{WWW: m.www, Token: string(token)}
	if err := m.mail(join(u.Email), join(), m.path+"/student_welcome.txt", d); err != nil {
		log.Printf("Unable to send invitation to '%s': %s\n", u.Fullname(), err.Error())
	}
}

func (m *SMTP) SendPasswordResetLink(u internship.User, token []byte) {
	d := InvitationData{WWW: m.www, Token: string(token)}
	if err := m.mail(join(u.Email), join(), m.path+"/reset.txt", d); err != nil {
		log.Printf("Unable to send the reset link to '%s': %s\n", u.Fullname(), err.Error())
	}
}

func (m *SMTP) SendAccountRemoval(u internship.User) {
	empty := struct{}{}
	if err := m.mail(join(u.Email), join(), m.path+"/delete.txt", empty); err != nil {
		log.Printf("Unable to send the account removal mail to '%s': %s\n", u.Fullname(), err.Error())
	}
}

func (m *SMTP) SendRoleUpdate(u internship.User) {
	d := struct {
		WWW  string
		Role string
	}{
		m.www,
		u.Role.String(),
	}
	if err := m.mail(join(u.Email), join(), m.path+"/privileges.txt", d); err != nil {
		log.Printf("Unable to send the role update mail to '%s': %s\n", u.Fullname(), err.Error())
	}
}

func (m *SMTP) SendTutorUpdate(s internship.User, old internship.User, now internship.User) {
	d := struct {
		WWW   string
		Tutor internship.User
	}{
		m.www,
		now,
	}
	if err := m.mail(join(s.Email, now.Email), join(old.Email), m.path+"/tutor_switch.txt", d); err != nil {
		log.Printf("Unable to send the role update mail to '%s': %s\n", now.Fullname(), err.Error())
	}
}

func (m *SMTP) SendReportUploaded(s internship.User, t internship.User, kind string) {
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
	if err := m.mail(join(t.Email), join(s.Email), m.path+"/report_uploaded.txt", d); err != nil {
		log.Printf("Unable to send the report uploaded notificationl to '%s': %s\n", t.Fullname(), err.Error())
	}

}

func (m *SMTP) SendReportDeadline(s internship.User, t internship.User, kind string, deadline time.Time) {
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
	if err := m.mail(join(s.Email), join(t.Email), m.path+"/report_deadline.txt", d); err != nil {
		log.Printf("Unable to send the report deadline notification to '%s': %s\n", s.Fullname(), err.Error())
	}
}

func (m *SMTP) SendGradeUploaded(s internship.User, t internship.User, kind string) {
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
	if err := m.mail(join(s.Email), join(t.Email), m.path+"/report_graded.txt", d); err != nil {
		log.Printf("Unable to send the report graded notification to '%s': %s\n", s.Fullname(), err.Error())
	}
}

func (m *SMTP) SendReportPrivate(s internship.User, t internship.User, kind string, b bool) {
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
	if err := m.mail(join(s.Email), join(t.Email), m.path+"/report_status.txt", d); err != nil {
		log.Printf("Unable to send the report status update to '%s': %s\n", s.Fullname(), err.Error())
	}
}

func (m *SMTP) SendTest(s string) {
	if err := m.mail(join(s), join(), m.path+"/test.txt", struct{}{}); err != nil {
		log.Fatalf("Unable to send the test mail: %s\n", err.Error())
	}
}
