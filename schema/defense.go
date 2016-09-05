package schema

import "time"

//Defense depicts the defense of a student
type Defense struct {
	Room      string
	SessionId string
	Public    bool
	Local     bool
	Grade     int
	Time      time.Time
	Student   Student
	Company   Company
}

//DefenseSession groups all the defenses that occurs during a session
type DefenseSession struct {
	Juries   []User
	Defenses []Defense
	Id       string
	Room     string
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

func (d *Defense) Anonymise() {
	d.Grade = -1
	d.Student.User.Person.Email = ""
	d.Student.User.LastVisit = nil
	d.Room = ""
	d.SessionId = ""
	d.Public = true
	d.Local = true
	d.Time = time.Now()
	d.Student.Alumni = nil
}

func (s DefenseSession) Anonymise() {
	for idx, d := range s.Defenses {
		d.Anonymise()
		s.Defenses[idx] = d
	}
	for idx, j := range s.Juries {
		j.Person.Email = ""
		s.Juries[idx] = j
	}
}
