// +build integration

package sqlstore

import (
	"testing"

	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/testutil"
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
	assert.Equal(t, 1, len(i2.Surveys))
	c, err := store.Convention(stuEm)
	assert.Nil(t, err)
	assert.Equal(t, i.Convention, c)

	//The setters
	unknown := string(randomBytes(12))
	cpy := schema.Company{}
	assert.Nil(t, store.SetCompany(stuEm, cpy))
	assert.Equal(t, schema.ErrUnknownInternship, store.SetCompany(unknown, cpy))

	sup := testutil.Person()
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

func TestInternships(t *testing.T) {
	assert.Nil(t, store.Install())

	ints, err := store.Internships()
	assert.Nil(t, err)
	assert.Empty(t, ints)

	s1 := newStudent(t)
	s2 := newStudent(t)
	tut1 := newTutor(t)
	i1 := newInternship(t, s1, tut1)
	i2 := newInternship(t, s2, tut1)
	ints, err = store.Internships()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(ints))
	assert.Contains(t, ints, i1, i2)

	cc, err := store.Conventions()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(cc))
	assert.Contains(t, cc, i1.Convention, i2.Convention)
}

func BenchmarkInternships(t *testing.B) {
	for i := 0; i < 100; i++ {
		stu := newStudent(nil)
		tut := newTutor(nil)
		newInternship(nil, stu, tut)
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		ints, err := store.Internships()
		assert.Nil(t, err)
		for _, int := range ints {
			assert.Equal(t, 1, len(int.Surveys))
		}
	}
}
