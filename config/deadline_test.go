package config

import (
	"testing"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	relative = `deadline = "1440h"`
	absolute = `deadline = "01/09/2015"`
)

type D struct {
	Deadline Deadline
}

func TestRelativeDeadline(t *testing.T) {
	var d D
	if _, err := toml.Decode(relative, &d); err != nil {
		t.Error(err.Error())
	}
	expected, _ := time.ParseDuration("1440h")
	if *d.Deadline.relative != expected {
		t.Errorf("Got %s but expected %s\n", *d.Deadline.relative, expected)
	}
	if d.Deadline.absolute != nil {
		t.Error("absolute value should be nil")
	}
	now := time.Now()
	v := d.Deadline.Value(now)
	if v != now.Add(expected) {
		t.Errorf("Got %s but expected %s\n", v, now.Add(expected))
	}
}

func TestAbsoluteDeadline(t *testing.T) {
	var d D
	if _, err := toml.Decode(absolute, &d); err != nil {
		t.Error(err.Error())
	}
	expected, _ := time.Parse(DateLayout, "01/09/2015")
	if *d.Deadline.absolute != expected {
		t.Errorf("Got %s but expected %s\n", *d.Deadline.absolute, expected)
	}
	if d.Deadline.relative != nil {
		t.Error("absolute value should be nil")
	}
	v := d.Deadline.Value(time.Now())
	if v != expected {
		t.Errorf("Got %s but expected %s\n", v, expected)
	}
}
