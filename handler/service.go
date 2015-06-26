package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/mail"
	"github.com/gorilla/mux"

	"github.com/daaku/go.httpgzip"
)

var path string

type Service struct {
	backend internship.Service
	mailer  mail.Mailer
	r       *mux.Router
}

var (
	ErrJSONWriting    = errors.New("Unable to serialize the JSON response")
	ErrMissingCookies = errors.New("Cookies 'session' or 'token' are missing or empty")
)

func NewService(backend internship.Service, mailer mail.Mailer, p string, j journal.Journal) Service {
	//Rest stuff
	r := mux.NewRouter()
	path = p
	s := Service{backend: backend, r: r, mailer: mailer}
	userMngt(s, mailer, j)
	internshipsMngt(s, mailer, j)
	reportMngt(s, mailer, j)
	conventionMgnt(s, mailer, j)
	surveyMngt(s, mailer, j)
	defenseMngt(s, mailer, j)
	fs := http.Dir(path + "/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/" + path + "/").Handler(httpgzip.NewHandler(http.StripPrefix("/"+path, fileHandler)))
	http.Handle("/", s.r)
	s.r.HandleFunc("/", mon(j, home(backend))).Methods("GET")
	s.r.HandleFunc("/statistics", mon(j, asset(filepath.Join(path, "statistics-new.html")))).Methods("GET")
	s.r.HandleFunc("/defense-program", mon(j, asset(filepath.Join(path, "defense-program.html")))).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}", surveyFromToken(backend)).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}", setSurveyContent(backend)).Methods("POST")
	s.r.HandleFunc("/api/v1/statistics/", mon(j, statistics(backend))).Methods("GET")
	s.r.HandleFunc("/api/v1/majors/", mon(j, majors(backend))).Methods("GET")
	return s
}

func majors(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := backend.Majors()
		writeJSONIfOk(nil, w, r, m)
	}
}

func statistics(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := backend.Statistics()
		writeJSONIfOk(err, w, r, stats)
	}
}

func home(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, err := authenticated(backend, w, r)
		if err != nil {
			http.ServeFile(w, r, path+"/login.html")
			return
		}
		u, err := backend.User(email)
		if err != nil {
			http.ServeFile(w, r, path+"/login.html")
			return
		}
		if u.Role == internship.NONE {
			http.ServeFile(w, r, path+"/student.html")
		} else {
			http.ServeFile(w, r, path+"/admin.html")
		}
	}
}

func (s *Service) Listen(host string, cert, pk string) error {
	return http.ListenAndServeTLS(host, cert, pk, nil)
}

func userMngt(s Service, mailer mail.Mailer, j journal.Journal) {
	s.r.HandleFunc("/api/v1/users/{email}", restHandler(user, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/", restHandler(users, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/{email}", restHandler(rmUser, j, s, mailer)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/profile", restHandler(setUserProfile, j, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/{email}/role", restHandler(setUserRole, j, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/", restHandler(newTutor, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/users/{email}/password", mon(j, resetPassword(j, s.backend, mailer))).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/password", restHandler(setPassword, j, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/sessions/", restHandler(sessions, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/newPassword", mon(j, newPassword(j, s.backend, mailer))).Methods("POST")
	s.r.HandleFunc("/api/v1/login", mon(j, login(j, s.backend))).Methods("POST")
	s.r.HandleFunc("/api/v1/logout", mon(j, logout(j, s.backend))).Methods("GET")
	s.r.HandleFunc("/resetPassword", mon(j, asset(filepath.Join(path, "new_password.html")))).Methods("GET")
}

type PasswordUpdate struct {
	Old string
	New string
}

type NewPassword struct {
	Token string
	New   string
}

func login(j journal.Journal, srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.PostFormValue("login")
		password := r.PostFormValue("password")
		t, err := srv.Registered(login, []byte(password))
		j.Log(login, "login", err)
		if err != nil {
			http.Redirect(w, r, "/#badLogin", 302)
			return
		}
		cookie := &http.Cookie{
			Name:  "session",
			Value: login,
			Path:  "/",
		}
		token := &http.Cookie{
			Name:  "token",
			Value: string(t),
			Path:  "/",
		}

		http.SetCookie(w, cookie)
		http.SetCookie(w, token)
		http.Redirect(w, r, "/", 302)
	}
}

func logout(j journal.Journal, srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := &http.Cookie{
			Name:   "session",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}

		token := &http.Cookie{
			Name:   "token",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, email)
		http.SetCookie(w, token)
		http.Redirect(w, r, "/", 302)
	}
}

func newPassword(j journal.Journal, srv internship.Service, mailer mail.Mailer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PostFormValue("token")
		password := r.PostFormValue("password")
		email, err := srv.NewPassword([]byte(token), []byte(password))
		j.Log(token, "reset his password (email='"+email+"')", err)
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
			http.Redirect(w, r, "/?email="+url.QueryEscape(email), 302)
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

func resetPassword(j journal.Journal, srv internship.Service, mailer mail.Mailer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		em := mux.Vars(r)["email"]
		_, err := srv.ResetPassword(em)
		j.Log(em, "initiate password reset", err)
		status(w, err)
	}
}

func sessions(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	res, err := srv.Sessions()
	return writeJSONIfOk(err, w, r, res)
}

func newTutor(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var u internship.User
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	_, err = srv.NewTutor(u)
	return err
}

func setUserRole(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var p internship.Privilege
	err := jsonRequest(w, r, &p)
	if err != nil {
		return err
	}
	em := mux.Vars(r)["email"]
	return srv.SetUserRole(em, p)
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
	return writeJSONIfOk(err, w, r, us)
}

func rmUser(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	u, err := srv.User(mux.Vars(r)["email"])
	if err != nil {
		return err
	}
	return srv.RmUser(u.Email)
}

func users(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.Users()
	return writeJSONIfOk(err, w, r, us)
}

func reportMngt(s Service, mailer mail.Mailer, j journal.Journal) {
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}", restHandler(report, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", restHandler(reportContent, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", restHandler(setReportContent, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/grade", restHandler(setReportGrade, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/deadline", restHandler(setReportDeadline, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/private", restHandler(setReportPrivate, j, s, mailer)).Methods("POST")
}

func report(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	h, err := srv.Report(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, r, h)
}

func reportContent(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	cnt, err := srv.ReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return fileReplyIfOk(err, w, "application/pdf", mux.Vars(r)["kind"]+".pdf", cnt)
}

func setReportContent(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	fi, fh, err := r.FormFile("report")
	if err != nil {
		return err
	}
	cnt, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}
	mime := fh.Header.Get("content-type")
	if mime != "application/pdf" {
		http.Error(w, "The report must be a PDF", http.StatusBadRequest)
		return errors.New("The report must be a PDF")
	}
	em := mux.Vars(r)["email"]
	k := mux.Vars(r)["kind"]
	err = srv.SetReportContent(k, em, cnt)
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
	return srv.SetReportGrade(k, em, n.Grade, n.Comment)
}

func setReportDeadline(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var t time.Time
	err := jsonRequest(w, r, &t)
	if err != nil {
		return err
	}
	return srv.SetReportDeadline(mux.Vars(r)["kind"], mux.Vars(r)["email"], t)
}

func setReportPrivate(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var b bool
	err := jsonRequest(w, r, &b)
	if err != nil {
		return err
	}
	return srv.SetReportPrivate(mux.Vars(r)["kind"], mux.Vars(r)["email"], b)
}

func conventionMgnt(s Service, mailer mail.Mailer, j journal.Journal) {
	s.r.HandleFunc("/api/v1/conventions/", restHandler(conventions, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/conventions/{email}/skip", restHandler(skipConvention, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/conventions/{email}", restHandler(deleteConvention, j, s, mailer)).Methods("DELETE")
}

func skipConvention(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var b bool
	err := jsonRequest(w, r, &b)
	if err != nil {
		return err
	}
	return srv.SkipConvention(mux.Vars(r)["email"], b)
}

func deleteConvention(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	return srv.DeleteConvention(mux.Vars(r)["email"])
}

func conventions(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	c, err := srv.Conventions()
	return writeJSONIfOk(err, w, r, c)
}

func internshipsMngt(s Service, mailer mail.Mailer, j journal.Journal) {
	s.r.HandleFunc("/api/v1/internships/{email}", restHandler(getInternship, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", restHandler(internships, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", restHandler(newInternship, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/supervisor", restHandler(setSupervisor, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/tutor", restHandler(setTutor, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/company", restHandler(setCompany, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/title", restHandler(setTitle, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/major", restHandler(setMajor, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/promotion", restHandler(setPromotion, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/alumni", restHandler(setAlumni, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/", restHandler(students, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/students/{email}", restHandler(insertStudents, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/{email}/internship", restHandler(alignStudentWithInternship, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/{email}/hidden", restHandler(hideStudent, j, s, mailer)).Methods("POST")
}

func setAlumni(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c internship.Alumni
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetAlumni(mux.Vars(r)["email"], c)
}

func newInternship(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c internship.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	_, err = srv.NewInternship(c)
	return err
}

func getInternship(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internship(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, r, i)
}

func internships(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internships()
	return writeJSONIfOk(err, w, r, i)
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

func setTitle(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c string
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetTitle(mux.Vars(r)["email"], c)
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
	_, err = srv.Internship(em)
	if err != nil {
		return err
	}
	return srv.SetTutor(em, s)
}

func surveyMngt(s Service, mailer mail.Mailer, j journal.Journal) {
	s.r.HandleFunc("/surveys/{kind}", surveyForm).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{student}/surveys/{kind}", restHandler(survey, j, s, mailer)).Methods("GET")
}

func surveyForm(w http.ResponseWriter, r *http.Request) {
	k := mux.Vars(r)["kind"]
	http.ServeFile(w, r, path+"/"+k+".html")
}

func survey(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	student := mux.Vars(r)["student"]
	kind := mux.Vars(r)["kind"]
	s, err := srv.Survey(student, kind)
	return writeJSONIfOk(err, w, r, s)
}

type LongSurvey struct {
	Student   internship.Person
	Tutor     internship.Person
	Kind      string
	Deadline  time.Time
	Timestamp time.Time
	Answers   map[string]string
}

func surveyFromToken(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tok := mux.Vars(r)["token"]
		student, kind, err := backend.SurveyToken(tok)
		if err != nil {
			if err == internship.ErrUnknownSurvey {
				http.Error(w, "No survey associated to that token. Maybe a broken link", http.StatusNotFound)
			} else {
				http.Error(w, "A possible bug to report", http.StatusInternalServerError)
			}
			return
		}
		s, err := backend.Survey(student, kind)
		if err != nil {
			log.Printf("Unable to send the survey ", err.Error())
			http.Error(w, "A possible bug to report", http.StatusInternalServerError)
			return
		}
		i, err := backend.Internship(student)
		if err != nil {
			log.Printf("Unable to send the survey ", err.Error())
			http.Error(w, "A possible bug to report", http.StatusInternalServerError)
			return
		}
		lv := LongSurvey{Student: i.Student.ToPerson(), Tutor: i.Tutor.ToPerson(), Kind: kind, Deadline: s.Deadline, Timestamp: s.Timestamp, Answers: s.Answers}
		//TODO: long version of the survey
		err = writeJSONIfOk(err, w, r, lv)
	}
}

func setSurveyContent(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tok := mux.Vars(r)["token"]
		//Check for suspicious content (too long for example)
		var cnt map[string]string
		err := jsonRequest(w, r, &cnt)
		if err != nil {
			http.Error(w, "Bad content", http.StatusBadRequest)
			return
		}
		err = backend.SetSurveyContent(tok, cnt)
		if err != nil {
			log.Println("Unable to store the survey: " + err.Error())
			http.Error(w, "Bad content", http.StatusBadRequest)
		}
	}
}

func students(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	st, err := srv.Students()
	return writeJSONIfOk(err, w, r, st)
}

func alignStudentWithInternship(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var intern string
	err := jsonRequest(w, r, &intern)
	if err != nil {
		return err
	}
	return srv.AlignWithInternship(mux.Vars(r)["email"], intern)
}

func insertStudents(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var all string
	err := jsonRequest(w, r, &all)
	if err != nil {
		return err
	}
	return srv.InsertStudents(all)
}

func hideStudent(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var flag bool
	err := jsonRequest(w, r, &flag)
	if err != nil {
		return err
	}
	return srv.HideStudent(mux.Vars(r)["email"], flag)
}

func defenseMngt(s Service, mailer mail.Mailer, j journal.Journal) {
	s.r.HandleFunc("/api/v1/defenses/", restHandler(getDefenses, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/defenses/", restHandler(postDefenses, j, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/program/", mon(j, getPublicSessions(s.backend))).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/defense", restHandler(getDefense, j, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/defense/grade", restHandler(setDefenseGrade, j, s, mailer)).Methods("POST")
}

func getDefenses(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	defs, err := srv.DefenseSessions()
	return writeJSONIfOk(err, w, r, defs)
}

func postDefenses(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var defs []internship.DefenseSession
	err := jsonRequest(w, r, &defs)
	if err != nil {
		return err
	}
	return srv.SetDefenseSessions(defs)
}

func getDefense(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	em := mux.Vars(r)["email"]
	def, err := srv.DefenseSession(em)
	return writeJSONIfOk(err, w, r, def)
}

func setDefenseGrade(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var g int
	err := jsonRequest(w, r, &g)
	if err != nil {
		return err
	}
	em := mux.Vars(r)["email"]
	return srv.SetDefenseGrade(em, g)
}

func getPublicSessions(srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessions, err := srv.PublicDefenseSessions()
		writeJSONIfOk(err, w, r, sessions)
	}
}
