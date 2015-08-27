package httpd

import (
	"net/http"

	"github.com/daaku/go.httpgzip"
	"github.com/fhermeni/wints/config"
)

//HTTPd just wrap a HTTP daemon that will serve both the static assets
//and the rest endpoints
type HTTPd struct {
	cfg config.HTTPd
}

//NewHTTPd makes a new HTTP daemon
func NewHTTPd(cfg config.HTTPd) HTTPd {

	//The assets
	fs := http.Dir(cfg.Assets + "/")
	fileHandler := http.FileServer(fs)
	http.Handle("/", Mon(httpgzip.NewHandler(fileHandler).ServeHTTP))

	//The rest endpoints
	rest := NewEndPoints(cfg.Rest.Prefix)
	http.Handle(cfg.Rest.Prefix, Mon(rest.router.ServeHTTP))

	return HTTPd{
		cfg: cfg,
	}
}

//Listen starts listening
func (d *HTTPd) Listen() error {
	return http.ListenAndServeTLS(d.cfg.Listen, d.cfg.Certificate, d.cfg.PrivateKey, nil)
}
