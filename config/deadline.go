package config

import "time"

//UnmarshalText parse either duration or date
func (d *Deadline) UnmarshalText(text []byte) error {
	var err error
	str := string(text)

	rel, err := time.ParseDuration(str)
	if err != nil {
		var abs time.Time
		abs, err = time.Parse(DateTimeLayout, str)
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

//AbsoluteDeadline returns a deadline from a timestamp
func AbsoluteDeadline(t time.Time) Deadline {
	return Deadline{absolute: &t}
}

//Value returns the concrete deadline depending on its type
func (d *Deadline) Value(from time.Time) time.Time {
	if d.absolute != nil {
		return *d.absolute
	}
	return from.Add(*d.relative)
}
