package mail

import (
	"fmt"
	"os"

	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/schema"
)

//Fake to mock the mailing system
type Fake struct {
	Config Config
	WWW    string
	Logger journal.Logger
}

//Send just print the mailing on stdout
func (fake *Fake) Send(to schema.Person, tpl string, data interface{}, cc ...schema.Person) error {
	path := fmt.Sprintf("%s%c%s", fake.Config.Path, os.PathSeparator, tpl)
	dta := metaData{
		WWW:      fake.WWW,
		Fullname: fake.Config.Fullname,
		Data:     data,
	}
	body, err := fill(path, dta)
	if err != nil {
		fake.Logger.Log("mailer", "filling template '"+tpl+"'", err)
		return err
	}
	buf := fmt.Sprintf("-----\nFrom: %s\nTo: %s\nCC: %s\n%s\n-----\n", fake.Config.Sender, to.Email, emails(cc...), body)
	fake.Logger.Log("mailer", buf, nil)
	return nil
}
