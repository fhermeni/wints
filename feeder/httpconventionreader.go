package feeder

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/fhermeni/wints/journal"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

//HTTPConventionReader reads conventions that are made available from the Http server
type HTTPConventionReader struct {
	url      string
	login    string
	password string
	Encoding string
	j        journal.Journal
}

//NewHTTPConventionReader creates a new reader.
//Default encoding is "windows-1252"
func NewHTTPConventionReader(url, login, password string, j journal.Journal) *HTTPConventionReader {
	return &HTTPConventionReader{
		url:      url,
		login:    login,
		password: password,
		Encoding: "windows-1252",
		j:        j,
	}
}

//Reader read all the conventions for a given year and promotion
func (h *HTTPConventionReader) Reader(year int, promotion string) (io.Reader, error) {
	client := &http.Client{}
	client.Timeout = time.Duration(5) * time.Second
	query := fmt.Sprintf("%s?action=down&filiere=%s&annee=%d", h.url, url.QueryEscape(promotion), year)
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(h.login, h.password)
	res, err := client.Do(req)
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			if err, ok := e.Err.(net.Error); ok && err.Timeout() {
				return nil, ErrTimeout
			}
		} else if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, ErrTimeout
		}
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, ErrAuthorization
	}
	return charset.NewReader(h.Encoding, res.Body)
}

var ErrAuthorization = errors.New("Permission denied")
var ErrTimeout = errors.New("Timeout")
var ErrNotFound = errors.New("No convention file")
