package backend

import (
	"time"
	"database/sql"
	"encoding/base64"
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
	err := SingleUpdate(db, sql, s, kind, d, -1)
	if err != nil {
		return ReportMetaData{}, err
	}
	return ReportMetaData{kind, s, d, -1, false}, nil
}

func GetReportMetaData(db *sql.DB, s string, kind string) (ReportMetaData, error) {
	sql := "select deadline, grade, kind from reports where student=$1 and kind=$2"
	var d time.Time
	var g int
	var k string

	err := db.QueryRow(sql, s, kind).Scan(&d, &g, &k)
	if err != nil {
		return ReportMetaData{}, err
	}
	rows, err := db.Query("select * from reports where student=$1 and kind=$2 and cnt is not null", s, kind)
	if err != nil {
		return ReportMetaData{}, err
	}
	defer rows.Close()
	m := ReportMetaData{k, s, d, g, rows.Next()}
 	return m, nil
}

func Report(db *sql.DB, s string, kind string) ([]byte, error) {
	sql := "select cnt from reports where student=$1 and kind=$2"
	var cnt []byte
	err := db.QueryRow(sql, s, kind).Scan(&cnt)
	if err != nil {
		return []byte{}, err
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func SetReport(db *sql.DB, s string, kind string, cnt []byte) error {
	sql := "update reports set cnt=$3 where student=$1 and kind=$2"
	enc := base64.StdEncoding
	err := SingleUpdate(db, sql, s, kind, enc.EncodeToString(cnt))
	return err
}

func UpdateGrade(db *sql.DB, s, kind string, g int) error {
	sql := "update reports set g=$3 where student=$1 and kind=$2"
	err := SingleUpdate(db, sql, s, kind, g)
	return err
}

func UpdateDeadline(db *sql.DB, s, kind string, t time.Time) error {
	sql := "update reports set deadline=$3 where student=$1 and kind=$2"
	err := SingleUpdate(db, sql, s, kind, t)
	return err
}

