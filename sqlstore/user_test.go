package sqlstore

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/fhermeni/wints/internship"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

//Adduser
//ToStudent
//SetMale
//Skip
//SetMajor
//SetPromotion

var store *Store

func init() {
	conn := "user=fhermeni dbname=wints_test host=localhost sslmode=disable"
	DB, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}
	store, _ = NewStore(DB)
	err = store.Install()
	if err != nil {
		log.Fatalln("Unable to install the DB: " + err.Error())
	}
}

func user(t *testing.T, email string) internship.User {
	us, err := store.Users()
	assert.Nil(t, err)
	for _, u := range us {
		if u.Person.Email == email {
			return u
		}
	}
	assert.Fail(t, "")
	return internship.User{}
}

func TestUserManagement(t *testing.T) {
	us, err := store.Users()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(us))

	assert.Nil(t, store.NewUser("foo", "bar", "baz", "foo@bar.com"))
	assert.Equal(t, internship.ErrUserExists, store.NewUser("foo", "bar", "baz", "foo@bar.com"))

	//Add the user
	us, err = store.Users()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(us))
	u := internship.User{
		Person: internship.Person{
			Firstname: "foo",
			Lastname:  "bar",
			Tel:       "baz",
			Email:     "foo@bar.com",
		},
		Role:      internship.NONE,
		LastVisit: nil,
	}
	assert.Equal(t, u, us[0])

	//Change Role
	assert.Equal(t, internship.ErrUnknownUser, store.SetUserRole("", internship.STUDENT))
	assert.Nil(t, store.SetUserRole("foo@bar.com", internship.STUDENT))

	//Change profile
	assert.Equal(t, internship.ErrUnknownUser, store.SetUserProfile("", "", "", ""))
	assert.Nil(t, store.SetUserProfile("foo@bar.com", "aa", "bb", "cc"))
	u.Person.Firstname = "aa"
	u.Person.Lastname = "bb"
	u.Person.Tel = "cc"
	u.Role = internship.STUDENT
	assert.Equal(t, u, user(t, "foo@bar.com"))

	//Visit
	assert.Nil(t, store.Visit("foo@bar.com"))
	assert.Equal(t, internship.ErrUnknownUser, store.Visit(""))
	assert.NotNil(t, user(t, "foo@bar.com").LastVisit)

	//Login stuff
	//Test if the credentials match a user, return a session token
	_, err = store.Login("foo@bar.com", []byte("foo"))
	assert.Equal(t, internship.ErrCredentials, err)
	_, err = store.Login("foo@baz.com", []byte("foo"))
	assert.Equal(t, internship.ErrCredentials, err)
	assert.Equal(t, internship.ErrCredentials, store.OpenedSession("foo@bar.com", []byte{}))
	assert.Equal(t, internship.ErrCredentials, store.OpenedSession("foo@baz.com", []byte{}))

	//Reset the password
	d, _ := time.ParseDuration("1h")
	_, err = store.ResetPassword("foo@baz.com", d)
	assert.Equal(t, internship.ErrUnknownUser, err)
	tok, err := store.ResetPassword("foo@bar.com", d)
	assert.Nil(t, err)

	em, err := store.NewPassword([]byte("dd"), []byte("lkjpp"))
	assert.Equal(t, internship.ErrNoPendingRequests, err)
	em, err = store.NewPassword(tok, []byte("lkjpp"))
	assert.Nil(t, err)
	assert.Equal(t, "foo@bar.com", em)
	_, err = store.NewPassword(tok, []byte("lkjpp"))
	assert.Equal(t, internship.ErrNoPendingRequests, err)

	//login
	token, err := store.Login("foo@bar.com", []byte("lkjpp"))
	assert.Nil(t, err)
	//opened session
	assert.Nil(t, store.OpenedSession("foo@bar.com", token))
	//logout
	assert.Nil(t, store.Logout("foo@bar.com", token))
	assert.NotNil(t, store.Logout("foo@bar.com", token))
	assert.NotNil(t, store.OpenedSession("foo@bar.com", token))

	//changePassword
	assert.Nil(t, store.SetPassword("foo@bar.com", []byte("lkjpp"), []byte("1111")))
	//login
	token, err = store.Login("foo@bar.com", []byte("1111"))
	assert.Nil(t, err)
	assert.Nil(t, store.Logout("foo@bar.com", token))
	//logout

	//profile

	//rmUser

}
