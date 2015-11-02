// +build integration

package sqlstore

import (
	"testing"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

func TestStudentWorkflow(t *testing.T) {
	student := newStudent(t)
	em := student.User.Person.Email
	got, err := store.Student(em)
	assert.Nil(t, err)
	assert.Equal(t, student, got)
	assert.Equal(t, schema.STUDENT, student.User.Role)
	//Default values
	assert.Equal(t, "al", student.Major)
	assert.Equal(t, "si", student.Promotion)
	assert.Equal(t, true, student.Male)
	assert.Equal(t, false, student.Skip)
	assert.Nil(t, student.Alumni)

	//Setters
	assert.Equal(t, schema.ErrInvalidMajor, store.SetMajor(em, "moo"))
	assert.Equal(t, schema.ErrInvalidPromotion, store.SetPromotion(em, "moo"))
	assert.Nil(t, store.SetMajor(em, "ihm"))
	assert.Nil(t, store.SetPromotion(em, "master"))
	assert.Nil(t, store.SetStudentSkippable(em, !student.Skip))
	assert.Nil(t, store.SetMale(em, !student.Male))
	al := schema.Alumni{
		Contact:     "foo",
		Position:    "here",
		France:      true,
		Permanent:   true,
		SameCompany: true,
	}
	assert.Equal(t, schema.ErrInvalidEmail, store.SetAlumni(em, al))
	al.Contact = "foo@bar.com"
	assert.Nil(t, store.SetAlumni(em, al))
	student.Alumni = &al
	student.Major = "ihm"
	student.Promotion = "master"
	student.Male = !student.Male
	student.Skip = !student.Skip
	student2, err := store.Student(em)
	assert.Equal(t, student, student2)
	_, err = store.Students()
	assert.Nil(t, err)
}

func TestStudents(t *testing.T) {
	assert.Nil(t, store.Install())
	stus, err := store.Students()
	assert.Nil(t, err)
	assert.Empty(t, stus)

	s1 := newStudent(t)
	s2 := newStudent(t)
	newTutor(t)
	stus, err = store.Students()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(stus))
	assert.Contains(t, stus, s1, s2)
}
