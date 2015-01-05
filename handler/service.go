package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/mail"
	"github.com/gorilla/mux"
)

type Service struct {
	backend internship.Service
	mailer  mail.Mailer
	r       *mux.Router
}

var (
	ErrJSONWriting = errors.New("Unable to serialize the JSON response")
)

func NewService(backend internship.Service, mailer mail.Mailer) Service {
	//Rest stuff
	r := mux.NewRouter()

	s := Service{backend: backend, r: r, mailer: mailer}
	userMngt(s, mailer)
	internshipsMngt(s, mailer)
	reportMngt(s, mailer)
	conventionMgnt(s, mailer)
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

		if u.Role == internship.NONE {
			http.Redirect(w, r, "student", 302)
		} else {
			http.Redirect(w, r, "admin", 302)
		}
	}
}

func (s *Service) Listen(host string, cert, pk string) error {
	return http.ListenAndServeTLS(host, cert, pk, nil)
}

func userMngt(s Service, mailer mail.Mailer) {

	s.r.HandleFunc("/api/v1/users/{email}", serviceHandler(user, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/", serviceHandler(users, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/{email}", serviceHandler(rmUser, s, mailer)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/profile", serviceHandler(setUserProfile, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/{email}/role", serviceHandler(setUserRole, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/", serviceHandler(newTutor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/users/{email}/password", resetPassword(s.backend, mailer)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/password", serviceHandler(setPassword, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/newPassword", newPassword(s.backend, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/login", login(s.backend)).Methods("POST")  //FAIL
	s.r.HandleFunc("/api/v1/logout", logout(s.backend)).Methods("GET") //FAIL
	s.r.HandleFunc("/admin", serviceHandler(admin, s, mailer)).Methods("GET")
	s.r.HandleFunc("/student", serviceHandler(student, s, mailer)).Methods("GET")
	s.r.HandleFunc("/resetPassword", password).Methods("GET")
}

func admin(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	http.ServeFile(w, r, "static/admin.html")
	return nil
}

func student(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
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

func newPassword(srv internship.Service, mailer mail.Mailer) http.HandlerFunc {
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

func setPassword(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var u PasswordUpdate
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	e := mux.Vars(r)["email"]
	return srv.SetUserPassword(e, []byte(u.Old), []byte(u.New))
}

func resetPassword(srv internship.Service, mailer mail.Mailer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := mux.Vars(r)["email"]
		token, err := srv.ResetPassword(e)
		if err != nil {
			log.Println("Unable to get the reset token for " + e + ": " + err.Error())
			return
		}
		if u, err := srv.User(e); err == nil {
			mailer.SendPasswordResetLink(u, token)
		}
	}
}

func newTutor(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var u internship.User
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	t, err := srv.NewTutor(u)
	if err == nil {
		mailer.SendAdminInvitation(u, t)
	}
	return err
}

func setUserRole(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var p internship.Privilege
	err := jsonRequest(w, r, &p)
	if err != nil {
		return err
	}
	em := mux.Vars(r)["email"]
	err = srv.SetUserRole(em, p)
	if err != nil {
		return err
	}
	if u, err := srv.User(em); err == nil {
		mailer.SendRoleUpdate(u)
	}
	return nil
}

type ProfileUpdate struct {
	Firstname string
	Lastname  string
	Tel       string
}

func setUserProfile(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var p ProfileUpdate
	err := jsonRequest(w, r, &p)
	if err != nil {
		return err
	}
	return srv.SetUserProfile(mux.Vars(r)["email"], p.Firstname, p.Lastname, p.Tel)
}

func user(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.User(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, us)
}

func rmUser(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	u, err := srv.User(mux.Vars(r)["email"])
	if err != nil {
		return err
	}
	err = srv.RmUser(u.Email)
	if err == nil {
		mailer.SendAccountRemoval(u)
	}
	return err
}

func users(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.Users()
	return writeJSONIfOk(err, w, us)
}

func reportMngt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}", serviceHandler(report, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", serviceHandler(reportContent, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", serviceHandler(setReportContent, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/grade", serviceHandler(setReportGrade, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/deadline", serviceHandler(setReportDeadline, s, mailer)).Methods("POST")
}

func report(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	h, err := srv.Report(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, h)
}

func reportContent(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	cnt, err := srv.ReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return fileReplyIfOk(err, w, "application/pdf", mux.Vars(r)["kind"]+".pdf", cnt)
}

func setReportContent(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	cnt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = srv.SetReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"], cnt)
	if err == nil {
		if i, err := srv.Internship(mux.Vars(r)["email"]); err == nil {
			mailer.SendReportUploaded(i.Student, i.Tutor, mux.Vars(r)["kind"])
		}
	}
	return err
}

type ReportGrade struct {
	Grade   int
	Comment string
}

func setReportGrade(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var n ReportGrade
	err := jsonRequest(w, r, &n)
	if err != nil {
		return err
	}
	k := mux.Vars(r)["kind"]
	em := mux.Vars(r)["email"]
	err = srv.SetReportGrade(k, em, n.Grade, n.Comment)

	if err == nil {
		if i, err := srv.Internship(em); err == nil {
			mailer.SendGradeUploaded(i.Student, i.Tutor, k)
		}
	}
	return err
}
func setReportDeadline(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var t time.Time
	err := jsonRequest(w, r, &t)
	if err != nil {
		return err
	}
	err = srv.SetReportDeadline(mux.Vars(r)["kind"], mux.Vars(r)["email"], t)
	if err == nil {
		if i, err := srv.Internship(mux.Vars(r)["email"]); err != nil {
			mailer.SendReportDeadline(i.Student, i.Tutor, mux.Vars(r)["kind"], t)
		}
	}
	return err
}

func conventionMgnt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/api/v1/conventions/", serviceHandler(conventions, s, mailer)).Methods("GET")
}

func conventions(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	c, err := srv.Conventions()
	return writeJSONIfOk(err, w, c)
}

func internshipsMngt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/api/v1/internships/{email}", serviceHandler(getInternship, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", serviceHandler(internships, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", serviceHandler(newInternship, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/supervisor", serviceHandler(setSupervisor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/tutor", serviceHandler(setTutor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/company", serviceHandler(setCompany, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/major", serviceHandler(setMajor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/promotion", serviceHandler(setPromotion, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/majors/", serviceHandler(majors, s, mailer)).Methods("GET")
}

func majors(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	m := srv.Majors()
	return writeJSONIfOk(nil, w, m)
}
func newInternship(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c internship.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	token, err := srv.NewInternship(c)
	if err == nil {
		mailer.SendStudentInvitation(internship.User{Firstname: c.Student.Firstname, Lastname: c.Student.Lastname, Email: c.Student.Email}, token)
	}
	return err
}

func getInternship(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internship(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, i)
}

func internships(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internships()
	return writeJSONIfOk(err, w, i)
}

func setCompany(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c internship.Company
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetCompany(mux.Vars(r)["email"], c)
}

func setMajor(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c string
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetMajor(mux.Vars(r)["email"], c)
}

func setPromotion(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c string
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetPromotion(mux.Vars(r)["email"], c)
}

func setSupervisor(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var s internship.Person
	err := jsonRequest(w, r, &s)
	if err != nil {
		return err
	}
	return srv.SetSupervisor(mux.Vars(r)["email"], s)
}

func setTutor(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var s string
	err := jsonRequest(w, r, &s)
	if err != nil {
		return err
	}
	em := mux.Vars(r)["email"]
	i, err := srv.Internship(em)
	if err != nil {
		return err
	}
	err = srv.SetTutor(em, s)
	if err == nil {
		if u, err := srv.User(s); err == nil {
			mailer.SendTutorUpdate(i.Student, i.Tutor, u)
		}
	}
	return err
}
