//Package config aggregates the necessary material to configure the wints daemon
package config

import "github.com/fhermeni/wints/mail"

var (
	//DateTimeLayout expresses the expected format for a date + time
	DateTimeLayout = "02/01/2006 15:04"
	//DateLayout expresses the expected format for a date
	DateLayout = "02/01/2006"
)

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

//Rest configures the rest service
type Rest struct {
	SessionLifeTime        Duration
	RenewalRequestLifetime Duration
	//The endpoints prefix
	Prefix string
}

//HTTPd configures the http daemon
type HTTPd struct {
	InsecureListen string
	WWW            string
	Assets         string
	Listen         string
	Certificate    string
	PrivateKey     string
	Rest           Rest
}

//Journal configures the logging system
type Journal struct {
	Path string
	Key  string
}

//Internships declare the internship organization
type Internships struct {
	Majors      []string
	Promotions  []string
	Reports     []Report
	Surveys     []Survey
	LatePenalty int
	Version     string
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

//ValidMajor tests if a given major is supported
func (i Internships) ValidMajor(m string) bool {
	return contains(i.Majors, m)
}

//ValidPromotion tests if a promotion is supported
func (i Internships) ValidPromotion(p string) bool {
	return contains(i.Promotions, p)
}

//Report configures a report definition
type Report struct {
	Kind     string
	Delivery Deadline
	Review   Duration
	Reminder Duration
	Grade    bool
}

//Survey configures a survey definition
type Survey struct {
	Invitation Deadline
	Deadline   Duration
	Kind       string
}

//Crons list the waiting periods for some period tasks
type Crons struct {
	//NewsLetters is the waiting time between two news letters
	NewsLetters string
	//Surveys is the waiting time between two scan for missing surveys
	Surveys string
	//Idles is the waiting time between two scan for missing student connections
	Idles string
}

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Feeder      Feeder
	Db          Db
	Mailer      mail.Config
	HTTPd       HTTPd
	Majors      []string
	Journal     Journal
	Internships Internships
	Crons       Crons
}
