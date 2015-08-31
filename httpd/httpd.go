package httpd

import (
	"errors"
	"net/http"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/sqlstore"
)

//HTTPd just wrap a HTTP daemon that will serve both the static assets
//and the rest endpoints
type HTTPd struct {
	cfg   config.HTTPd
	store *sqlstore.Store
}

//NewHTTPd makes a new HTTP daemon
func NewHTTPd(store *sqlstore.Store, cfg config.HTTPd, org config.Internships) HTTPd {

	//The assets
	fileHandler := http.FileServer(http.Dir(cfg.Assets))
	http.Handle("/"+cfg.Assets, http.StripPrefix("/"+cfg.Assets, fileHandler))
	//The rest endpoints
	rest := NewEndPoints(store, cfg.Rest, org)
	http.Handle(cfg.Rest.Prefix, Mon(rest.router.ServeHTTP))

	httpd := HTTPd{
		cfg:   cfg,
		store: store,
	}
	//the pages
	for _, p := range []string{"home", "login", "password"} {
		http.HandleFunc("/"+p, httpd.page(p+".html"))
	}
	http.HandleFunc("/", httpd.home)
	return httpd
}

func (ed *HTTPd) page(path string) func(http.ResponseWriter, *http.Request) {
	return Mon(func(w http.ResponseWriter, r *http.Request) {
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
	return http.ListenAndServeTLS(d.cfg.Listen, d.cfg.Certificate, d.cfg.PrivateKey, nil)
}

var (
	ErrMalformedJSON = errors.New("Malformed json message")
)
