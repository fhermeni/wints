package report

import (
	"database/sql"
	"encoding/base64"
	"github.com/fhermeni/wints/db"
	"time"
)

type DbService struct {
	db *sql.DB
}

func (s *DbService) New(kind, email string, date time.Time) (MetaData, error) {
	sql := "insert into reports(student, kind, deadline) values($1,$2,$3)"
	if err := db.SingleUpdate(s.db, ErrExists, sql, email, kind, date); err != nil {
		return MetaData{}, err
	}
	return MetaData{kind, email, date, -1}, nil
}

func (srv *DbService) Get(kind, email string) (MetaData, error) {
	q := "select deadline, grade from reports where student=$1 and kind=$2"
	var d time.Time
	var g sql.NullInt64

	if err := srv.db.QueryRow(q, email, kind).Scan(&d, &g); err != nil {
		return MetaData{}, err
	}
	if !g.Valid {
		g.Int64 = -1
	}
	return MetaData{kind, email, d, int(g.Int64)}, nil
}

func (s *DbService) Content(kind, email string) ([]byte, error) {
	sql := "select cnt from reports where student=$1 and kind=$2"
	var cnt []byte
	if err := s.db.QueryRow(sql, email, kind).Scan(&cnt); err != nil {
		return []byte{}, err
	}
	return base64.StdEncoding.DecodeString(string(cnt))
}

func (srv *DbService) SetContent(kind, email string, cnt []byte) error {
	sql := "update reports set cnt=$3 where student=$1 and kind=$2"
	enc := base64.StdEncoding
	return db.SingleUpdate(srv.db, ErrUnknown, sql, email, kind, enc.EncodeToString(cnt))

}

func (s *DbService) SetGrade(kind, email string, g int) error {
	if g < 0 || g > 20 {
		return ErrInvalidGrade
	}
	sql := "update reports set grade=$3 where student=$1 and kind=$2"
	return db.SingleUpdate(s.db, ErrUnknown, sql, email, kind, g)
}

func (s *DbService) SetDeadline(kind, email string, t time.Time) error {
	sql := "update reports set deadline=$3 where student=$1 and kind=$2"
	return db.SingleUpdate(s.db, ErrUnknown, sql, email, kind, t)
}

func (s *DbService) List(email string) ([]MetaData, error) {
	q := "select deadline, grade, kind from reports where student=$1"
	rows, err := s.db.Query(q, email)
	reports := make([]MetaData, 0, 0)
	if err != nil {
		return reports, err
	}
	defer rows.Close()
	var date time.Time
	var grade sql.NullInt64
	var kind string

	for rows.Next() {
		rows.Scan(&date, &grade, &kind)
		if !grade.Valid {
			grade.Int64 = -1
		}
		m := MetaData{kind, email, date, int(grade.Int64)}
		reports = append(reports, m)
	}
	return reports, nil
}
