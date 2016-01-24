package session

import "github.com/fhermeni/wints/schema"

//RmUser removes an account if the emitter is at least an admin and not himself
func (s *Session) RmUser(email string) error {
	if s.Role().Level() >= schema.AdminLevel && !s.Myself(email) {
		return s.store.RmUser(email)
	}
	return ErrPermission
}

//Users lists all the users if the emitter is at least an admin
func (s *Session) Users() ([]schema.User, error) {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.Users()
	}
	return []schema.User{}, ErrPermission
}

//User returns a given user if the emitter is himself or an admin
func (s *Session) User(em string) (schema.User, error) {
	if s.Myself(em) || s.Role().Level() >= schema.AdminLevel {
		return s.store.User(em)
	}
	return schema.User{}, ErrPermission
}

//SetUserPerson set the user profile if the emitter is the targeted user
func (s *Session) SetUserPerson(p schema.Person) error {
	if s.Myself(p.Email) || s.Role().Level() >= schema.AdminLevel {
		return s.store.SetUserPerson(p)
	}
	return ErrPermission
}

//SetUserRole changes the user privileges if the emitter is an admin at minimum and not himself
func (s *Session) SetUserRole(email string, priv schema.Role) error {
	if s.Role().Level() >= schema.AdminLevel && !s.Myself(email) {
		return s.store.SetUserRole(email, priv)
	}
	return ErrPermission
}

//NewUser creates a new user account if the emitter is an admin at least
func (s *Session) NewUser(p schema.Person, role schema.Role) ([]byte, error) {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.NewUser(p, role)
	}
	return []byte{}, ErrPermission
}

//NewStudent creates a new student account if the emitter is an admin at least
func (s *Session) NewStudent(p schema.Person, major, promotion string, male bool) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.NewStudent(p, major, promotion, male)
	}
	return ErrPermission
}

//ReplaceUserWith allowed if the emitter is an admin at least
func (s *Session) ReplaceUserWith(src, dst string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.ReplaceUserWith(src, dst)
	}
	return ErrPermission
}

//SetEmail change the person email if the emitter is at least an admin
func (s *Session) SetEmail(old, cur string) error {
	if s.Role().Level() >= schema.AdminLevel {
		return s.store.SetEmail(old, cur)
	}
	return ErrPermission
}
