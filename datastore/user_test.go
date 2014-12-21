package datastore

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func getDB(t *testing.T) *sql.DB {
	if len(os.Getenv("TESTING_DB")) == 0 {
		t.Fatalf("Missing environment variable TESTING_DB")
	}
	url := "dbname=" + os.Getenv("TESTING_DB") + " sslmode=disable"
	DB, err := sql.Open("postgres", url)
	if err != nil {
		t.Fatalf("Unable to connect to the testing Database: " + err.Error())
	}
	return DB
}
func TestWorkflow(t *testing.T) {
	db := getDB(t)
	s, err := NewService(db)
	assert.NoError(t, err)
	//NewUser
	u1 := internship.User{Firstname: "foo", Lastname: "bar", Email: "foo@bar.com", Tel: "0123456", Role: internship.STUDENT}
	_, err = s.NewUser(u1)
	assert.NoError(t, err)
	_, err = s.NewUser(u1)
	assert.Equal(t, internship.ErrUserExists, err)
	u2 := internship.User{Firstname: "foo", Lastname: "baz", Email: "foo@baz.com", Tel: "0123456", Role: internship.MAJOR}
	_, err = s.NewUser(u2)
	assert.NoError(t, err)

	//User()
	u, err := s.User("foo@bar.com")
	assert.NoError(t, err)
	assert.Equal(t, u1, u)
	u, err = s.User("foo@bib.com")
	assert.Equal(t, internship.ErrUnknownUser, err)

	//Users()
	users, err := s.Users()
	assert.NoError(t, err)
	assert.Contains(t, users, u1)
	assert.Contains(t, users, u2)

	//SetUserRole
	assert.Equal(t, internship.ErrUnknownUser, s.SetUserRole("foo.bar.com", internship.STUDENT))
	assert.NoError(t, s.SetUserRole("foo@bar.com", internship.ADMIN))
	u, err = s.User("foo@bar.com")
	assert.NoError(t, err)
	assert.Equal(t, internship.ADMIN, u.Role)

	//SetUserProfile
	assert.Equal(t, internship.ErrUnknownUser, s.SetUserProfile("foo.bar.com", "toto", "tata", "titi"))
	assert.NoError(t, s.SetUserProfile("foo@bar.com", "toto", "tata", "987654"))
	u, err = s.User("foo@bar.com")
	assert.NoError(t, err)
	assert.Equal(t, "toto", u.Firstname)
	assert.Equal(t, "tata", u.Lastname)
	assert.Equal(t, "987654", u.Tel)

	//Login
	_, err = s.Registered(u1.Email, []byte{})
	assert.Equal(t, internship.ErrCredentials, err)

	//ResetRootAccount -- to be aware of the current password
	assert.NoError(t, s.ResetRootAccount())
	u, err = s.Registered(DEFAULT_LOGIN, []byte(DEFAULT_PASSWORD))
	assert.NoError(t, err)
	assert.NotNil(t, u)

	//SetUserPassword()
	assert.NoError(t, s.SetUserPassword(DEFAULT_LOGIN, []byte(DEFAULT_PASSWORD), []byte("test")))
	_, err = s.Registered(DEFAULT_LOGIN, []byte("test"))
	assert.NoError(t, err)
	_, err = s.Registered("foo", []byte("test"))
	assert.Equal(t, internship.ErrCredentials, err)

	//NewPassword && ResetPassword
	_, err = s.ResetPassword("foo")
	assert.Equal(t, internship.ErrUnknownUser, err)
	_, err = s.NewPassword([]byte{}, []byte("foo"))
	assert.Equal(t, internship.ErrNoPendingRequests, err)
	tok, err := s.ResetPassword("foo@bar.com")
	assert.NoError(t, err)
	l, err := s.NewPassword(tok, []byte("baz"))
	assert.Equal(t, "foo@bar.com", l)
	u, err = s.Registered("foo@bar.com", []byte("baz"))
	assert.NoError(t, err)
	assert.Equal(t, "toto", u.Firstname)

	//RmUsers
	//Need an internship to play with conflicts
	c := internship.Company{Name: "foo", WWW: "bar"}
	sup := internship.Supervisor{Firstname: "foo", Lastname: "bar", Email: "toto@bac.com", Tel: "12345"}
	i := internship.Internship{Student: u1, Tutor: u2, Start: time.Now(), End: time.Now, Title: "My Internship", Cpy: c, Sup: sup}

}