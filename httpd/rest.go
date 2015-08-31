//Package rest provides all the endpoints
package httpd

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/session"
	"github.com/fhermeni/wints/sqlstore"
)

//EndPoints is a wrapper to embeds a set of Rest endpoints
type EndPoints struct {
	router       *httptreemux.TreeMux
	prefix       string
	store        *sqlstore.Store
	cfg          config.Rest
	organization config.Internships
}

//NewEndPoints creates new prefixed endpoints
func NewEndPoints(store *sqlstore.Store, cfg config.Rest, org config.Internships) EndPoints {
	ed := EndPoints{
		store:        store,
		router:       httptreemux.New(),
		prefix:       strings.TrimSuffix(cfg.Prefix, "/"),
		cfg:          cfg,
		organization: org,
	}
	ed.router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotFound)
	}
	//Set the endpoints
	ed.get("/users/", users)
	ed.post("/users/", newUser)
	ed.post("/users/:u/password", setPassword)
	ed.post("/users/:u/person", setUserPerson)
	ed.post("/users/:u/role", setUserRole)
	ed.del("/users/:u", delUser)
	ed.get("/users/:u", user)
	ed.del("/users/:u/session", ed.delSession)
	ed.post("/students/:s/major", setMajor)
	ed.post("/students/:s/promotion", setPromotion)
	ed.post("/students/:s/male", setMale)
	ed.post("/students/:s/alumni/contact", setNextContact)
	ed.post("/students/:s/alumni/position", setNextPosition)
	ed.post("/students/:s/skip", setStudentSkippable)
	ed.get("/students/", students)

	ed.get("/internships/", internships)
	ed.get("/internships/:s/", internship)

	ed.get("/reports/:s/:k/content", reportContent)
	ed.post("/reports/:s/:k/content", setReportContent)
	ed.post("/reports/:s/:k/deadline", setReportDeadline)
	ed.post("/reports/:s/:k/grade", setReportGrade)
	ed.post("/reports/:s/:k/private", setReportPrivacy)

	ed.get("/conventions/", conventions)
	ed.post("/conventions/:s/company", setCompany)
	ed.post("/conventions/:s/skip", setConventionSkippable)
	ed.post("/conventions/:s/supervisor", setSupervisor)
	ed.post("/conventions/:s/title", setTitle)
	ed.post("/conventions/:s/tutor", setTutor)
	ed.post("/conventions/:s/valid", validateConvention)

	ed.router.POST(ed.prefix+"/signin", ed.signin)
	ed.router.POST(ed.prefix+"/resetPassword", ed.resetPassword)
	ed.router.POST(ed.prefix+"/newPassword", ed.newPassword)
	ed.router.GET(ed.prefix+"/config", ed.config)
	return ed
}

//EndPoint aliases the call to an endpoint
type EndPoint func(Exchange) error

func (ed *EndPoints) get(path string, handler EndPoint) {
	ed.router.GET(ed.prefix+path, ed.wrap(handler))
}

func (ed *EndPoints) post(path string, handler EndPoint) {
	ed.router.POST(ed.prefix+path, ed.wrap(handler))
}

func (ed *EndPoints) del(path string, handler EndPoint) {
	ed.router.DELETE(ed.prefix+path, ed.wrap(handler))
}

func (ed *EndPoints) openSession(w http.ResponseWriter, r *http.Request) (session.Session, error) {
	token, _ := r.Cookie("token")
	s, err := ed.store.Session([]byte(token.Value))
	if err != nil {
		return session.Session{}, err
	}
	if s.Expire.Before(time.Now()) {
		return session.Session{}, schema.ErrSessionExpired
	}
	user, err := ed.store.User(s.Email)
	if err != nil {
		return session.Session{}, err
	}
	return session.NewSession(user, ed.store), err
}

func (ed *EndPoints) wrap(fn EndPoint) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {
		//Create a session
		s, err := ed.openSession(w, r)
		if err != nil {
			status(w, err)
		}
		ex := Exchange{
			w:  w,
			r:  r,
			ps: ps,
			s:  s,
		}
		status(w, fn(ex))
	}
}

func users(ex Exchange) error {
	us, err := ex.s.Users()
	return ex.outJSON(us, err)
}

func newUser(ex Exchange) error {
	var u schema.User
	if err := ex.inJSON(&u); err != nil {
		return err
	}
	return ex.outJSON(u, ex.s.NewUser(u.Person, u.Role))
}
func delUser(ex Exchange) error {
	return ex.s.RmUser(ex.V("u"))
}

func user(ex Exchange) error {
	u, err := ex.s.User(ex.V("u"))
	return ex.outJSON(u, err)
}

func setPassword(ex Exchange) error {
	var req struct {
		Current string
		Now     string
	}
	if err := ex.inJSON(&req); err != nil {
		return err
	}
	return ex.s.SetPassword(ex.V("u"), []byte(req.Current), []byte(req.Now))
}

func setUserPerson(ex Exchange) error {
	var p schema.Person
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	return ex.outJSON(p, ex.s.SetUserPerson(p))
}

func setUserRole(ex Exchange) error {
	var p schema.Privilege
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	return ex.s.SetUserRole(ex.V("u"), p)
}

func setMajor(ex Exchange) error {
	var m string
	if err := ex.inJSON(&m); err != nil {
		return err
	}
	return ex.s.SetMajor(ex.V("s"), m)
}

func setPromotion(ex Exchange) error {
	var p string
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	return ex.s.SetPromotion(ex.V("s"), p)
}

func setMale(ex Exchange) error {
	var m bool
	if err := ex.inJSON(&m); err != nil {
		return err
	}
	return ex.s.SetMale(ex.V("s"), m)
}

func setNextContact(ex Exchange) error {
	var c string
	if err := ex.inJSON(&c); err != nil {
		return err
	}
	return ex.s.SetNextContact(ex.V("s"), c)
}

func setNextPosition(ex Exchange) error {
	var p int
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	return ex.s.SetNextPosition(ex.V("s"), p)
}

func setStudentSkippable(ex Exchange) error {
	var b bool
	if err := ex.inJSON(&b); err != nil {
		return err
	}
	return ex.s.SetStudentSkippable(ex.V("s"), b)
}

func students(ex Exchange) error {
	ss, err := ex.s.Students()
	return ex.outJSON(ss, err)
}

func internships(ex Exchange) error {
	ss, err := ex.s.Internships()
	return ex.outJSON(ss, err)
}

func internship(ex Exchange) error {
	ss, err := ex.s.Internship(ex.V("s"))
	return ex.outJSON(ss, err)
}

func reportContent(ex Exchange) error {
	r, err := ex.s.ReportContent(ex.V("k"), ex.V("s"))
	return ex.outJSON(r, err)
}

func setReportContent(ex Exchange) error {
	cnt, err := ex.inBytes("application/pdf")
	if err != nil {
		return err
	}
	return ex.s.SetReportContent(ex.V("k"), ex.V("s"), cnt)
}

func setReportDeadline(ex Exchange) error {
	var t time.Time
	ex.inJSON(&t)
	return ex.s.SetReportDeadline(ex.V("k"), ex.V("s"), t)
}

func setReportGrade(ex Exchange) error {
	var report struct {
		Grade   int
		Comment string
	}
	ex.inJSON(&report)
	return ex.s.SetReportGrade(ex.V("k"), ex.V("s"), report.Grade, report.Comment)
}

func setReportPrivacy(ex Exchange) error {
	var p bool
	ex.inJSON(&p)
	return ex.s.SetReportPrivacy(ex.V("k"), ex.V("s"), p)
}

func conventions(ex Exchange) error {
	cc, err := ex.s.Conventions()
	return ex.outJSON(cc, err)
}

func setCompany(ex Exchange) error {
	var c schema.Company
	ex.inJSON(&c)
	return ex.s.SetCompany(ex.V("s"), c)
}

func setConventionSkippable(ex Exchange) error {
	var b bool
	ex.inJSON(&b)
	return ex.s.SetConventionSkippable(ex.V("s"), b)
}

func setSupervisor(ex Exchange) error {
	var p schema.Person
	ex.inJSON(&p)
	return ex.s.SetSupervisor(ex.V("s"), p)
}

func setTitle(ex Exchange) error {
	var t string
	ex.inJSON(&t)
	return ex.s.SetTitle(ex.V("s"), t)
}

func setTutor(ex Exchange) error {
	var t string
	ex.inJSON(&t)
	return ex.s.SetTutor(ex.V("s"), t)
}

func validateConvention(ex Exchange) error {
	var b bool
	ex.inJSON(&b)
	return ex.s.ValidateConvention(ex.V("s"), config.Internships{})
}

func (ed *EndPoints) delSession(ex Exchange) error {
	token := &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	login := &http.Cookie{
		Name:   "login",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(ex.w, token)
	http.SetCookie(ex.w, login)
	http.Redirect(ex.w, ex.r, "/", http.StatusTemporaryRedirect)
	return nil
}

func (ed *EndPoints) signin(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	var cred struct {
		Login    string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		status(w, ErrMalformedJSON)
		return
	}
	s, err := ed.store.NewSession(cred.Login,
		[]byte(cred.Password),
		ed.cfg.SessionLifeTime.Duration)
	if err != nil {
		status(w, err)
		return
	}

	token := &http.Cookie{
		Name:  "token",
		Value: string(s.Token),
		Path:  "/",
	}
	login := &http.Cookie{
		Name:  "login",
		Value: string(s.Email),
		Path:  "/",
	}
	http.SetCookie(w, token)
	http.SetCookie(w, login)
	http.Redirect(w, r, "/", 302)
	ed.store.Visit(cred.Login)
}

func (ed *EndPoints) resetPassword(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	var email string
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		status(w, ErrMalformedJSON)
		return
	}
	s, err := ed.store.ResetPassword(email, ed.cfg.RenewalRequestLifetime.Duration)
	if err != nil {
		status(w, err)
		return
	}
	status(w, json.NewEncoder(w).Encode(s))
}

func (ed *EndPoints) config(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	status(w, json.NewEncoder(w).Encode(ed.organization))
}

func (ed *EndPoints) newPassword(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	var req struct {
		Token    string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		status(w, ErrMalformedJSON)
		return
	}
	s, err := ed.store.NewPassword([]byte(req.Token), []byte(req.Password))
	if err != nil {
		status(w, err)
		return
	}
	status(w, json.NewEncoder(w).Encode(s))
}
