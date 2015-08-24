package config

import "time"

//Duration allows to parse duration expressed using toml
type Duration struct {
	time.Duration
}

//UnmarshalText parse durations
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
