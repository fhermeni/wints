package httpd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fhermeni/wints/logger"
	"github.com/stathat/go"
)

//MonResponseWriter embeds a responsewriter to save the status code
//then to provide a proper monitoring
type MonResponseWriter struct {
	status int
	http.ResponseWriter
}

//NewMonResponseWriter makes a new monitoring response writer
func NewMonResponseWriter(res http.ResponseWriter) *MonResponseWriter {
	// Default the status code to 200
	return &MonResponseWriter{200, res}
}

//Status returns the status code provided by WriteHeader
func (w *MonResponseWriter) Status() int {
	return w.status
}

//Header delegates to the underlyning writer
func (w *MonResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

//Writer delegates to the underlyning writer
func (w *MonResponseWriter) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

//WriteHeader saves the status code and delegates to the underlyning writer
func (w *MonResponseWriter) WriteHeader(statusCode int) {
	// Store the status code
	w.status = statusCode
	// Write the status code onward.
	w.ResponseWriter.WriteHeader(statusCode)
}

//Mon monitores the requests
func Mon(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myRw := NewMonResponseWriter(w)
		start := time.Now()
		defer func() {
			ms := int(time.Since(start).Nanoseconds() / 1000000)
			msg := fmt.Sprintf("\"%s\" %d %d", r.URL.String(), myRw.Status(), ms)
			logger.Log("access", r.Method, msg, nil)
			stathat.PostEZValue("API latency", "fabien.hermenier@unice.fr", float64(ms))
			if myRw.Status() >= 500 {
				stathat.PostEZCount("HTTP 5xx", "fabien.hermenier@unice.fr", 1)
			} else if myRw.Status() >= 400 {
				stathat.PostEZCount("HTTP 4xx", "fabien.hermenier@unice.fr", 1)
			} else if myRw.Status() >= 200 {
				stathat.PostEZCount("HTTP 2xx", "fabien.hermenier@unice.fr", 1)
			}

		}()
		h(myRw, r)
	}
}
