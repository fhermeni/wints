// +build integration

package sqlstore

import (
	"testing"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestStudents(t *testing.T) {
	pStudent, _ := preparePerson(t)
	//log.Println(pStudent)
	store.ToStudent(pStudent.Email, "ma1", "p1", true)

	student, err := store.Student(pStudent.Email)
	assert.Nil(t, err)
	assert.Equal(t, user(t, pStudent.Email), student.User)
	assert.Equal(t, internship.STUDENT, student.User.Role)
	assert.Equal(t, "ma1", student.Major)
	assert.Equal(t, "p1", student.Promotion)
	assert.Equal(t, true, student.Male)
	assert.Equal(t, false, student.Skip) //default value
	assert.Equal(t, 0, student.Alumni.Position)
	assert.Equal(t, "", student.Alumni.Contact)

	//Setters
	assert.Nil(t, store.SetMajor(pStudent.Email, "moo"))
	assert.Nil(t, store.SetPromotion(pStudent.Email, "ppp"))
	assert.Nil(t, store.SetStudentSkippable(pStudent.Email, !student.Skip))
	assert.Nil(t, store.SetNextPosition(pStudent.Email, 3))
	assert.Nil(t, store.SetNextContact(pStudent.Email, "baz"))
	assert.Nil(t, store.SetMale(pStudent.Email, !student.Male))
	student2, err := store.Student(pStudent.Email)
	assert.Equal(t, "moo", student2.Major)
	assert.Equal(t, "ppp", student2.Promotion)
	assert.Equal(t, !student.Skip, student2.Skip)
	assert.Equal(t, !student.Male, student2.Male)
	assert.Equal(t, 3, student2.Alumni.Position)
	assert.Equal(t, "baz", student2.Alumni.Contact)
	_, err = store.Students()
	assert.Nil(t, err)
}
