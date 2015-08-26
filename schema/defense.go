package internship

import "time"

//Defense depicts the defense of a student
type Defense struct {
	Private bool
	Local   bool
	Grade   int
	Time    time.Time
	Room    string
}

//DefenseSession groups all the defenses that occurs during a session
type DefenseSession struct {
	Juries   []User
	Defenses map[string]Defense
}

//InJury check if a user is a part of a jury
func (s DefenseSession) InJury(em string) bool {
	for _, j := range s.Juries {
		if j.Person.Email == em {
			return true
		}
	}
	return false
}
