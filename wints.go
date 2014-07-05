package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/fhermeni/wints/backend"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"encoding/json"
	"log"
	"os"
	"errors"
	"strconv"
	"math/rand"
	"io/ioutil"
	"time"
)

const (
	ROOT_API = "/api/v1"
)
var DB *sql.DB
var store = sessions.NewCookieStore([]byte("wints"))

func reportIfError(w http.ResponseWriter, header string, err error) bool {
	if err == nil {
		return false
	}
	if be, ok := err.(*backend.BackendError); ok {
		http.Error(w, be.Error(), be.Status)
	} else {
		http.Error(w, "", http.StatusInternalServerError)
		log.Printf("%s: %s\n", header, err.Error())
	}
	return true
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "static/login.html", 301)
}

func jsonReply(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-type","application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	err := enc.Encode(j)
	reportIfError(w, "Error while serialising response", err)
}

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&j)
	if err != nil {
		log.Printf("Bad JSON message format: %s", err)
		http.Error(w, "Bad JSON message format", http.StatusBadRequest)
		return  err
	}
	return nil
}

type PendingConventionMsg struct {
	C backend.Convention
	Known []backend.User
	Pending int
}

func RandomPendingConvention(w http.ResponseWriter, r *http.Request, email string) {
	cc, err := backend.GetRawConventions(DB)
	nb := len(cc)
	if nb == 0 {
		jsonReply(w, PendingConventionMsg{backend.Convention{}, nil, 0})
		return
	}
	c := cc[rand.Intn(nb)];
	known, err := backend.Admins(DB)
	if !reportIfError(w, "Unable to get the possible tutors", err) {
		jsonReply(w, PendingConventionMsg{c, known, nb})
	}
}

func GetAllConventions(w http.ResponseWriter, r *http.Request, email string) {
	conventions, err := backend.GetConventions(DB)
	if !reportIfError(w, "Unable to get the conventions ", err) {
		jsonReply(w, conventions)
	}
}


func GetAdmins(w http.ResponseWriter, r *http.Request, email string) {
	admins, err := backend.Admins(DB)
	if !reportIfError(w, "Unable to get the possible admins", err) {
		jsonReply(w, admins)
	}
}

type NewUser struct {
	Firstname string
	Lastname string
	Tel string
	Email string
	Priv string
}

func NewAdmin(w http.ResponseWriter, r *http.Request, email string) {
	var nu NewUser
	err := jsonRequest(w, r, &nu)
	if err != nil {
		return
	}
	p := backend.User{nu.Firstname, nu.Lastname, nu.Email, nu.Tel, nu.Priv}
	err = backend.NewUser(DB, p)
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, " User '" + nu.Email + "'(" + nu.Priv + ") created");
	}
}

func RmUser(w http.ResponseWriter, r *http.Request, email string) {
	target,_ := mux.Vars(r)["email"]
	err := backend.RmUser(DB, target)
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, " User '" + target + "' deleted");
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred backend.Credential
	if jsonRequest(w, r, &cred) != nil {
		return
	}
	u, err := backend.Register(DB, cred)
	if err != nil {
		http.Error(w, "Unknown username or incorrect password", http.StatusUnauthorized)
		return
	}
	tok, err := backend.OpenSession(DB, cred.Email)
	if reportIfError(w, "Unable to open the session", err) {
		return
	}
	s, err := store.Get(r, "wints")
	if reportIfError(w, "Unable to get the session identifier", err) {
		return
	}
	s.Values["token"] = tok
	s.Values["email"] = cred.Email
	s.Save(r, w)
	w.Header().Add("X-auth-token", string(tok))
	jsonReply(w, u)
	backend.LogActionInfo(cred.Email, " log in");
}

func Logout(w http.ResponseWriter, r *http.Request, email string) {
	token := r.Header.Get("X-auth-token")
	if (len(token) == 0) {
		http.Error(w, "X-auth-token is missing", http.StatusUnauthorized)
		return
	}
	err := backend.CloseSession(DB, token)
	if !reportIfError(w, "Unable to logout", err) {
		backend.LogActionInfo(email, " log out");
	}
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
		err = backend.NewUser(DB, c.Tutor)
		if reportIfError(w, "Unable to register the tutor " + c.Tutor.String(), err) {
			return
		}
	}
	err = backend.RegisterInternship(DB, c, true);
	if reportIfError(w, "Unable to register the internship", err) {
		return
	}
	backend.LogActionInfo(email, "validate the convention of '" + c.Stu.P.Email + "'");
	backend.RescanPending(DB, c)
}

type PersonUpdate struct {
	Firstname string
	Lastname string
	Tel string
}

func ChangeProfile(w http.ResponseWriter, r *http.Request, email string) {
	target,_ := mux.Vars(r)["email"]
	if target != email {
		http.Error(w, "You cannot update the profile of another person", http.StatusForbidden)
		return
	}
	var p PersonUpdate
	if jsonRequest(w, r, &p) != nil {
		return
	}
	err := backend.SetProfile(DB, email, p.Firstname, p.Lastname, p.Tel)
	u, err := backend.GetUser(DB, email)
	if !reportIfError(w, "", err) {
		jsonReply(w, u)
	}
}

type PasswordRenewal struct {
	OldPassword string
	NewPassword string
}

func ChangePassword(w http.ResponseWriter, r *http.Request, email string) {
	target,_ := mux.Vars(r)["email"]
	if target != email {
		http.Error(w, "You cannot update the profile of another person", http.StatusForbidden)
		return
	}

	var p PasswordRenewal
	if jsonRequest(w, r, &p) != nil {
		return
	}

	err := backend.NewPassword(DB, email, []byte(p.OldPassword), []byte(p.NewPassword))
	reportIfError(w, "", err)
}

func UpdateMajor(w http.ResponseWriter, r *http.Request, email string) {
	target,_ := mux.Vars(r)["email"]
	m, err := requiredContent(w, r)
	if err != nil {
		return
	}
	err = backend.SetMajor(DB, target, string(m))
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, "Major of '" + target + "' set to '" + string(m) + "'");
	}
}

func UpdateMidtermDeadline(w http.ResponseWriter, r *http.Request, email string) {
	d, err := requiredContent(w, r)
	if err != nil {
		return
	}
	t, err := time.Parse("2/1/2006", string(d))
	if reportIfError(w, "Invalid format", err) {
		return
	}
	target,_ := mux.Vars(r)["email"]
	u, err := backend.GetUser(DB, target);
	if u.Role == "" {
		c, err := backend.GetConvention(DB, target)
		if reportIfError(w, "", err) {
			return
		}
		if c.Tutor.Email != email {
			http.Error(w, "You cannot update the deadline on a student not under your responsibility", http.StatusForbidden)
			return
		}
	}
	//Other levels are ok
	err = backend.SetMidtermDeadline(DB, target, t)
	reportIfError(w, "", err);
}


func requiredContent(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	cnt, err := ioutil.ReadAll(r.Body)
	if reportIfError(w, "", err) {
		return []byte{}, err
	}
	if len(cnt) == 0 {
		reportIfError(w, "", errors.New("Missing request body"))
	}
	return cnt, nil
}

func GrantRole(w http.ResponseWriter, r *http.Request, email string) {
	target,_ := mux.Vars(r)["email"]
	role, err := requiredContent(w, r)
	if err != nil {
		return
	}
	err = backend.GrantPrivilege(DB, target, string(role))
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, "'" + target + "' granted to role '" + string(role) + "'");
	}
}

func RequireToken(cb func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-auth-token")
		if (token == "") {
			http.Error(w, "Header 'X-auth-token' is missing", http.StatusForbidden)
		}
		email, err := backend.CheckSession(DB, token)
		if err != nil {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
		}
		cb(w, r, email)
	}
}

func RequireRole(cb func(http.ResponseWriter, *http.Request, string), role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-auth-token")
		if (token == "") {
			http.Error(w, "Header 'X-auth-token' is missing", http.StatusForbidden)
		}
		email, err := backend.CheckSession(DB, token)
		if err != nil {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
		}

		u, _ := backend.GetUser(DB, email)
		if u.CompatibleRole(role) {
			cb(w, r, email)
		} else {
			http.Error(w, "This operation requires '" + role + "' right but got only '" + u.Role + "'", http.StatusUnauthorized)
		}
	}
}

func main() {

	cfg, err := backend.ReadConfig("./wints.conf")
	if err != nil {
		log.Fatalln(err.Error())
	}

	backend.InitLogger(cfg.Logfile)
	defer backend.CloseLogger();
	DB, err = sql.Open("postgres", cfg.DBurl)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if len(os.Args) > 1 && os.Args[1] == "-i" {
		err := backend.NewUser(DB, backend.User{"R", "oot", "root@localhost", "0662496455", "root"})
		if err != nil {
			backend.LogError("Unable to create the root account: " + err.Error());
		} else {
			backend.LogInfo("Root account created");
		}
		return
	}

	/*delay, err := time.ParseDuration(cfg.RefreshPeriod)
	if err != nil {
		log.Fatalln(err.Error())
	}
	backend.DaemonConventionsPuller(DB, cfg.WWWConventionsURL,
										cfg.WWWConventionsLogin,
										cfg.WWWConventionsPassword,
										delay);
	*/
	//Rest stuff
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	r.HandleFunc(ROOT_API + "/pending/_random", RequireRole(RandomPendingConvention, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API + "/conventions/", RequireRole(GetAllConventions,"major")).Methods("GET")
	r.HandleFunc(ROOT_API + "/conventions/", RequireRole(CommitPendingConvention,"admin")).Methods("POST")
	r.HandleFunc(ROOT_API + "/conventions/{email}/major", RequireRole(UpdateMajor, "major")).Methods("POST")
	r.HandleFunc(ROOT_API + "/conventions/{email}/midterm/deadline", RequireToken(UpdateMidtermDeadline)).Methods("POST")
	r.HandleFunc(ROOT_API + "/users/", RequireRole(GetAdmins, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API + "/users/", RequireRole(NewAdmin, "root")).Methods("POST")
	r.HandleFunc(ROOT_API + "/users/{email}/role", RequireRole(GrantRole, "root")).Methods("POST")
	r.HandleFunc(ROOT_API + "/users/{email}", RequireRole(RmUser, "root")).Methods("DELETE")
	r.HandleFunc(ROOT_API + "/users/{email}/password", RequireToken(ChangePassword)).Methods("POST")
	r.HandleFunc(ROOT_API + "/users/{email}/", RequireToken(ChangeProfile)).Methods("POST")
	r.HandleFunc(ROOT_API + "/login", Login).Methods("POST")
	r.HandleFunc(ROOT_API + "/logout", RequireToken(Logout)).Methods("POST")
	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", r)
	log.Println("Daemon started")
	err = http.ListenAndServe(":" + strconv.Itoa(cfg.Port), nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer DB.Close()
}
