package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

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
	return Service{backend: backend, r: r}
}

func userMngt(s Service) {

	s.r.HandleFunc("/users/{email}", serviceHandler(user, s)).Methods("GET")
	s.r.HandleFunc("/users/", serviceHandler(users, s)).Methods("GET")
	s.r.HandleFunc("/users/{email}", serviceHandler(rmUser, s)).Methods("DELETE")
	s.r.HandleFunc("/users/{email}/profile", serviceHandler(setProfile, s)).Methods("PUT")
	//The routes
	/*
		TODO
		PUT /users/{email}/role -> SetRole()
		POST /users/ -> NewUser()
		DELETE /users/{email}/password -> ResetPassword()
		PUT /users/{email}/password -> SetPassword()
		POST /users/{email}/password -> NewPassword()
		POST /users/{email}/login -> Registered()
	*/
}

type ProfileUpdate struct {
	Firstname string
	Lastname  string
	Tel       string
}

func setProfile(srv Service, w http.ResponseWriter, r *http.Request) error {
	var p ProfileUpdate
	err := jsonRequest(w, r, &p)
	if err != nil {
		return err
	}
	return srv.backend.SetUserProfile(mux.Vars(r)["email"], p.Firstname, p.Lastname, p.Tel)
}

func user(srv Service, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.backend.User(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, us)
}

func rmUser(srv Service, w http.ResponseWriter, r *http.Request) error {
	return srv.backend.RmUser(mux.Vars(r)["email"])
}

func users(srv Service, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.backend.Users()
	return writeJSONIfOk(err, w, us)
}

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(&j)
}

func writeJSONIfOk(e error, w http.ResponseWriter, j interface{}) error {
	if e != nil {
		return nil
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	return enc.Encode(j)
}

func serviceHandler(cb func(Service, http.ResponseWriter, *http.Request) error, srv Service) http.HandlerFunc {
	//Get the session
	return func(w http.ResponseWriter, r *http.Request) {
		e := cb(srv, w, r)
		switch e {
		case internship.ErrUnknownUser, internship.ErrUnknownReport, internship.ErrUnknownInternship:
			http.Error(w, e.Error(), http.StatusNotFound)
			return
		case internship.ErrReportExists, internship.ErrUserExists, internship.ErrInternshipExists:
			http.Error(w, e.Error(), http.StatusConflict)
			return
		default:
			log.Printf("Unsupported error: %s\n", e.Error())
		}
	}
}
