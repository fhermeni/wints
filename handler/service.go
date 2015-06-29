package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/fhermeni/wints/filter"
	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/mail"
	"github.com/gorilla/mux"

	"github.com/daaku/go.httpgzip"
)

var path string

type Service struct {
	backend   internship.Service
	mailer    mail.Mailer
	r         *mux.Router
	assetPath string
	j         journal.Journal
}

var (
	ErrJSONWriting    = errors.New("Unable to serialize the JSON response")
	ErrMissingCookies = errors.New("Cookies 'session' or 'token' are missing or empty")
)

func NewService(backend internship.Service, mailer mail.Mailer, p string, j journal.Journal) Service {
	//Rest stuff
	r := mux.NewRouter()
	path = p
	s := Service{backend: backend, r: r, mailer: mailer, assetPath: p, j: j}
	fs := http.Dir(path + "/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/" + path + "/").Handler(httpgzip.NewHandler(http.StripPrefix("/"+path, fileHandler)))
	http.Handle("/", s.r)
	s.r.HandleFunc("/", mon(j, home(backend))).Methods("GET")
	s.r.HandleFunc("/login", mon(j, s.asset("login.html"))).Methods("GET")
	s.r.HandleFunc("/statistics", mon(j, s.asset("statistics-new.html"))).Methods("GET")
	s.r.HandleFunc("/defense-program", mon(j, s.asset("defense-program.html"))).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}", mon(j, surveyFromToken(j, backend))).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}", mon(j, setSurveyContent(j, mailer, backend))).Methods("POST")
	s.r.HandleFunc("/api/v1/statistics/", mon(j, statistics(backend))).Methods("GET")
	s.r.HandleFunc("/api/v1/majors/", mon(j, majors(backend))).Methods("GET")
	s.r.HandleFunc("/api/v1/users/{email}", s.rest(user)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/", s.rest(users)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/{email}", s.rest(rmUser)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/profile", s.rest(setUserProfile)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/{email}/role", s.rest(setUserRole)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/", s.rest(newTutor)).Methods("POST")
	s.r.HandleFunc("/api/v1/users/{email}/password", mon(j, resetPassword(j, s.backend, mailer))).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/password", s.rest(setPassword)).Methods("PUT")
	s.r.HandleFunc("/api/v1/sessions/", s.rest(sessions)).Methods("GET")
	s.r.HandleFunc("/api/v1/newPassword", mon(j, newPassword(j, s.backend, mailer))).Methods("POST")
	s.r.HandleFunc("/api/v1/login", mon(j, login(j, s.backend))).Methods("POST")
	s.r.HandleFunc("/api/v1/logout", mon(j, logout(j, s.backend))).Methods("GET")
	s.r.HandleFunc("/resetPassword", mon(j, s.asset("new_password.html"))).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}", s.rest(report)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", s.rest(reportContent)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", s.rest(setReportContent)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/grade", s.rest(setReportGrade)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/deadline", s.rest(setReportDeadline)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/private", s.rest(setReportPrivate)).Methods("POST")
	s.r.HandleFunc("/api/v1/conventions/", s.rest(conventions)).Methods("GET")
	s.r.HandleFunc("/api/v1/conventions/{email}/skip", s.rest(skipConvention)).Methods("POST")
	s.r.HandleFunc("/api/v1/conventions/{email}", s.rest(deleteConvention)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/internships/{email}", s.rest(getInternship)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", s.rest(internships)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", s.rest(newInternship)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/supervisor", s.rest(setSupervisor)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/tutor", s.rest(setTutor)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/company", s.rest(setCompany)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/title", s.rest(setTitle)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/major", s.rest(setMajor)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/promotion", s.rest(setPromotion)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/alumni", s.rest(setAlumni)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/", s.rest(students)).Methods("GET")
	s.r.HandleFunc("/api/v1/students/{email}", s.rest(insertStudents)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/{email}/internship", s.rest(alignStudentWithInternship)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/{email}/hidden", s.rest(hideStudent)).Methods("POST")
	s.r.HandleFunc("/api/v1/defenses/", s.rest(getDefenses)).Methods("GET")
	s.r.HandleFunc("/api/v1/defenses/", s.rest(postDefenses)).Methods("POST")
	s.r.HandleFunc("/api/v1/program/", mon(j, getPublicSessions(s.backend))).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/defense", s.rest(getDefense)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/defense/grade", s.rest(setDefenseGrade)).Methods("POST")
	s.r.HandleFunc("/surveys/{kind}", mon(j, surveyForm)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{student}/surveys/{kind}", s.rest(survey)).Methods("GET")

	return s
}

func (s *Service) rest(cb func(internship.Service, http.ResponseWriter, *http.Request) error) http.HandlerFunc {

	return mon(s.j, func(w http.ResponseWriter, r *http.Request) {

		email, err := authenticated(s.backend, w, r)
		if err == ErrMissingCookies || err == internship.ErrSessionExpired {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err == internship.ErrCredentials {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			s.j.Log("ERROR", "Unsupported error: %s\n", err)
			return
		}
		u, err := s.backend.User(email)
		if err == internship.ErrUnknownUser {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		wrapper, _ := filter.NewService(s.j, s.backend, u, s.mailer)
		e := cb(wrapper, w, r)
		status(w, e)
	})
}

func (s *Service) asset(file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(s.assetPath, file))
	}
}

func majors(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSONIfOk(nil, w, r, backend.Majors())
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
		r.Header.Del("Cache-Control")
		r.Header.Set("Pragma", "no-cache")
		r.Header.Del("If-Modified-Since")
		if _, err := authenticated(backend, w, r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		w.Header().Set("Pragma", "no-cache")
		http.ServeFile(w, r, filepath.Join(path, "home.html"))
	}
}

func (s *Service) Listen(host string, cert, pk string) error {
	return http.ListenAndServeTLS(host, cert, pk, nil)
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
		http.Redirect(w, r, "/login", 302)
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
			http.Error(w, "Unable to reset the password", http.StatusInternalServerError)
			return
		}
	}
}

func setPassword(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var u PasswordUpdate
	if err := jsonRequest(w, r, &u); err != nil {
		return err
	}
	return srv.SetUserPassword(mux.Vars(r)["email"], []byte(u.Old), []byte(u.New))
}

func resetPassword(j journal.Journal, srv internship.Service, mailer mail.Mailer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		em := mux.Vars(r)["email"]
		_, err := srv.ResetPassword(em)
		j.Log(em, "initiate password reset", err)
		status(w, err)
	}
}

func sessions(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	res, err := srv.Sessions()
	return writeJSONIfOk(err, w, r, res)
}

func newTutor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var u internship.User
	err := jsonRequest(w, r, &u)
	if err != nil {
		return err
	}
	_, err = srv.NewTutor(u)
	return err
}

func setUserRole(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var p internship.Privilege
	if err := jsonRequest(w, r, &p); err != nil {
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
	if err := jsonRequest(w, r, &p); err != nil {
		return err
	}
	return srv.SetUserProfile(mux.Vars(r)["email"], p.Firstname, p.Lastname, p.Tel)
}

func user(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.User(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, r, us)
}

func rmUser(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	u, err := srv.User(mux.Vars(r)["email"])
	if err != nil {
		return err
	}
	return srv.RmUser(u.Email)
}

func users(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.Users()
	return writeJSONIfOk(err, w, r, us)
}

func report(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	h, err := srv.Report(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, r, h)
}

func reportContent(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	cnt, err := srv.ReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"])
	return fileReplyIfOk(err, w, "application/pdf", mux.Vars(r)["kind"]+".pdf", cnt)
}

func setReportContent(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
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
	return srv.SetReportContent(mux.Vars(r)["kind"], mux.Vars(r)["email"], cnt)
}

type ReportGrade struct {
	Grade   int
	Comment string
}

func setReportGrade(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var n ReportGrade
	if err := jsonRequest(w, r, &n); err != nil {
		return err
	}
	return srv.SetReportGrade(mux.Vars(r)["kind"], mux.Vars(r)["email"], n.Grade, n.Comment)
}

func setReportDeadline(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var t time.Time
	if err := jsonRequest(w, r, &t); err != nil {
		return err
	}
	return srv.SetReportDeadline(mux.Vars(r)["kind"], mux.Vars(r)["email"], t)
}

func setReportPrivate(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var b bool
	if err := jsonRequest(w, r, &b); err != nil {
		return err
	}
	return srv.SetReportPrivate(mux.Vars(r)["kind"], mux.Vars(r)["email"], b)
}

func skipConvention(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var b bool
	if err := jsonRequest(w, r, &b); err != nil {
		return err
	}
	return srv.SkipConvention(mux.Vars(r)["email"], b)
}

func deleteConvention(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	return srv.DeleteConvention(mux.Vars(r)["email"])
}

func conventions(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	c, err := srv.Conventions()
	return writeJSONIfOk(err, w, r, c)
}

func setAlumni(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c internship.Alumni
	if err := jsonRequest(w, r, &c); err != nil {
		return err
	}
	return srv.SetAlumni(mux.Vars(r)["email"], c)
}

func newInternship(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c internship.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	_, err = srv.NewInternship(c)
	return err
}

func getInternship(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internship(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, r, i)
}

func internships(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	i, err := srv.Internships()
	return writeJSONIfOk(err, w, r, i)
}

func setCompany(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c internship.Company
	if err := jsonRequest(w, r, &c); err != nil {
		return err
	}
	return srv.SetCompany(mux.Vars(r)["email"], c)
}

func setMajor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c string
	if err := jsonRequest(w, r, &c); err != nil {
		return err
	}
	return srv.SetMajor(mux.Vars(r)["email"], c)
}

func setTitle(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c string
	if err := jsonRequest(w, r, &c); err != nil {
		return err
	}
	return srv.SetTitle(mux.Vars(r)["email"], c)
}

func setPromotion(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var c string
	if err := jsonRequest(w, r, &c); err != nil {
		return err
	}
	return srv.SetPromotion(mux.Vars(r)["email"], c)
}

func setSupervisor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var s internship.Person
	if err := jsonRequest(w, r, &s); err != nil {
		return err
	}
	return srv.SetSupervisor(mux.Vars(r)["email"], s)
}

func setTutor(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
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

func surveyForm(w http.ResponseWriter, r *http.Request) {
	k := mux.Vars(r)["kind"]
	http.ServeFile(w, r, path+"/"+k+".html")
}

func survey(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
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

func surveyFromToken(j journal.Journal, backend internship.Service) http.HandlerFunc {
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
			j.Log(tok, "Sending the long survey", err)
			http.Error(w, "A possible bug to report", http.StatusInternalServerError)
			return
		}
		i, err := backend.Internship(student)
		if err != nil {
			j.Log(tok, "Sending the long survey", err)
			http.Error(w, "A possible bug to report", http.StatusInternalServerError)
			return
		}
		lv := LongSurvey{Student: i.Student.ToPerson(), Tutor: i.Tutor.ToPerson(), Kind: kind, Deadline: s.Deadline, Timestamp: s.Timestamp, Answers: s.Answers}
		//TODO: long version of the survey
		err = writeJSONIfOk(err, w, r, lv)
	}
}

func setSurveyContent(j journal.Journal, m mail.Mailer, backend internship.Service) http.HandlerFunc {
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
			http.Error(w, "Bad content", http.StatusBadRequest)
			return
		}

		j.Log(tok, "uploaded the survey", err)
		if err == nil {
			stu, kind, err2 := backend.SurveyToken(tok)
			if err2 != nil {
				j.Log(tok, "Unable to mail about survey", err2)
				return
			}
			i, err2 := backend.Internship(stu)
			if err2 != nil {
				j.Log(tok, "Unable to mail about survey", err2)
				return
			}
			m.SendSurveyUploaded(i.Tutor, i.Student, kind)
		}
	}
}

func students(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	st, err := srv.Students()
	return writeJSONIfOk(err, w, r, st)
}

func alignStudentWithInternship(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var intern string
	if err := jsonRequest(w, r, &intern); err != nil {
		return err
	}
	return srv.AlignWithInternship(mux.Vars(r)["email"], intern)
}

func insertStudents(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var all string
	if err := jsonRequest(w, r, &all); err != nil {
		return err
	}
	return srv.InsertStudents(all)
}

func hideStudent(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var flag bool
	if err := jsonRequest(w, r, &flag); err != nil {
		return err
	}
	return srv.HideStudent(mux.Vars(r)["email"], flag)
}

func getDefenses(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	defs, err := srv.DefenseSessions()
	return writeJSONIfOk(err, w, r, defs)
}

func postDefenses(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var defs []internship.DefenseSession
	if err := jsonRequest(w, r, &defs); err != nil {
		return err
	}
	return srv.SetDefenseSessions(defs)
}

func getDefense(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	def, err := srv.DefenseSession(mux.Vars(r)["email"])
	return writeJSONIfOk(err, w, r, def)
}

func setDefenseGrade(srv internship.Service, w http.ResponseWriter, r *http.Request) error {
	var g int
	if err := jsonRequest(w, r, &g); err != nil {
		return err
	}
	return srv.SetDefenseGrade(mux.Vars(r)["email"], g)
}

func getPublicSessions(srv internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessions, err := srv.PublicDefenseSessions()
		writeJSONIfOk(err, w, r, sessions)
	}
}
