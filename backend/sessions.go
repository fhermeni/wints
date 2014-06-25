package backend

import (
	_ "github.com/lib/pq"
	"time"
	"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	"errors"
)

const (
	sessionDuration = "24h"
	wipeSessions = "delete from sessions"
	deleteSession = "delete from sessions where uid=$1"
	makeSession = "insert into sessions(email, token, last) values($1, $2, $3)"
	updateSession = "update sessions set token=$2, last=$3 where email=$1"
	selectUIDFromToken = "select uid, last from sessions where token=$1"
)

type Session struct {
	Token string
	Last time.Time
	Role string
	Uid int
}

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
	res, err := db.Exec(updateSession, email, tok, last)
	if err != nil {
		return nil, err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if nb == 1 {
		//session already existed so it was a refresh
		return tok, nil
	}

	res, err = db.Exec(makeSession, email, tok, last)
	if err != nil {
		return nil, err
	}
	nb, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if nb != 1 {
		return nil, errors.New("Unable to store the session")
	}
	return tok, nil
}
                               /*
func CloseSession(db *sql.DB, uid int) error {
	res, err := db.Exec(deleteSession, uid)
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

func CheckSession(db *sql.DB, token string) (int, error) {
	rows, err := db.Query(selectUIDFromToken, token)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, &BackendError{http.StatusUnauthorized, "Invalid session token"}
	}
	var uid int
	var last time.Time
	rows.Scan(&uid, &last)
	//Check if the token is fresh enough
	if (time.Now().After(last)) {
		return 0, &BackendError{http.StatusGone, "The session expired"}
	}
	return uid, nil
}
                     */
