package httpd

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/klauspost/compress/gzip"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/session"
)

//Exchange is a simple utility struct to gather the element
//manipulated while communicating with a client within a session
type Exchange struct {
	w   http.ResponseWriter
	r   *http.Request
	ps  map[string]string
	s   session.Session
	cfg config.Rest
	err error
	not *notifier.Notifier
}

// Create a Pool that contains previously used Writers and
// can create new ones if we run out.
var zippers = sync.Pool{New: func() interface{} {
	return gzip.NewWriter(nil)
}}

//V returns the value of the given parameter.
func (ex *Exchange) V(id string) string {
	v, ok := ex.ps[id]
	if !ok && ex.err == nil {
		ex.err = errors.New("Missing request parameter '" + id + "'")
	}
	return v
}

func (ex *Exchange) outJSON(j interface{}, e error) error {
	if e != nil {
		ex.err = e
	}
	if ex.err != nil {
		return ex.err
	}
	ex.w.Header().Set("Content-type", "application/json; charset=utf-8")
	ex.w.Header().Add("Vary", "Accept-Encoding")
	ex.w.Header().Set("Content-Encoding", "gzip")
	// Get a Writer from the Pool
	gz := zippers.Get().(*gzip.Writer)

	// When done, put the Writer back in to the Pool
	defer zippers.Put(gz)

	// We use Reset to set the writer we want to use.
	gz.Reset(ex.w)
	defer gz.Close()

	enc := json.NewEncoder(gz)
	ex.err = enc.Encode(j)
	return ex.err
}

func (ex *Exchange) inBytes(mime string) ([]byte, error) {
	if mime != ex.r.Header.Get("content-type") {
		err := errors.New("Expected content-type '" + mime + "'")
		http.Error(ex.w, err.Error(), http.StatusBadRequest)
		return []byte{}, err
	}
	return ioutil.ReadAll(ex.r.Body)
}

func (ex *Exchange) inJSON(j interface{}) error {
	if ex.err != nil {
		return ex.err
	}
	dec := json.NewDecoder(ex.r.Body)
	ex.err = dec.Decode(&j)
	return ex.err
}

func (ex *Exchange) out(mime string, in io.ReadCloser, e error) error {
	if e != nil {
		return e
	}
	defer in.Close()
	ex.w.Header().Set("Content-type", mime)

	ex.w.Header().Add("Vary", "Accept-Encoding")
	ex.w.Header().Set("Content-Encoding", "gzip")
	// Get a Writer from the Pool
	gz := zippers.Get().(*gzip.Writer)

	// When done, put the Writer back in to the Pool
	defer zippers.Put(gz)

	// We use Reset to set the writer we want to use.
	gz.Reset(ex.w)
	defer gz.Close()

	_, ex.err = io.Copy(gz, in)
	return ex.err
}

func (ex *Exchange) outFile(mime, filename string, cnt []byte, e error) error {
	if e != nil {
		return e
	}
	ex.w.Header().Set("Content-type", mime)
	ex.w.Header().Set("Content-disposition", "attachment; filename="+filename)
	_, err := ex.w.Write(cnt)
	if err != nil {
		logger.Log("event", ex.s.Me().Person.Email, "Unable to send '"+filename+"'", err)
		http.Error(ex.w, "", http.StatusInternalServerError)
	}
	return err
}
