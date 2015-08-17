package internship

import "time"

//Stat gathers anonimzed statistics
type Stat struct {
	Male           bool
	Creation       time.Time
	Major          string
	Promotion      string
	Future         Alumni
	ForeignCountry bool
	Lab            bool
	Gratification  int
	Begin          time.Time
	End            time.Time
	Cpy            Company
	//Grade for each report
	Reports []ReportHeader
	//Answers for each survey
	Surveys []Survey
}
