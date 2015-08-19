type User struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
	Role Privilege
}

type Student struct {
	U User
	Major string
	Promotion string		
	Future         Alumni		
}

type Convention struct {
	Male       bool
	Creation   time.Time //always remember the last convention
	Student    Student
	Tutor      User
	Supervisor User
	Cpy        Company
	Begin      time.Time
	End        time.Time
	Title      string	
	//Ignore this convention
	Skip           bool
	ForeignCountry bool
	Lab            bool
	Gratification  int
}

type Internship struct {
	C Convention	
	//The headers for each
	Reports map[string]Report
	//The surveys
	Surveys map[string]Survey
	//Defense
	Defense DefenseSession
}

/*
user(email, firstname, lastname, tel, password, role)
student(email, major, promotion, future)
convention(male, creation, student, tutor, supervisor, cpy, begin, end, title, foreigner, lab, gratification, ignore)
surveys()
reports()
defense()
sessions()
password_renewal()
*/

//Convention scanning
-> fullfill convention table with last convention versions only
	-> tutor email might not be referring to a user

	-> "skip" means I don't manage this internship

	-> pending convention is those without validated flag
		validated: tutor student aligned, reports & surveys created


