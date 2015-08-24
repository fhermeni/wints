package sqlstore

import (
	"database/sql"
	"encoding/base64"
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

func (s *Service) Reports(email string) (map[string]internship.ReportHeader, error) {
	res := make(map[string]internship.ReportHeader)
	st, err := s.stmt(selectReports)
	if err != nil {
		return res, nil
	}
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
	if delivery.Valid {
		hdr.Delivery = &delivery.Time
	}
	if reviewed.Valid {
		hdr.Reviewed = &reviewed.Time
	}
	if grade.Valid {
		hdr.Grade = int(grade.Int64)
	}
	return hdr, mapCstrToError(err)
}

func (s *Service) Report(k, email string) (internship.ReportHeader, error) {
	hdr := internship.ReportHeader{Grade: -1}

	st, err := s.stmt(selectReport)
	if err != nil {
		return hdr, mapCstrToError(err)
	}
	rows, err := st.Query(email, k)
	if err != nil {
		return hdr, mapCstrToError(err)
	}
	defer rows.Close()
	return scanReport(rows)
}

func (s *Service) ReportContent(kind, email string) ([]byte, error) {
	var cnt []byte
	var delivery pq.NullTime
	st, err := s.stmt(selectReportCnt)
	if err != nil {
		return cnt, err
	}
	if err := st.QueryRow(email, kind).Scan(&cnt, &delivery); err != nil {
		return []byte{}, mapCstrToError(err)
	}
	if !delivery.Valid {
		return []byte{}, mapCstrToError(err)
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func (s *Service) SetReportContent(kind, email string, cnt []byte) error {
	var deadline time.Time
	var delivery, reviewed pq.NullTime
	tx := newTxErr(s.DB)
	tx.err = tx.tx.QueryRow(selectReport, email, kind).Scan(&deadline, &delivery, &reviewed)
	tx.err = noRowsTo(tx.err, internship.ErrUnknownReport)
	if tx.err != nil {
		return tx.Done()
	}
	if time.Now().After(deadline) {
		//We allow 1 delivery once the deadline passed
		if delivery.Valid {
			//You had your chance
			tx.err = internship.ErrDeadlinePassed
			return tx.Done()
		}
	}
	if reviewed.Valid {
		tx.err = internship.ErrGradedReport
	}
	tx.Update(setReportCnt, email, kind, base64.StdEncoding.EncodeToString(cnt), time.Now())
	//No need to check if nb == 1, ensured by the transaction context
	return tx.Done()

}

func (s *Service) SetReportGrade(kind, email string, g int, comment string) error {
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	return s.singleUpdate(setGrade, internship.ErrUnknownReport, email, kind, g, comment, time.Now())
}

func (s *Service) SetReportDeadline(kind, email string, t time.Time) error {
	return s.singleUpdate(setReportDeadline, internship.ErrUnknownReport, email, kind, t)
}

func (s *Service) SetReportPrivacy(kind, email string, p bool) error {
	return s.singleUpdate(setReportPrivacy, internship.ErrUnknownReport, email, kind, p)
}
