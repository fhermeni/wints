package sqlstore

import (
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	selectReport      = "select student, kind, deadline, delivery, reviewed, grade, comment, private, toGrade from reports where student=$1 and kind=$2"
	selectReports     = "select student, kind, deadline, delivery, reviewed, grade, comment, private, toGrade from reports where student=$1 order by deadline asc"
	selectAllReports  = "select student, kind, deadline, delivery, reviewed, grade, comment, private, toGrade from reports order by deadline asc"
	selectReportCnt   = "select cnt from reports where student=$1 and kind=$2"
	setReportCnt      = "update reports set cnt=$3, delivery=$4 where student=$1 and kind=$2"
	setGrade          = "update reports set grade=$3, comment=$4,reviewed=$5 where student=$1 and kind=$2"
	setReportDeadline = "update reports set deadline=$3 where student=$1 and kind=$2"
	setReportPrivacy  = "update reports set private=$3 where student=$1 and kind=$2"
)

//Reports returns all the reports for a given student
func (s *Store) Reports(email string) ([]schema.ReportHeader, error) {
	res := make([]schema.ReportHeader, len(s.config.Reports))
	st := s.stmt(selectReports)
	rows, err := st.Query(email)
	if err != nil {
		return res, err
	}
	defer rows.Next()
	for rows.Next() {
		var hdr schema.ReportHeader
		if _, hdr, err = scanReport(rows); err != nil {
			return res, err
		}
		//place at the right offset
		for idx, r := range s.config.Reports {
			if r.Kind == hdr.Kind {
				res[idx] = hdr
				break
			}
		}
	}
	return res, err
}

//Reports returns all the reports for a given student
func (s *Store) allReports() (map[string][]schema.ReportHeader, error) {
	res := make(map[string][]schema.ReportHeader)
	st := s.stmt(selectAllReports)
	rows, err := st.Query()
	if err != nil {
		return res, err
	}
	defer rows.Next()
	for rows.Next() {
		var hdr schema.ReportHeader
		stu, hdr, err := scanReport(rows)
		if err != nil {
			return res, err
		}
		reports, ok := res[stu]
		if !ok {
			reports = []schema.ReportHeader{}
		}
		reports = append(reports, hdr)
		res[stu] = reports
	}
	return res, err
}

func scanReport(rows *sql.Rows) (string, schema.ReportHeader, error) {
	hdr := schema.ReportHeader{Grade: -1}
	var comment sql.NullString
	var delivery, reviewed pq.NullTime
	var grade sql.NullInt64
	var student string
	err := rows.Scan(
		&student,
		&hdr.Kind,
		&hdr.Deadline,
		&delivery,
		&reviewed,
		&grade,
		&comment,
		&hdr.Private,
		&hdr.ToGrade)
	if err != nil {
		return "", hdr, mapCstrToError(err)
	}
	if comment.Valid {
		hdr.Comment = comment.String
	}
	hdr.Deadline = hdr.Deadline.Truncate(time.Minute).UTC()
	hdr.Delivery = nullableTime(delivery)
	hdr.Reviewed = nullableTime(reviewed)
	if grade.Valid {
		hdr.Grade = int(grade.Int64)
	}
	return student, hdr, mapCstrToError(err)
}

//Report returns the given report
func (s *Store) Report(k, email string) (schema.ReportHeader, error) {
	hdr := schema.ReportHeader{Grade: -1}

	st := s.stmt(selectReport)
	rows, err := st.Query(email, k)
	if err != nil {
		return hdr, mapCstrToError(err)
	}
	defer rows.Close()
	if !rows.Next() {
		return hdr, schema.ErrUnknownReport
	}
	_, r, e := scanReport(rows)
	return r, e
}

//ReportContent returns the content of a given report
func (s *Store) ReportContent(kind, email string) ([]byte, error) {
	var cnt []byte

	st := s.stmt(selectReportCnt)
	err := st.QueryRow(email, kind).Scan(&cnt)
	if err != nil {
		return []byte{}, noRowsTo(err, schema.ErrUnknownReport)
	}
	res, err := base64.StdEncoding.DecodeString(string(cnt))
	return res, err
}

//SetReportContent saves the content of a given report if the report has not been already graded
//If the deadline passed, it is however possible to store the report once.
func (s *Store) SetReportContent(kind, email string, cnt []byte) (time.Time, error) {
	//Check if it has already been graded
	now := time.Now().Truncate(time.Minute).UTC()
	r, err := s.Report(kind, email)
	if err != nil {
		return now, err
	}
	if r.Reviewed != nil {
		//No re-submission once graded
		return now, schema.ErrGradedReport
	}
	if r.Delivery != nil && r.Deadline.Before(now) {
		//deadline passed and already submitted. no other chance
		return now, schema.ErrDeadlinePassed
	}
	res := base64.StdEncoding.EncodeToString(cnt)
	return now, s.singleUpdate(setReportCnt, schema.ErrUnknownReport, email, kind, res, now)
}

//SetReportGrade stores the given report grade
func (s *Store) SetReportGrade(kind, email string, g int, comment string) (time.Time, error) {
	if g < 0 || g > 20 {
		return time.Time{}, schema.ErrInvalidGrade
	}
	now := time.Now().Truncate(time.Minute).UTC()
	return now, s.singleUpdate(setGrade, schema.ErrUnknownReport, email, kind, g, comment, now)
}

//SetReportDeadline change a report deadline
func (s *Store) SetReportDeadline(kind, email string, t time.Time) error {
	t = t.Truncate(time.Minute).UTC()
	return s.singleUpdate(setReportDeadline, schema.ErrUnknownReport, email, kind, t)
}

//SetReportPrivacy changes the privacy level of a report
func (s *Store) SetReportPrivacy(kind, email string, p bool) error {
	return s.singleUpdate(setReportPrivacy, schema.ErrUnknownReport, email, kind, p)
}
