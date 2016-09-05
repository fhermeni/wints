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

//Anonymise removes the internship subject title
func (c *Company) Anonymise() {
	c.Title = ""
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

//Anonymise the company, the person
func (c *Convention) Anonymise() {
	c.Company.Anonymise()
	c.Student.User.Person.Anonymise()
	c.Student.Alumni = nil
	c.Tutor.Person.Anonymise()
	c.Supervisor.Anonymise()
}

//Internship is the core type to specify required data related to an internship
type Internship struct {
	Convention Convention
	//The headers for each
	Reports []ReportHeader
	//The surveys
	Surveys []SurveyHeader
	//Defense
	Defense Defense
}

//Anonymise the convention, the reports and the surveys
func (i *Internship) Anonymise() {
	i.Convention.Anonymise()
	for idx, hdr := range i.Reports {
		hdr.Anonymise()
		i.Reports[idx] = hdr
	}

	for idx, hdr := range i.Surveys {
		hdr.Anonymise()
		i.Surveys[idx] = hdr
	}
	i.Defense.Anonymise()
}

//Internships aliases slices of internship to exhibit filtering methods
type Internships []Internship

//Filter returns the internships that pass the given filters
func (ss Internships) Filter(filters ...func(Internship) bool) Internships {
	var res Internships
	for _, i := range ss {
		keep := false
		for _, f := range filters {
			keep = keep || f(i)
		}
		if keep {
			res = append(res, i)
		}
	}
	return res
}

//Anonymise every internships
func (ss Internships) Anonymise() {
	for idx, i := range ss {
		i.Anonymise()
		ss[idx] = i
	}
}

//Tutoring is a filter that keep only the internships tutored by the given user
func Tutoring(tut string) func(Internship) bool {
	return func(i Internship) bool {
		return i.Convention.Tutor.Person.Email == tut
	}
}

//InMajor is a filter that keep only the internships in the given major
func InMajor(major string) func(Internship) bool {
	return func(i Internship) bool {
		return i.Convention.Student.Major == major
	}
}
