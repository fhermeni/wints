package testutil

import (
	"crypto/rand"
	"time"

	"github.com/fhermeni/wints/schema"
)

func Person() schema.Person {
	p := schema.Person{
		Firstname: string(randomBytes(12)),
		Lastname:  string(randomBytes(12)),
		Tel:       string(randomBytes(12)),
		Email:     string(randomBytes(12)) + "@fr",
	}
	return p
}

func randomBytes(s int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, s)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func convention(student schema.Student, tutor schema.User) schema.Convention {
	creation := time.Now().Truncate(time.Minute).UTC()
	return schema.Convention{
		Student:    student,
		Tutor:      tutor,
		Supervisor: Person(),
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
}
