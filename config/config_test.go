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
[EndPoint]
Listen = ":8080"
WWW = "https://localhost:8080"
Certificate = "chain-wints.polytech.unice.fr.pem"
PrivateKey = "wints.polytech.unice.fr.key"
Assets = "static"
SessionLifetime = "24h"
RenewalRequestLifetime = "48h"

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

[Reports.dow]
Deadline= "504h"
Grade = false

[Reports.midterm]
Deadline = "1440h"
Grade = true

[Reports.final]
Deadline = "01/09/2015"
Grade = true

[Surveys.midterm]
Deadline = "1440h"

[Surveys.final]
Deadline = "01/09/2015"
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
		EndPoint: EndPoint{
			Listen:      ":8080",
			WWW:         "https://localhost:8080",
			Certificate: "chain-wints.polytech.unice.fr.pem",
			PrivateKey:  "wints.polytech.unice.fr.key",
			Assets:      "static",
		},
		Surveys: map[string]Survey{
			"midterm": {},
			"final":   {},
		},
		Reports: map[string]Report{
			"dow":     {Grade: false},
			"midterm": {Grade: true},
			"final":   {Grade: true},
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
	expected.EndPoint.SessionLifeTime.Duration, _ = time.ParseDuration("24h")
	expected.EndPoint.RenewalRequestLifetime.Duration, _ = time.ParseDuration("48h")

	rel, _ := time.ParseDuration("1440h")
	sm := expected.Surveys["midterm"]
	sm.Deadline.relative = &rel
	expected.Surveys["midterm"] = sm

	abs, _ := time.Parse(DateLayout, "01/09/2015")
	sf := expected.Surveys["final"]
	sf.Deadline.absolute = &abs
	expected.Surveys["final"] = sf

	rel, _ = time.ParseDuration("504h")
	rd := expected.Reports["dow"]
	rd.Deadline.relative = &rel
	expected.Reports["dow"] = rd

	rel, _ = time.ParseDuration("1440h")
	rm := expected.Reports["midterm"]
	rm.Deadline.relative = &rel
	//expected.Reports["midterm"] = rm

	abs, _ = time.Parse(DateLayout, "01/09/2015")
	rf := expected.Reports["final"]
	rf.Deadline.absolute = &abs
	//expected.Reports["final"] = rf

	assert.Equal(t, expected.Mailer, cfg.Mailer)
	assert.Equal(t, expected.Feeder, cfg.Feeder)
	assert.Equal(t, expected.EndPoint, cfg.EndPoint)
	assert.Equal(t, expected.db, cfg.db)
}
