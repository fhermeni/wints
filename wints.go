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
	"time"
	"os"
	"io/ioutil"
)

var DB *sql.DB
var store = sessions.NewCookieStore([]byte("wints"))

func reportError(w http.ResponseWriter, header string, err error) {
	if err == nil {
		return
	}
	if be, ok := err.(*backend.BackendError); ok {
		http.Error(w, be.Error(), be.Status)
	} else {
		http.Error(w, "", http.StatusInternalServerError)
		log.Printf("%s: %s\n", header, err.Error())
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "static/login.html", 301)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

func jsonReply(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-type","application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	err := enc.Encode(j)
	if err != nil {
		reportError(w, "Error while serialising response", err)
	}
}

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&j)
	if err != nil {
		log.Printf("Bad JSON message format: %s", r.Body)
		http.Error(w, "Bad JSON message format", http.StatusBadRequest)
		return  err
	}
	return nil
}

func pullConventions() {
	conventions, err := backend.GetAllRawConventions()
	if err != nil {
		log.Printf("%s\n", err)
	} else {
		for _,c := range conventions {
			backend.InspectRawConvention(DB, c)
		}
	}
}

func daemonConventionsPuller() {
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				pullConventions();
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

type PendingConventionMsg struct {
	C backend.Convention
	Known []backend.User
	Pending int
	Total int
}

func RandomPendingConvention(w http.ResponseWriter, r *http.Request, email string) {
	c, err := backend.PeekRawConvention(DB)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error2: %s\n", err)
		return
	}

	nbPending, err := backend.CountPending(DB)
	if err != nil {
		log.Printf("Error at counting: %s\n", err)
		return
	}
	cc, err := backend.GetConventions(DB)
	if err != nil {
		log.Printf("Error at getting: %s\n", err)
		return
	}
	nbCommitted := len(cc)
	known, err := backend.AvailableTutors(DB)
	if err != nil {
		log.Printf("Error at getting tutors: %s\n", err)
		return
	}
	jsonReply(w, PendingConventionMsg{c, known, nbPending, nbPending + nbCommitted})
}

func GetAllConventions(w http.ResponseWriter, r *http.Request, email string) {
	conventions, err := backend.GetConventions(DB)
	if err != nil {
		log.Printf("Error here: %s\n", err)
		jsonReply(w,make([]backend.Convention, 0, 0))
		return
	}
	jsonReply(w, conventions)
}


func GetAdmins(w http.ResponseWriter, r *http.Request, email string) {
	admins, err := backend.Admins(DB)
	if err != nil {
		log.Printf("Error admins: %s\n", err)
		return
	}
	jsonReply(w, admins)
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
	if err != nil {
		reportError(w, "", err)
		return
	}
}

func RmUser(w http.ResponseWriter, r *http.Request, email string) {
	target,_ := mux.Vars(r)["email"]
	err := backend.RmUser(DB, target)
	if err != nil {
		reportError(w, "", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred backend.Credential
	if jsonRequest(w, r, &cred) != nil {
		return
	}
	u, err := backend.Register(DB, cred)
	log.Printf("Login %s\n", u)
	if err != nil {
		http.Error(w, "Unknown username or incorrect password", http.StatusUnauthorized)
		return
	}
	tok, err := backend.OpenSession(DB, cred.Email)
	if err != nil {
		reportError(w, "Error at opening session", err)
		return
	}
	s, _ := store.Get(r, "wints")
	s.Values["token"] = tok
	s.Values["email"] = cred.Email
	s.Save(r, w)
	w.Header().Add("X-auth-token", string(tok))
	jsonReply(w, u)
}

func Logout(w http.ResponseWriter, r *http.Request, email string) {
	token := r.Header.Get("X-auth-token")
	if (len(token) == 0) {
		http.Error(w, "X-auth-token is missing", http.StatusUnauthorized)
		return
	}
	err := backend.CloseSession(DB, token)
	if err != nil {
		log.Printf("Unable to logout: %s\n", err)
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
		log.Printf("Register the new tutor %s\n", c.Tutor)
		if err != nil {
			log.Printf("Unable to register the new tutor %s: %s\n", c.Tutor, err)
			return
		}
	} else {
		log.Println("No need to register the tutor. Already in")
	}
	err = backend.RegisterInternship(DB, c, true);
	if err != nil {
		log.Printf("Unable to register the internship: %s\n", err)
		return
	}
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	backend.RescanPending(DB, c)
}

type PersonUpdate struct {
	Firstname string
	Lastname string
	Tel string
}

func ChangeProfile(w http.ResponseWriter, r *http.Request, email string) {
	log.Printf("Permission check")
	var p PersonUpdate
	if jsonRequest(w, r, &p) != nil {
		return
	}
	err := backend.SetProfile(DB, email, p.Firstname, p.Lastname, p.Tel)
	u, err := backend.GetUser(DB, email)
	reportError(w, "", err)

	jsonReply(w, u)
}

func UpdatePromotion(w http.ResponseWriter, r *http.Request, email string) {
	//TODO: permission
	p, err := ioutil.ReadAll(r.Body)
	target,_ := mux.Vars(r)["email"]
	if err != nil {
		reportError(w, "", err)
	}
	err = backend.SetPromotion(DB, target, string(p))
	if err != nil {
		reportError(w, "", err)
	}
}

func UpdateMajor(w http.ResponseWriter, r *http.Request, email string) {
	//TODO: permission
	target,_ := mux.Vars(r)["email"]
	m, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reportError(w, "", err)
	}
	err = backend.SetMajor(DB, target, string(m))
	if err != nil {
		reportError(w, "", err)
	}
}

type PasswordRenewal struct {
	OldPassword []byte
	NewPassword []byte
}

/*func ChangePassword(w http.ResponseWriter, r *http.Request, uid int) {
	var p PasswordRenewal
	if jsonRequest(w, r, &p) != nil {
		return
	}

	err := backend.NewPassword(DB, uid, p.OldPassword, p.NewPassword)
	if err != nil {
		reportError(w, "", err)
	}
}         */

func GrantRole(w http.ResponseWriter, r *http.Request, email string) {
	//TODO: permission
	target,_ := mux.Vars(r)["email"]
	role, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reportError(w, "", err)
	}
	err = backend.GrantPrivilege(DB, target, string(role))
	if err != nil {
		reportError(w, "", err)
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

func main() {

	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		dbURL = "user=fhermeni dbname=wints host=localhost password=wints sslmode=disable"
	}
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	backend.NewUser(DB, backend.User{"Fabien", "Hermenier", "fabien.hermenier@unice.fr", "0662496455", "root"})

	go pullConventions()
	daemonConventionsPuller();

	//Rest stuff
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	r.HandleFunc("/conventions/_random", RequireToken(RandomPendingConvention)).Methods("GET")
	r.HandleFunc("/conventions/", RequireToken(GetAllConventions)).Methods("GET")
	r.HandleFunc("/conventions/", RequireToken(CommitPendingConvention)).Methods("POST")
	r.HandleFunc("/conventions/{email}/major", RequireToken(UpdateMajor)).Methods("POST")
	r.HandleFunc("/admins/", RequireToken(GetAdmins)).Methods("GET")
	r.HandleFunc("/users/", RequireToken(NewAdmin)).Methods("POST")
	r.HandleFunc("/users/{email}/roles/", RequireToken(GrantRole)).Methods("POST")
	r.HandleFunc("/users/{email}", RequireToken(RmUser)).Methods("DELETE")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/profile", RequireToken(ChangeProfile)).Methods("POST")
	r.HandleFunc("/logout", RequireToken(Logout)).Methods("POST")
	// handle all requests by serving a file of the same name
	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", r)
	log.Println("Daemon started")
	err = http.ListenAndServe(":" + os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer DB.Close()
}
