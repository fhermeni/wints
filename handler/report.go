package handler;

import (		
	"github.com/fhermeni/wints/report"	
	"github.com/fhermeni/wints/user"	
	"net/http"
	"github.com/gorilla/mux"	
	"log"	
	"strconv"			
	"time"
	"strings"	
	"bytes"
	"archive/tar"
	"compress/gzip"
)

func reportError(w http.ResponseWriter, msg string, e error) bool {		
	if e != nil {
		switch (e) {
		case report.ErrUnknown:
			http.Error(w, e.Error(), http.StatusNotFound)
			return true
		case report.ErrExists:
			http.Error(w, e.Error(), http.StatusConflict)
			return true			
		case report.ErrInvalidGrade:
			http.Error(w, e.Error(), http.StatusForbidden)
			return true
		default:
			log.Printf("%s: %s\n", msg, e.Error());
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return true
		}
	}
	return false
}

func GetReport(us user.UserService, srv report.ReportService, w http.ResponseWriter, r *http.Request)  {
	kind := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	cnt, err := srv.Content(kind, student)
	if reportError(w, "Unable to get '" + kind + "' report metadata for " + student, err) {
		return
	}
	u, _ := us.Get(student)
	fileReply(w, "application/pdf", u.Lastname + "-" + kind + ".pdf", cnt)
}

func GetReports(us user.UserService, srv report.ReportService, w http.ResponseWriter, r *http.Request) {
	kind, _ := mux.Vars(r)["kind"]
	tmp := r.FormValue("students")
	if len(tmp) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	students := strings.Split(tmp, ",")

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	var missing string
	for _, stu := range students {
		c, err := us.Get(stu)		
		if err != nil {
			http.Error(w, "Unknown student '" + stu + "'", http.StatusBadRequest)
			return			
		}
		report, err := srv.Content(kind, stu)
		if err != nil {
			missing = missing + c.Fullname() + "\n";
			continue;			
		}
		hdr := &tar.Header{
			Name: c.Fullname() + ".pdf",
			Mode: 0644,
			Size: int64(len(report))}
		if err := tw.WriteHeader(hdr); err != nil {
			reportError(w, "Unable to archive " + stu + " " + kind + " header", err)
			return
		}
		if _, err := tw.Write(report); err != nil {
			reportError(w, "Unable to archive " + stu + " " + kind + " content", err)
			return
		}
	}
	if len(missing) > 0 {
		hdr := &tar.Header{
			Name: "missing_" + kind + "_reports.txt",
			Mode: 0644,
			Size: int64(len(missing))}
		if err := tw.WriteHeader(hdr); err != nil {
			reportError(w, "Unable to archive the missing header", err)
			return
		}
		if _, err := tw.Write([]byte(missing)); err != nil {
			reportError(w, "Unable to archive the missing content", err)
			return
		}
	}
	if err := tw.Close(); err != nil {
		reportError(w, "Unable to flush the archive", err)
		return
	}		
	w2 := gzip.NewWriter(w)
	w2.Write(buf.Bytes())
	w2.Close()
	//fileReply(w, "application/x-tar", "reports-" + kind + ".tar.gz", b.Bytes())
}

func UploadReport(srv report.ReportService, w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	d, err := requiredContent(w, r)
	if reportError(w, "Unable to upload '" + kind + "' report content for " + student, err) {
		return
	}
	err = srv.SetContent(student, kind, d)
	if reportError(w, "Unable to store '" + kind + "' report content for " + student, err) {
		return
	}
	w.Write([]byte("1"))
}


func UpdateDeadline(srv report.ReportService, w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	d, err := requiredContent(w, r)
	if reportError(w, "Unable to read the deadline", err) {
		return
	}
	t, err := time.Parse("2/1/2006", string(d))
	if reportError(w, "Invalid date format", err) {
		return
	}

	err = srv.SetDeadline(student, kind, t)
	reportError(w, "Unable to update the deadline of '" + kind + "' for " + student, err);
}

func UpdateMark(srv report.ReportService, w http.ResponseWriter, r *http.Request, email string) {
	kind, _ := mux.Vars(r)["kind"]
	student := mux.Vars(r)["student"]
	d, err := requiredContent(w, r)
	if reportError(w, "Unable to read the mark", err) {
		return
	}
	i, err := strconv.Atoi(string(d))
	if err != nil {
		http.Error(w, "Invalid grade", http.StatusForbidden)
		return
	}
	err = srv.SetGrade(student, kind, i)
	reportError(w, "Unable to upgrate '" + kind + "' grade for " + student, err);
}