package datastore

import (
	"crypto/rand"
	"database/sql"
)

//SingleUpdate executes a query that aims are affecting only 1 row.
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

func randomBytes(s int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, s)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}
