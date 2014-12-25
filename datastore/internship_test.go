package datastore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestInternshipWorkflow(t *testing.T) {
	db := getDB(t)
	defer db.Close()
	s, err := NewService(db)
	assert.NoError(t, err)
	assert.NoError(t, s.Clean())
	assert.NoError(t, s.Install())

	x, err := s.Internships()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(x))
	_, err = s.Internship("foo")
	assert.Equal(t, internship.ErrUnknownInternship, err)
	assert.Equal(t, internship.ErrUnknownInternship, s.SetCompany("foo", internship.Company{}))
	assert.Equal(t, internship.ErrUnknownInternship, s.SetSupervisor("foo", internship.Supervisor{}))

	//Create an internship
	stu := internship.User{Firstname: "fn", Lastname: "ln", Tel: "tel", Email: "email", Role: internship.STUDENT}
	sup := internship.Supervisor{Firstname: "supfn", Lastname: "supln", Tel: "suptel", Email: "supEmail"}
	tut := internship.User{Firstname: "tutfn", Lastname: "tutln", Tel: "tuttel", Email: "tutEmail"}
	c := internship.Company{Name: "cname", WWW: "www"}
	from := time.Now()
	to := time.Now().Add(1000)
	err = s.NewInternship(stu.Email, tut.Email, from, to, c, sup, "Title")
	assert.Equal(t, internship.ErrUnknownUser, err)
	s.NewUser(stu)
	err = s.NewInternship(stu.Email, tut.Email, from, to, c, sup, "Title")
	assert.Equal(t, internship.ErrUnknownUser, err) //the tutor
	s.NewUser(tut)
	err = s.NewInternship(stu.Email, tut.Email, from, to, c, sup, "Title")
	assert.NoError(t, err)
	//Invalid dates

	err = s.NewInternship(stu.Email, tut.Email, from, to, c, sup, "Title")
	assert.Equal(t, internship.ErrInternshipExists, err)

	i, err := s.Internship(stu.Email)
	assert.NoError(t, err)
	assert.Equal(t, stu, i.Student)
	assert.Equal(t, tut, i.Tutor)
	assert.Equal(t, sup, i.Sup)
	assert.Equal(t, c, i.Cpy)
	assert.Equal(t, "Title", i.Title)
	assert.Equal(t, from, i.Begin)
	assert.Equal(t, to, i.End)
	assert.Equal(t, 0, len(i.Reports))
	is, err := s.Internships()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(is))
	assert.Equal(t, i, is[0])

	//setters
	c.Name = "fa"
	assert.NoError(t, s.SetCompany(stu.Email, c))
	i, err = s.Internship(stu.Email)
	assert.NoError(t, err)
	assert.Equal(t, c, i.Cpy)

	sup.Firstname = "bar"
	assert.NoError(t, s.SetSupervisor(stu.Email, sup))
	i, err = s.Internship(stu.Email)
	assert.NoError(t, err)
	assert.Equal(t, sup, i.Sup)

	//reports
}
