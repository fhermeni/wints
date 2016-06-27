package httpd

import (
	"errors"
	"net/http"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/sqlstore"

	"github.com/daaku/go.httpgzip"
)

//HTTPd just wrap a HTTP daemon that will serve both the static assets
//and the rest endpoints
type HTTPd struct {
	cfg   config.HTTPd
	store *sqlstore.Store
}

//NewHTTPd makes a new HTTP daemon
func NewHTTPd(not *notifier.Notifier, store *sqlstore.Store, conventions feeder.Conventions, cfg config.HTTPd, org config.Internships) HTTPd {

	//The assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/"+cfg.Assets, httpgzip.NewHandler(fs)))
	//The rest endpoints
	rest := NewEndPoints(not, store, conventions, cfg.Rest, org)
	http.HandleFunc(cfg.Rest.Prefix, Mon(rest.router.ServeHTTP))

	httpd := HTTPd{
		cfg:   cfg,
		store: store,
	}
	//the pages
	for _, p := range []string{"home", "login", "password", "statistics", "defense-program"} {
		http.HandleFunc("/"+p, httpd.page(p+".html"))
	}
	http.HandleFunc("/", httpd.home)
	http.HandleFunc("/survey", httpd.survey)
	return httpd
}

func (ed *HTTPd) page(path string) func(http.ResponseWriter, *http.Request) {
	return Mon(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=157680000; includeSubDomains; preload")
		http.ServeFile(w, r, ed.cfg.Assets+path)
	})
}

func (ed *HTTPd) home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	s, err := ed.store.Session([]byte(c.Value))
	if err != nil || s.Expire.Before(time.Now()) {
		msg := "session expired"
		if err != nil {
			msg = "invalid session token"
		}
		logger.Log("event", s.Email, msg, err)
		wipeCookies(w)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

//Listen starts listening
func (ed *HTTPd) Listen() error {
	return http.ListenAndServeTLS(ed.cfg.Listen, ed.cfg.Certificate, ed.cfg.PrivateKey, nil)
}

//Wipe the cookies
func wipeCookies(w http.ResponseWriter) {
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
	http.SetCookie(w, token)
	http.SetCookie(w, login)
}

var (
	//ErrMalformedJSON reports a JSON message that cannot be mapped to a struct
	ErrMalformedJSON = errors.New("Malformed json message")
	//ErrNotActivatedAccount refines a schema.ErrCredential when the user never logged before.
	ErrNotActivatedAccount = errors.New("You did not activate the account. Check for the invitation mail in this email mailbox")
)
