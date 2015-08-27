package httpd

import (
	"log"
	"net/http"
	"time"
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
			log.Printf("%s %s %d %d", r.Method, r.URL.String(), myRw.Status(), int(time.Since(start).Nanoseconds()/1000000))
		}()
		h(myRw, r)
	}
}
