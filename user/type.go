package user

import (
	"crypto/rand"
	"errors"
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Tel       string
	Role      string
}

var (
	ErrExists            = errors.New("User already exists")
	ErrNotFound          = errors.New("User not found")
	ErrUserTutoring      = errors.New("The user is tutoring students")
	ErrCredentials       = errors.New("Incorrect credentials")
	ErrNoPendingRequests = errors.New("No password renewable request pending")
)

func (p User) String() string {
	return p.Firstname + " " + p.Lastname + " (" + p.Email + ")"
}

func (p User) Fullname() string {
	return p.Firstname + " " + p.Lastname
}

type UserService interface {
	Register(email, password string) (User, error)
	New(p User) error
	Rm(email string) error
	Get(email string) (User, error)
	List(roles ...string) ([]User, error)
	SetPassword(email string, oldP, newP []byte) error
	SetProfile(email, fn, ln, tel string) error
	SetRole(email, priv string) error
	NewPassword(token string, newP []byte) (string, error)
	ResetPassword(email string) (string, error)
}

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
