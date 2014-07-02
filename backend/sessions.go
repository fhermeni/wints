package backend

import (
	_ "github.com/lib/pq"
	"time"
	"code.google.com/p/go.crypto/bcrypt"
	"net/http"
	"database/sql"
)

const (
	sessionDuration = "24h"
	wipeSessions = "delete from sessions"
	deleteSession = "delete from sessions where token=$1"
	makeSession = "insert into sessions(email, token, last) values($1, $2, $3)"
	updateSession = "update sessions set token=$2, last=$3 where email=$1"
	selectUIDFromToken = "select email, last from sessions where token=$1"
)

type Credential struct {
	Email string
	Password string
}

func OpenSession(db *sql.DB, email string) ([]byte,error) {

	//Get the token and the new deadline
	tok, err := bcrypt.GenerateFromPassword([]byte(time.Now().String()), bcrypt.MinCost)
	d,_ := time.ParseDuration(sessionDuration)
	last := time.Now().Add(d)

	//Update if possible
	err = SingleUpdate(db, updateSession, email, tok, last)
	if err == nil {
		return tok, err
	}

	err = SingleUpdate(db, makeSession, email, tok, last)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func CloseSession(db *sql.DB, token string) error {
	res, err := db.Exec(deleteSession, token)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return &BackendError{http.StatusUnauthorized, "Invalid session token"}
	}
	return nil
}

func WipeSessions(db *sql.DB) error {
	_, err := db.Exec(wipeSessions)
	return err
}

func CheckSession(db *sql.DB, token string) (string, error) {
	rows, err := db.Query(selectUIDFromToken, token)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	if !rows.Next() {
		return "", &BackendError{http.StatusUnauthorized, "Invalid session token"}
	}
	var email string
	var last time.Time
	rows.Scan(&email, &last)
	//Check if the token is fresh enough
	if (time.Now().After(last)) {
		return "", &BackendError{http.StatusGone, "The session expired"}
	}
	return email, nil
}

