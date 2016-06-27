package httpd

import (
	"net/http"
	"os"

	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/session"
)

func status(not *notifier.Notifier, w http.ResponseWriter, r *http.Request, e error) {
	switch e {
	case os.ErrNotExist, schema.ErrUnknownConvention, schema.ErrUnknownStudent, schema.ErrUnknownSurvey, schema.ErrUnknownUser, schema.ErrUnknownReport, schema.ErrUnknownInternship, schema.ErrNoPendingRequests, schema.ErrUnknownDefense:
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	case schema.ErrDefenseConflict, schema.ErrReportExists, schema.ErrUserExists, schema.ErrInternshipExists, schema.ErrUserTutoring, schema.ErrConventionExists, schema.ErrDefenseJuryConflict, schema.ErrDefenseExists:
		http.Error(w, e.Error(), http.StatusConflict)
		return
	case schema.ErrCredentials, schema.ErrInvalidToken, schema.ErrSessionExpired, ErrNotActivatedAccount:
		http.Error(w, e.Error(), http.StatusUnauthorized)
		return
	case feeder.ErrAuthorization:
		http.Error(w, "Unable to fetch the conventions from the server: "+e.Error(), http.StatusUnauthorized)
		return
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
		logger.Log("event", "httpd", "unsupported error ", e)
	}
}
