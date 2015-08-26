// +build integration

package sqlstore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func prepareStudent(t *testing.T, student internship.User) internship.Student {
	stu := internship.Student{
		User:      student,
		Major:     string(randomBytes(3)),
		Promotion: string(randomBytes(3)),
		Male:      true,
	}
	assert.Nil(t, store.ToStudent(student.Person.Email, stu.Major, stu.Promotion, stu.Male))
	return stu
}

func prepareConvention(t *testing.T, student internship.Student, tutor internship.User) internship.Convention {
	creation := time.Now()
	c := internship.Convention{
		Student: student,
		Tutor:   tutor,
		Supervisor: internship.Person{
			Firstname: string(randomBytes(10)),
			Lastname:  string(randomBytes(10)),
			Email:     string(randomBytes(10)),
			Tel:       string(randomBytes(10)),
		},
		Creation:       creation.Truncate(time.Minute),
		Begin:          creation.Add(time.Minute),
		End:            creation.Add(time.Hour),
		Title:          string(randomBytes(10)),
		Gratification:  10000,
		Lab:            false,
		ForeignCountry: true,
		Skip:           false, //default value
		Valid:          false, //default value
		Company: internship.Company{
			Name: string(randomBytes(10)),
			WWW:  string(randomBytes(10)),
		},
	}

	ok, err := store.NewConvention(student.User.Person.Email,
		c.Begin,
		c.End,
		tutor.Person.Email,
		c.Company,
		c.Supervisor,
		c.Title,
		c.Creation,
		c.ForeignCountry,
		c.Lab,
		c.Gratification)
	assert.Nil(t, err)
	assert.Equal(t, false, ok) //creation, not updated
	return c
}

func convention(t *testing.T, p internship.Person) internship.Convention {
	c2, err := store.Convention(p.Email)
	assert.Nil(t, err)
	return c2
}

func TestConvention(t *testing.T) {
	pStudent, _ := preparePerson(t)
	pTutor, _ := preparePerson(t)
	student := prepareStudent(t, user(t, pStudent.Email))
	tutor := user(t, pTutor.Email)
	prepareConvention(t, student, tutor)
	c, err := store.Convention(student.User.Person.Email)
	assert.Nil(t, err)

	//Fullfil with the default values
	//assert.Equal(t, c, c2)

	//The setters
	cpy := internship.Company{}
	title := string(randomBytes(10))
	sup := internship.Person{}
	pTut2, _ := preparePerson(t)
	tut2 := user(t, pTut2.Email)
	assert.Nil(t, store.SetCompany(pStudent.Email, cpy))
	assert.Nil(t, store.SetTitle(pStudent.Email, title))
	assert.Nil(t, store.SetSupervisor(pStudent.Email, sup))
	assert.Nil(t, store.SetConventionSkippable(pStudent.Email, !c.Skip))
	assert.Nil(t, store.SetTutor(pStudent.Email, pTut2.Email))
	c2 := convention(t, pStudent)
	assert.Equal(t, cpy, c2.Company)
	assert.Equal(t, title, c2.Title)
	assert.Equal(t, sup, c2.Supervisor)
	assert.Equal(t, !c.Skip, c2.Skip)
	assert.Equal(t, tut2, c2.Tutor)
}
