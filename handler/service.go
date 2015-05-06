package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/fhermeni/wints/internship"
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

func NewService(backend internship.Service, mailer mail.Mailer, p string) Service {
	//Rest stuff
	r := mux.NewRouter()
	path = p
	s := Service{backend: backend, r: r, mailer: mailer}
	userMngt(s, mailer)
	internshipsMngt(s, mailer)
	reportMngt(s, mailer)
	conventionMgnt(s, mailer)
	surveyMngt(s, mailer)
	fs := http.Dir(path + "/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/" + path + "/").Handler(httpgzip.NewHandler(http.StripPrefix("/"+path, fileHandler)))
	http.Handle("/", s.r)
	s.r.HandleFunc("/", mon(home(backend))).Methods("GET")
	s.r.HandleFunc("/statistics", mon(stats())).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}", surveyFromToken(backend)).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}", setSurveyContent(backend)).Methods("POST")
	s.r.HandleFunc("/api/v1/statistics/", mon(statistics(backend))).Methods("GET")
	return s
}

func statistics(backend internship.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := backend.Statistics()
		writeJSONIfOk(err, w, r, stats)
	}
}
func stats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path+"/statistics-new.html")
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

func userMngt(s Service, mailer mail.Mailer) {

	s.r.HandleFunc("/api/v1/users/{email}", restHandler(user, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/", restHandler(users, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/users/{email}", restHandler(rmUser, s, mailer)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/profile", restHandler(setUserProfile, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/{email}/role", restHandler(setUserRole, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/users/", restHandler(newTutor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/users/{email}/password", mon(resetPassword(s.backend, mailer))).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/password", restHandler(setPassword, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/sessions/", restHandler(sessions, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/newPassword", mon(newPassword(s.backend, mailer))).Methods("POST")
	s.r.HandleFunc("/api/v1/login", mon(login(s.backend))).Methods("POST")
	s.r.HandleFunc("/api/v1/logout", mon(logout(s.backend))).Methods("GET")
	s.r.HandleFunc("/resetPassword", mon(password)).Methods("GET")
}

func password(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path+"/new_password.html")
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
		t, err := srv.Registered(login, []byte(password))
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

func logout(srv internship.Service) http.HandlerFunc {
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

func resetPassword(srv internship.Service, mailer mail.Mailer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := mux.Vars(r)["email"]
		token, err := srv.ResetPassword(e)
		if err != nil {
			log.Println("Unable to get the reset token for " + e + ": " + err.Error())
			return
		}
		b := r.URL.Query().Get("invite")
		go func() {
			if b == "true" {
				if u, err := srv.User(e); err == nil {
					mailer.SendAdminInvitation(u, token)
				}
			} else {
				if u, err := srv.User(e); err == nil {
					mailer.SendPasswordResetLink(u, token)
				}
			}
		}()
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
	t, err := srv.NewTutor(u)
	if err == nil {
		go mailer.SendAdminInvitation(u, t)
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
	go func() {
		if u, err := srv.User(em); err == nil {
			mailer.SendRoleUpdate(u)
		}
	}()
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
	return writeJSONIfOk(err, w, r, us)
}

func rmUser(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	u, err := srv.User(mux.Vars(r)["email"])
	if err != nil {
		return err
	}
	err = srv.RmUser(u.Email)
	if err == nil {
		go mailer.SendAccountRemoval(u)
	}
	return err
}

func users(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	us, err := srv.Users()
	return writeJSONIfOk(err, w, r, us)
}

func reportMngt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}", restHandler(report, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", restHandler(reportContent, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/content", restHandler(setReportContent, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/grade", restHandler(setReportGrade, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/deadline", restHandler(setReportDeadline, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/reports/{kind}/private", restHandler(setReportPrivate, s, mailer)).Methods("POST")
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
	go func() {
		if err == nil {
			if i, err := srv.Internship(em); err == nil {
				mailer.SendReportUploaded(i.Student, i.Tutor, k)
			} else {
				log.Println("Bwa: " + err.Error() + em)
			}
		} else {
			log.Println("No mail since: " + err.Error())
		}
	}()
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

	go func() {
		if err == nil {
			if i, err := srv.Internship(em); err == nil {
				mailer.SendGradeUploaded(i.Student, i.Tutor, k)
			}
		}
	}()
	return err
}

func setReportDeadline(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var t time.Time
	err := jsonRequest(w, r, &t)
	if err != nil {
		return err
	}
	err = srv.SetReportDeadline(mux.Vars(r)["kind"], mux.Vars(r)["email"], t)
	go func() {
		if err == nil {
			if i, err := srv.Internship(mux.Vars(r)["email"]); err != nil {
				mailer.SendReportDeadline(i.Student, i.Tutor, mux.Vars(r)["kind"], t)
			}
		}
	}()
	return err
}

func setReportPrivate(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var b bool
	err := jsonRequest(w, r, &b)
	if err != nil {
		return err
	}
	err = srv.SetReportPrivate(mux.Vars(r)["kind"], mux.Vars(r)["email"], b)
	go func() {
		if err == nil {
			if i, err := srv.Internship(mux.Vars(r)["email"]); err != nil {
				mailer.SendReportPrivate(i.Student, i.Tutor, mux.Vars(r)["kind"], b)
			}
		}
	}()
	return err
}

func conventionMgnt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/api/v1/conventions/", restHandler(conventions, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/conventions/{email}/skip", restHandler(skipConvention, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/conventions/{email}", restHandler(deleteConvention, s, mailer)).Methods("DELETE")
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

func internshipsMngt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/api/v1/internships/{email}", restHandler(getInternship, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", restHandler(internships, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/", restHandler(newInternship, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/supervisor", restHandler(setSupervisor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/tutor", restHandler(setTutor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/company", restHandler(setCompany, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/title", restHandler(setTitle, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/major", restHandler(setMajor, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/promotion", restHandler(setPromotion, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/internships/{email}/alumni", restHandler(setAlumni, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/majors/", restHandler(majors, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/students/", restHandler(students, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/students/{email}", restHandler(insertStudents, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/{email}/internship", restHandler(alignStudentWithInternship, s, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/students/{email}/hidden", restHandler(hideStudent, s, mailer)).Methods("POST")
}

func setAlumni(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c internship.Alumni
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	return srv.SetAlumni(mux.Vars(r)["email"], c)
}

func majors(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	m := srv.Majors()
	return writeJSONIfOk(nil, w, r, m)
}
func newInternship(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var c internship.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return err
	}
	token, err := srv.NewInternship(c)
	if err == nil {
		go func() {
			s := internship.User{Firstname: c.Student.Firstname, Lastname: c.Student.Lastname, Email: c.Student.Email}
			mailer.SendStudentInvitation(s, token)
			mailer.SendTutorNotification(c.Student, c.Tutor)
		}()
	}
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
	i, err := srv.Internship(em)
	if err != nil {
		return err
	}
	err = srv.SetTutor(em, s)
	go func() {
		if err == nil {
			if u, err := srv.User(s); err == nil {
				mailer.SendTutorUpdate(i.Student, i.Tutor, u)
			}
		}
	}()
	return err
}

func surveyMngt(s Service, mailer mail.Mailer) {
	s.r.HandleFunc("/surveys/{kind}", surveyForm).Methods("GET")
	s.r.HandleFunc("/api/v1/internships/{student}/surveys/{kind}", restHandler(survey, s, mailer)).Methods("GET")
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
