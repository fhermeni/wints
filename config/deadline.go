package config

import "time"

//UnmarshalText parse either duration or date
func (d *Deadline) UnmarshalText(text []byte) error {
	var err error
	str := string(text)
	/*if strings.HasSuffix(str, "d") {
		str = strings.TrimSuffix(str, "d")
		days, r := strconv.Atoi(str)
		if r == nil {
			d.relative = &time.Duration(time.Hour * time.Duration(24*days))
		}
	}*/

	rel, err := time.ParseDuration(str)
	if err != nil {
		var abs time.Time
		abs, err = time.Parse(DateLayout, str)
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
