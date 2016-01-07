package config

import (
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fhermeni/wints/mail"
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

[[Internships.Reports]]
Kind = "dow"
Deadline= "504h"
Grade = false

[[Internships.Reports]]
Kind = "midterm"
Delivery = "1440h"
Review = "504h"
Grade = true

[[Internships.Reports]]
Kind = "final"
Delivery = "01/09/2015 15:00"
Review = "168h"
Grade = true

[[Internships.Surveys]]
Kind = "midterm"
Deadline = "168h"
Invitation = "1440h"

[[Internships.Surveys]]
Kind = "final"
Invitation = "01/09/2015 15:00"
Deadline = "168h"

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
		Mailer: mail.Config{
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
			Majors:      []string{"caspar", "ubinet", "al", "web", "gmd", "iam", "ihm", "imafa", "inum"},
			Surveys:     make([]Survey, 0),
			Reports:     make([]Report, 0),
			LatePenalty: 2,
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
	oneWeek, _ := time.ParseDuration("168h")
	sm := Survey{Kind: "midterm",
		Invitation: Deadline{relative: &rel},
		Deadline:   Duration{Duration: oneWeek},
	}

	abs, _ := time.Parse(DateLayout, "01/09/2015 15:00")
	sf := Survey{Kind: "final",
		Invitation: Deadline{
			absolute: &abs,
		},
		Deadline: Duration{Duration: oneWeek},
	}

	expected.Internships.Surveys = []Survey{sm, sf}

	rel, _ = time.ParseDuration("504h")
	rd := Report{
		Kind: "dow",
		Delivery: Deadline{
			relative: &rel,
		},
	}

	rel, _ = time.ParseDuration("1440h")
	rm := Report{
		Kind: "midterm",
		Delivery: Deadline{
			relative: &rel,
		},
	}

	abs, _ = time.Parse(DateLayout, "01/09/2015 15:00")
	rf := Report{
		Kind: "final",
		Delivery: Deadline{
			absolute: &abs,
		},
	}

	expected.Internships.Reports = []Report{rd, rm, rf}

	assert.Equal(t, expected.Mailer, cfg.Mailer)
	assert.Equal(t, expected.Feeder, cfg.Feeder)
	assert.Equal(t, expected.HTTPd, cfg.HTTPd)
	assert.Equal(t, expected.Db, cfg.Db)
	assert.Equal(t, 2, expected.Internships.LatePenalty)
}
