//Utilities for database related stuff
package datastore

import (
	"crypto/rand"
	"database/sql"
)

func SingleUpdate(db *sql.DB, errNoUpdate error, q string, args ...interface{}) error {
	res, err := db.Exec(q, args...)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errNoUpdate
	}
	return nil
}

func rand_str(str_size int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}
