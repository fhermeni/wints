package datastore

import (
	"database/sql"
	"strings"
	"time"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	AddUser                      = "insert into users(firstname, lastname, tel, email, password, role) values ($1,$2,$3,$4,$5,$6)"
	UpdateLastVisit              = "update users set lastVisit=$1 where email=$2"
	AllUsers                     = "select firstname, lastname, email, tel, role, lastVisit from users"
	UpdateUserProfile            = "update users set firstname=$1, lastname=$2, tel=$3 where email=$4"
	UpdateUserRole               = "update users set role=$2 where email=$1"
	DeleteSession                = "delete sessions where email=$1 and token=$2"
	SelectPassword               = "select password from users where email=$1"
	DeletePasswordRenewalRequest = "delete from password_renewal where user=$1"
	UpdateUserPassword           = "update users set password=$2 where email=$1"
	SelectExpire                 = "select expire from sessions where email=$1 and token=$2"
	StartPasswordRenewall        = "insert into password_renewal(email,token,deadline) values($1,$2,$3)"
	NewPasswordReq               = "update users set password=$2 where email=$1"
	EmailFromRenewableToken      = "select email from password_renewal where token=$1"
)

type Rollbackable struct {
	tx  *sql.Tx
	err error
}

func newRollbackable(db *sql.DB) Rollbackable {
	tx, err := db.Begin()
	return Rollbackable{tx: tx, err: err}
}

func (r *Rollbackable) Done() error {
	if r.err != nil {
		return r.tx.Rollback()
	}
	return r.tx.Commit()
}

func (r *Rollbackable) Exec(query string, args ...interface{}) {
	if r.err != nil {
		return
	}
	_, err := r.tx.Exec(query, args...)
	if err != nil {
		r.err = err
	}
}

func (r *Rollbackable) Update(query string, args ...interface{}) int64 {
	if r.err != nil {
		return -1
	}
	res, err := r.tx.Exec(query, args...)
	if err != nil {
		r.err = err
		return -1
	}
	nb, err := res.RowsAffected()
	if err != nil {
		r.err = err
		return -1
	}
	return nb
}

//addUser add the given user.
//Every strings are turned into their lower case version
func (s *Service) addUser(tx *sql.Tx, u internship.User) error {
	_, err := tx.Exec(AddUser, strings.ToLower(u.Firstname), strings.ToLower(u.Lastname), u.Tel, strings.ToLower(u.Email), randomBytes(32), u.Role)
	return violationAsErr(err, "pk_email", internship.ErrUserExists)
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
func (s *Service) Logout(email, token string) error {
	return s.singleUpdate(DeleteSession, internship.ErrUnknownUser, email, token)
}

//SetPassword changes the user password if the old one is the current one.
//Any pending renewable requests are deleted on success
/*func (s *Service) SetPassword(email string, oldP, newP []byte) error {
	//Get the password
	var p []byte
	st, err := s.stmt(SelectPassword)
	if err != nil {
		return err
	}
	if err := st.QueryRow(email).Scan(&p); err != nil {
		if err == sql.ErrNoRows {
			return internship.ErrCredentials //hide unknown user
		}
		return err
	}
	if bcrypt.CompareHashAndPassword(p, oldP) != nil {
		return internship.ErrCredentials
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(DeletePasswordRenewalRequest, email)
	if err != nil {
		return rollback(err, tx)
	}

	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return rollback(err, tx)
	}
	res, err := tx.Exec(UpdateUserPassword, email, hash)
	if err != nil {
		return rollback(err, tx)
	}
	if nb, err := res.RowsAffected(); err != nil {
		return rollback(err, tx)
	} else {
		if nb != 1 {
			return rollback(internship.ErrUnknownUser, tx)
		}
		return tx.Commit()
	}
}*/

func (s *Service) SetPassword(email string, oldP, newP []byte) error {
	//Get the password
	var p []byte
	st, err := s.stmt(SelectPassword)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return err
	}

	if err := st.QueryRow(email).Scan(&p); err != nil {
		if err == sql.ErrNoRows {
			return internship.ErrCredentials //hide unknown user
		}
		return err
	}
	if bcrypt.CompareHashAndPassword(p, oldP) != nil {
		return internship.ErrCredentials
	}

	rb := newRollbackable(s.DB)
	rb.Exec(DeletePasswordRenewalRequest, email)
	nb := rb.Update(UpdateUserPassword, email, hash)
	if rb.err == nil && nb != 1 {
		rb.err = internship.ErrUnknownUser
	}
	return rb.Done()
}

//OpenedSession check if a session is currently open for the user and is not expired
func (s *Service) OpenedSession(email, token string) error {
	var last time.Time
	st, err := s.stmt(SelectExpire)
	if err != nil {
		return err
	}
	err = st.QueryRow(email, token).Scan(&last)
	if err != nil {
		if err == sql.ErrNoRows {
			return internship.ErrCredentials
		}
		return err

	}
	if time.Now().After(last) {
		return internship.ErrSessionExpired
	}
	return err
}

//ResetPassword initiates a password renewable request
//The request token is returned upon success
/*func (s *Service) ResetPassword(email string) ([]byte, error) {
	//In case a request already exists
	token := randomBytes(32)
	d, _ := time.ParseDuration("48h")

	tx, err := s.DB.Begin()
	if err != nil {
		return []byte{}, err
	}
	_, err = tx.Exec(DeletePasswordRenewalRequest, email)
	if err != nil {
		return []byte{}, rollback(err, tx)
	}
	_, err = tx.Exec(StartPasswordRenewall, email, token, d)
	if err != nil {
		if violated(err, "fk_password_renewal_email") {
			return []byte{}, rollback(internship.ErrUnknownUser, tx)
		}
		return []byte{}, rollback(err, tx)
	}
	if err := tx.Commit(); err != nil {
		return []byte{}, err
	}
	return token, nil
}*/

func (s *Service) ResetPassword(email string) ([]byte, error) {
	//In case a request already exists
	token := randomBytes(32)
	d, _ := time.ParseDuration("48h")

	rb := newRollbackable(s.DB)
	rb.Exec(DeletePasswordRenewalRequest, email)
	rb.Exec(StartPasswordRenewall, email, token, d)
	rb.err = violationAsErr(rb.err, "fk_password_renewal_email", internship.ErrUnknownUser)
	return token, rb.Done()
}

func (s *Service) NewPassword(token, newP []byte) (string, error) {
	//Delete possible renew requests
	var email string
	err := s.DB.QueryRow(EmailFromRenewableToken, token).Scan(&email)
	if err != nil {
		return "", internship.ErrNoPendingRequests
	}
	//Make the new one
	hash, err := hash(newP)
	if err != nil {
		return "", err
	}
	rb := newRollbackable(s.DB)
	nb := rb.Update(NewPasswordReq, email, hash)
	if nb != 1 && rb.err == nil {
		rb.err = internship.ErrUnknownUser
	}
	rb.Exec(DeletePasswordRenewalRequest, token)
	return email, rb.Done()
}

func (s *Service) Login(email string, password []byte) ([]byte, error) {
	var p []byte
	if err := s.DB.QueryRow("select password from users where email=$1", email).Scan(&p); err != nil {
		return []byte{}, internship.ErrCredentials
	}
	if err := bcrypt.CompareHashAndPassword(p, password); err != nil {
		return []byte{}, internship.ErrCredentials
	}
	_, err := s.DB.Exec("delete from sessions where email=$1", email)
	if err != nil {
		return []byte{}, err
	}
	//token
	token := randomBytes(32)
	err = SingleUpdate(s.DB, internship.ErrUnknownUser, "insert into sessions(email, token, expire) values ($1,$2,$3)", email, token, time.Now().Add(time.Hour*24)) //1 day
	if err != nil {
		return []byte{}, err
	}
	return token, err
}

/*



func (s *Service) NewTutor(p internship.User) ([]byte, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return []byte{}, err
	}
	_, err = tx.Exec("insert into users(email, firstname, lastname, tel, password, role) values ($1,$2,$3,$4,$5,$6)", p.Email, p.Firstname, p.Lastname, p.Tel, randomBytes(32), p.Role)
	if err != nil {
		return []byte{}, internship.ErrUserExists
	}

	token := randomBytes(32)
	d, _ := time.ParseDuration("48h")
	_, err = tx.Exec("insert into password_renewal(email,token,deadline) values($1,$2,$3)", p.Email, token, time.Now().Add(d))
	if err != nil {
		err = rollback(internship.ErrUnknownUser, tx)
		return []byte{}, err
	}
	err = tx.Commit()
	return token, err
}

func (s *Service) RmUser(email string) error {
	err := SingleUpdate(s.DB, internship.ErrUnknownUser, "DELETE FROM users where email=$1", email)
	if e, ok := err.(*pq.Error); ok {
		if e.Constraint == "internships_tutor_fkey" {
			return internship.ErrUserTutoring
		}
	}
	return err
}
*/
