package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/schema"
	"github.com/lib/pq"
)

var (
	selectReport      = "select kind, deadline, delivery, reviewed, grade, comment, private, toGrade from reports where student=$1 and kind=$2"
	selectReports     = "select kind, deadline, delivery, reviewed, grade, comment, private, toGrade from reports where student=$1"
	selectReportCnt   = "select delivery, cnt from reports where student=$1 and kind=$2"
	setReportCnt      = "update reports set cnt=$3, delivery=$4 where student=$1 and kind=$2"
	setGrade          = "update reports set grade=$3, comment=$4,reviewed=$5 where student=$1 and kind=$2"
	setReportDeadline = "update reports set deadline=$3 where student=$1 and kind=$2"
	setReportPrivacy  = "update reports set private=$3 where student=$1 and kind=$2"
)

//Reports returns all the reports for a given student
func (s *Store) Reports(email string) (map[string]schema.ReportHeader, error) {
	res := make(map[string]schema.ReportHeader)
	st := s.stmt(selectReports)
	rows, err := st.Query(email)
	if err != nil {
		return res, nil
	}
	for rows.Next() {
		var hdr schema.ReportHeader
		if hdr, err = scanReport(rows); err != nil {
			return res, nil
		}
		res[hdr.Kind] = hdr
	}
	return res, nil

}

func scanReport(rows *sql.Rows) (schema.ReportHeader, error) {
	hdr := schema.ReportHeader{Grade: -1}
	var comment sql.NullString
	var delivery, reviewed pq.NullTime
	var grade sql.NullInt64
	err := rows.Scan(
		&hdr.Deadline,
		&delivery,
		&reviewed,
		&grade,
		&comment,
		&hdr.Private,
		&hdr.ToGrade)
	if err != nil {
		return hdr, mapCstrToError(err)
	}
	if comment.Valid {
		hdr.Comment = comment.String
	}
	hdr.Delivery = nullableTime(delivery)
	hdr.Reviewed = nullableTime(reviewed)
	if grade.Valid {
		hdr.Grade = int(grade.Int64)
	}
	return hdr, mapCstrToError(err)
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
		return hdr, schema.ErrUnknownStudent
	}
	return scanReport(rows)
}

//ReportContent returns the content of a given report
func (s *Store) ReportContent(kind, email string) ([]byte, error) {
	var cnt []byte
	st := s.stmt(selectReportCnt)
	err := st.QueryRow(email, kind).Scan(&cnt)
	return cnt, noRowsTo(err, schema.ErrUnknownReport)
}

//SetReportContent saves the content of a given report
func (s *Store) SetReportContent(kind, email string, cnt []byte) error {
	return s.singleUpdate(setReportCnt, schema.ErrUnknownReport, email, kind, cnt, time.Now())
}

//SetReportGrade stores the given report grade
func (s *Store) SetReportGrade(kind, email string, g int, comment string) error {
	return s.singleUpdate(setGrade, schema.ErrUnknownReport, email, kind, g, comment, time.Now())
}

//SetReportDeadline change a report deadline
func (s *Store) SetReportDeadline(kind, email string, t time.Time) error {
	return s.singleUpdate(setReportDeadline, schema.ErrUnknownReport, email, kind, t)
}

//SetReportPrivacy changes the privacy level of a report
func (s *Store) SetReportPrivacy(kind, email string, p bool) error {
	return s.singleUpdate(setReportPrivacy, schema.ErrUnknownReport, email, kind, p)
}
