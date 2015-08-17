package internship

import "time"

//Defense
type Defense struct {
	Student      User
	Major        string
	Promotion    string
	Cpy          Company
	Title        string
	Private      bool
	Remote       bool
	Grade        int
	Offset       int
	Surveys      []Survey
	NextPosition int
}

type DefenseSession struct {
	Date     time.Time
	Room     string
	Juries   []User
	Defenses []Defense
}

func (s DefenseSession) InJury(em string) bool {
	for _, j := range s.Juries {
		if j.Email == em {
			return true
		}
	}
	return false
}
