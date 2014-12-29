package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

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
	s := Service{backend: backend, r: r}
	userMngt(s)
	internshipsMngt(s)
	reportMngt(s)
	return s
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

func reportMngt(s Service) {
	s.r.HandleFunc("/internships/{email}/reports/{kind}", serviceHandler(report, s)).Methods("GET")
	s.r.HandleFunc("/internships/{email}/reports/{kind}/content", serviceHandler(reportContent, s)).Methods("GET")
	s.r.HandleFunc("/internships/{email}/reports/{kind}/content", serviceHandler(setReportContent, s)).Methods("POST")
	s.r.HandleFunc("/internships/{email}/reports/{kind}/grade", serviceHandler(setReportGrade, s)).Methods("POST")
	s.r.HandleFunc("/internships/{email}/reports/{kind}/dealine", serviceHandler(setReportDeadline, s)).Methods("POST")
	/*
		PlanReport(student string, r ReportHeader) error -> NO REST CALL
	*/
}

func report(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	h, err := srv.Report(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, h)
}

func reportContent(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	cnt, err := srv.ReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return fileReplyIfOk(err, w, "application/pdf", mux.Vars(r)["kind"]+".pdf", cnt)
}

func setReportContent(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	cnt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = srv.SetReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"], cnt)
	return err
}

func setReportGrade(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var n int
	err := jsonRequest(w, r, &n)
	if err != nil {
		return err
	}
	err = srv.SetReportGrade(mux.Vars(r)["kind"], mux.Vars(r)["email"], n)
	return err
}
func setReportDeadline(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var t time.Time
	err := jsonRequest(w, r, &t)
	if err != nil {
		return err
	}
	err = srv.SetReportDeadline(mux.Vars(r)["kind"], mux.Vars(r)["email"], t)
	return err
}

func internshipsMngt(s Service) {
	s.r.HandleFunc("/internships/{email}", serviceHandler(getInternship, s)).Methods("GET")
	s.r.HandleFunc("/internships/", serviceHandler(internships, s)).Methods("GET")
	s.r.HandleFunc("/internships/{email}/supervisor", serviceHandler(setSupervisor, s)).Methods("POST")
	s.r.HandleFunc("/internships/{email}/company", serviceHandler(setCompany, s)).Methods("POST")
}

func getInternship(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internship(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, i)
}

func internships(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internships()
	return writeJSONIfOk(err, w, i)
}

func setCompany(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c internship.Company
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetCompany(mux.Vars(r)["email"], c)
}

func setSupervisor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var s internship.Supervisor
	err := jsonRequest(w, r, &s)
	if err != nil {
		return err
	}
	return srv.SetSupervisor(mux.Vars(r)["email"], s)
}
