package httpd

import (
	"errors"
	"net/http"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/sqlstore"
)

//HTTPd just wrap a HTTP daemon that will serve both the static assets
//and the rest endpoints
type HTTPd struct {
	cfg   config.HTTPd
	store *sqlstore.Store
	not   *notifier.Notifier
}

//NewHTTPd makes a new HTTP daemon
func NewHTTPd(not *notifier.Notifier, store *sqlstore.Store, conventions feeder.Conventions, cfg config.HTTPd, org config.Internships) HTTPd {

	//The assets
	fileHandler := http.FileServer(http.Dir(cfg.Assets))
	http.Handle("/"+cfg.Assets, http.StripPrefix("/"+cfg.Assets, fileHandler))
	//The rest endpoints
	rest := NewEndPoints(not, store, conventions, cfg.Rest, org)
	http.HandleFunc(cfg.Rest.Prefix, Mon(not, rest.router.ServeHTTP))

	httpd := HTTPd{
		cfg:   cfg,
		store: store,
		not:   not,
	}
	//the pages
	for _, p := range []string{"home", "login", "password"} {
		http.HandleFunc("/"+p, httpd.page(p+".html"))
	}
	http.HandleFunc("/", httpd.home)
	http.HandleFunc("/survey", httpd.survey)
	return httpd
}

func (ed *HTTPd) page(path string) func(http.ResponseWriter, *http.Request) {
	return Mon(ed.not, func(w http.ResponseWriter, r *http.Request) {
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
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

//Listen starts listening
func (d *HTTPd) Listen() error {
	//return http.ListenAndServe(d.cfg.Listen, d.cfg.Certificate, d.cfg.PrivateKey, nil)
	return http.ListenAndServe(d.cfg.Listen, nil) //, d.cfg.Certificate, d.cfg.PrivateKey, nil)
}

var (
	ErrMalformedJSON = errors.New("Malformed json message")
)
