package sqlstore

import (
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/fhermeni/wints/internship"
	"github.com/lib/pq"
)

var (
	SelectReport      = "select deadline, delivery, reviewed, grade, comment, private, toGrade from reports where student=$1 and kind=$2"
	SelectReportCnt   = "select delivery, cnt from reports where student=$1 and kind=$2"
	SetReportCnt      = "update reports set cnt=$3, delivery=$4 where student=$1 and kind=$2"
	SetGrade          = "update reports set grade=$3, comment=$4,reviewed=$5 where student=$1 and kind=$2"
	SetReportDeadline = "update reports set deadline=$3 where student=$1 and kind=$2"
	SetReportPrivacy  = "update reports set private=$3 where student=$1 and kind=$2"
)

func (s *Service) Report(k, email string) (internship.ReportHeader, error) {
	hdr := internship.ReportHeader{Grade: -1}

	st, err := s.stmt(SelectReport)
	if err != nil {
		return hdr, err
	}
	var comment sql.NullString
	var delivery, reviewed pq.NullTime
	var grade sql.NullInt64
	if err := st.QueryRow(email, k).Scan(
		&hdr.Deadline,
		&delivery,
		&reviewed,
		&grade,
		&comment,
		&hdr.Private,
		&hdr.ToGrade); err != nil {
		return internship.ReportHeader{}, internship.ErrUnknownReport
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
	return hdr, nil
}

func (s *Service) ReportContent(kind, email string) ([]byte, error) {
	var cnt []byte
	var delivery pq.NullTime
	st, err := s.stmt(SelectReportCnt)
	if err != nil {
		return cnt, err
	}
	if err := st.QueryRow(email, kind).Scan(&cnt, &delivery); err != nil {
		return []byte{}, internship.ErrUnknownReport
	}
	if !delivery.Valid {
		return []byte{}, internship.ErrReportConflict
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func (s *Service) SetReportContent(kind, email string, cnt []byte) error {
	var deadline time.Time
	var delivery, reviewed pq.NullTime
	tx := newTxErr(s.DB)
	tx.err = tx.tx.QueryRow(SelectReport, email, kind).Scan(&deadline, &delivery, &reviewed)
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
	tx.Update(SetReportCnt, email, kind, base64.StdEncoding.EncodeToString(cnt), time.Now())
	//No need to check if nb == 1, ensured by the transaction context
	return tx.Done()

}

func (s *Service) SetReportGrade(kind, email string, g int, comment string) error {
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	return s.singleUpdate(SetGrade, internship.ErrUnknownReport, email, kind, g, comment, time.Now())
}

func (s *Service) SetReportDeadline(kind, email string, t time.Time) error {
	return s.singleUpdate(SetReportDeadline, internship.ErrUnknownReport, email, kind, t)
}

func (s *Service) SetReportPrivacy(kind, email string, p bool) error {
	return s.singleUpdate(SetReportPrivacy, internship.ErrUnknownReport, email, kind, p)
}

/*
func (s *Service) PlanReport(email string, r internship.ReportHeader) error {
	sql := "insert into reports(student, kind, deadline, grade, private) values($1,$2,$3)"
	return SingleUpdate(s.DB, internship.ErrReportExists, sql, email, r.Kind, r.Deadline, -2, false)
}

func (s *Service) ReportDefs() []internship.ReportDef {
	return s.reportDefs
}



*/
