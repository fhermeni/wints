package sqlstore

import (
	"errors"
	"strings"
	"time"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	insertUser                   = "insert into users(firstname, lastname, tel, email, role, password) values ($1,$2,$3,$4,$5,$6)"
	startPasswordRenewal         = "insert into password_renewal(email,token,deadline) values($1,$2,$3)"
	newSession                   = "insert into sessions(email, token, expire) values ($1,$2,$3)"
	updateLastVisit              = "update users set lastVisit=$1 where email=$2"
	updateUserProfile            = "update users set firstname=$1, lastname=$2, tel=$3 where email=$4"
	updateUserRole               = "update users set role=$2 where email=$1"
	updateUserPassword           = "update users set password=$2 where email=$1"
	deleteSession                = "delete from sessions where email=$1"
	deleteSessionFromToken       = "delete from sessions where token=$1"
	deletePasswordRenewalRequest = "delete from password_renewal where email=$1"
	deleteUser                   = "DELETE FROM users where email=$1"
	allUsers                     = "select firstname, lastname, email, tel, role, lastVisit from users"
	selectPassword               = "select password from users where email=$1"
	selectExpire                 = "select expire from sessions where token=$1"
	emailFromRenewableToken      = "select email from password_renewal where token=$1"
	replaceTutorInConventions    = "update conventions set tutor=$2 where tutor=$1"
	replaceJuryInDefenses        = "update defenseJuries set jury=$2 where tutor=$1"
)

//addUser add the given user.
//Every strings are turned into their lower case version
func (s *Store) addUser(tx *TxErr, u internship.User) {
	tx.Exec(insertUser,
		strings.ToLower(u.Person.Firstname),
		strings.ToLower(u.Person.Lastname),
		u.Person.Tel,
		strings.ToLower(u.Person.Email),
		u.Role,
		randomBytes(32),
	)
}

//Visit writes the current time for the given user
func (s *Store) Visit(u string) error {
	return s.singleUpdate(updateLastVisit, internship.ErrUnknownUser, time.Now(), u)
}

//Users list all the registered users
func (s *Store) Users() ([]internship.User, error) {
	users := make([]internship.User, 0, 0)
	st := s.stmt(allUsers)
	rows, err := st.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		u := internship.User{Person: internship.Person{}}
		var last pq.NullTime
		rows.Scan(
			&u.Person.Firstname,
			&u.Person.Lastname,
			&u.Person.Email,
			&u.Person.Tel,
			&u.Role,
			&last,
		)
		u.LastVisit = nullableTime(last)
		//Get the role if exists
		users = append(users, u)
	}
	return users, nil
}

//SetUserProfile changes the user profile if exists
func (s *Store) SetUserProfile(email, fn, ln, tel string) error {
	return s.singleUpdate(updateUserProfile, internship.ErrUnknownUser, fn, ln, tel, email)
}

//SetUserRole updates the user privilege
func (s *Store) SetUserRole(email string, priv internship.Privilege) error {
	return s.singleUpdate(updateUserRole, internship.ErrUnknownUser, email, priv)
}

//SetPassword changes the user password
//Any pending renewable requests are deleted on success
func (s *Store) SetPassword(email string, oldP, newP []byte) error {
	hash, err := hash(newP)
	if err != nil {
		return err
	}
	tx := newTxErr(s.db)
	var p []byte
	tx.err = tx.tx.QueryRow(selectPassword, strings.ToLower(email)).Scan(&p)
	tx.err = noRowsTo(tx.err, internship.ErrCredentials)
	if tx.err != nil {
		return tx.Done()
	}
	if err := bcrypt.CompareHashAndPassword(p, oldP); err != nil {
		tx.err = internship.ErrCredentials
	}
	tx.Exec(deletePasswordRenewalRequest, email)
	tx.Update(updateUserPassword, email, hash)
	//no need to check update, we know the user exists in the transaction context
	return tx.Done()
}

//OpenedSession check if a session is currently open for the user and is not expired
func (s *Store) OpenedSession(token []byte) error {
	var last time.Time
	st := s.stmt(selectExpire)
	err := st.QueryRow(token).Scan(&last)
	err = noRowsTo(err, internship.ErrCredentials)
	if err != nil {
		return err
	}
	if time.Now().After(last) {
		return internship.ErrSessionExpired
	}
	return err
}

//ResetPassword starts a reset procedure, valid for a given period
func (s *Store) ResetPassword(email string, d time.Duration) ([]byte, error) {
	//In case a request already exists
	token := randomBytes(32)

	tx := newTxErr(s.db)
	tx.Exec(deletePasswordRenewalRequest, email)
	tx.Exec(startPasswordRenewal, email, token, time.Now().Add(d))
	return token, tx.Done()
}

//NewPassword commits a password renewall request.
//From a request token and a new password, it returns upon success the target user email
func (s *Store) NewPassword(token, newP []byte) (string, error) {
	hash, err := hash(newP)
	if err != nil {
		return "", err
	}

	var email string
	tx := newTxErr(s.db)
	tx.err = tx.tx.QueryRow(emailFromRenewableToken, token).Scan(&email)
	tx.err = noRowsTo(tx.err, internship.ErrNoPendingRequests)
	tx.Update(updateUserPassword, email, hash)
	//no need to check updated rows as it is sure the user exists in the tx context
	nb := tx.Update(deletePasswordRenewalRequest, email)
	if tx.err == nil && nb != 1 {
		tx.err = errors.New("Unable to clean the password renewable request of " + email)
	}
	return email, tx.Done()
}

//Login log a user.
//Upon success, it generates a session token that will be valid for the next 24 hours.
func (s *Store) Login(email string, password []byte) (internship.Session, error) {
	var p []byte
	ss := internship.Session{
		Email:  email,
		Token:  randomBytes(32),
		Expire: time.Now().Add(time.Hour * 24),
	}
	tx := newTxErr(s.db)
	tx.err = tx.tx.QueryRow(selectPassword, strings.ToLower(email)).Scan(&p)
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

//Logout destroy the current user session if exists
func (s *Store) Logout(token []byte) error {
	return s.singleUpdate(deleteSessionFromToken, internship.ErrUnknownUser, token)
}

//NewUser add a user
//Basically, calls addUser
func (s *Store) NewUser(fn, ln, tel, email string) error {
	return s.singleUpdate(insertUser, internship.ErrUserExists, fn, ln, tel, email, internship.NONE, randomBytes(32))
}

//RmUser removes a user from the database
func (s *Store) RmUser(email string) error {
	return s.singleUpdate(deleteUser, internship.ErrUnknownUser, email)
}

//ReplaceUserWith the account referred by src by the account referred by dst
func (s *Store) ReplaceUserWith(src, dst string) error {
	tx := newTxErr(s.db)
	tx.Update(replaceTutorInConventions, src, dst)
	tx.Update(replaceJuryInDefenses, src, dst)
	tx.Update(deleteUser, src)
	return tx.Done()
}
