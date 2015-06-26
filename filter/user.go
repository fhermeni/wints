package filter

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

func (v *Service) Registered(email string, password []byte) ([]byte, error) {
	cnt, err := v.srv.Registered(email, password)
	v.log.Log(email, "log in", err)
	return cnt, err
}

func (v *Service) OpenedSession(email, token string) error {
	err := v.srv.OpenedSession(email, token)
	v.log.Log(email, "continues his session", err)
	return err
}

func (v *Service) Logout(email, token string) error {
	err := v.srv.Logout(email, token)
	v.UserLog("logout '"+email+"'", err)
	return err
}

func (v *Service) Sessions() (map[string]time.Time, error) {
	err := ErrPermission
	sessions := map[string]time.Time{}
	if v.my.Role == internship.ROOT {
		sessions, err = v.srv.Sessions()
	}
	return sessions, err
}
func (v *Service) NewTutor(p internship.User) ([]byte, error) {
	err := ErrPermission
	token := []byte{}
	if v.my.Role == internship.ROOT {
		token, err = v.srv.NewTutor(p)
	}
	v.UserLog("new tutor '"+p.Email+"'", err)
	if err == nil {
		v.mailer.SendAdminInvitation(p, token)
	}

	return token, err
}

func (v *Service) RmUser(email string) error {
	err := ErrPermission
	if v.my.Role == internship.ROOT && email != v.my.Email {
		err = v.srv.RmUser(email)
	}
	v.UserLog("delete user '"+email+"'", err)
	if err == nil {
		u, err := v.srv.User(email)
		if err == nil {
			v.mailer.SendAccountRemoval(u)
		}
	}
	return err
}

func (v *Service) User(email string) (internship.User, error) {
	u := internship.User{}
	err := ErrPermission
	if v.mine(email) || v.my.Role >= internship.MAJOR {
		u, err = v.srv.User(email)
	}
	return u, err
}

func (v *Service) Users() ([]internship.User, error) {
	if v.my.Role >= internship.MAJOR {
		return v.srv.Users()
	}
	return []internship.User{v.my}, nil
}

//Change the user password
func (v *Service) SetUserPassword(email string, oldP, newP []byte) error {
	err := ErrPermission
	if v.my.Email == email {
		err = v.srv.SetUserPassword(email, oldP, newP)
	}
	v.UserLog("password changed", err)
	return err
}

//Change user profile
func (v *Service) SetUserProfile(email, fn, ln, tel string) error {
	err := ErrPermission
	if v.my.Email == email {
		err = v.srv.SetUserProfile(email, fn, ln, tel)
	}
	v.UserLog("profil update", err)
	return err
}

//Change user role
func (v *Service) SetUserRole(email string, priv internship.Privilege) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN && email != v.my.Email {
		err = v.srv.SetUserRole(email, priv)
	}
	v.UserLog("set role of '"+email+"' to '"+priv.String()+"'", err)
	if u, err := v.srv.User(email); err == nil {
		v.mailer.SendRoleUpdate(u)
	}
	return err
}

func (v *Service) ResetPassword(email string) ([]byte, error) {
	cnt, err := v.srv.ResetPassword(email)
	v.log.Log(email, "password reset", err)
	return cnt, err
}

func (v *Service) NewPassword(token, newP []byte) (string, error) {
	cnt, err := v.srv.NewPassword(token, newP)
	v.log.Log(string(token), "new password", err)
	return cnt, err
}
