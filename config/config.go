//Package config aggregate the necessary material to configure the wints daemon
package config

import (
	"encoding/json"
	"fmt"
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
	Encoding   string
}

//MailerConfig allows to configure the mailing service
type MailerConfig struct {
	Server   string
	Login    string
	Password string
	Sender   string
	Path     string
}

//DbConfig allows to configure the access to the wints database
type DbConfig struct {
	URL string
}

//HTTPConfig allows to configure the Http endpoint
type HTTPConfig struct {
	WWW         string
	Listen      string
	Certificate string
	PrivateKey  string
	Path        string
}

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Puller  PullerConfig
	Reports []internship.ReportDef
	Surveys []internship.SurveyDef
	DB      DbConfig
	Mailer  MailerConfig
	HTTP    HTTPConfig
	Logfile string
	Majors  []string
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

func Blank() error {
	puller := PullerConfig{Login: "****",
		Password:   "******",
		URL:        "http://conventions.polytech.unice.fr/admin/admin.cgi",
		Promotions: []string{"Master%20IFI", "Master%20IMAFA", "MAM%205", "SI%205", "Master%202%20SSTIM-images"},
		Period:     "1h",
		Encoding:   "windows-1252",
	}
	mailer := MailerConfig{Server: "*******",
		Login:    "******",
		Password: "*******",
		Sender:   "******",
		Path:     "static/mails",
	}
	db := DbConfig{URL: "user=******* dbname=wints host=localhost sslmode=disable"}
	http := HTTPConfig{Listen: ":8080",
		WWW:         "https://localhost:8080",
		Certificate: "cert.pem",
		PrivateKey:  "key.pem",
		Path:        "static",
	}
	reports := []internship.ReportDef{
		internship.ReportDef{Name: "DoW", Deadline: "relative", Value: "504h", ToGrade: false},
		internship.ReportDef{Name: "Midterm", Deadline: "relative", Value: "1440h", ToGrade: true},
		internship.ReportDef{Name: "Final", Deadline: "absolute", Value: "01/09/2015 23:59", ToGrade: true},
	}

	surveys := []internship.SurveyDef{
		internship.SurveyDef{Name: "midterm", Deadline: "relative", Value: "1440h"},
		internship.SurveyDef{Name: "final", Deadline: "absolute", Value: "01/09/2015 23:59"},
	}
	cfg := Config{
		Puller:  puller,
		Reports: reports,
		Surveys: surveys,
		DB:      db,
		Mailer:  mailer,
		HTTP:    http,
		Majors:  []string{},
	}
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)
	return err
}
