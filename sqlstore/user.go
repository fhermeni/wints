package sqlstore

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	selectUser                   = "select firstname, lastname, email, tel, role, lastVisit from users where email=$1"
	insertUser                   = "insert into users(firstname, lastname, tel, email, role, password) values ($1,$2,$3,$4,$5,$6)"
	startPasswordRenewal         = "insert into password_renewal(email,token,deadline) values($1,$2,$3)"
	updateLastVisit              = "update users set lastVisit=$1 where email=$2"
	updateUserProfile            = "update users set firstname=$1, lastname=$2, tel=$3 where email=$4"
	updateUserRole               = "update users set role=$2 where email=$1"
	updateUserPassword           = "update users set password=$2 where email=$1"
	deletePasswordRenewalRequest = "delete from password_renewal where email=$1"
	deleteUser                   = "DELETE FROM users where email=$1"
	allUsers                     = "select firstname, lastname, email, tel, role, lastVisit from users"
	selectPassword               = "select password from users where email=$1"
	emailFromRenewableToken      = "select email from password_renewal where token=$1"
	replaceTutorInConventions    = "update conventions set tutor=$2 where tutor=$1"
	replaceJuryInDefenses        = "update defenseJuries set jury=$2 where tutor=$1"
)

//addUser add the given user.
//Every strings are turned into their lower case version
func (s *Store) addUser(tx *TxErr, u schema.User) {
	tx.Exec(insertUser,
		u.Person.Firstname,
		u.Person.Lastname,
		u.Person.Tel,
		u.Person.Email,
		u.Role,
		randomBytes(32),
	)
}

//Visit writes the current time for the given user
func (s *Store) Visit(u string) error {
	return s.singleUpdate(updateLastVisit, schema.ErrUnknownUser, time.Now(), u)
}

func scanUser(row *sql.Rows) (schema.User, error) {
	u := schema.User{Person: schema.Person{}}
	var last pq.NullTime
	err := row.Scan(
		&u.Person.Firstname,
		&u.Person.Lastname,
		&u.Person.Email,
		&u.Person.Tel,
		&u.Role,
		&last,
	)
	u.LastVisit = nullableTime(last)
	return u, err
}

//User returns the given user account
func (s *Store) User(email string) (schema.User, error) {
	st := s.stmt(selectUser)
	rows, err := st.Query(email)
	if err != nil {
		return schema.User{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return schema.User{}, schema.ErrUnknownUser
	}
	return scanUser(rows)
}

//Users list all the registered users
func (s *Store) Users() ([]schema.User, error) {
	users := make([]schema.User, 0, 0)
	st := s.stmt(allUsers)
	rows, err := st.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		u := schema.User{Person: schema.Person{}}
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

//SetUserPerson changes the user profile if exists
func (s *Store) SetUserPerson(p schema.Person) error {
	return s.singleUpdate(updateUserProfile, schema.ErrUnknownUser, p.Firstname, p.Lastname, p.Tel, p.Email)
}

//SetUserRole updates the user privilege
func (s *Store) SetUserRole(email string, priv schema.Privilege) error {
	return s.singleUpdate(updateUserRole, schema.ErrUnknownUser, email, priv)
}

//SetPassword changes the user password
//Any pending renewable requests are deleted on success
func (s *Store) SetPassword(email string, oldP, newP []byte) error {
	if len(newP) < 8 {
		return schema.ErrPasswordTooShort
	}
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return err
	}
	tx := newTxErr(s.db)
	var p []byte
	tx.QueryRow(selectPassword, email).Scan(&p)
	tx.err = noRowsTo(tx.err, schema.ErrCredentials)
	if tx.err != nil {
		return tx.Done()
	}
	if err := bcrypt.CompareHashAndPassword(p, oldP); err != nil {
		tx.err = schema.ErrCredentials
	}
	tx.Exec(deletePasswordRenewalRequest, email)
	tx.Update(updateUserPassword, email, hash)
	//no need to check update, we know the user exists in the transaction context
	return tx.Done()
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
	if len(newP) < 8 {
		return "", schema.ErrPasswordTooShort
	}
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	var email string
	tx := newTxErr(s.db)
	tx.err = tx.QueryRow(emailFromRenewableToken, token).Scan(&email)
	tx.err = noRowsTo(tx.err, schema.ErrNoPendingRequests)
	tx.Update(updateUserPassword, email, hash)
	//no need to check updated rows as it is sure the user exists in the tx context
	nb := tx.Update(deletePasswordRenewalRequest, email)
	if tx.err == nil && nb != 1 {
		tx.err = errors.New("Unable to clean the password renewable request of " + email)
	}
	return email, tx.Done()
}

//NewUser add a user
//Basically, calls addUser
func (s *Store) NewUser(p schema.Person, role schema.Privilege) error {
	if !strings.Contains(p.Email, "@") {
		return schema.ErrInvalidEmail
	}
	return s.singleUpdate(insertUser, schema.ErrUserExists, p.Firstname, p.Lastname, p.Tel, p.Email, role, randomBytes(32))
}

//RmUser removes a user from the database
func (s *Store) RmUser(email string) error {
	return s.singleUpdate(deleteUser, schema.ErrUnknownUser, email)
}

//ReplaceUserWith the account referred by src by the account referred by dst
func (s *Store) ReplaceUserWith(src, dst string) error {
	tx := newTxErr(s.db)
	tx.Update(replaceTutorInConventions, src, dst)
	tx.Update(replaceJuryInDefenses, src, dst)
	tx.Update(deleteUser, src)
	return tx.Done()
}
