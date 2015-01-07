package datastore

import (
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/fhermeni/wints/internship"
)

func (s *Service) PlanReport(email string, r internship.ReportHeader) error {
	sql := "insert into reports(student, kind, deadline, grade, private) values($1,$2,$3)"
	return SingleUpdate(s.DB, internship.ErrReportExists, sql, email, r.Kind, r.Deadline, -2, false)
}

func (s *Service) ReportDefs() []internship.ReportDef {
	return s.reportDefs
}
func (srv *Service) Report(k, email string) (internship.ReportHeader, error) {
	q := "select deadline, grade, comment, private, toGrade from reports where student=$1 and kind=$2"
	var d time.Time
	var g int
	var priv, toGrade bool
	var comment sql.NullString
	if err := srv.DB.QueryRow(q, email, k).Scan(&d, &g, &comment, &priv, &toGrade); err != nil {
		return internship.ReportHeader{}, internship.ErrUnknownReport
	}
	hdr := internship.ReportHeader{Kind: k, Deadline: d, Grade: g, Private: priv, ToGrade: toGrade}
	if comment.Valid {
		hdr.Comment = comment.String
	}
	return hdr, nil
}

func (s *Service) ReportContent(kind, email string) ([]byte, error) {
	sql := "select cnt from reports where student=$1 and kind=$2"
	var cnt []byte
	if err := s.DB.QueryRow(sql, email, kind).Scan(&cnt); err != nil {
		return []byte{}, internship.ErrUnknownReport
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func (srv *Service) SetReportContent(kind, email string, cnt []byte) error {
	sql := "update reports set grade=-1 and cnt=$3 where student=$1 and kind=$2"
	enc := base64.StdEncoding
	return SingleUpdate(srv.DB, internship.ErrUnknownReport, sql, email, kind, enc.EncodeToString(cnt))

}

func (s *Service) SetReportGrade(kind, email string, g int, comment string) error {
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	sql := "update reports set grade=$3 comment=$4 where student=$1 and kind=$2"
	return SingleUpdate(s.DB, internship.ErrUnknownReport, sql, email, kind, g, comment)
}

func (s *Service) SetReportDeadline(kind, email string, t time.Time) error {
	sql := "update reports set deadline=$3 where student=$1 and kind=$2"
	return SingleUpdate(s.DB, internship.ErrUnknownReport, sql, email, kind, t)
}

func (s *Service) SetReportPrivate(kind, email string, p bool) error {
	sql := "update reports set private=$3 where student=$1 and kind=$2"
	return SingleUpdate(s.DB, internship.ErrUnknownReport, sql, email, kind, p)
}
