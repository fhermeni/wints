package sqlstore

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/fhermeni/wints/schema"
)

var (
	selectSession = "select email, token, expire from sessions where token=$1"
	deleteSession = "delete from sessions where email=$1"

	newSession = "insert into sessions(email, token, expire) values ($1,$2,$3)"
)

//Session get the token related session
func (s *Store) Session(token []byte) (schema.Session, error) {
	st := s.stmt(selectSession)
	ss := schema.Session{}
	err := st.QueryRow(token).Scan(&ss.Email, &ss.Token, &ss.Expire)
	ss.Expire = ss.Expire.UTC()
	return ss, noRowsTo(err, schema.ErrCredentials)
}

//NewSession creates a new session if the email and the password matches a user
func (s *Store) NewSession(email string, password []byte, expire time.Duration) (schema.Session, error) {
	var p []byte
	ss := schema.Session{
		Email:  email,
		Token:  randomBytes(32),
		Expire: time.Now().Add(expire).Truncate(time.Minute).UTC(),
	}
	tx := newTxErr(s.db)
	tx.err = tx.QueryRow(selectPassword, email).Scan(&p)
	tx.err = noRowsTo(tx.err, schema.ErrUnknownUser)
	if tx.err != nil {
		return ss, tx.Done()
	}
	if err := bcrypt.CompareHashAndPassword(p, password); err != nil {
		tx.err = schema.ErrCredentials
	}
	tx.Exec(deleteSession, email)
	nb := tx.Update(newSession, ss.Email, ss.Token, ss.Expire)
	if tx.err == nil && nb != 1 {
		tx.err = schema.ErrUnknownUser
	}
	return ss, tx.Done()
}

//RmSession destroy the current user session if exists
func (s *Store) RmSession(email string) error {
	return s.singleUpdate(deleteSession, schema.ErrUnknownUser, email)
}
