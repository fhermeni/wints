// +build integration

package sqlstore

import (
	"testing"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

func TestInternship(t *testing.T) {
	stu := newStudent(t)
	stuEm := stu.User.Person.Email
	tut := newTutor(t)
	//tutEm := tut.Person.Email
	i := newInternship(t, stu, tut)
	i2, err := store.Internship(stuEm)
	assert.Nil(t, err)
	assert.Equal(t, i, i2)

	//The setters
	unknown := string(randomBytes(12))
	cpy := schema.Company{}
	assert.Nil(t, store.SetCompany(stuEm, cpy))
	assert.Equal(t, schema.ErrUnknownInternship, store.SetCompany(unknown, cpy))

	sup := person()
	assert.Nil(t, store.SetSupervisor(stuEm, sup))
	assert.Equal(t, schema.ErrUnknownInternship, store.SetSupervisor(unknown, sup))

	tut2 := newTutor(t)
	assert.Nil(t, store.SetTutor(stuEm, tut2.Person.Email))
	assert.Equal(t, schema.ErrUnknownInternship, store.SetTutor(unknown, tut2.Person.Email))
	assert.Equal(t, schema.ErrUnknownUser, store.SetTutor(stuEm, unknown))

	i.Convention.Company = cpy
	i.Convention.Tutor = tut2
	i.Convention.Supervisor = sup

	i2, err = store.Internship(stuEm)
	assert.Nil(t, err)
	assert.Equal(t, i, i2)

	//Remove the tutor
	assert.Equal(t, schema.ErrUserTutoring, store.RmUser(tut2.Person.Email))
}
