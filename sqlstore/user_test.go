// +build integration

package sqlstore

import (
	"testing"

	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/testutil"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var store *Store

func TestUserManagement(t *testing.T) {
	u := newTutor(t)

	//Test token generation
	_, err := store.ResetPassword(string(randomBytes(12)))
	assert.Equal(t, schema.ErrUnknownUser, err)
	token, err := store.ResetPassword(u.Person.Email)
	assert.Nil(t, err)
	password := randomBytes(32)

	//Validate the new password
	_, err = store.NewPassword(token, randomBytes(3))
	assert.Equal(t, schema.ErrPasswordTooShort, err)
	em, err := store.NewPassword(token, password)
	assert.Nil(t, err)
	assert.Equal(t, u.Person.Email, em)

	//Cannot change again, the token is consumed
	_, err = store.NewPassword(token, password)
	assert.Equal(t, schema.ErrNoPendingRequests, err)

	//With a bad token
	_, err = store.NewPassword(randomBytes(12), randomBytes(12))
	assert.Equal(t, schema.ErrNoPendingRequests, err)

	//User()
	u2, err := store.User(em)
	assert.Nil(t, err)
	assert.Equal(t, u, u2)
	assert.Nil(t, u2.LastVisit)
	_, err = store.User(string(randomBytes(10)))
	assert.Equal(t, schema.ErrUnknownUser, err)

	//Visit
	assert.Nil(t, store.Visit(em))
	assert.Equal(t, schema.ErrUnknownUser, store.Visit(string(randomBytes(10))))
	assert.NotNil(t, user(t, em).LastVisit)

	//Person update
	p2 := schema.Person{Email: em}
	assert.Nil(t, store.SetUserPerson(p2))
	u3 := user(t, p2.Email)
	assert.Equal(t, p2, u3.Person)

	//Change Role
	assert.Equal(t, schema.ErrUnknownUser, store.SetUserRole(string(randomBytes(10)), schema.STUDENT))
	assert.Nil(t, store.SetUserRole(em, schema.STUDENT))

	//Delete account
	assert.Equal(t, schema.ErrUnknownUser, store.RmUser(string(randomBytes(1))))
	assert.Nil(t, store.RmUser(u3.Person.Email))
	_, err = store.User(u3.Person.Email)
	assert.Equal(t, schema.ErrUnknownUser, err)
}

func TestUserCreation(t *testing.T) {
	p1 := testutil.Person()
	_, err := store.NewUser(p1, schema.ROOT)
	assert.Nil(t, err)
	_, err = store.NewUser(p1, schema.TUTOR)
	assert.Equal(t, schema.ErrUserExists, err)

	p1.Email = "foo"
	_, err = store.NewUser(p1, schema.TUTOR)
	assert.Equal(t, schema.ErrInvalidEmail, err)

}
func TestUsers(t *testing.T) {
	assert.Nil(t, store.Install())

	us, err := store.Users()
	assert.Nil(t, err)
	assert.Empty(t, us)

	t1 := newTutor(t)
	t2 := newTutor(t)
	s1 := newStudent(t)
	us, err = store.Users()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(us))
	assert.Contains(t, us, t1, t2, s1.User)
}
