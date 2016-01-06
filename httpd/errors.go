package httpd

import (
	"net/http"

	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/session"
)

func status(not *notifier.Notifier, w http.ResponseWriter, r *http.Request, e error) {
	switch e {
	case schema.ErrUnknownConvention, schema.ErrUnknownStudent, schema.ErrUnknownSurvey, schema.ErrUnknownUser, schema.ErrUnknownReport, schema.ErrUnknownInternship, schema.ErrNoPendingRequests:
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	case schema.ErrReportExists, schema.ErrUserExists, schema.ErrInternshipExists, schema.ErrUserTutoring:
		http.Error(w, e.Error(), http.StatusConflict)
		return
	case schema.ErrCredentials, schema.ErrInvalidToken, schema.ErrSessionExpired:
		http.Error(w, e.Error(), http.StatusUnauthorized)
	case ErrMalformedJSON, schema.ErrInvalidSurvey, schema.ErrInvalidAlumniEmail, schema.ErrInvalidEmail, schema.ErrDeadlinePassed, schema.ErrInvalidGrade, schema.ErrInvalidMajor, schema.ErrInvalidPromotion, schema.ErrInvalidPeriod, schema.ErrGradedReport, schema.ErrPasswordTooShort:
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	case session.ErrPermission:
		http.Error(w, e.Error(), http.StatusForbidden)
		return
	case feeder.ErrTimeout:
		http.Error(w, e.Error(), http.StatusRequestTimeout)
		return
	case nil:
		return
	default:
		http.Error(w, "Internal server error. A possible bug to report", http.StatusInternalServerError)
		not.Log.Log("wintsd", "unsupported error ", e)
	}
}
