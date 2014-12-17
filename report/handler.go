package report

import (
	"github.com/fhermeni/wints/handler"
	"github.com/fhermeni/wints/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"	
	"time"
)

func Routes(rs ReportService, us user.UserService) []handler.Route{
	return []handler.Route{
		handler.Route{"POST", "/reports/{stu}/{kind}/mark",withReportService(SetMark, rs)},
		handler.Route{"POST", "/reports/{stu}/{kind}/deadline", withReportService(SetDeadline, rs)},
		handler.Route{"POST", "/reports/{stu}/{kind}/content", withReportService(SetContent, rs)},
		handler.Route{"GET", "/reports/{stu}/{kind}/content", withServices(Content, us, rs)},
		handler.Route{"GET", "/reports/{stu}", withReportService(List, rs)},
	}
}

func withReportService(cb func(ReportService, http.ResponseWriter, *http.Request), rs ReportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cb(rs, w, r)
	}
}

func withServices(cb func(user.UserService, ReportService, http.ResponseWriter, *http.Request), us user.UserService, rs ReportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cb(us, rs, w, r)
	}
}

func reportError(w http.ResponseWriter, msg string, e error) bool {
	if e != nil {
		switch e {
		case ErrUnknown:
			http.Error(w, e.Error(), http.StatusNotFound)
			return true
		case ErrExists,ErrConflict:
			http.Error(w, e.Error(), http.StatusConflict)
			return true
		case ErrInvalidGrade:
			http.Error(w, e.Error(), http.StatusForbidden)
			return true
		default:
			log.Printf("%s: %s\n", msg, e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return true
		}
	}
	return false
}

func SetDeadline(srv ReportService, w http.ResponseWriter, r *http.Request) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["stu"]
	var t time.Time
	err := handler.JsonRequest(w, r, &t)
	if err != nil {
		return
	}
	err = srv.SetDeadline(kind, student, t)
	reportError(w, "Unable to update the deadline of '"+kind+"' for "+student, err)
}

func SetMark(srv ReportService, w http.ResponseWriter, r *http.Request) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["stu"]
	var d int	 
	if handler.JsonRequest(w, r, &d) != nil {
		log.Println("bad request")
		return
	}
	err := srv.SetGrade(kind, student, d)
	reportError(w, "Unable to upgrate '"+kind+"' grade for "+student, err)
}

func List(srv ReportService, w http.ResponseWriter, r *http.Request) {
	stu := mux.Vars(r)["stu"]
	reports, err := srv.List(stu)
	if reportError(w, "Unable to list reports for '"+stu+"'", err) {
		return
	}
	handler.JsonReply(w, false, reports)
}

func Content(us user.UserService, srv ReportService, w http.ResponseWriter, r *http.Request) {
	kind := mux.Vars(r)["kind"]
	student := mux.Vars(r)["stu"]
	cnt, err := srv.Content(kind, student)
	if reportError(w, "Unable to get '"+kind+"' report metadata for "+student, err) {
		return
	}
	u, _ := us.Get(student)
	handler.FileReply(w, "application/pdf", u.Fullname()+" - "+kind+".pdf", cnt)
}

func SetContent(srv ReportService, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/pdf" {
		http.Error(w, "PDF expected", http.StatusUnsupportedMediaType)
		return
	}
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["stu"]
	d, err := handler.RequiredContent(w, r)
	if reportError(w, "Unable to upload '"+kind+"' report content for "+student, err) {
		return
	}
	err = srv.SetContent(kind, student, d)
	if reportError(w, "Unable to store '"+kind+"' report content for "+student, err) {
		return
	}
	w.Write([]byte("1"))
}
