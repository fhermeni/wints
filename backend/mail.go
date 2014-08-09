package backend

import (
	"text/template"
	"net/smtp"
	"bytes"
	"crypto/tls"
	"net"
	"log"
)

func Mail(cfg Config, to string, input string, data interface{}) error {

	tpl, err := template.ParseFiles(input)
	var out bytes.Buffer

	if tpl.Execute(&out, data); err != nil {
		return err
	}
	hostname,_,err := net.SplitHostPort(cfg.SmtpServer)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", cfg.SmtpUsername, cfg.SmtpPassword, hostname)
	tls := tls.Config{ServerName: hostname, InsecureSkipVerify: true}
	go func() {
		err := SendMail(cfg.SmtpServer,
			auth,
			tls,
			cfg.SmtpSender, []string{to},
			out.Bytes())
		if err != nil {
			log.Printf("Unable to send a mail: %s\n", err.Error())
		}
	}()
	return err
}


func SendMail(addr string, a smtp.Auth, tlsCfg tls.Config, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if ok, _ := c.Extension("STARTTLS"); ok {
		if err = c.StartTLS(&tlsCfg); err != nil {
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
