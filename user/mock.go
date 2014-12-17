package user

type MockService struct {	
	Users map[string]User
	Passwords map[string]string
	Tokens map[string]string
}

func NewMock() *MockService {
	return &MockService{make(map[string]User),
		make(map[string]string),
		make(map[string]string)}
}

func (s *MockService) Register(email, password string) (User, error) {	
	p, ok := s.Passwords[email]
	if !ok || p != password {
		return s.Users[email], ErrCredentials
	}
	return s.Users[email],nil
}

func (s *MockService) New(p User) error {
	_, ok := s.Users[p.Email]
	if ok  {
		return ErrExists
	}
	newPassword := p.Firstname + p.Lastname
	s.Passwords[p.Email] = newPassword
	return nil
}

func (s *MockService) Get(email string) (User, error) {
	u, ok := s.Users[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

func (s *MockService) List(roles ...string) ([]User, error) {	
	users := make([]User, 0, 0)
	return users, nil
}

func (s *MockService) SetPassword(email string, oldP, newP []byte) error {
	p, ok := s.Passwords[email]
	if !ok || p != string(oldP) {
		return ErrCredentials
	}
	s.Passwords[email] = string(newP)	
	return nil
}

func (s *MockService) SetProfile(email, fn, ln, tel string) error {
	u, ok := s.Users[email]
	if !ok {
		return ErrNotFound
	}
	u.Firstname = fn
	u.Lastname = ln
	u.Tel = tel
	s.Users[email] = u
	return nil
}

func (s *MockService) SetRole(email, priv string) error {
	u, ok := s.Users[email]
	if !ok {
		return ErrNotFound
	}
	u.Role = priv
	s.Users[email] = u
	return nil	
}

func (s *MockService) NewPassword(token string, newP []byte) (string, error) {
	_, ok := s.Tokens[token]
	if !ok {
		return "", ErrNotFound
	}
	s.Passwords[token] = string(newP)
	return token, nil
}

func (s *MockService) Rm(email string) error {
	_, ok := s.Users[email]
	if !ok {
		return ErrNotFound
	}
	return ErrNotFound
}

func (s *MockService) ResetPassword(email string) (string, error) {
	_, ok := s.Users[email]
	if !ok {
		return "",ErrNotFound
	}
	s.Tokens[email] = email
	return email, nil
}
