package datastore

import (
	"crypto/rand"
	"database/sql"
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

func rollback(e error, tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return e
}

func violated(got error, cstr string) bool {
	if got == nil {
		return false
	}
	if e, ok := got.(*pq.Error); ok {
		return e.Constraint == cstr
	}
	return false
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
