package session

import "github.com/fhermeni/wints/schema"

//RmUser removes an account if the emitter is at least a root
func (s *Session) RmUser(email string) error {
	if s.Role() >= schema.ROOT {
		return s.store.RmUser(email)
	}
	return ErrPermission
}

//Users lists all the users if the emitter is at least an admin
func (s *Session) Users() ([]schema.User, error) {
	if s.Role() >= schema.ADMIN {
		return s.store.Users()
	}
	return []schema.User{}, ErrPermission
}

//SetPassword changes the user password if the emitter is the targeted user
func (s *Session) SetPassword(email string, oldP, newP []byte) error {
	if s.Myself(email) {
		return s.store.SetPassword(email, oldP, newP)
	}
	return ErrPermission
}

//SetUserPerson set the user profile if the emitter is the targeted user
func (s *Session) SetUserPerson(p schema.Person) error {
	if s.Myself(p.Email) {
		return s.store.SetUserPerson(p)
	}
	return ErrPermission
}

//SetUserRole changes the user privileges if the emitter is an admin at minimum
func (s *Session) SetUserRole(email string, priv schema.Privilege) error {
	if s.Role() >= schema.ADMIN {
		return s.store.SetUserRole(email, priv)
	}
	return ErrPermission
}
