//Package config aggregates the necessary material to configure the wints daemon
package config

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

//Mailer configure the mailing service
type Mailer struct {
	Server   string
	Login    string
	Password string
	Sender   string
	Path     string
}

//Db configures the database connection string
type Db struct {
	ConnectionString string
}

//EndPoint configures the rest endpoint
type EndPoint struct {
	WWW                    string
	Listen                 string
	Certificate            string
	PrivateKey             string
	Assets                 string
	SessionLifeTime        Duration
	RenewalRequestLifetime Duration
}

//Journal configures the logging system
type Journal struct {
	Path string
}

//Report configures a report definition
type Report struct {
	Deadline Deadline
	Grade    bool
}

//Survey configures a survey definition
type Survey struct {
	Deadline Deadline
}

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Feeder   Feeder
	Reports  map[string]Report
	Surveys  map[string]Survey
	Db       Db
	Mailer   Mailer
	EndPoint EndPoint
	Majors   []string
	Journal  Journal
}
