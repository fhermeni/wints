package handler

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/fhermeni/wints/filter"
	"github.com/fhermeni/wints/internship"
)

// Create a Pool that contains previously used Writers and
// can create new ones if we run out.
var zippers = sync.Pool{New: func() interface{} {
	return gzip.NewWriter(nil)
}}

func authenticated(backend internship.Service, w http.ResponseWriter, r *http.Request) (string, error) {
	var err error
	cookie, err := r.Cookie("session")
	if err != nil || len(cookie.Value) == 0 {
		return "", ErrMissingCookies
	}
	token, err := r.Cookie("token")
	if err != nil || len(token.Value) == 0 {
		return "", ErrMissingCookies
	}

	err = backend.OpenedSession(cookie.Value, token.Value)
	return cookie.Value, err
}

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}

func writeJSONIfOk(e error, w http.ResponseWriter, r *http.Request, j interface{}) error {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Add("Vary", "Accept-Encoding")
	w.Header().Set("Content-Encoding", "gzip")
	// Get a Writer from the Pool
	gz := zippers.Get().(*gzip.Writer)

	// When done, put the Writer back in to the Pool
	defer zippers.Put(gz)

	// We use Reset to set the writer we want to use.
	gz.Reset(w)
	defer gz.Close()

	enc := json.NewEncoder(gz)
	return enc.Encode(j)
}

func status(w http.ResponseWriter, e error) {
	switch e {
	case internship.ErrInvalidToken, internship.ErrSessionExpired:
		return
	case internship.ErrUnknownSurvey, internship.ErrUnknownUser, internship.ErrUnknownReport, internship.ErrUnknownInternship:
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	case internship.ErrReportExists, internship.ErrUserExists, internship.ErrInternshipExists, internship.ErrCredentials, internship.ErrUserTutoring:
		http.Error(w, e.Error(), http.StatusConflict)
		return
	case internship.ErrInvalidSurvey, internship.ErrInvalidAlumniEmail, internship.ErrDeadlinePassed, internship.ErrInvalidGrade, internship.ErrInvalidMajor, internship.ErrInvalidPeriod, internship.ErrGradedReport:
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	case filter.ErrPermission:
		http.Error(w, e.Error(), http.StatusForbidden)
		return
	case nil:
		return
	default:
		http.Error(w, "Internal server error. A possible bug to report", http.StatusInternalServerError)
		log.Printf("Unsupported error: %s\n", e.Error())
	}
}

func fileReplyIfOk(e error, w http.ResponseWriter, mime, filename string, cnt []byte) error {
	if e != nil {
		return e
	}
	w.Header().Set("Content-type", mime)
	w.Header().Set("Content-disposition", "attachment; filename="+filename)
	_, err := w.Write(cnt)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println("Unable to send file '" + filename + "': " + err.Error())
	}
	return err
}
