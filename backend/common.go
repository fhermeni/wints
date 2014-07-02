package backend

import (
	"database/sql"
	"time"
	"log"
	"errors"
)

func StoreLog(db *sql.DB, uid int, message string) {
	_, err := db.Exec("insert into logs(date, uid, message) values($1,$2,$3)", time.Now(), uid, message)
	if (err != nil) {
		log.Printf("Unable to store a log entry: %s", err)
	}
	log.Printf("%d - %s\n", uid, message)
}


type BackendError struct {
	Status int
	message string

}

func (err *BackendError) Error() string {
	return err.message
}

func SingleUpdate(db *sql.DB, q string, args ...interface{}) error {
	res, err := db.Exec(q, args...)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errors.New("No updates")
	}
	return nil
}

/*
func SetMajor(db *sql.DB, email string, m string) error {
	res, err := db.Exec("update students set major=$2 where email=$1", email, m)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errors.New("Unknown student '" + email + "'\n")
	}
	return nil
}
 */
