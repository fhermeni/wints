package testutil

import (
	"crypto/rand"
	"time"

	"github.com/fhermeni/wints/schema"
)

//Person returns a random person
func Person() schema.Person {
	p := schema.Person{
		Firstname: string(randomBytes(12)),
		Lastname:  string(randomBytes(12)),
		Tel:       string(randomBytes(12)),
		Email:     string(randomBytes(12)) + "@fr",
	}
	return p
}

//Student returns a random student
func Student(major, promotion string) schema.Student {
	return schema.Student{
		User: schema.User{
			Person: Person(),
			Role:   schema.STUDENT,
		},
		Major:     major,
		Promotion: promotion,
		Skip:      false,
	}
}

//Tutor returns a random tutor
func Tutor() schema.User {
	return schema.User{
		Person: Person(),
		Role:   schema.TUTOR,
	}
}

//Internship returns an internship for a random student and supervisor
func Internship(tutor schema.User) schema.Internship {
	creation := time.Now()
	c := schema.Convention{
		Student:    Student("al", "si"),
		Tutor:      tutor,
		Supervisor: Person(),
		Creation:   creation,
		Begin:      creation.Add(time.Minute),
		End:        creation.Add(time.Hour),
		Company: schema.Company{
			Title: string(randomBytes(10)),
			WWW:   string(randomBytes(10)),
			Name:  string(randomBytes(10)),
		},
		//tons of default values
		Gratification:  10000,
		Lab:            false,
		ForeignCountry: true,
	}
	return schema.Internship{
		Convention: c,
		Reports:    make([]schema.ReportHeader, 0),
		Surveys:    make([]schema.SurveyHeader, 0),
	}
}

func randomBytes(s int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, s)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func convention(student schema.Student, tutor schema.User) schema.Convention {
	creation := time.Now().Truncate(time.Minute).UTC()
	return schema.Convention{
		Student:    student,
		Tutor:      tutor,
		Supervisor: Person(),
		Creation:   creation,
		Begin:      creation.Add(time.Minute),
		End:        creation.Add(time.Hour),
		Company: schema.Company{
			Title: string(randomBytes(10)),
			WWW:   string(randomBytes(10)),
			Name:  string(randomBytes(10)),
		},
		//tons of default values
		Gratification:  10000,
		Lab:            false,
		ForeignCountry: true,
		Skip:           false,
		Valid:          false,
	}
}

//ReportHeader returns a report
func ReportHeader(k string, deadline time.Time, others ...time.Time) schema.ReportHeader {
	hdr := schema.ReportHeader{
		Kind:     k,
		Deadline: deadline,
	}
	if len(others) >= 1 {
		hdr.Delivery = &others[0]
	}
	if len(others) == 2 {
		hdr.Reviewed = &others[1]
	}
	return hdr
}

//FakeInternshipser mocks Internshipser
type FakeInternshipser struct {
	Ints schema.Internships
}

//Internships returns all the intenrships in FakeInternshipser
func (f *FakeInternshipser) Internships() (schema.Internships, error) {
	return f.Ints, nil
}
