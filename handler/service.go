package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fhermeni/wints/internship"
	"github.com/gorilla/mux"
)

type Service struct {
	backend internship.Service
	r       *mux.Router
}

var (
	ErrJSONWriting = errors.New("Unable to serialize the JSON response")
)

func NewService(backend internship.Service) Service {
	//Rest stuff
	r := mux.NewRouter()

	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", r)
	return Service{backend: backend, r: r}
}

func (s *Service) Listen(p int, cert, pk string) error {
	return http.ListenAndServeTLS(":"+strconv.Itoa(p), cert, pk, nil)
}

func userMngt(s Service) {

	s.r.HandleFunc("/users/{email}", serviceHandler(user, s)).Methods("GET")
	s.r.HandleFunc("/users/", serviceHandler(users, s)).Methods("GET")
	s.r.HandleFunc("/users/{email}", serviceHandler(rmUser, s)).Methods("DELETE")
	s.r.HandleFunc("/users/{email}/profile", serviceHandler(setUserProfile, s)).Methods("PUT")
	s.r.HandleFunc("/users/{email}/role", serviceHandler(setUserRole, s)).Methods("PUT")
	s.r.HandleFunc("/users/", serviceHandler(newUser, s)).Methods("POST")
	s.r.HandleFunc("/users/{email}/password", serviceHandler(resetPassword, s)).Methods("DELETE")
	s.r.HandleFunc("/users/{email}/password", serviceHandler(setPassword, s)).Methods("PUT")
	s.r.HandleFunc("/newPassword", serviceHandler(newPassword, s)).Methods("POST")
	s.r.HandleFunc("/login", serviceHandler(login, s)).Methods("POST") //FAIL
}

type PasswordUpdate struct {
	Old []byte
	New []byte
}

type NewPassword struct {
	Token []byte
	New   []byte
}

func login(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var email string
	var password []byte
	_, err := srv.Registered(email, password)
	return err
}

func newPassword(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var u NewPassword
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	_, err = srv.NewPassword(u.Token, u.New) //Email of the guy
	return err
}

func setPassword(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	e := mux.Vars(r)["email"]
	var u PasswordUpdate
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	return srv.SetUserPassword(e, u.Old, u.New)
}

func resetPassword(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	e := mux.Vars(r)["email"]
	_, err := srv.ResetPassword(e) //skip the token, should go to the mail
	return err
}

func newUser(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var u internship.User
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	_, err = srv.NewUser(u) //ignore the password
	return err
}

func setUserRole(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var p internship.Privilege
	err := jsonRequest(w, r, &p)
	if err != nil {
		return err
	}
	return srv.SetUserRole(mux.Vars(r)["email"], p)
}

type ProfileUpdate struct {
	Firstname string
	Lastname  string
	Tel       string
}

func setUserProfile(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var p ProfileUpdate
	err := jsonRequest(w, r, &p)
	if err != nil {
		return err
	}
	return srv.SetUserProfile(mux.Vars(r)["email"], p.Firstname, p.Lastname, p.Tel)
}

func user(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.User(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, us)
}

func rmUser(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	return srv.RmUser(mux.Vars(r)["email"])
}

func users(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.Users()
	return writeJSONIfOk(err, w, us)
}
