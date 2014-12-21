package datastore

import (
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/fhermeni/wints/internship"
)

func (s *Service) PlanReport(k, email string, d time.Time) (internship.ReportHeader, error) {
	sql := "insert into reports(student, kind, deadline) values($1,$2,$3)"
	if err := SingleUpdate(s.Db, internship.ErrReportExists, sql, email, k, d); err != nil {
		return internship.ReportHeader{}, err
	}
	return internship.ReportHeader{Kind: k, Deadline: d, Grade: -1}, nil
}

func (srv *Service) Report(k, email string) (internship.ReportHeader, error) {
	q := "select deadline, grade from reports where student=$1 and kind=$2"
	var d time.Time
	var g sql.NullInt64

	if err := srv.Db.QueryRow(q, email, k).Scan(&d, &g); err != nil {
		return internship.ReportHeader{}, err
	}
	if !g.Valid {
		g.Int64 = -1
	}
	return internship.ReportHeader{Kind: k, Deadline: d, Grade: int(g.Int64)}, nil
}

func (s *Service) ReportContent(kind, email string) ([]byte, error) {
	sql := "select cnt from reports where student=$1 and kind=$2"
	var cnt []byte
	if err := s.Db.QueryRow(sql, email, kind).Scan(&cnt); err != nil {
		return []byte{}, err
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func (srv *Service) SetReportContent(kind, email string, cnt []byte) error {
	sql := "update reports set cnt=$3 where student=$1 and kind=$2"
	enc := base64.StdEncoding
	return SingleUpdate(srv.Db, internship.ErrUnknownReport, sql, email, kind, enc.EncodeToString(cnt))

}

func (s *Service) SetReportGrade(kind, email string, g int) error {
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	sql := "update reports set grade=$3 where student=$1 and kind=$2"
	return SingleUpdate(s.Db, internship.ErrUnknownReport, sql, email, kind, g)
}

func (s *Service) SetReportDeadline(kind, email string, t time.Time) error {
	sql := "update reports set deadline=$3 where student=$1 and kind=$2"
	return SingleUpdate(s.Db, internship.ErrUnknownReport, sql, email, kind, t)
}