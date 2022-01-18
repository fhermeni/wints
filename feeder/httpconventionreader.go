package feeder

import (
	"crypto/tls"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"io"
	"net/http"
	"time"

	"golang.org/x/text/encoding/charmap"
)

//HTTPConventionReader reads conventions that are made available from the Http server
type HTTPConventionReader struct {
	url      string
	login    string
	password string
	Encoding string
}

//NewHTTPConventionReader creates a new reader.
//Default encoding is "windows-1252"
func NewHTTPConventionReader(url, login, password string) *HTTPConventionReader {
	return &HTTPConventionReader{
		url:      url,
		login:    login,
		password: password,
	}
}

//Reader read all the conventions for a given year and promotion
func (h *HTTPConventionReader) Reader(year int, promotion string) (io.Reader, error) {
	//only in dev mode
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport : tr}
	/* in non dev mode}
	else{
		client := &http.Client{}
	}*/
	client.Timeout = time.Duration(10) * time.Second
	query := fmt.Sprintf("%s%s&parcours=*", h.url, promotion)
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(h.login, h.password)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Unexpected response status: " + res.Status)
	}
	if h.Encoding == "windows-1252" {
		return charmap.Windows1252.NewDecoder().Reader(res.Body), nil
	}
	if h.Encoding == "UTF-8" {
		return unicode.UTF8.NewDecoder().Reader(res.Body), nil
	}
	return nil, errors.New("Unsupported encoding: " + h.Encoding)
}

var (
	//ErrUnsupportedEncoding signals an encoding that cannot be converted
	ErrUnsupportedEncoding = errors.New("Unsupported encoding")
	//ErrAuthorization declares a permission issue
	ErrAuthorization = errors.New("Permission denied")
	//ErrTimeout indicates a timeout
	ErrTimeout = errors.New("Timeout")
	//ErrNotFound states there is no convention file
	ErrNotFound = errors.New("No convention file")
)
