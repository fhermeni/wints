package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/fhermeni/wints/backend"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"math/rand"
	"io/ioutil"
	"compress/gzip"
	"time"
	"strings"
	"bytes"
)

const (
	ROOT_API = "/api/v1"
)

var DB *sql.DB
var cfg backend.Config

func report500OnError(w http.ResponseWriter, header string, err error) bool {

	msg := "The operation did not complete. A possible bug"
	code := http.StatusInternalServerError

	if err != nil {
		switch (err) {
		case backend.ErrUnknownConvention, backend.ErrUnknownDefense, backend.ErrUnknownStudent, backend.ErrUserNotFound, backend.ErrUnknownReport:
			code = http.StatusNotFound
			msg = err.Error()
		case backend.ErrConventionExists, backend.ErrDefenseExists, backend.ErrStudentExists, backend.ErrReportExists, backend.ErrUserExists:
			code = http.StatusConflict
			msg = err.Error()
		case backend.ErrInvalidGrade, backend.ErrInvalidTutor, backend.ErrUserTutoring, backend.ErrCredentials, backend.ErrNoPendingRequests:
			code = http.StatusForbidden
			msg = err.Error()
		default:
			log.Printf("%s: %s\n", header, err.Error());
		}
		http.Error(w, msg, code)

		return true
	}
	return false
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	email, err := backend.ExtractEmail(r)
	if err != nil {
		http.ServeFile(w, r, "static/login.html")
	} else {
		_, err := backend.GetConvention(DB, email)
		if err == nil {
			http.Redirect(w, r, "/student", 302)
		} else {
			http.Redirect(w, r, "/admin", 302)
		}
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	n, err := backend.ExtractEmail(r)
	if report500OnError(w, "Unable to get the email from the token", err) {
		return
	}
	if len(n) == 0 {
		http.Redirect(w, r, "/", 302)
		return
	}
	_, err = backend.GetUser(DB, n)
	if report500OnError(w, "No user having email '" + n + "'", err) {
		return
	}
	http.ServeFile(w, r, "static/admin.html")
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	n, err := backend.ExtractEmail(r)
	if report500OnError(w, "Unable to get the email from the token", err) {
		return
	}
	if len(n) == 0 {
		http.Redirect(w, r, "/", 302)
		return
	}
	_, err = backend.GetUser(DB, n)
	if report500OnError(w, "No user having email '" + n + "'", err) {
		return
	}
	http.ServeFile(w, r, "static/student.html")
}

func jsonReply(w http.ResponseWriter, gz bool, j interface{}) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if gz {
		w.Header().Set("Content-Encoding", "gzip")
	}
	var err error
	if (gz) {
		var w2 *gzip.Writer
		w2 = gzip.NewWriter(w)
		defer w2.Close()
		enc := json.NewEncoder(w2)
		err = enc.Encode(j)
	} else {
		enc := json.NewEncoder(w)
		err = enc.Encode(j)
	}
	report500OnError(w, "Unable to serialize the response", err)
}

func fileReply(w http.ResponseWriter, mime, filename string, cnt []byte) {
	w.Header().Set("Content-type", mime)
	w.Header().Set("Content-disposition", "attachment; filename=" + filename)
	_, err := w.Write(cnt)
	report500OnError(w, "Unable to serve file '" + filename + "'", err)
}

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(&j)
}

type PendingConventionMsg struct {
	C       backend.Convention
	Known   []backend.User
	Pending int
}

func RandomPendingConvention(w http.ResponseWriter, r *http.Request, email string) {
	cc, err := backend.GetRawConventions(DB)
	nb := len(cc)
	if nb == 0 {
		jsonReply(w, false, PendingConventionMsg{backend.Convention{}, nil, 0})
		return
	}
	c := cc[rand.Intn(nb)];
	known, err := backend.Admins(DB)
	if report500OnError(w, "Unable to get the possible tutors", err) {
		return
	}
	jsonReply(w, false, PendingConventionMsg{c, known, nb})
}

func GetAllConventions(w http.ResponseWriter, r *http.Request, email string) {
	conventions, err := backend.GetConventions2(DB, email)
	if report500OnError(w, "Unable to get the conventions ", err) {
		return
	}
	jsonReply(w, true, conventions)
}


func GetAdmins(w http.ResponseWriter, r *http.Request, email string) {
	admins, err := backend.Admins(DB)
	if report500OnError(w, "Unable to get the possible admins", err) {
		return
	}
	jsonReply(w, true, admins)
}

func NewAdmin(w http.ResponseWriter, r *http.Request, email string) {
	var nu backend.User
	err := jsonRequest(w, r, &nu)
	if report500OnError(w, "Unable to extract the user profile", err) {
		return
	}
	_, err = backend.NewUser(DB, nu)
	if report500OnError(w, "Unable to create the user " + nu.String(), err) {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, email string) {
	u, err := backend.GetUser(DB, email)
	if report500OnError(w, "Unable to get user '" + email + "'", err) {
		return
	}
	jsonReply(w, false, u)
}

func RmUser(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	err := backend.RmUser(DB, target)
	if report500OnError(w, "Unable to remove user '" + email + "'", err) {
		return
	}
}

func GetDefense(w http.ResponseWriter, r *http.Request) {
	fmt := r.FormValue("fmt")
	if (len(fmt) == 0) {
		fmt = "embedded"
	}
	var cnt string
	var err error
	switch (fmt) {
	case "public":
		cnt, err = backend.LongDefense(DB, "si/ifi")
	case "embedded":
		email, _ := backend.ExtractEmail(r)
		if (len(email) == 0) {
			http.Redirect(w, r, "/", 302)
	    	return
		}
		u, err := backend.GetUser(DB, email)
		if report500OnError(w, "Unknown user '" + email + "'", err) {
			return
		}
		if u.CompatibleRole("admin") {
			cnt, err = backend.Defense(DB, "si/ifi")
		} else {
			http.Error(w, "Unsufficient permission", http.StatusForbidden)
			return
		}
	}
	//Compress
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	w2 := gzip.NewWriter(w)
	defer w2.Close()
	_, err = w2.Write([]byte(cnt));
	report500OnError(w, "Unable to send the defense content", err)
}

type DefensesPost struct {
	Short string
	Long  string
}

func PostDefense(w http.ResponseWriter, r *http.Request, email string) {
	var dp DefensesPost
	err := jsonRequest(w, r, &dp)
	if report500OnError(w, "Unable to parse the defense content", err) {
		return
	}
	err = backend.SaveDefense(DB, "si/ifi", string(dp.Short), string(dp.Long))
	report500OnError(w, "Unable to save the defense", err)
}

func Login(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	_, err := backend.Register(DB, login, password)
	if err != nil {
		http.Redirect(w, r, "/#badLogin", 302);
		return
	}
	backend.SetSession(login, w)
	http.Redirect(w, r, "/", 302);
}

func PasswordReset(w http.ResponseWriter, r *http.Request) {
	email, err := requiredContent(w, r)
	if report500OnError(w, "", err) {
		return
	}
	token, err := backend.PasswordRenewalRequest(DB, string(email))
	if report500OnError(w, "Unable to generate a token for a password reset", err) {
		return
	}
	err = backend.Mail(cfg, []string{string(email)}, "mails/account_reset.txt", struct {
				WWW   string
				Token string
			}{cfg.WWW, token})
	report500OnError(w, "Unable to send the token for a password reset by mail", err)
}

func Logout(w http.ResponseWriter, r *http.Request, email string) {
	backend.ClearSession(w);
	http.Redirect(w, r, "/", 302);
}

func CommitPendingConvention(w http.ResponseWriter, r *http.Request, email string) {
	var c backend.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return
	}
	tEmail := c.Tutor.Email
	_, err = backend.GetUser(DB, tEmail)
	if err != nil {
		_, err = backend.NewUser(DB, c.Tutor)
		if report500OnError(w, "Unable to register the tutor "+c.Tutor.String(), err) {
			return
		}
	}
	err = backend.RegisterInternship(DB, c, true);
	if report500OnError(w, "Unable to register the internship", err) {
		return
	}
	backend.RescanPending(DB, c)
}


func ChangeProfile(w http.ResponseWriter, r *http.Request, email string) {
	var p backend.User
	if jsonRequest(w, r, &p) != nil {
		return
	}
	err := backend.SetProfile(DB, email, p.Firstname, p.Lastname, p.Tel)
	if report500OnError(w, "Unable to update the profile", err) {
		return
	}
	u, err := backend.GetUser(DB, email)
	if report500OnError(w, "Unable to get user '" + email +"'", err) {
		return
	}
	jsonReply(w, false, u)
}

type PasswordRenewal struct {
	OldPassword string
	NewPassword string
}

func ChangePassword(w http.ResponseWriter, r *http.Request, email string) {
	var p PasswordRenewal
	err := jsonRequest(w, r, &p)
	if report500OnError(w, "Unable to extract the passwords", err) {
		return
	}
	err = backend.ChangePassword(DB, email, []byte(p.OldPassword), []byte(p.NewPassword))
	report500OnError(w, "Unable to change the password", err)
}

func NewPassword(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	password := r.FormValue("password")
	email, err := backend.NewPassword(DB, token, []byte(password))
	if report500OnError(w, "Unable to create a new password from the token", err) {
		return
	}
	_, err = backend.Register(DB, email, password)
	if err != nil {
		http.Redirect(w, r, "/#badLogin", 302);
		return
	}
	backend.SetSession(email, w)
	http.Redirect(w, r, "/", 302);
}


func UpdateMajor(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	m, err := requiredContent(w, r)
	if err != nil {
		return
	}
	err = backend.SetMajor(DB, target, string(m))
	if report500OnError(w, "Unable to update the major of '" + target + "' to " + string(m), err) {
		return
	}
}

func UpdateTutor(w http.ResponseWriter, r *http.Request, email string) {
	student, _ := mux.Vars(r)["email"]
	newTutor, err := requiredContent(w, r)
	if err != nil {
		return
	}
	log.Println(newTutor)

	c, err := backend.GetConvention(DB, student)
	//oldTutor := c.Tutor.Email
	if report500OnError(w, "", err) {
		return
	}
	err = backend.UpdateTutor(DB, string(student), string(newTutor))
	if report500OnError(w, "", err) {
		return
	}
	c, err = backend.GetConvention(DB, string(student))
	if report500OnError(w, "", err) {
		return
	}
	//TODO: mail both tutors and the student
	//backend.Mail(cfg, string[]{string(oldTutor), string(newTutor), string(c.Student.P.Email)}, "tutor_update.txt", c.Tutor);
	jsonReply(w, false, c)
}


func MarkKind(cb func(http.ResponseWriter, *http.Request, string), kind string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cb(w, r, kind)
	}
}

func GetReport(w http.ResponseWriter, r *http.Request, email string)  {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	cnt, err := backend.Report(DB, student, kind)
	if report500OnError(w, "", err) {
		return
	}
	u, _ := backend.GetUser(DB, student)
	fileReply(w, "application/pdf", u.Lastname + "-" + kind + ".pdf", cnt)
}

func UploadReport(w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	d, err := requiredContent(w, r)
	if report500OnError(w, "", err) {
		return
	}
	err = backend.SetReport(DB, student, kind, d)
	if !report500OnError(w, "", err) {
		w.Write([]byte("1"))
	}
}


func UpdateDeadline(w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	d, err := requiredContent(w, r)
	if err != nil {
		return
	}
	t, err := time.Parse("2/1/2006", string(d))
	if report500OnError(w, "Invalid format", err) {
		return
	}

	err = backend.UpdateDeadline(DB, student, kind, t)
	report500OnError(w, "", err);
}

func UpdateMark(w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	d, err := requiredContent(w, r)
	if err != nil {
		return
	}
	i, err := strconv.Atoi(string(d))
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = backend.UpdateGrade(DB, student, kind, i)
	report500OnError(w, "", err);
}

type Company struct {
	Name string
	WWW string
	Title string
}
func UpdateCompany(w http.ResponseWriter, r *http.Request, student string) {
	var c Company
	if err := jsonRequest(w, r, &c); err != nil {
		report500OnError(w, "", err)
		return
	}

	err := backend.UpdateCompany(DB, student, c.Name, c.WWW, c.Title)
	if report500OnError(w, "", err) {
		return
	}
}

func UpdateSupervisor(w http.ResponseWriter, r *http.Request, student string) {
	var u backend.User
	if err := jsonRequest(w, r, &u); err != nil {
		report500OnError(w, "", err)
		return
	}
	err:= backend.UpdateSupervisor(DB, student, u)
	if report500OnError(w, "", err) {
		return
	}
}

func myStudent(w http.ResponseWriter, mine, target string) bool {
	u, err := backend.GetUser(DB, target);
	if err != nil {
		return false
	}
	if u.Role == "" {
		c, err := backend.GetConvention(DB, target)
		if err != nil {
			return false
		}
		if c.Tutor.Email != mine {
			http.Error(w, "The student is not under your responsibility", http.StatusForbidden)
			return false
		}
	}
	return true
}

func requiredContent(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	cnt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}
	if len(cnt) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return []byte{}, err
	}
	return cnt, nil
}

func GrantRole(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	role, err := requiredContent(w, r)
	if report500OnError(w, "", err) {
		return
	}
	err = backend.GrantPrivilege(DB, target, string(role))
	if report500OnError(w, "", err) {
		return
	}
	backend.Mail(cfg, []string{target}, "mails/priv_" + string(role) + ".txt", nil)
}

func GetReports(w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	emails := r.FormValue("students")
	if len(emails) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	cnt, err := backend.Reports(DB, kind, strings.Split(emails, ","))
	if report500OnError(w, "Unable to make the archive", err) {
		return
	}
	var b bytes.Buffer
	w2 := gzip.NewWriter(&b)
	w2.Write(cnt)
	w2.Close()
	fileReply(w, "application/x-tar", "reports-" + kind + ".tar.gz", b.Bytes())
}

func RequireToken(cb func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n, _ := backend.ExtractEmail(r)
		if (len(n) == 0) {
			http.Redirect(w, r, "/", 302)
			return
		}
		cb(w, r, n)
	}
}

func RequireRole(cb func(http.ResponseWriter, *http.Request, string), role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, _ := backend.ExtractEmail(r)
		if (len(email) == 0) {
			http.Redirect(w, r, "/", 302)
			return
		}

		u, _ := backend.GetUser(DB, email)
		if u.CompatibleRole(role) {
			cb(w, r, email)
		} else {
			http.Error(w, "Unsufficient permission", http.StatusUnauthorized)
		}
	}
}

func GetConvention(w http.ResponseWriter, r *http.Request, email string) {
	c, err := backend.GetConvention(DB, email)
	if report500OnError(w, "unable to get the convention of '" + email + "'", err) {
		return
	}
	jsonReply(w, true, c)
}

func main() {
	cfg, err := backend.ReadConfig("./wints.conf")
	if err != nil {
		log.Fatalln("Error while parsing wints.conf: " + err.Error())
	}
	dbUrl := os.Getenv("DATABASE_URL")

	if len(dbUrl) == 0 {
		dbUrl = cfg.DBurl
	}
	DB, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer DB.Close()
	for _, v := range os.Args {
		if v == "-s" {
			delay, err := time.ParseDuration(cfg.RefreshPeriod)
			if err != nil {
				log.Fatalln(err.Error())
			}
			backend.DaemonConventionsPuller(DB, cfg.WWWConventionsURL,
				cfg.WWWConventionsLogin,
				cfg.WWWConventionsPassword,
				delay);
		}
	}

	//Rest stuff
	r := mux.NewRouter()

	//the dashboards
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/admin", adminHandler)
	r.HandleFunc("/student", studentHandler)

	//Convention management
	r.HandleFunc(ROOT_API+"/pending/_random", RequireRole(RandomPendingConvention, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API+"/conventions/", RequireToken(GetAllConventions)).Methods("GET")
	r.HandleFunc(ROOT_API+"/conventions/", RequireRole(CommitPendingConvention, "admin")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/major", RequireRole(UpdateMajor, "major")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/tutor", RequireRole(UpdateTutor, "admin")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/midterm/mark", MarkKind(UpdateMark, "midterm")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/final/mark", MarkKind(UpdateMark, "final")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/supReport/mark", MarkKind(UpdateMark, "supReport")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/major", RequireRole(UpdateMajor, "major")).Methods("POST")

	//Student stuff
	r.HandleFunc(ROOT_API+"/convention", RequireToken(GetConvention)).Methods("GET")
	r.HandleFunc(ROOT_API+"/convention/company", RequireToken(UpdateCompany)).Methods("POST")
	r.HandleFunc(ROOT_API+"/convention/supervisor", RequireToken(UpdateSupervisor)).Methods("POST")

	//Reports
	r.HandleFunc(ROOT_API+"/reports/{kind}/", RequireRole(GetReports,"admin")).Methods("GET")
	r.HandleFunc(ROOT_API+"/reports/{kind}/{student}/document", RequireToken(GetReport)).Methods("GET")
	r.HandleFunc(ROOT_API+"/reports/{kind}/{student}/document", RequireToken(UploadReport)).Methods("POST")
	r.HandleFunc(ROOT_API+"/reports/{kind}/{student}/deadline", RequireToken(UpdateDeadline)).Methods("POST")
	r.HandleFunc(ROOT_API+"/reports/{kind}/{student}/mark", RequireToken(UpdateMark)).Methods("POST")

	//Profile management
	r.HandleFunc(ROOT_API+"/profile", RequireToken(GetUser)).Methods("GET")
	r.HandleFunc(ROOT_API+"/profile", RequireToken(ChangeProfile)).Methods("PUT")
	r.HandleFunc(ROOT_API+"/profile/password", RequireToken(ChangePassword)).Methods("PUT")
	r.HandleFunc(ROOT_API+"/profile/password", NewPassword).Methods("POST")
	r.HandleFunc(ROOT_API+"/profile/password", PasswordReset).Methods("DELETE")

	r.HandleFunc(ROOT_API+"/users/", RequireRole(GetAdmins, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API+"/users/", RequireRole(NewAdmin, "root")).Methods("POST")
	r.HandleFunc(ROOT_API+"/users/{email}/role", RequireRole(GrantRole, "root")).Methods("POST")
	r.HandleFunc(ROOT_API+"/users/{email}", RequireRole(RmUser, "root")).Methods("DELETE")
	r.HandleFunc(ROOT_API+"/login", Login).Methods("POST")

	r.HandleFunc(ROOT_API+"/logout", RequireToken(Logout)).Methods("GET")
	r.HandleFunc(ROOT_API+"/defenses", GetDefense).Methods("GET")
	r.HandleFunc(ROOT_API+"/defenses", RequireRole(PostDefense, "admin")).Methods("POST")
	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", r)
	log.Println("Daemon started")
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}
