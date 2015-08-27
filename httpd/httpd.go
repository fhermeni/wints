package httpd

import (
	"net/http"

	"github.com/daaku/go.httpgzip"
	"github.com/fhermeni/wints/config"
)

type HTTPd struct {
	cfg config.Config
}

func NewHTTPd(cfg config.Config) HTTPd {

	//The rest endpoints
	rest := NewEndPoints("/api/v1")
	http.Handle(rest.prefix, rest.router)

	//The assets
	fs := http.Dir(cfg.HTTPd.Assets + "/")
	fileHandler := http.FileServer(fs)
	http.Handle("/", httpgzip.NewHandler(http.StripPrefix("/"+cfg.HTTPd.Assets, fileHandler)))

	return HTTPd{
		cfg: cfg,
	}
}
