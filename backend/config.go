package backend

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	WWWConventionsLogin string
	WWWConventionsPassword string
	WWWConventionsURL string
	RefreshPeriod string
	MidtermReportDeadline string
	DBurl string
	Port int
	Logfile string
}

func ReadConfig(path string) (Config, error) {
	var c Config
 	cnt, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(cnt, &c)
	return c, err
}
