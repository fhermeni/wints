package sqlstore

import (
	"crypto/rand"
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

func randomBytes(s int) []byte {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, s)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func nullableTime(t pq.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func noRowsTo(got error, with error) error {
	if got == sql.ErrNoRows {
		return with
	}
	return got
}
func violationAsErr(got error, cstr string, with error) error {
	if got == nil {
		return got
	}
	if e, ok := got.(*pq.Error); ok {
		if e.Constraint == cstr {
			return with
		}
	}
	return got
}

func mapCstrToError(err error) error {
	if err == nil {
		return err
	}
	if e, ok := err.(*pq.Error); ok {
		switch e.Constraint {
		case "pk_email":
			return internship.ErrUserExists
		case "fk_sessions_email", "fk_student_email", "fk_password_renewal_email":
			return internship.ErrUnknownUser
		case "pk_students_email":
			return internship.ErrStudentExists
		case "pk_conventions_student":
			return internship.ErrConventionExists
		case "fk_conventions_student", "fk_reports_student", "fk_surveys_student", "fk_defenses_student":
			return internship.ErrUnknownStudent
		case "fk_conventions_tutor":
			return internship.ErrUserTutoring
		}
	}
	return err
}
