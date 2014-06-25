package backend

import (
	"database/sql"
	"time"
	"log"
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
