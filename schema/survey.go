package schema

import "time"

//SurveyHeader provides the metadata associated to a student survey
type SurveyHeader struct {
	Kind string
	//Invitation the moment where the request must be sent
	Invitation time.Time
	//LastInvitation the moment where the last request has been emitted
	LastInvitation time.Time
	//The deadline where the survey shall be uploaded
	Deadline time.Time
	//The moment the survey was committed
	Delivery *time.Time `,json:"omitempty"`
	//The token to access in write mode
	Token string
	//Survey content
	Cnt map[string]string
}

type Survey struct {
	En string
	Fr string

	Categories []Category
}

type Category struct {
	En string
	Fr string
	Q  []Question
}

type Question struct {
	En       string
	Fr       string
	Type     string
	Required bool
	IsMark   bool
}
