package feeder

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
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
	if res.StatusCode != 200 {
		return nil, ErrAuthorization
	}
	if err != nil {
		return nil, err
	}
	//Convert to utf8
	//buf, err := ioutil.ReadAll(res.Body)
	//log.Println(string(buf))
	//return res.Body, err
	return charset.NewReader(h.Encoding, res.Body)
}

var ErrAuthorization = errors.New("Permission denied")
