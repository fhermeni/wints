package filter

import (
	"testing"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestUserManagement(t *testing.T) {
	var mock MockBackend

	f, _ := NewService(&mock, internship.User{Email: "toto", Role: internship.STUDENT})
	//Update _my_ password
	assert.NoError(t, f.SetUserPassword("toto", []byte{}, []byte{}))
	assert.Equal(t, ErrPermission, f.SetUserPassword("titi", []byte{}, []byte{}))

	//Change _my_ profile
	assert.NoError(t, f.SetUserProfile("toto", "fn", "ln", "0123"))
	assert.Equal(t, ErrPermission, f.SetUserProfile("titi", "fn", "ln", "0123"))

	//Manage role
	assert.Equal(t, ErrPermission, f.SetUserRole("titi", internship.ADMIN)) //no enough rights

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ADMIN})
	assert.Equal(t, ErrPermission, f.SetUserRole("toto", internship.ROOT)) //no self-promotion
	assert.NoError(t, f.SetUserRole("titi", internship.ADMIN))             //okok

	//Reset password
	_, err := f.ResetPassword("titi")
	assert.Equal(t, ErrPermission, err)
	_, err = f.ResetPassword("toto")
	assert.NoError(t, err)

	//User()
	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.STUDENT})
	_, err = f.User("toto")
	assert.NoError(t, err)
	_, err = f.User("titi")
	assert.Equal(t, ErrPermission, err)

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ADMIN})
	_, err = f.User("titi")
	assert.NoError(t, err)

	//Users()
	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.STUDENT})
	_, err = f.Users()
	assert.Equal(t, ErrPermission, err)

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ADMIN})
	_, err = f.Users()
	assert.NoError(t, err)

	//RmUser()
	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.STUDENT})
	assert.Equal(t, ErrPermission, f.RmUser("foo"))

	f, _ = NewService(&mock, internship.User{Email: "toto", Role: internship.ROOT})
	assert.NoError(t, f.RmUser("foo"))

	//no self-destruction
	assert.Equal(t, ErrPermission, f.RmUser("toto"))
}
