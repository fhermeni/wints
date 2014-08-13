package backend

import (
	"time"
	"database/sql"
	"encoding/base64"
	"errors"
)

var (
	errReportExists = errors.New("Report already exists")
	errUnknownReport = errors.New("Unknown report")
	errInvalidGrade = errors.New("The grade must be between 0 and 20 (inclusive)")
)
type ReportMetaData struct {
	Kind string
	Email string
	Deadline time.Time
	Grade int
	IsIn bool
}

func NewReport(db *sql.DB, s, kind string, d time.Time) (ReportMetaData, error) {
	sql := "insert into reports(student, kind, deadline, grade) values($1,$2,$3,$4)"
	if err := SingleUpdate(db,  errReportExists, sql, s, kind, d, -1); err != nil {
		return ReportMetaData{}, err
	}
	return ReportMetaData{kind, s, d, -1, false}, nil
}

func GetReportMetaData(db *sql.DB, s string, kind string) (ReportMetaData, error) {
	sql := "select deadline, grade, kind from reports where student=$1 and kind=$2"
	var d time.Time
	var g int
	var k string

	if err := db.QueryRow(sql, s, kind).Scan(&d, &g, &k); err != nil {
		return ReportMetaData{}, err
	}
	rows, err := db.Query("select * from reports where student=$1 and kind=$2 and cnt is not null", s, kind)
	if err != nil {
		return ReportMetaData{}, err
	}
	defer rows.Close()
 	return ReportMetaData{k, s, d, g, rows.Next()}, nil
}

func Report(db *sql.DB, s string, kind string) ([]byte, error) {
	sql := "select cnt from reports where student=$1 and kind=$2"
	var cnt []byte
	if err := db.QueryRow(sql, s, kind).Scan(&cnt); err != nil {
		return []byte{}, err
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func SetReport(db *sql.DB, s string, kind string, cnt []byte) error {
	sql := "update reports set cnt=$3 where student=$1 and kind=$2"
	enc := base64.StdEncoding
	return SingleUpdate(db, errUnknownReport, sql, s, kind, enc.EncodeToString(cnt))
}

func UpdateGrade(db *sql.DB, s, kind string, g int) error {
	if g < 0 || g > 20 {
		return errInvalidGrade
	}
	sql := "update reports set grade=$3 where student=$1 and kind=$2"
	return SingleUpdate(db, errUnknownReport, sql, s, kind, g)
}

func UpdateDeadline(db *sql.DB, s, kind string, t time.Time) error {
	sql := "update reports set deadline=$3 where student=$1 and kind=$2"
	return SingleUpdate(db, errUnknownReport, sql, s, kind, t)
}

