// +build integration

package sqlstore

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/fhermeni/wints/schema"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

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

func user(t *testing.T, email string) schema.User {
	u, err := store.User(email)
	assert.Nil(t, err)
	return u
}

func TestUserManagement(t *testing.T) {
	p, passwd := preparePerson(t)

	//User()
	u, err := store.User(p.Email)
	assert.Nil(t, err)
	assert.Equal(t, schema.NONE, u.Role)
	assert.Nil(t, u.LastVisit)
	_, err = store.User(string(randomBytes(10)))
	assert.Equal(t, schema.ErrUnknownUser, err)

	//Visit
	assert.Nil(t, store.Visit(p.Email))
	assert.Equal(t, schema.ErrUnknownUser, store.Visit(string(randomBytes(10))))
	assert.NotNil(t, user(t, p.Email).LastVisit)

	//Person update
	p2 := schema.Person{Email: p.Email}
	assert.Nil(t, store.SetUserPerson(p2))
	u2 := user(t, p2.Email)
	assert.Equal(t, p2, u2.Person)

	//Change Role
	assert.Equal(t, schema.ErrUnknownUser, store.SetUserRole(string(randomBytes(10)), schema.STUDENT))
	assert.Nil(t, store.SetUserRole(p.Email, schema.STUDENT))

	//SetPassword
	passwd2 := randomBytes(10)
	assert.Equal(t, schema.ErrCredentials, store.SetPassword(string(randomBytes(10)), passwd, passwd2))
	assert.Nil(t, store.SetPassword(p.Email, passwd, passwd2))
	_, err = store.NewSession(p.Email, passwd2, time.Hour)
	assert.Nil(t, err)
	_, err = store.NewSession(p.Email, passwd, time.Hour)
	assert.Equal(t, schema.ErrCredentials, err)
}
