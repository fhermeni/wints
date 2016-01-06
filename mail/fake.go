package mail

import (
	"fmt"
	"log"
	"os"

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
		Fullname: fake.Config.Fullname,
		Data:     data,
	}
	body, err := fill(path, dta)
	if err != nil {
		return err
	}
	log.Printf("From: %s\n", fake.Config.Sender)
	log.Printf("To: %s\n", to.Email)
	log.Printf("Cc: %s\n", emails(cc...))
	log.Println(string(body))
	return nil
}
