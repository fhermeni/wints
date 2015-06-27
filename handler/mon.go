package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fhermeni/wints/journal"
)

// Create our own MyResponseWriter to wrap a standard http.ResponseWriter
// so we can store the status code.
type MyResponseWriter struct {
	status int
	http.ResponseWriter
}

func NewMyResponseWriter(res http.ResponseWriter) *MyResponseWriter {
	// Default the status code to 200
	return &MyResponseWriter{200, res}
}

// Give a way to get the status
func (w MyResponseWriter) Status() int {
	return w.status
}

// Satisfy the http.ResponseWriter interface
func (w MyResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w MyResponseWriter) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

func (w MyResponseWriter) WriteHeader(statusCode int) {
	// Store the status code
	w.status = statusCode
	log.Println("Store status " + strconv.Itoa(statusCode))
	// Write the status code onward.
	w.ResponseWriter.WriteHeader(statusCode)
}

func mon(j journal.Journal, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myRw := w //NewMyResponseWriter(w)
		start := time.Now()
		defer func() {
			log.Println("st: " + strconv.Itoa(myRw.Status()))
			j.Access(r.Method, r.URL.String(), 0, int(time.Since(start).Nanoseconds()/1000000))
		}()
		log.Println(r.URL.String())
		h(myRw, r)
	}
}
