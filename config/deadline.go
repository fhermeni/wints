package config

import "time"

//UnmarshalText parse either duration or date
func (d *Deadline) UnmarshalText(text []byte) error {
	var err error
	rel, err := time.ParseDuration(string(text))
	if err != nil {
		var abs time.Time
		abs, err = time.Parse(DateLayout, string(text))
		if err == nil {
			d.absolute = &abs
		}
	} else {
		d.relative = &rel
	}
	return err
}

//Deadline is a simple wrapper to express either durations or date
type Deadline struct {
	relative *time.Duration
	absolute *time.Time
}

//Value returns the concrete deadline depending on its type
func (d *Deadline) Value(from time.Time) time.Time {
	if d.absolute != nil {
		return *d.absolute
	}
	return from.Add(*d.relative)
}
