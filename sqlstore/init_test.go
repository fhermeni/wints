// +build integration

package sqlstore

import (
	"database/sql"
	"flag"
	"log"
	"testing"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

var dbUrl = flag.String("db-url", "user=fhermeni2 dbname=wints_next host=localhost sslmode=disable", "database connexion string")

var cfg = config.Internships{
	Majors:     []string{"al", "ihm"},
	Promotions: []string{"si", "master"},
	Surveys:    []config.Survey{},
	Reports: []config.Report{
		config.Report{
			Kind:     "foo",
			Delivery: config.AbsoluteDeadline(time.Now().Add(time.Minute)),
		},
	},
	LatePenalty: 2,
}

func person() schema.Person {
	p := schema.Person{
		Firstname: string(randomBytes(12)),
		Lastname:  string(randomBytes(12)),
		Tel:       string(randomBytes(12)),
		Email:     string(randomBytes(12)) + "@fr",
	}
	return p
}

func newTutor(t *testing.T) schema.User {
	//Preparation
	p := person()
	_, err := store.NewUser(p, schema.TUTOR)
	assert.Nil(t, err)
	return schema.User{
		Person: p,
		Role:   schema.TUTOR,
	}
}

func newPassword(t *testing.T, em string) []byte {
	token, err := store.ResetPassword(em)
	assert.Nil(t, err)
	password := randomBytes(32)
	em2, err := store.NewPassword(token, password)
	assert.Equal(t, em, em2)
	assert.Nil(t, err)
	return password
}

func newStudent(t *testing.T) schema.Student {
	p := person()
	u := schema.User{
		Person: p,
		Role:   schema.STUDENT,
	}
	stu := schema.Student{
		User:      u,
		Major:     "al",
		Promotion: "si",
		Male:      true,
	}
	assert.Nil(t, store.NewStudent(p, stu.Major, stu.Promotion, stu.Male))
	return stu
}

func newInternship(t *testing.T, student schema.Student, tutor schema.User) schema.Internship {
	creation := time.Now()
	c := schema.Convention{
		Student:    student,
		Tutor:      tutor,
		Supervisor: person(),
		Creation:   creation,
		Begin:      creation.Add(time.Minute),
		End:        creation.Add(time.Hour),
		Company: schema.Company{
			Title: string(randomBytes(10)),
			WWW:   string(randomBytes(10)),
			Name:  string(randomBytes(10)),
		},
		//tons of default values
		Gratification:  10000,
		Lab:            false,
		ForeignCountry: true,
		Skip:           false,
		Valid:          false,
	}
	i, _, err := store.NewInternship(c)
	assert.Nil(t, err)
	c.Creation = c.Creation.Truncate(time.Minute).UTC()
	c.Begin = c.Begin.Truncate(time.Minute).UTC()
	c.End = c.End.Truncate(time.Minute).UTC()
	i.Convention = c
	return i
}

func convention(t *testing.T, u schema.User) schema.Convention {
	c2, err := store.Convention(u.Person.Email)
	assert.Nil(t, err)
	return c2
}

func init() {
	log.Println("Initiate database connexion using url '" + *dbUrl + "'")
	DB, err := sql.Open("postgres", *dbUrl)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}

	store, _ = NewStore(DB, cfg)
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
