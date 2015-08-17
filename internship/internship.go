package internship

import "time"

//Company is just a composite to store meaningful information for a company hosting a student.
type Company struct {
	//The company name
	Name string
	//The company website
	WWW string
}

type Convention struct {
	Male       bool
	Creation   time.Time
	Student    Student
	Tutor      User
	Supervisor User
	Cpy        Company
	Begin      time.Time
	End        time.Time
	Title      string
	//Ignore this convention
	Skip           bool
	Valid          bool
	ForeignCountry bool
	Lab            bool
	Gratification  int
}

//Internship is the core type to specify required data related to an internship
type Internship struct {
	C Convention
	//The headers for each
	Reports map[string]ReportHeader
	//The surveys
	Surveys map[string]Survey
	//Defense
	Defense DefenseSession
}
