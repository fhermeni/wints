package feeder

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"code.google.com/p/go-charset/charset"
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
		Encoding: "windows-1252",
	}
}

//Reader read all the conventions for a given year and promotion
func (h *HTTPConventionReader) Reader(year int, promotion string) (io.Reader, error) {
	client := &http.Client{}
	query := fmt.Sprintf("%s?action=down&filiere=%s&annee=%d", h.url, url.QueryEscape(promotion), year)
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(h.login, h.password)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//Convert to utf8
	return charset.NewReader(h.Encoding, res.Body)
}
