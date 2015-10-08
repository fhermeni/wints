package config

import (
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	toParse = `
[HTTPd]
Listen = ":8080"
WWW = "https://localhost:8080"
Certificate = "chain-wints.polytech.unice.fr.pem"
PrivateKey = "wints.polytech.unice.fr.key"
Assets = "static"

[HTTPd.Rest]
SessionLifetime = "24h"
RenewalRequestLifetime = "48h"
Prefix = "api/v2"

[DB]
ConnectionString = "user=fhermeni dbname=wints host=localhost sslmode=disable"

[Feeder]
Login = "loginFeeder"
Password = "passwordFeeder"
URL = "http://conventions.polytech.unice.fr/admin/admin.cgi"
Encoding = "windows-1252"
Frequency = "1h"
Promotions = ["SI 5", "Master IFI", "Master IMAFA", "MAM 5", "Master 2 SSTIM-images"]

[Mailer]
Server = "smtp.inria.fr:507"
Login = "fooMailer"
Password = "barMailer"
Sender = "myMail"
Path = "assets/mails"

[Internships]
Majors = ["caspar","ubinet","al","web","gmd","iam","ihm","imafa","inum"]

[Internships.Reports.dow]
Deadline= "504h"
Grade = false

[Internships.Reports.midterm]
Delivery = "1440h"
Review = "504h"
Grade = true

[Internships.Reports.final]
Delivery = "01/09/2015"
Review = "168h"
Grade = true

[Internships.Surveys.midterm]
Deadline = "1440h"

[Internships.Surveys.final]
Deadline = "01/09/2015"

[[Spies]]
Kind = "tutor"
Cron = "0 0 9 * * 1"

[[Spies]]
Kind = "student"
Cron = "0 0 9 * * 1"
`

	expected = Config{
		Feeder: Feeder{
			Login:      "loginFeeder",
			Password:   "passwordFeeder",
			URL:        "http://conventions.polytech.unice.fr/admin/admin.cgi",
			Encoding:   "windows-1252",
			Promotions: []string{"SI 5", "Master IFI", "Master IMAFA", "MAM 5", "Master 2 SSTIM-images"},
		},
		Mailer: Mailer{
			Server:   "smtp.inria.fr:507",
			Login:    "fooMailer",
			Password: "barMailer",
			Sender:   "myMail",
			Path:     "assets/mails",
		},
		Db: Db{
			ConnectionString: "user=fhermeni dbname=wints host=localhost sslmode=disable",
		},
		HTTPd: HTTPd{
			Listen:      ":8080",
			WWW:         "https://localhost:8080",
			Certificate: "chain-wints.polytech.unice.fr.pem",
			PrivateKey:  "wints.polytech.unice.fr.key",
			Assets:      "static",
			Rest: Rest{
				Prefix: "api/v2",
			},
		},
		Internships: Internships{
			Majors: []string{"caspar", "ubinet", "al", "web", "gmd", "iam", "ihm", "imafa", "inum"},
			Surveys: map[string]Survey{
				"midterm": {},
				"final":   {},
			},
			Reports: map[string]Report{
				"dow":     {Grade: false},
				"midterm": {Grade: true},
				"final":   {Grade: true},
			},
		},
	}
)

func TestLoadConfig(t *testing.T) {
	var cfg Config
	if _, err := toml.Decode(toParse, &cfg); err != nil {
		t.Error(err)
	}
	//Fullfill with only-parseable stuff
	expected.Feeder.Frequency.Duration, _ = time.ParseDuration("1h")
	expected.HTTPd.Rest.SessionLifeTime.Duration, _ = time.ParseDuration("24h")
	expected.HTTPd.Rest.RenewalRequestLifetime.Duration, _ = time.ParseDuration("48h")

	rel, _ := time.ParseDuration("1440h")
	sm := expected.Internships.Surveys["midterm"]
	sm.Deadline.relative = &rel
	expected.Internships.Surveys["midterm"] = sm

	abs, _ := time.Parse(DateLayout, "01/09/2015")
	sf := expected.Internships.Surveys["final"]
	sf.Deadline.absolute = &abs
	expected.Internships.Surveys["final"] = sf

	rel, _ = time.ParseDuration("504h")
	rd := expected.Internships.Reports["dow"]
	rd.Delivery.relative = &rel
	expected.Internships.Reports["dow"] = rd

	rel, _ = time.ParseDuration("1440h")
	rm := expected.Internships.Reports["midterm"]
	rm.Delivery.relative = &rel
	//expected.Reports["midterm"] = rm

	abs, _ = time.Parse(DateLayout, "01/09/2015")
	rf := expected.Internships.Reports["final"]
	rf.Delivery.absolute = &abs
	//expected.Reports["final"] = rf

	assert.Equal(t, expected.Mailer, cfg.Mailer)
	assert.Equal(t, expected.Feeder, cfg.Feeder)
	assert.Equal(t, expected.HTTPd, cfg.HTTPd)
	assert.Equal(t, expected.Db, cfg.Db)
}
