package datastore

import (
	"database/sql"
	"os"
	"testing"

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
	u1 := internship.User{Firstname: "foo", Lastname: "bar", Email: "foo@bar.com", Tel: "0123456", Role: internship.STUDENT}
	assert.NoError(t, s.NewUser(u1))
	//Get the user
	u, err := s.User("foo@bar.com")
	assert.NoError(t, err)
	assert.Equal(t, u1, u)

	//Unknown user
	u, err = s.User("foo@bib.com")
	assert.Equal(t, internship.ErrUnknownUser, err)

	//Same user, expect already exits
	assert.Equal(t, internship.ErrUserExists, s.NewUser(u1))

	u2 := internship.User{Firstname: "foo", Lastname: "baz", Email: "foo@baz.com", Tel: "0123456", Role: internship.MAJOR}
	assert.NoError(t, s.NewUser(u2))

	//List users
	users, err := s.Users()
	assert.NoError(t, err)
	assert.Contains(t, users, u1)
	assert.Contains(t, users, u2)

	//Change role
	assert.Equal(t, internship.ErrUnknownUser, s.SetUserRole("foo.bar.com", internship.STUDENT))
	assert.NoError(t, s.SetUserRole("foo@bar.com", internship.ADMIN))
	u, err = s.User("foo@bar.com")
	assert.NoError(t, err)
	assert.Equal(t, internship.ADMIN, u.Role)

	//Change profile
	assert.Equal(t, internship.ErrUnknownUser, s.SetUserProfile("foo.bar.com", "toto", "tata", "titi"))
	assert.NoError(t, s.SetUserProfile("foo@bar.com", "toto", "tata", "987654"))
	u, err = s.User("foo@bar.com")
	assert.NoError(t, err)
	assert.Equal(t, "toto", u.Firstname)
	assert.Equal(t, "tata", u.Lastname)
	assert.Equal(t, "987654", u.Tel)

}
