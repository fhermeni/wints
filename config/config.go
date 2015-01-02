//Package config aggregate the necessary material to configure the wints daemon
package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fhermeni/wints/internship"
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

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Puller  PullerConfig
	Reports []internship.ReportDef
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
