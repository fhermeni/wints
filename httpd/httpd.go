package httpd

import (
	"net/http"

	"github.com/daaku/go.httpgzip"
	"github.com/fhermeni/wints/config"
)

type HTTPd struct {
	cfg config.HTTPd
}

func NewHTTPd(cfg config.HTTPd) HTTPd {

	//The assets
	fs := http.Dir(cfg.Assets + "/")
	fileHandler := http.FileServer(fs)
	http.Handle("/", httpgzip.NewHandler(fileHandler))

	//The rest endpoints
	rest := NewEndPoints(cfg.Rest.Prefix)
	http.Handle(cfg.Rest.Prefix, rest.router)

	return HTTPd{
		cfg: cfg,
	}
}

func (d *HTTPd) Listen() error {
	return http.ListenAndServeTLS(d.cfg.Listen, d.cfg.Certificate, d.cfg.PrivateKey, nil)
}
