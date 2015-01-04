package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fhermeni/wints/filter"
	"github.com/fhermeni/wints/internship"
)

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}

func writeJSONIfOk(e error, w http.ResponseWriter, j interface{}) error {
	if e != nil {
		return e
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	return enc.Encode(j)
}

func serviceHandler(cb func(internship.Service, http.ResponseWriter, *http.Request) error, srv Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.URL.String())
		var err error
		cookie, err := r.Cookie("session")
		if err != nil || len(cookie.Value) == 0 {
			http.Redirect(w, r, "/", 302)
			return
		}
		u, err := srv.backend.User(cookie.Value)
		if err == internship.ErrUnknownUser {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else if err != nil {
			log.Println("Unable to get the user behind the cookie: " + err.Error())
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		wrapper, _ := filter.NewService(srv.backend, u)
		e := cb(wrapper, w, r)
		//Error management
		if e != nil {
			log.Println("Reply with error " + e.Error())
		}
		switch e {
		case internship.ErrUnknownUser, internship.ErrUnknownReport, internship.ErrUnknownInternship:
			http.Error(w, e.Error(), http.StatusNotFound)
			return
		case internship.ErrReportExists, internship.ErrUserExists, internship.ErrInternshipExists, internship.ErrCredentials:
			http.Error(w, e.Error(), http.StatusConflict)
			return
		case filter.ErrPermission:
			http.Error(w, e.Error(), http.StatusForbidden)
			return
		case nil:
			return
		default:
			http.Error(w, e.Error(), http.StatusInternalServerError)
			log.Printf("Unsupported error: %s\n", e.Error())
		}
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
