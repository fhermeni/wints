package internship

import "time"

//Company stores meaningful information for a company hosting a student.
type Company struct {
	//The company name
	Name string
	//The company website
	WWW string
}

//Convention declares a student convention
type Convention struct {
	Creation   time.Time
	Student    Student
	Tutor      User
	Supervisor Person
	Cpy        Company
	Begin      time.Time
	End        time.Time
	Title      string
	//Ignore this convention (don't managed by wints if true)
	Skip           bool
	Valid          bool
	ForeignCountry bool
	Lab            bool
	Gratification  int
}

//Internship is the core type to specify required data related to an internship
type Internship struct {
	Convention Convention
	//The headers for each
	Reports map[string]ReportHeader
	//The surveys
	Surveys map[string]SurveyHeader
	//Defense
	Defense Defense
}
