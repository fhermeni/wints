// +build integration

package sqlstore

import (
	"testing"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

func preparePerson(t *testing.T) (schema.Person, []byte) {
	//Preparation
	p := schema.Person{
		Firstname: string(randomBytes(12)),
		Lastname:  string(randomBytes(12)),
		Tel:       string(randomBytes(12)),
		Email:     string(randomBytes(12)),
	}
	password := randomBytes(32)
	assert.Nil(t, store.NewUser(p))
	resetToken, err := store.ResetPassword(p.Email, time.Hour)
	assert.Nil(t, err)
	em, err := store.NewPassword(resetToken, password)
	assert.Nil(t, err)
	assert.Equal(t, em, p.Email)
	return p, password
}

//TestSessions creates a session, get it and delete it
func TestSessions(t *testing.T) {

	//Preparation
	p, password := preparePerson(t)
	//New session
	_, err := store.NewSession(string(randomBytes(12)), randomBytes(12), time.Hour)
	assert.Equal(t, schema.ErrCredentials, err)

	s, err := store.NewSession(p.Email, password, time.Hour)
	assert.Nil(t, err)

	//Get the session
	/*s2, err := store.Session(s.Token)
	assert.Equal(t, s, s2)*/
	_, err = store.Session(randomBytes(12))
	assert.Equal(t, schema.ErrCredentials, err)

	//Delete the session
	assert.Nil(t, store.RmSession(s.Token))
	assert.Equal(t, schema.ErrUnknownUser, store.RmSession(s.Token))
}
