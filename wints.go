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
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Refreshing conventions")
				conventions, err := backend.GetAllRawConventions()
				if err != nil {
					log.Printf("%s\n", err)
				} else {
					for _,c := range conventions {
						backend.InspectRawConvention(DB, c)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

type PendingConventionMsg struct {
	C backend.Convention
	Known []backend.Person
	Pending int
	Total int
}
func RandomPendingConvention(w http.ResponseWriter, r *http.Request) {
	c, err := backend.PeekRawConvention(DB)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error2: %s\n", err)
		return
	}

	nbPending, err := backend.CountPending(DB)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	cc, err := backend.GetConventions(DB)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	nbCommitted := len(cc)
	known, err := backend.Tutors(DB)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	jsonReply(w, PendingConventionMsg{c, known, nbPending, nbPending + nbCommitted})
}

func GetAllConventions(w http.ResponseWriter, r *http.Request) {
	conventions, err := backend.GetConventions(DB)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	jsonReply(w, conventions)
}


func GetAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := backend.Admins(DB)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	jsonReply(w, admins)
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
	if err != nil {
		reportError(w, "Error at opening session", err)
		return
	}
	s, _ := store.Get(r, "wints")
	s.Values["token"] = tok
	s.Values["email"] = cred.Email
	s.Save(r, w)
	w.Header().Add("X-AUTH-TOKEN", string(tok))
	jsonReply(w, u)
}

func CommitPendingConvention(w http.ResponseWriter, r *http.Request) {
	var c backend.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return
	}
	tEmail := c.Tutor.Email
	_, err = backend.GetPerson(DB, tEmail)
	if err != nil {
		err = backend.NewTutor(DB, c.Tutor)
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

func main() {

	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		dbURL = "user=fhermeni dbname=wints host=localhost password=wints sslmode=disable"
	}
	log.Println(dbURL)
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	pullConventions()

	//Rest stuff
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	//User mgnt
	r.HandleFunc("/conventions/_random", RandomPendingConvention).Methods("GET")
	r.HandleFunc("/conventions/", GetAllConventions).Methods("GET")
	r.HandleFunc("/conventions/", CommitPendingConvention).Methods("POST")
	r.HandleFunc("/admins/", GetAdmins).Methods("GET")
	r.HandleFunc("/login", Login).Methods("POST")
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
