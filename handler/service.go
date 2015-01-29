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
	s.r.HandleFunc("/", home(backend)).Methods("GET")
	return s
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
	s.r.HandleFunc("/api/v1/users/{email}/password", resetPassword(s.backend, mailer)).Methods("DELETE")
	s.r.HandleFunc("/api/v1/users/{email}/password", restHandler(setPassword, s, mailer)).Methods("PUT")
	s.r.HandleFunc("/api/v1/newPassword", newPassword(s.backend, mailer)).Methods("POST")
	s.r.HandleFunc("/api/v1/login", login(s.backend)).Methods("POST")
	s.r.HandleFunc("/api/v1/logout", logout(s.backend)).Methods("GET")
	s.r.HandleFunc("/resetPassword", password).Methods("GET")
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
		go func() {
			if u, err := srv.User(e); err == nil {
				mailer.SendPasswordResetLink(u, token)
			}
		}()
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
}

func skipConvention(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	var b bool
	err := jsonRequest(w, r, &b)
	if err != nil {
		return err
	}
	return srv.SkipConvention(mux.Vars(r)["email"], b)
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
	s.r.HandleFunc("/api/v1/majors/", restHandler(majors, s, mailer)).Methods("GET")
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
	s.r.HandleFunc("/api/v1/surveys/{token}", restHandler(surveyFromToken, s, mailer)).Methods("GET")
	s.r.HandleFunc("/api/v1/surveys/{token}/content", restHandler(surveyContent, s, mailer)).Methods("GET")
}

func surveyForm(w http.ResponseWriter, r *http.Request) {
	k := mux.Vars(r)["kind"]
	http.ServeFile(w, r, path+"/surveys/"+k+".html")
}

func survey(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	student := mux.Vars(r)["student"]
	kind := mux.Vars(r)["kind"]
	s, err := srv.Survey(student, kind)
	return writeJSONIfOk(err, w, r, s)
}

func surveyContent(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	student := mux.Vars(r)["student"]
	kind := mux.Vars(r)["kind"]
	s, err := srv.SurveyContent(student, kind)
	return writeJSONIfOk(err, w, r, s)
}

func surveyFromToken(srv internship.Service, mailer mail.Mailer, w http.ResponseWriter, r *http.Request) error {
	tok := mux.Vars(r)["token"]
	student, kind, err := srv.SurveyToken(tok)
	if err != nil {
		return err
	}
	s, err := srv.Survey(student, kind)
	return writeJSONIfOk(err, w, r, s)
}
