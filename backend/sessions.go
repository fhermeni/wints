package backend

import (
	_ "github.com/lib/pq"
	"time"
)

const (
	sessionDuration = "24h"
	wipeSessions = "delete from sessions"
	deleteSession = "delete from sessions where uid=$1"
	makeSession = "insert into sessions(uid, token, last) values($1, $2, $3)"
	updateSession = "update sessions set token=$2, last=$3 where uid=$1"
	selectUIDFromToken = "select uid, last from sessions where token=$1"
)

type Session struct {
	Token string
	Last time.Time
	Role string
	Uid int
}

/*func OpenSession(db *sql.DB, p Profile) (Session,error) {

	//Get the token and the new deadline
	tok, err := bcrypt.GenerateFromPassword([]byte(time.Now().String()), bcrypt.MinCost)
	d,_ := time.ParseDuration(sessionDuration)
	last := time.Now().Add(d)

	//Update if possible
	res, err := db.Exec(updateSession, p.UID, tok, last)
	if err != nil {
		return Session{}, err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return Session{}, err
	}
	if nb == 1 {
		//session already existed so it was a refresh
		return Session{string(tok), last, p.Role, p.UID}, nil
	}

	res, err = db.Exec(makeSession, p.UID, tok, last)
	if err != nil {
		return Session{}, err
	}
	nb, err = res.RowsAffected()
	if err != nil {
		return Session{}, err
	}
	if nb != 1 {
		return Session{}, errors.New("Unable to store the session")
	}
	return Session{string(tok), last, p.Role, p.UID}, nil
}

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
