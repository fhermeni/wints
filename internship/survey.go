package internship

import "time"

//SurveyHeader provides the metadata associated to a student survey
type SurveyHeader struct {
	Kind string
	//The deadline to submit
	Deadline time.Time
	//The moment the survey was committed
	Delivery *time.Time `,json:"omitempty"`
	//The token to access in write mode
	Token string
	//Supervisor answers
	Answers []byte
}
