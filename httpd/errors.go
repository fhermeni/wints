package httpd

import (
	"log"
	"net/http"

	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/session"
)

func status(w http.ResponseWriter, e error) {
	switch e {
	case schema.ErrInvalidToken, schema.ErrSessionExpired:
		return
	case schema.ErrUnknownSurvey, schema.ErrUnknownUser, schema.ErrUnknownReport, schema.ErrUnknownInternship, schema.ErrNoPendingRequests:
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	case schema.ErrReportExists, schema.ErrUserExists, schema.ErrInternshipExists, schema.ErrUserTutoring:
		http.Error(w, e.Error(), http.StatusConflict)
		return
	case schema.ErrCredentials:
		http.Error(w, e.Error(), http.StatusUnauthorized)
	case ErrMalformedJSON, schema.ErrInvalidSurvey, schema.ErrInvalidAlumniEmail, schema.ErrDeadlinePassed, schema.ErrInvalidGrade, schema.ErrInvalidMajor, schema.ErrInvalidPeriod, schema.ErrGradedReport, schema.ErrPasswordTooShort:
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	case session.ErrPermission, session.ErrConfidentialReport:
		http.Error(w, e.Error(), http.StatusForbidden)
		return
	case nil:
		return
	default:
		http.Error(w, "Internal server error. A possible bug to report", http.StatusInternalServerError)
		log.Printf("Unsupported error: %s\n", e.Error())
	}
}
