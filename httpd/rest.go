//Package httpd provides all the endpoints
package httpd

import (
	"net/http"
	"strings"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/session"
	"github.com/fhermeni/wints/sqlstore"
)

//EndPoints is a wrapper to embeds a set of Rest endpoints
type EndPoints struct {
	router       *httptreemux.TreeMux
	prefix       string
	store        *sqlstore.Store
	conventions  feeder.Conventions
	cfg          config.Rest
	organization config.Internships
	notifier     *notifier.Notifier
}

//NewEndPoints creates new prefixed endpoints
func NewEndPoints(not *notifier.Notifier, store *sqlstore.Store, convs feeder.Conventions, cfg config.Rest, org config.Internships) EndPoints {
	ed := EndPoints{
		store:        store,
		router:       httptreemux.New(),
		prefix:       strings.TrimSuffix(cfg.Prefix, "/"),
		cfg:          cfg,
		organization: org,
		conventions:  convs,
		notifier:     not,
	}
	ed.router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotFound)
	}
	//Set the endpoints
	ed.get("/users/", users)
	ed.post("/users/", newUser)
	ed.post("/users/:u/email", setEmail)
	ed.post("/users/:u/person", setUserPerson)
	ed.post("/users/:u/role", setUserRole)
	ed.del("/users/:u", delUser)
	ed.get("/users/:u", user)
	ed.del("/users/:u/session", ed.delSession)
	ed.post("/students/:s/major", setMajor)
	ed.post("/students/:s/promotion", setPromotion)
	ed.post("/students/:s/male", setMale)
	ed.post("/students/:s/alumni", setAlumni)
	ed.post("/students/:s/skip", setStudentSkippable)
	ed.get("/students/", students)
	ed.post("/students/", newStudent)

	ed.router.GET(ed.prefix+"/internships/", ed.anon(internships))
	ed.get("/internships/:s", internship)
	ed.post("/internships/:s/company", setCompany)
	ed.post("/internships/:s/supervisor", setSupervisor)
	ed.post("/internships/:s/tutor", setTutor)
	ed.del("/internships/:s/defense", rmDefense)
	ed.post("/internships/:s/defense", setStudentDefense)
	ed.put("/internships/:s/defense", updateStudentDefense)
	ed.get("/internships/:s/defense", studentDefense)
	ed.post("/internships/:s/defense/grade", setDefenseGrade)
	ed.post("/internships/", ed.newInternship)

	ed.del("/surveys/:s/:k", resetSurvey)
	ed.post("/surveys/:s/:k", requestSurvey)
	ed.get("/reports/:s/:k/", report)
	ed.get("/reports/:s/:k/content", reportContent)
	ed.post("/reports/:s/:k/content", setReportContent)
	ed.post("/reports/:s/:k/deadline", setReportDeadline)
	ed.post("/reports/:s/:k/grade", setReportGrade)
	ed.post("/reports/:s/:k/private", setReportPrivacy)

	ed.get("/conventions/", conventions)
	ed.get("/conventions/:s", convention)

	ed.get("/logs/:k", streamLog)
	ed.get("/logs/", logs)

	ed.get("/defenses/", defenseSessions)
	ed.post("/defenses/", newDefenseSession)
	ed.get("/defenses/:id/:r/", defenseSession)
	ed.del("/defenses/:id/:r", rmDefenseSession)
	ed.post("/defenses/:id/:r/jury/", newDefenseSessionJury)
	ed.del("/defenses/:id/:r/jury/:j", rmDefenseSessionJury)

	ed.router.GET(ed.prefix+"/program/", ed.anon(ed.program))
	ed.router.POST(ed.prefix+"/signin", ed.anon(ed.signin))
	ed.router.POST(ed.prefix+"/resetPassword", ed.anon(ed.resetPassword))
	ed.router.POST(ed.prefix+"/newPassword", ed.anon(ed.newPassword))
	ed.router.GET(ed.prefix+"/config", ed.anon(ed.config))
	ed.router.GET(ed.prefix+"/surveys/:t", ed.anon(ed.survey))
	ed.router.POST(ed.prefix+"/surveys/:t", ed.anon(ed.setSurvey))
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

func (ed *EndPoints) put(path string, handler EndPoint) {
	ed.router.PUT(ed.prefix+path, ed.wrap(handler))
}

func (ed *EndPoints) del(path string, handler EndPoint) {
	ed.router.DELETE(ed.prefix+path, ed.wrap(handler))
}

func (ed *EndPoints) openSession(w http.ResponseWriter, r *http.Request) (session.Session, error) {
	token, err := r.Cookie("token")
	if err != nil {
		token = &http.Cookie{}
	}
	s, err := ed.store.Session([]byte(token.Value))
	if err != nil {
		wipeCookies(w)
		return session.Session{}, err
	}
	if s.Expire.Before(time.Now()) {
		logger.Log("event", s.Email, "session expired at "+s.Expire.Format(config.DateTimeLayout), nil)
		wipeCookies(w)
		return session.Session{}, schema.ErrSessionExpired
	}
	user, err := ed.store.User(s.Email)
	if err != nil {
		wipeCookies(w)
		return session.Session{}, err
	}
	ed.store.Visit(user.Person.Email)
	return session.NewSession(user, ed.store, ed.conventions), err
}

func (ed *EndPoints) anon(fn EndPoint) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {
		//We create a session if possible. Otherwise we move to an anonymous session
		s, err := ed.openSession(w, r)
		if err != nil {
			s = session.AnonSession(ed.store)
		}
		ex := Exchange{
			w:   w,
			r:   r,
			ps:  ps,
			s:   s,
			cfg: ed.cfg,
			not: ed.notifier,
		}
		status(ed.notifier, w, r, fn(ex))
	}
}

func (ed *EndPoints) wrap(fn EndPoint) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {

		//Create a session
		s, err := ed.openSession(w, r)
		if err != nil {
			status(ed.notifier, w, r, err)
			return
		}
		ex := Exchange{
			w:   w,
			r:   r,
			ps:  ps,
			s:   s,
			cfg: ed.cfg,
			not: ed.notifier,
		}
		status(ed.notifier, w, r, fn(ex))
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
	token, err := ex.s.NewUser(u.Person, u.Role)
	err = ex.not.InviteTeacher(ex.s.Me(), u, string(token), err)
	return ex.outJSON(u, err)
}
func delUser(ex Exchange) error {
	err := ex.s.RmUser(ex.V("u"))
	return ex.not.RmAccount(ex.s.Me(), ex.V("u"), err)
}

func user(ex Exchange) error {
	u, err := ex.s.User(ex.V("u"))
	return ex.outJSON(u, err)
}

func setUserPerson(ex Exchange) error {
	var p schema.Person
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	err := ex.s.SetUserPerson(p)
	ex.not.ProfileEdited(ex.s.Me(), p, err)
	if err == nil {
		u, e := ex.s.User(ex.V("u"))
		return ex.outJSON(u, e)
	}
	return err
}

func setEmail(ex Exchange) error {
	var em string
	if err := ex.inJSON(&em); err != nil {
		return err
	}
	return ex.s.SetEmail(ex.V("u"), em)
}

func setUserRole(ex Exchange) error {
	var p string
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	r := schema.Role(p)
	err := ex.s.SetUserRole(ex.V("u"), r)
	ex.not.PrivilegeUpdated(ex.s.Me(), ex.V("u"), r, err)
	if err == nil {
		u, e := ex.s.User(ex.V("u"))
		return ex.outJSON(u, e)
	}
	return err
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

func setAlumni(ex Exchange) error {
	var a schema.Alumni
	if err := ex.inJSON(&a); err != nil {
		return err
	}
	err := ex.s.SetAlumni(ex.V("s"), a)
	ex.not.AlumniUpdated(ex.s.Me(), a, err)
	return ex.outJSON(a, err)
}

func setStudentSkippable(ex Exchange) error {
	var b bool
	if err := ex.inJSON(&b); err != nil {
		return err
	}
	err := ex.s.SetStudentSkippable(ex.V("s"), b)
	ex.not.SkipStudent(ex.s.Me(), ex.V("s"), b, err)
	if err != nil {
		return err
	}
	stu, err := ex.s.Student(ex.V("s"))
	return ex.outJSON(stu, err)
}

func students(ex Exchange) error {
	ss, err := ex.s.Students()
	return ex.outJSON(ss, err)
}

func newStudent(ex Exchange) error {
	var s schema.Student
	if err := ex.inJSON(&s); err != nil {
		return err
	}
	s.User.Role = schema.STUDENT
	err := ex.s.NewStudent(s.User.Person, s.Major, s.Promotion, s.Male)
	ex.not.NewStudent(ex.s.Me(), s, err)
	return ex.outJSON(s, err)
}

func internships(ex Exchange) error {
	start := time.Now()

	ss, err := ex.s.Internships()

	ms := int(time.Since(start).Nanoseconds() / 1000000)
	e := ex.outJSON(ss, err)
	logger.ReportValue("list internships", float64(ms))
	return e
}

func internship(ex Exchange) error {
	start := time.Now()
	ss, err := ex.s.Internship(ex.V("s"))
	ms := int(time.Since(start).Nanoseconds() / 1000000)
	e := ex.outJSON(ss, err)

	logger.ReportValue("one internship", float64(ms))
	return e
}

func reportContent(ex Exchange) error {
	r, err := ex.s.ReportContent(ex.V("k"), ex.V("s"))
	return ex.outFile("application/pdf", ex.V("k")+".pdf", r, err)
}

func report(ex Exchange) error {
	r, err := ex.s.Report(ex.V("k"), ex.V("s"))
	return ex.outJSON(r, err)
}

func setReportContent(ex Exchange) error {
	cnt, err := ex.inBytes("application/pdf")
	if err != nil {
		return err
	}
	_, err = ex.s.SetReportContent(ex.V("k"), ex.V("s"), cnt)
	i, err := ex.s.Internship(ex.V("s"))
	if err != nil {
		return err
	}
	ex.not.ReportUploaded(ex.s.Me(), i.Convention.Tutor, ex.V("k"), err)
	if err != nil {
		return err
	}
	hdr, err := ex.s.Report(ex.V("k"), ex.V("s"))
	return ex.outJSON(hdr, err)
}

func setReportDeadline(ex Exchange) error {
	var t time.Time
	if err := ex.inJSON(&t); err != nil {
		return err
	}
	err := ex.s.SetReportDeadline(ex.V("k"), ex.V("s"), t)
	err = ex.not.ReportDeadlineUpdated(ex.s.Me(), ex.V("s"), ex.V("k"), t, err)
	return err
}

func setReportGrade(ex Exchange) error {
	var report struct {
		Grade   int
		Comment string
	}
	if err := ex.inJSON(&report); err != nil {
		return err
	}
	i, err := ex.s.Internship(ex.V("s"))
	if err != nil {
		return err
	}
	_, err = ex.s.SetReportGrade(ex.V("k"), ex.V("s"), report.Grade, report.Comment)
	return ex.not.ReportReviewed(ex.s.Me(), i.Convention.Student.User, i.Convention.Tutor, ex.V("k"), err)
}

func setReportPrivacy(ex Exchange) error {
	var p bool
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	err := ex.s.SetReportPrivacy(ex.V("k"), ex.V("s"), p)
	ex.not.ReportPrivacyUpdated(ex.s.Me(), ex.V("s"), ex.V("k"), p, err)
	return err
}

func conventions(ex Exchange) error {
	cc, err := ex.s.Conventions()
	var all = struct {
		Conventions []schema.Convention
		Errors      *feeder.ImportError
	}{
		Conventions: cc,
		Errors:      err,
	}
	return ex.outJSON(all, err.Fatal)
}

func convention(ex Exchange) error {
	c, err := ex.s.Convention(ex.V("s"))
	return ex.outJSON(c, err)
}

func setCompany(ex Exchange) error {
	var c schema.Company
	ex.inJSON(&c)
	err := ex.s.SetCompany(ex.V("s"), c)
	ex.not.CompanyUpdated(ex.s.Me(), c, err)
	return ex.outJSON(c, err)
}

func setSupervisor(ex Exchange) error {
	var p schema.Person
	if err := ex.inJSON(&p); err != nil {
		return err
	}
	err := ex.s.SetSupervisor(ex.V("s"), p)
	ex.not.SupervisorUpdated(ex.s.Me(), p, err)
	return ex.outJSON(p, err)
}

func resetSurvey(ex Exchange) error {
	stu := ex.V("s")
	kind := ex.V("k")
	err := ex.s.ResetSurvey(stu, kind)
	return ex.not.SurveyReseted(ex.s.Me(), stu, kind, err)
}

func requestSurvey(ex Exchange) error {
	if ex.s.Me().Role.Level() < schema.AdminLevel {
		return session.ErrPermission
	}
	stu := ex.V("s")
	kind := ex.V("k")
	i, err := ex.s.Internship(stu)
	if err != nil {
		return err
	}
	for _, s := range i.Surveys {
		if s.Kind == kind {
			_, err := ex.s.SetSurveyInvitation(stu, kind)
			return ex.not.SurveyRequest(i.Convention.Supervisor, i.Convention.Tutor, i.Convention.Student, s, err)
		}
	}
	return schema.ErrUnknownSurvey
}

func setTutor(ex Exchange) error {
	var t string
	if err := ex.inJSON(&t); err != nil {
		return err
	}
	stu := ex.V("s")
	i, err := ex.s.Internship(stu)
	if err != nil {
		return err
	}
	now, err := ex.s.User(t)
	if err != nil {
		return err
	}
	err = ex.s.SetTutor(stu, t)
	err = ex.not.NewTutor(ex.s.Me(), stu, i.Convention.Tutor, now, err)
	return ex.outJSON(now, err)
}

func (ed *EndPoints) newInternship(ex Exchange) error {
	var c schema.Convention
	if err := ex.inJSON(&c); err != nil {
		return err
	}
	i, token, err := ex.s.NewInternship(c)
	err = ed.notifier.InviteStudent(ex.s.Me(), i, string(token), err)
	return ex.outJSON(i, err)
}

func (ed *EndPoints) delSession(ex Exchange) error {
	err := ex.s.RmSession(ex.s.Me().Person.Email)
	ex.not.Logout(ex.s.Me(), err)
	if err != nil {
		return err
	}
	wipeCookies(ex.w)
	http.Redirect(ex.w, ex.r, "/", http.StatusTemporaryRedirect)
	return nil
}

func (ed *EndPoints) signin(ex Exchange) error {
	var cred struct {
		Login    string
		Password string
	}
	if err := ex.inJSON(&cred); err != nil {
		return err
	}
	s, err := ed.store.NewSession(cred.Login, []byte(cred.Password), ed.cfg.SessionLifeTime.Duration)
	if err != nil {
		if err == schema.ErrCredentials {
			u, e := ed.store.User(cred.Login)
			if e == nil && u.LastVisit == nil {
				//Never visited, so account might be not activated
				err = ErrNotActivatedAccount
			}
		}
	}
	ed.notifier.Login(s, err)
	if err != nil {
		return err
	}
	token := &http.Cookie{
		Name:   "token",
		Value:  string(s.Token),
		Path:   "/",
		MaxAge: int(ed.cfg.SessionLifeTime.Duration.Seconds()),
	}
	login := &http.Cookie{
		Name:   "login",
		Value:  s.Email,
		Path:   "/",
		MaxAge: int(ed.cfg.SessionLifeTime.Duration.Seconds()),
	}
	http.SetCookie(ex.w, token)
	http.SetCookie(ex.w, login)
	http.Redirect(ex.w, ex.r, "/", http.StatusTemporaryRedirect)
	return ed.store.Visit(cred.Login)
}

func (ed *EndPoints) resetPassword(ex Exchange) error {
	var email string
	if err := ex.inJSON(&email); err != nil {
		return err
	}
	token, err := ed.store.ResetPassword(email)
	return ed.notifier.AccountReseted(email, token, err)
}

func (ed *EndPoints) config(ex Exchange) error {
	return ex.outJSON(ed.organization, nil)
}

func (ed *EndPoints) newPassword(ex Exchange) error {
	var req struct {
		Token    string
		Password string
	}
	if err := ex.inJSON(&req); err != nil {
		return err
	}
	email, err := ed.store.NewPassword([]byte(req.Token), []byte(req.Password))
	ex.not.PasswordChanged(email, err)
	if err == nil {
		ed.store.Visit(email)
	}
	return ex.outJSON(email, err)
}

func (ed *EndPoints) survey(ex Exchange) error {
	tok := ex.V("t")
	student, s, err := ed.store.SurveyFromToken(tok)
	if err != nil {
		return err
	}
	i, err := ed.store.Internship(student)
	if err != nil {
		return err
	}
	req := struct {
		Student schema.Person
		Tutor   schema.Person
		Survey  schema.SurveyHeader
	}{Student: i.Convention.Student.User.Person, Tutor: i.Convention.Tutor.Person, Survey: s}
	return ex.outJSON(req, err)
}

func (ed *EndPoints) setSurvey(ex Exchange) error {
	tok := ex.V("t")
	answers := make(map[string]interface{})
	if err := ex.inJSON(&answers); err != nil {
		return err
	}
	if _, err := ed.store.SetSurveyContent(tok, answers); err != nil {
		return err
	}

	stu, s, err := ed.store.SurveyFromToken(tok)
	if err != nil {
		return err
	}
	i, err := ed.store.Internship(stu)
	if err != nil {
		return err
	}
	err = ex.not.SurveyUploaded(i.Convention.Student.User, i.Convention.Tutor, i.Convention.Supervisor, s.Kind, err)
	return ex.outJSON(s, err)
}

//Defense management
func defenseSessions(ex Exchange) error {
	return ex.outJSON(ex.s.DefenseSessions())
}

func (ed *EndPoints) program(ex Exchange) error {
	p, err := ex.s.DefenseProgram()
	if err == nil {
		for i, session := range p {
			for ii, d := range session.Defenses {
				d.Grade = -10
				d.Student.Alumni = nil
				session.Defenses[ii] = d
			}
			p[i] = session
		}
	}
	return ex.outJSON(p, err)
}

func newDefenseSession(ex Exchange) error {
	var d schema.DefenseSession
	err := ex.inJSON(&d)
	if err != nil {
		return err
	}
	d, err = ex.s.NewDefenseSession(d.Room, d.Id)
	return ex.outJSON(d, err)
}

func rmDefenseSession(ex Exchange) error {
	id := ex.V("id")
	room := ex.V("r")
	return ex.s.RmDefenseSession(room, id)
}

func newDefenseSessionJury(ex Exchange) error {
	id := ex.V("id")
	room := ex.V("r")
	var jury string
	if err := ex.inJSON(&jury); err != nil {
		return err
	}
	if err := ex.s.AddJuryToDefenseSession(room, id, jury); err != nil {
		return err
	}
	return ex.outJSON(ex.s.DefenseSession(room, id))
}

func defenseSession(ex Exchange) error {
	id := ex.V("id")
	room := ex.V("r")
	return ex.outJSON(ex.s.DefenseSession(room, id))
}

func rmDefenseSessionJury(ex Exchange) error {
	id := ex.V("id")
	room := ex.V("r")
	jury := ex.V("j")
	if err := ex.s.DelJuryToDefenseSession(room, id, jury); err != nil {
		return err
	}
	return ex.outJSON(ex.s.DefenseSession(room, id))
}

func setDefenseGrade(ex Exchange) error {
	stu := ex.V("s")
	var grade int
	if err := ex.inJSON(&grade); err != nil {
		return err
	}
	if err := ex.s.SetDefenseGrade(stu, grade); err != nil {
		return err
	}
	return nil
}

func rmDefense(ex Exchange) error {
	stu := ex.V("s")
	return ex.s.RmDefense(stu)
}

func studentDefense(ex Exchange) error {
	stu := ex.V("s")
	def, err := ex.s.Defense(stu)
	return ex.outJSON(def, err)
}

func setStudentDefense(ex Exchange) error {
	stu := ex.V("s")
	def := schema.Defense{}
	if err := ex.inJSON(&def); err != nil {
		return err
	}
	if err := ex.s.SetStudentDefense(def.SessionId, def.Room, stu, def.Time, def.Public, def.Local); err != nil {
		return err
	}
	return ex.outJSON(ex.s.Defense(stu))
}

func updateStudentDefense(ex Exchange) error {
	stu := ex.V("s")
	def := schema.Defense{}
	if err := ex.inJSON(&def); err != nil {
		return err
	}
	if err := ex.s.UpdateStudentDefense(stu, def.Time, def.Public, def.Local); err != nil {
		return err
	}
	return ex.outJSON(ex.s.Defense(stu))
}
