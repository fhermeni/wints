//Package config aggregates the necessary material to configure the wints daemon
package config

import "github.com/fhermeni/wints/mail"

var (
	//DateTimeLayout expresses the expected format for a date + time
	DateTimeLayout = "02/01/2006 15:04"
	//DateLayout expresses the expected format for a date
	DateLayout = "02/01/2006"
)

type Spy struct {
	Kind string
	Cron string
}

//Feeder configures the feeder than scan conventions
type Feeder struct {
	Login      string
	Password   string
	URL        string
	Frequency  Duration
	Promotions []string
	Encoding   string
}

//Db configures the database connection string
type Db struct {
	ConnectionString string
}

//EndPoint configures the rest endpoints
type Rest struct {
	SessionLifeTime        Duration
	RenewalRequestLifetime Duration
	Prefix                 string
}

//HTTPd configures the http daemon
type HTTPd struct {
	WWW         string
	Assets      string
	Listen      string
	Certificate string
	PrivateKey  string
	Rest        Rest
}

//Journal configures the logging system
type Journal struct {
	Path string
}

//Internships declare the internship organization
type Internships struct {
	Majors     []string
	Promotions []string
	Reports    map[string]Report
	Surveys    map[string]Survey
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
func (i Internships) ValidMajor(m string) bool {
	return contains(i.Majors, m)
}

func (i Internships) ValidPromotion(p string) bool {
	return contains(i.Promotions, p)
}

//Report configures a report definition
type Report struct {
	Delivery Deadline
	Review   Duration
	Reminder Duration
	Grade    bool
}

//Survey configures a survey definition
type Survey struct {
	Deadline Deadline
}

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Spy         Spy
	Feeder      Feeder
	Db          Db
	Mailer      mail.Config
	HTTPd       HTTPd
	Majors      []string
	Journal     Journal
	Internships Internships
	Spies       []Spy
}
