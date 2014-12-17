package config

import (
	"encoding/json"
	"io/ioutil"
)

type PullerConfig struct {
	Login    string
	Password string
	Url      string
	Period   string
}

type MailerConfig struct {
	Server   string
	Login    string
	Password string
	Sender   string
}

type DbConfig struct {
	Url string
}

type HttpConfig struct {
	Host        string
	Certificate string
	PrivateKey  string
}

type ReportsConfig struct {
	MidtermAfter string
}
type Config struct {
	Puller  PullerConfig
	Reports ReportsConfig
	Db      DbConfig
	Mailer  MailerConfig
	Http    HttpConfig
	Logfile string
}

func Load(path string) (Config, error) {
	var c Config
	cnt, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(cnt, &c)
	return c, err
}
