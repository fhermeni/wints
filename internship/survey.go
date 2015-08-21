package internship

import (
	"errors"
	"time"
)

//SurveyDef allows to define internship reports
type SurveyDef struct {
	Name     string
	Deadline string
	Value    string
}

//ReportHeader allows to instantiate a report definition to a header
func (def *SurveyDef) Instantiate(from time.Time) (time.Time, error) {
	switch def.Deadline {
	case "relative":
		shift, err := time.ParseDuration(def.Value)
		if err != nil {
			return time.Now(), err
		}
		return from.Add(shift), nil
	case "absolute":
		at, err := time.Parse(DateLayout, def.Value)
		if err != nil {
			return time.Now(), err
		}
		return at, nil
	}
	return time.Now(), errors.New("Unsupported time definition '" + def.Deadline + "'")
}

//SurveyHeader provides the metadata associated to a student survey
type Survey struct {
	Kind string
	//The deadline to submit
	Deadline time.Time
	//The moment the survey was committed
	Delivery *time.Time `,json:"omitempty"`
	//The token to access in write mode
	Token string
	//Supervisor answers
	Answers map[string]string
}
