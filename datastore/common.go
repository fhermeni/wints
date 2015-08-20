package datastore

import (
	"crypto/rand"
	"time"

	"github.com/lib/pq"
)

func randomBytes(s int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, s)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func nullableTime(t pq.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func violationAsErr(got error, cstr string, with error) error {
	if got == nil {
		return got
	}
	if e, ok := got.(*pq.Error); ok {
		if e.Constraint == cstr {
			return with
		}
	}
	return got
}
