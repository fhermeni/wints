package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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

	s := Service{backend: backend, r: r}
	userMngt(s)
	internshipsMngt(s)
	reportMngt(s)
	conventionMgnt(s)
	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", s.r)
	s.r.HandleFunc("/", home(backend)).Methods("GET")
	return s
}

func home(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		cookie, err := r.Cookie("session")
		if err != nil || len(cookie.Value) == 0 {
			http.ServeFile(w, r, "static/login.html")
			return
		}
		u, err := backend.User(cookie.Value)
		if err == internship.ErrUnknownUser {
			http.ServeFile(w, r, "static/login.html")
			return
		} else if err != nil {
			log.Println("Unable to get the user behind the cookie: " + err.Error())
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if u.Role == internship.STUDENT {
			http.Redirect(w, r, "student", 302)
		} else {
			http.Redirect(w, r, "admin", 302)
		}
	}
}

func (s *Service) Listen(host string, cert, pk string) error {
	return http.ListenAndServeTLS(host, cert, pk, nil)
}

func userMngt(s Service) {

	s.r.HandleFunc("/api/v1/users/{email}", serviceHandler(user, s)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/", serviceHandler(users, s)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/{email}", serviceHandler(rmUser, s)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/profile", serviceHandler(setUserProfile, s)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/{email}/role", serviceHandler(setUserRole, s)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/", serviceHandler(newUser, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/users/{email}/password", resetPassword(s.backend)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/password", serviceHandler(setPassword, s)).Methods("PUT")
	s.r.HandleFunc("/api/v1/newPassword", newPassword(s.backend)).Methods("POST")
	s.r.HandleFunc("/api/v1/login", login(s.backend)).Methods("POST")  //FAIL
	s.r.HandleFunc("/api/v1/logout", logout(s.backend)).Methods("GET") //FAIL
	s.r.HandleFunc("/admin", serviceHandler(admin, s)).Methods("GET")
	s.r.HandleFunc("/student", serviceHandler(student, s)).Methods("GET")
	s.r.HandleFunc("/resetPassword", password).Methods("GET")
}

func admin(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	http.ServeFile(w, r, "static/admin.html")
	return nil
}

func student(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	http.ServeFile(w, r, "static/student.html")
	return nil
}

func password(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new_password.html")
}

type PasswordUpdate struct {
	Old string
	New string
}

type NewPassword struct {
	Token string
	New   string
}

func login(srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.PostFormValue("login")
		password := r.PostFormValue("password")
		_, err := srv.Registered(login, []byte(password))
		if err != nil {
			http.Redirect(w, r, "/#badLogin", 302)
			return
		}
		cookie := &http.Cookie{
			Name:  "session",
			Value: login,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", 302)
	}
}

func logout(srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", 302)
	}
}

func newPassword(srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PostFormValue("token")
		password := r.PostFormValue("password")
		email, err := srv.NewPassword([]byte(token), []byte(password))
		switch err {
		case internship.ErrNoPendingRequests:
			http.Redirect(w, r, "/resetPassword?token="+token+"#noRequest", 302)
			return
		case internship.ErrUnknownUser:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case nil:
			//log the user
			cookie := &http.Cookie{
				Name:  "session",
				Value: email,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", 302)
		default:
			log.Println("Unable to reset the password for token " + token + ": " + err.Error())
			http.Error(w, "Unable to reset the password", http.StatusInternalServerError)
			return
		}
	}
}

func setPassword(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var u PasswordUpdate
	err := jsonRequest(w, r, &u)
	e := mux.Vars(r)["email"]
	if err != nil {
		return err
	}
	return srv.SetUserPassword(e, []byte(u.Old), []byte(u.New))
}

func resetPassword(srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := mux.Vars(r)["email"]
		token, err := srv.ResetPassword(e)
		if err != nil {
			log.Println("Unable to get the reset token for " + e + ": " + err.Error())
			return
		}
		log.Println("Reset token for " + e + " " + string(token))
	}
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
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}", serviceHandler(report, s)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", serviceHandler(reportContent, s)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", serviceHandler(setReportContent, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/grade", serviceHandler(setReportGrade, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/dealine", serviceHandler(setReportDeadline, s)).Methods("POST")
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

func conventionMgnt(s Service) {
	s.r.HandleFunc("/api/v1/conventions/", serviceHandler(conventions, s)).Methods("GET")
}

func conventions(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	c, err := srv.Conventions()
	return writeJSONIfOk(err, w, c)
}

func internshipsMngt(s Service) {
	s.r.HandleFunc("/api/v1/internships/{email}", serviceHandler(getInternship, s)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", serviceHandler(internships, s)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", serviceHandler(newInternship, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/supervisor", serviceHandler(setSupervisor, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/company", serviceHandler(setCompany, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/major", serviceHandler(setMajor, s)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/promotion", serviceHandler(setPromotion, s)).Methods("POST")
}

func newInternship(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c internship.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.NewInternship(c)
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

func setMajor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c string
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetMajor(mux.Vars(r)["email"], c)
}

func setPromotion(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c string
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetPromotion(mux.Vars(r)["email"], c)
}

func setSupervisor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var s internship.Person
	err := jsonRequest(w, r, &s)
	if err != nil {
		return err
	}
	return srv.SetSupervisor(mux.Vars(r)["email"], s)
}
