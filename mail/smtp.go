package mail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"os"

	"github.com/fhermeni/wints/schema"
)

//SMTP depicts a SMTP-based mailer
type SMTP struct {
	server   string
	sender   string
	path     string
	auth     smtp.Auth
	tls      tls.Config
	www      string
	fullname string
}

//Config configures the SMTP-based mailer
type Config struct {
	Server   string
	Login    string
	Password string
	Sender   string
	Path     string
	Fullname string
}

type metaData struct {
	WWW      string
	Fullname string
	Data     interface{}
}

//NewSMTP creates a new SMTP mailer
func NewSMTP(cfg Config, www string) (*SMTP, error) {
	hostname, _, err := net.SplitHostPort(cfg.Server)
	if err != nil {
		return &SMTP{}, err
	}
	m := SMTP{
		server:   cfg.Server,
		sender:   cfg.Sender,
		path:     cfg.Path,
		auth:     smtp.PlainAuth("", cfg.Login, cfg.Password, hostname),
		tls:      tls.Config{ServerName: hostname, InsecureSkipVerify: true},
		www:      www,
		fullname: cfg.Fullname,
	}
	return &m, err
}

//Send the mail to the smtp server
func (m *SMTP) Send(to schema.Person, tpl string, data interface{}, cc ...schema.Person) error {
	path := fmt.Sprintf("%s%c%s", m.path, os.PathSeparator, tpl)
	dta := metaData{
		WWW:      m.www,
		Fullname: m.fullname,
		Data:     data,
	}
	body, err := fill(path, dta)
	if err != nil {
		return err
	}
	return m.sendMail(to.Email, emails(cc...), body)
}

func (m *SMTP) sendMail(to string, cc []string, msg []byte) error {
	c, err := smtp.Dial(m.server)
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

	if err = c.Mail(m.sender); err != nil {
		return err
	}
	if err = c.Rcpt(to); err != nil {
		return err
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
