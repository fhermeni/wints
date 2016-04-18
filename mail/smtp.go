package mail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"os"
	"strings"

	"github.com/fhermeni/wints/logger"
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
	To       schema.Person
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
		To:       to,
		WWW:      m.www,
		Fullname: m.fullname,
		Data:     data,
	}
	body, err := fill(path, dta)
	if err != nil {
		return err
	}
	if err == nil {
		err = m.sendMail(to.Email, emails(cc...), body)
	}
	buf := fmt.Sprintf("sending '%s' to %s (cc %s)", tpl, to.Email, emails(cc...))
	logger.Log("event", "mailer", buf, err)
	return err

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

	for _, other := range cc {
		if err = c.Rcpt(other); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	write := func(buf string) {
		if err != nil {
			return
		}
		_, err = w.Write([]byte(buf))
	}
	write("To: " + to + "\n")
	if len(cc) > 0 {
		write("Cc: ")
		write(strings.Join(cc, ","))
		write("\n")
	}
	write(string(msg))
	if err != nil {
		return nil
	}
	if err = w.Close(); err != nil {
		return err
	}
	return c.Quit()
}
