package internship

import "time"

//Defense
type Defense struct {
	Student User
	Private bool
	Remote  bool
	Grade   int
	Offset  int
}

type DefenseSession struct {
	Date     time.Time
	Room     string
	Juries   []User
	Defenses []Defense
}

func (s DefenseSession) InJury(em string) bool {
	for _, j := range s.Juries {
		if j.Person.Email == em {
			return true
		}
	}
	return false
}
