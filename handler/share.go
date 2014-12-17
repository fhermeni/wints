package handler

import (
	"compress/gzip"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Route struct {
	Method string
	Path string
	Handler http.HandlerFunc
}

func NewRouter() *mux.Router {
	//Rest stuff
	r := mux.NewRouter()

	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	return r
}

func FileReply(w http.ResponseWriter, mime, filename string, cnt []byte) {
	w.Header().Set("Content-type", mime)
	w.Header().Set("Content-disposition", "attachment; filename="+filename)
	_, err := w.Write(cnt)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println("Unable to send file '" + filename + "': " + err.Error())
	}
}

func RequiredContent(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	cnt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}
	if len(cnt) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return []byte{}, err
	}
	return cnt, nil
}

func JsonReply(w http.ResponseWriter, gz bool, j interface{}) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if gz {
		w.Header().Set("Content-Encoding", "gzip")
	}
	var err error
	if gz {
		var w2 *gzip.Writer
		w2 = gzip.NewWriter(w)
		defer w2.Close()
		enc := json.NewEncoder(w2)
		err = enc.Encode(j)
	} else {
		enc := json.NewEncoder(w)
		err = enc.Encode(j)
	}
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println("Unable to send json " + err.Error())
	}
}

func JsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(&j)
}

func SetSession(email string, response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: email,
		Path:  "/",
	}
	http.SetCookie(response, cookie)
}

func Session(w http.ResponseWriter, r *http.Request) (string, error) {
	var err error
	cookie, err := r.Cookie("session")
	if len(cookie.Value) == 0 {
		http.Redirect(w, r, "/", 302)
		return "", err
	}
	return cookie.Value, err
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
