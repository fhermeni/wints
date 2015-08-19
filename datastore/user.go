package datastore

import (
	"database/sql"
	"strings"
	"time"

	"github.com/fhermeni/wints/internship"
)

var (
	AddUser         = "insert into users(firstname, lastname, tel, email, password, role) values ($1,$2,$3,$4,$5,$6)"
	UpdateLastVisit = "update users set lastVisit=$1 where email=$2"
)

//addUser add the given user.
//Every strings are turned into their lower case version
func (s *Service) addUser(tx *sql.Tx, u internship.User) error {
	_, err := tx.Exec(AddUser, strings.ToLower(u.Firstname), strings.ToLower(u.Lastname), u.Tel, strings.ToLower(u.Email), randomBytes(32), u.Role)
	if violated(err, "pk_email") {
		return internship.ErrUserExists
	}
	return err
}

//Visit writes the current time for the given user
func (s *Service) Visit(u string) error {
	return s.singleUpdate(UpdateLastVisit, internship.ErrUnknownUser, time.Now(), u)
}

//Prepared requests
/*

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

func (s *Service) OpenedSession(email, token string) error {
	var last time.Time
	err := s.DB.QueryRow("select expire from sessions where email=$1 and token=$2", email, token).Scan(&last)
	if err != nil {
		return internship.ErrCredentials
	}
	if time.Now().After(last) {
		return internship.ErrSessionExpired
	}
	return err
}

func (s *Service) Logout(email, token string) error {
	err := SingleUpdate(s.DB, internship.ErrUnknownUser, "update sessions set token=$3 where email=$1 and token=$2", token, email, randomBytes(32)) //unknown token (to remind last visit)
	return err
}

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

func (s *Service) Users() ([]internship.User, error) {
	rows, err := s.DB.Query("select firstname, lastname, email, tel, role, lastVisit from users")
	users := make([]internship.User, 0, 0)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var fn, ln, email, tel string
		var role internship.Privilege
		var last pq.NullTime
		rows.Scan(&fn, &ln, &email, &tel, &role, &last)
		u := internship.User{Firstname: fn, Lastname: ln, Email: email, Tel: tel, Role: role}
		if last.Valid {
			u.LastLogin = last.Value()
		}
		//Get the role if exists
		users = append(users, u)
	}
	return users, nil
}

func (s *Service) SetUserPassword(email string, oldP, newP []byte) error {
	//Get the password
	var p []byte
	if err := s.DB.QueryRow("select password from users where email=$1", email).Scan(&p); err != nil {
		return internship.ErrCredentials
	}
	if bcrypt.CompareHashAndPassword(p, oldP) != nil {
		return internship.ErrCredentials
	}

	//Delete possible renew requests
	s.DB.QueryRow("delete from password_renewal where user=$1", email)

	//Make the new one
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return err
	}
	return SingleUpdate(s.DB, internship.ErrUnknownUser, "update users set password=$2 where email=$1", email, hash)
}

func (s *Service) SetUserProfile(email, fn, ln, tel string) error {
	return SingleUpdate(s.DB, internship.ErrUnknownUser, "update users set firstname=$1, lastname=$2, tel=$3 where email=$4", fn, ln, tel, email)
}

func (s *Service) SetUserRole(email string, priv internship.Privilege) error {
	return SingleUpdate(s.DB, internship.ErrUnknownUser, "update users set role=$2 where email=$1", email, priv)
}

func (s *Service) NewPassword(token, newP []byte) (string, error) {
	//Delete possible renew requests
	var email string
	err := s.DB.QueryRow("select email from password_renewal where token=$1", token).Scan(&email)
	if err != nil {
		return "", internship.ErrNoPendingRequests
	}
	//Make the new one
	hash, err := hash(newP)
	if err != nil {
		return "", err
	}
	if err := SingleUpdate(s.DB, internship.ErrUnknownUser, "update users set password=$2 where email=$1", email, hash); err != nil {
		return "", err
	}
	_, err = s.DB.Exec("delete from password_renewal where token=$1", token)
	return email, err
}

func (s *Service) ResetPassword(email string) ([]byte, error) {
	//In case a request already exists
	s.DB.QueryRow("delete from password_renewal where email=$1", email)
	token := randomBytes(32)
	d, _ := time.ParseDuration("48h")
	_, err := s.DB.Exec("insert into password_renewal(email,token,deadline) values($1,$2,$3)", email, token, time.Now().Add(d))
	if err != nil {
		log.Println(err.Error())
		//Not very sure, might be a SQL error or a connexion error as well.
		return []byte{}, internship.ErrUnknownUser
	}
	return token, err
}
*/
