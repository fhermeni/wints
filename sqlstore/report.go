package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
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

func (s *Store) Reports(email string) (map[string]internship.ReportHeader, error) {
	res := make(map[string]internship.ReportHeader)
	st := s.stmt(selectReports)
	rows, err := st.Query(email)
	if err != nil {
		return res, nil
	}
	for rows.Next() {
		if hdr, err := scanReport(rows); err != nil {
			return res, nil
		} else {
			res[hdr.Kind] = hdr
		}
	}
	return res, nil

}

func scanReport(rows *sql.Rows) (internship.ReportHeader, error) {
	hdr := internship.ReportHeader{Grade: -1}
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

func (s *Store) Report(k, email string) (internship.ReportHeader, error) {
	hdr := internship.ReportHeader{Grade: -1}

	st := s.stmt(selectReport)
	rows, err := st.Query(email, k)
	if err != nil {
		return hdr, mapCstrToError(err)
	}
	defer rows.Close()
	return scanReport(rows)
}

func (s *Store) ReportContent(kind, email string) ([]byte, error) {
	var cnt []byte
	st := s.stmt(selectReportCnt)
	err := st.QueryRow(email, kind).Scan(&cnt)
	return cnt, err
}

func (s *Store) SetReportContent(kind, email string, cnt []byte) error {
	return s.singleUpdate(setReportCnt, internship.ErrUnknownReport, email, kind, cnt, time.Now())
}

func (s *Store) SetReportGrade(kind, email string, g int, comment string) error {
	return s.singleUpdate(setGrade, internship.ErrUnknownReport, email, kind, g, comment, time.Now())
}

func (s *Store) SetReportDeadline(kind, email string, t time.Time) error {
	return s.singleUpdate(setReportDeadline, internship.ErrUnknownReport, email, kind, t)
}

func (s *Store) SetReportPrivacy(kind, email string, p bool) error {
	return s.singleUpdate(setReportPrivacy, internship.ErrUnknownReport, email, kind, p)
}
