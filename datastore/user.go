package datastore

import (
	"errors"
	"strings"
	"time"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	AddUser                      = "insert into users(firstname, lastname, tel, email, password, role) values ($1,$2,$3,$4,$5,$6)"
	StartPasswordRenewall        = "insert into password_renewal(email,token,deadline) values($1,$2,$3)"
	NewSession                   = "insert into sessions(email, token, expire) values ($1,$2,$3)"
	UpdateLastVisit              = "update users set lastVisit=$1 where email=$2"
	UpdateUserProfile            = "update users set firstname=$1, lastname=$2, tel=$3 where email=$4"
	UpdateUserRole               = "update users set role=$2 where email=$1"
	UpdateUserPassword           = "update users set password=$2 where email=$1"
	DeleteSession                = "delete from sessions where email=$1"
	DeletePasswordRenewalRequest = "delete from password_renewal where user=$1"
	AllUsers                     = "select firstname, lastname, email, tel, role, lastVisit from users"
	SelectPassword               = "select password from users where email=$1"
	SelectExpire                 = "select expire from sessions where email=$1 and token=$2"
	EmailFromRenewableToken      = "select email from password_renewal where token=$1"
)

//addUser add the given user.
//Every strings are turned into their lower case version
func (s *Service) addUser(tx TxErr, u internship.User) {
	tx.Exec(AddUser, strings.ToLower(u.Firstname), strings.ToLower(u.Lastname), u.Tel, strings.ToLower(u.Email), randomBytes(32), u.Role)
}

//Visit writes the current time for the given user
func (s *Service) Visit(u string) error {
	return s.singleUpdate(UpdateLastVisit, internship.ErrUnknownUser, time.Now(), u)
}

//Users list all the registered users
func (s *Service) Users() ([]internship.User, error) {
	users := make([]internship.User, 0, 0)
	st, err := s.stmt(AllUsers)
	if err != nil {
		return users, err
	}
	rows, err := st.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		u := internship.User{}
		var last pq.NullTime
		rows.Scan(
			&u.Firstname,
			&u.Lastname,
			&u.Email,
			&u.Tel,
			&u.Role,
			&last)
		if last.Valid {
			u.LastVisit = &last.Time
		}
		//Get the role if exists
		users = append(users, u)
	}
	return users, nil
}

//SetUserProfile changes the user profile if exists
func (s *Service) SetUserProfile(email, fn, ln, tel string) error {
	return s.singleUpdate(UpdateUserProfile, internship.ErrUnknownUser, fn, ln, tel, email)
}

//SetUserRole updates the user privilege
func (s *Service) SetUserRole(email string, priv internship.Privilege) error {
	return s.singleUpdate(UpdateUserRole, internship.ErrUnknownUser, email, priv)
}

//Logout destroy the current user session if exists
func (s *Service) Logout(email string) error {
	return s.singleUpdate(DeleteSession, internship.ErrUnknownUser, email)
}

//SetPassword changes the user password if the old one is the current one.
//Any pending renewable requests are deleted on success
func (s *Service) SetPassword(email string, oldP, newP []byte) error {
	//Get the password
	var p []byte
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return err
	}

	tx := newTxErr(s.DB)
	tx.err = tx.tx.QueryRow(SelectPassword, email).Scan(&p)
	tx.err = noRowsTo(tx.err, internship.ErrCredentials)
	if bcrypt.CompareHashAndPassword(p, oldP) != nil {
		tx.err = internship.ErrCredentials
	}
	tx.Exec(DeletePasswordRenewalRequest, email)
	tx.Update(UpdateUserPassword, email, hash)
	//no need to check update, we know the user exists in the transaction context
	return tx.Done()
}

//OpenedSession check if a session is currently open for the user and is not expired
func (s *Service) OpenedSession(email, token string) error {
	var last time.Time
	st, err := s.stmt(SelectExpire)
	if err != nil {
		return err
	}
	err = st.QueryRow(email, token).Scan(&last)
	err = noRowsTo(err, internship.ErrCredentials)
	if err != nil {
		return err
	}
	if time.Now().After(last) {
		return internship.ErrSessionExpired
	}
	return err
}

//ResetPassword starts a reset procedure
//Upon success, it generates a token that will be valid for the next 48h.
func (s *Service) ResetPassword(email string) ([]byte, error) {
	//In case a request already exists
	token := randomBytes(32)
	d, _ := time.ParseDuration("48h")

	tx := newTxErr(s.DB)
	tx.Exec(DeletePasswordRenewalRequest, email)
	tx.Exec(StartPasswordRenewall, email, token, d)
	return token, tx.Done()
}

//NewPassword commits a password renewall request.
//From a request token and a new password, it returns upon success the target user email
func (s *Service) NewPassword(token, newP []byte) (string, error) {
	//Delete possible renew requests
	var email string
	tx := newTxErr(s.DB)
	tx.err = tx.tx.QueryRow(EmailFromRenewableToken, token).Scan(&email)
	tx.err = noRowsTo(tx.err, internship.ErrNoPendingRequests)
	if tx.err != nil {
		return "", tx.Done()
	}
	//Make the new one
	hash, err := hash(newP)
	if err != nil {
		return "", err
	}
	tx.Update(UpdateUserPassword, email, hash)
	//no need to check updated rows as it is sure the user exists in the tx context
	tx.Exec(DeletePasswordRenewalRequest, token)
	return email, tx.Done()
}

//Login log a user.
//Upon success, it generates a session token that will be valid for the next 24 hours.
func (s *Service) Login(email string, password []byte) ([]byte, error) {
	var p []byte
	if err := s.DB.QueryRow(SelectPassword, email).Scan(&p); err != nil {
		return []byte{}, internship.ErrCredentials
	}
	if err := bcrypt.CompareHashAndPassword(p, password); err != nil {
		return []byte{}, internship.ErrCredentials
	}
	tx := newTxErr(s.DB)
	tx.Exec(DeleteSession, email)
	token := randomBytes(32)
	nb := tx.Update(NewSession, email, token, time.Now().Add(time.Hour*24))
	if tx.err == nil && nb != 1 {
		tx.err = internship.ErrUnknownUser
	}
	return token, tx.Done()
}

//AddUser add a user
//Basically, calls addUser
func (s *Service) AddUser(u internship.User) error {
	tx := newTxErr(s.DB)
	s.addUser(tx, u)
	return tx.Done()
}

func (s *Service) RmUser(email string) error {
	/*err := SingleUpdate(s.DB, internship.ErrUnknownUser, "DELETE FROM users where email=$1", email)
	if e, ok := err.(*pq.Error); ok {
		if e.Constraint == "internships_tutor_fkey" {
			return internship.ErrUserTutoring
		}
	}*/
	return errors.New("Unsupported")
}
