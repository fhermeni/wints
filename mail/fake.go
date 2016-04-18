package mail

import (
	"fmt"
	"os"

	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/schema"
)

//Fake to mock the mailing system
type Fake struct {
	Config Config
	WWW    string
}

//Send just print the mailing on stdout
func (fake *Fake) Send(to schema.Person, tpl string, data interface{}, cc ...schema.Person) error {
	path := fmt.Sprintf("%s%c%s", fake.Config.Path, os.PathSeparator, tpl)
	dta := metaData{
		WWW:      fake.WWW,
		To:       to,
		Fullname: fake.Config.Fullname,
		Data:     data,
	}
	body, err := fill(path, dta)
	if err != nil {
		logger.Log("mailer", tpl, "filling template", err)
		return err
	}
	buf := fmt.Sprintf("-----\nFrom: %s\nTo: %s\nCC: %s\n%s\n-----\n", fake.Config.Sender, to.Email, emails(cc...), body)
	logger.Log("mailer", tpl, buf, nil)
	return nil
}
