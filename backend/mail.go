package backend

import (
	"text/template"
	"net/smtp"
	"bytes"
	"crypto/tls"
	"net"
	"log"
	"io/ioutil"
)

func Mail(cfg Config, to []string, input string, data interface{}) error {

	var out []byte

	if data == nil {
		var err error
		out, err = ioutil.ReadFile(input)
		if err != nil {
			return err
		}
	} else {
		tpl, err := template.ParseFiles(input)
		var b bytes.Buffer
		if tpl.Execute(&b, data); err != nil {
			return err
		}
		out = b.Bytes()
	}

	hostname,_,err := net.SplitHostPort(cfg.SMTPServer)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, hostname)
	tls := &tls.Config{ServerName: hostname, InsecureSkipVerify: true}
	go func() {
		err := SendMail(cfg.SMTPServer,
			auth,
			tls,
			cfg.SMTPSender, to,
			out)
		if err != nil {
			log.Printf("Unable to send a mail: %s\n", err.Error())
		}
	}()
	return err
}


func SendMail(addr string, a smtp.Auth, tlsCfg *tls.Config, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if ok, _ := c.Extension("STARTTLS"); ok {
		if err = c.StartTLS(tlsCfg); err != nil {
			return err
		}
	}
	if err = c.Auth(a); err != nil {
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
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
