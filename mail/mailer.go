package mail

import (
	"bytes"
	"sync"

	"text/template"

	"github.com/fhermeni/wints/schema"
)

var (
	//NONE points a default person
	NONE = []schema.Person{}

	//cache the templates to speed up spaming
	cache = make(map[string]*template.Template)
	lock  = sync.Mutex{}
)

//Mailer is an interface to specify a mail must be send
type Mailer interface {
	//Send a mail to all the given person in to with cc'ing persons in cc.
	//The mail body is created using the given template and data
	Send(to schema.Person, tpl string, data interface{}, cc ...schema.Person) error
}

func getTemplate(path string) (*template.Template, error) {
	lock.Lock()
	defer lock.Unlock()
	tpl, ok := cache[path]
	var err error
	if !ok {
		tpl, err = template.ParseFiles(path)
		if err != nil {
			return tpl, err
		}
		cache[path] = tpl
	}
	return tpl, err
}

func fill(path string, data interface{}) ([]byte, error) {
	tpl, err := getTemplate(path)
	if err != nil {
		return []byte{}, err
	}
	var b bytes.Buffer
	if err := tpl.Execute(&b, data); err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}

func emails(persons ...schema.Person) []string {
	emails := make([]string, 0, 0)
	for _, p := range persons {
		emails = append(emails, p.Email)
	}
	return emails
}
