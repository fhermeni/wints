//Package rest provides all the endpoints
package httpd

import (
	"log"
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
	router *httptreemux.TreeMux
	prefix string
	store  *sqlstore.Store
}

//NewEndPoints creates new prefixed endpoints
func NewEndPoints(prefix string) EndPoints {
	ed := EndPoints{
		router: httptreemux.New(),
		prefix: strings.TrimSuffix(prefix, "/"),
	}
	ed.router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		log.Println("rest call not found: " + r.URL.String())
		http.Error(w, "", http.StatusNotFound)
	}
	//Set the endpoints
	ed.get("/users/", users)
	ed.post("/users/", newUser)
	ed.post("/users/:u/password", setPassword)
	ed.post("/users/:u/person", setUserPerson)
	ed.post("/users/:u/role", setUserRole)
	ed.del("/users/:u", delUser)
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
	token := r.Header.Get("X-Wints-Token")
	s, err := ed.store.Session([]byte(token))
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
		start := time.Now()
		myRw := NewMyResponseWriter(w)

		//Create a session
		s, err := ed.openSession(w, r)
		if err != nil {
			status(myRw, err)
		}

		ex := Exchange{
			w:  myRw,
			r:  r,
			ps: ps,
			s:  s,
		}
		defer func() {
			log.Printf("%s %s %d %d", r.Method, r.URL.String(), myRw.Status(), int(time.Since(start).Nanoseconds()/1000000))
		}()
		status(myRw, fn(ex))
	}
}

func users(ex Exchange) error {
	us, err := ex.s.Users()
	return ex.outJSON(us, err)
}

func newUser(ex Exchange) error {
	var p schema.Person
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	return ex.s.NewUser(p)
}
func delUser(ex Exchange) error {
	return ex.s.RmUser(ex.V("u"))
}

func setPassword(ex Exchange) error {
	var req struct {
		current []byte
		now     []byte
	}
	if err := ex.inJSON(&req); err != nil {
		return err
	}
	return ex.s.SetPassword(ex.V("u"), req.current, req.now)
}

func setUserPerson(ex Exchange) error {
	var p schema.Person
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	return ex.s.SetUserPerson(p)
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
		grade   int
		comment string
	}
	ex.inJSON(&report)
	return ex.s.SetReportGrade(ex.V("k"), ex.V("s"), report.grade, report.comment)
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
	return ex.s.ValidateConvention(ex.V("s"), config.Config{})
}
