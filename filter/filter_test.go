package filter

import (
	"testing"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestSetUserPassword(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto"}
	f, _ := NewService(&mock, u)
	assert.NoError(t, f.SetUserPassword("toto", []byte{}, []byte{}))
	assert.Equal(t, ErrPermission, f.SetUserPassword("titi", []byte{}, []byte{}))
}

func TestSetUserProfile(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto"}
	f, _ := NewService(&mock, u)
	assert.NoError(t, f.SetUserProfile("toto", "fn", "ln", "0123"))
	assert.Equal(t, ErrPermission, f.SetUserProfile("titi", "fn", "ln", "0123"))
}

func TestSetUserRole(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto", Role: internship.STUDENT}
	f, _ := NewService(&mock, u)
	assert.Equal(t, ErrPermission, f.SetUserRole("titi", internship.ADMIN)) //no enough rights

	//no self-promotion
	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ADMIN})
	assert.Equal(t, ErrPermission, f.SetUserRole("toto", internship.ROOT))

	//okok
	assert.NoError(t, f.SetUserRole("titi", internship.ADMIN))
}

func TestResetPassword(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto", Role: internship.STUDENT}
	f, _ := NewService(&mock, u)
	_, err := f.ResetPassword("titi")
	assert.Equal(t, ErrPermission, err)

	_, err = f.ResetPassword("toto")
	assert.NoError(t, err)
}

func TestUser(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto", Role: internship.STUDENT}
	f, _ := NewService(&mock, u)
	_, err := f.User("toto")
	assert.NoError(t, err)
	_, err = f.User("titi")
	assert.Equal(t, ErrPermission, err)

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ADMIN})
	_, err = f.User("titi")
	assert.NoError(t, err)
}

func TestUsers(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto", Role: internship.STUDENT}
	f, _ := NewService(&mock, u)
	_, err := f.Users()
	assert.Equal(t, ErrPermission, err)

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ADMIN})
	_, err = f.Users()
	assert.NoError(t, err)
}

func TestRmUser(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto", Role: internship.STUDENT}
	f, _ := NewService(&mock, u)
	assert.Equal(t, ErrPermission, f.RmUser("foo"))

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ROOT})
	assert.NoError(t, f.RmUser("foo"))

	//no self-destruction
	assert.Equal(t, ErrPermission, f.RmUser("toto"))
}
