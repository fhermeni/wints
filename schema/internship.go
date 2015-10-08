package schema

import "time"

//Company stores meaningful information for a company hosting a student.
type Company struct {
	//The company name
	Name string
	//The company website
	WWW string
	//The job title
	Title string
}

//Convention declares a student convention
type Convention struct {
	Creation   time.Time
	Student    Student
	Tutor      User
	Supervisor Person
	Company    Company
	Begin      time.Time
	End        time.Time
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

//Internships aliases slices of internship to exhibit filtering methods
type Internships []Internship

//Filter returns the internships that pass the given filter
func (ss Internships) Filter(filter func(Internship) bool) Internships {
	res := make(Internships, 0)
	for _, i := range ss {
		if filter(i) {
			res = append(res, i)
		}
	}
	return res
}

//Tutoring is a filter that keep only the internships tutored by the given user
func Tutoring(tut string) func(Internship) bool {
	return func(i Internship) bool {
		return i.Convention.Tutor.Person.Email == tut
	}
}
