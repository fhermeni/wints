//Package config aggregate the necessary material to configure the wints daemon
package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/fhermeni/wints/internship"
)

const (
	//DateLayout indicates the date format
	DateLayout = "02/01/2006 15:04"
)

//PullerConfig allows to configure the puller that browse the convention database
type PullerConfig struct {
	Login      string
	Password   string
	URL        string
	Period     string
	Promotions []string
}

//MailerConfig allows to configure the mailing service
type MailerConfig struct {
	Server   string
	Login    string
	Password string
	Sender   string
}

//DbConfig allows to configure the access to the wints database
type DbConfig struct {
	URL string
}

//HTTPConfig allows to configure the Http endpoint
type HTTPConfig struct {
	Host        string
	Certificate string
	PrivateKey  string
}

//ReportsDef allows to define internship reports
type ReportsDef struct {
	Name     string
	Deadline string
	Value    string
}

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Puller  PullerConfig
	Reports []ReportsDef
	DB      DbConfig
	Mailer  MailerConfig
	HTTP    HTTPConfig
	Logfile string
}

//Load allows to parse a configuration for a JSON message
func Load(path string) (Config, error) {
	var c Config
	cnt, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(cnt, &c)
	return c, err
}

//ReportHeader allows to instantiate a report definition to a header
func ReportHeader(from time.Time, h ReportsDef) (internship.ReportHeader, error) {
	switch h.Deadline {
	case "relative":
		shift, err := time.ParseDuration(h.Value)
		if err != nil {
			return internship.ReportHeader{}, err
		}
		return internship.ReportHeader{h.Name, from.Add(shift), -1}, nil
	case "absolute":
		at, err := time.Parse(DateLayout, h.Value)
		if err != nil {
			return internship.ReportHeader{}, err
		}
		return internship.ReportHeader{h.Name, at, -1}, nil
	}
	return internship.ReportHeader{}, errors.New("Unsupported time definition '" + h.Deadline + "'")
}
