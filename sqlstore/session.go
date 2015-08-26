package sqlstore

import (
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/fhermeni/wints/internship"
)

var (
	selectSession          = "select email, token, expire from sessions where token=$1"
	deleteSession          = "delete from sessions where email=$1"
	newSession             = "insert into sessions(email, token, expire) values ($1,$2,$3)"
	deleteSessionFromToken = "delete from sessions where token=$1"
)

//Session get the token related session
func (s *Store) Session(token []byte) (internship.Session, error) {
	st := s.stmt(selectSession)
	ss := internship.Session{}
	err := st.QueryRow(token).Scan(&ss.Email, &ss.Token, &ss.Expire)
	return ss, noRowsTo(err, internship.ErrCredentials)
}

//NewSession creates a new session if the email and the password matches a user
func (s *Store) NewSession(email string, password []byte, expire time.Duration) (internship.Session, error) {
	var p []byte
	ss := internship.Session{
		Email:  email,
		Token:  randomBytes(32),
		Expire: time.Now().Add(expire),
	}
	tx := newTxErr(s.db)
	tx.err = tx.QueryRow(selectPassword, email).Scan(&p)
	tx.err = noRowsTo(tx.err, internship.ErrCredentials)
	if tx.err != nil {
		return ss, tx.Done()
	}
	if err := bcrypt.CompareHashAndPassword(p, password); err != nil {
		tx.err = internship.ErrCredentials
	}
	tx.Exec(deleteSession, email)
	nb := tx.Update(newSession, ss.Email, ss.Token, ss.Expire)
	if tx.err == nil && nb != 1 {
		tx.err = internship.ErrUnknownUser
	}
	return ss, tx.Done()
}

//RmSession destroy the current user session if exists
func (s *Store) RmSession(token []byte) error {
	return s.singleUpdate(deleteSessionFromToken, internship.ErrUnknownUser, token)
}
