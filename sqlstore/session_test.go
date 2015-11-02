// +build integration

package sqlstore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

//TestSessions creates a session, get it and delete it
func TestSessions(t *testing.T) {

	//Preparation
	u := newTutor(t)
	em := u.Person.Email
	password := newPassword(t, em)

	//Try to create sessions

	//Unknown user
	_, err := store.NewSession(string(randomBytes(12)), randomBytes(12), time.Hour)
	assert.Equal(t, schema.ErrUnknownUser, err)

	//Incorrect password
	_, err = store.NewSession(em, randomBytes(12), time.Hour)
	assert.Equal(t, schema.ErrCredentials, err)

	//A good one
	s, err := store.NewSession(em, password, time.Hour)
	assert.Nil(t, err)
	assert.Equal(t, em, s.Email)
	assert.NotEmpty(t, s.Token)
	assert.NotNil(t, s.Expire)

	//Get the session
	s2, err := store.Session(s.Token)
	assert.Nil(t, err)
	assert.EqualValues(t, s, s2)
	//Get unknown session
	_, err = store.Session(randomBytes(12))
	assert.Equal(t, schema.ErrCredentials, err)

	//Delete an existing session
	assert.Nil(t, store.RmSession(s.Email))
	//Multiple delete fail
	assert.Equal(t, schema.ErrUnknownUser, store.RmSession(s.Email))
	//Delete an uknown session
	assert.Equal(t, schema.ErrUnknownUser, store.RmSession("lkj"))
}
