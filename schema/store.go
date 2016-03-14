package schema

//Internshipser is a basic interface to browse internships
type Internshipser interface {
	Internships() (Internships, error)
}

type Anonymiser interface {
	Anonymise()
}
