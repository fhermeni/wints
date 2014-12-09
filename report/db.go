package report;

import (
	"time"
	"database/sql"
	"encoding/base64"	
	"github.com/fhermeni/wints/db"
)

type DbService struct {
	db *sql.DB
}

func (s *DbService) New(kind, email string, date time.Time) (MetaData, error) {
	sql := "insert into reports(student, kind, deadline) values($1,$2,$3)"
	if err := db.SingleUpdate(s.db,  ErrExists, sql, email, kind, date); err != nil {
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
 	return MetaData{kind, email, d, int(g.Int64) }, nil
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

/*
func (s *DbService) Tar(us *UserService, kind string, emails [] string) ([]byte, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	var missing string
	for _, student := range emails {
		c, err := user.Get(s.db, student)
		if err != nil {
			return []byte{}, err
		}
		report, err := s.Content(student, kind)
		if err != nil {
			missing = missing + c.Fullname() + "\n";
			continue;			
		}
		hdr := &tar.Header{
			Name: c.Lastname + "-" + kind + ".pdf",
			Mode: 0644,
			Size: int64(len(report))}
		if err := tw.WriteHeader(hdr); err != nil {
			return []byte{}, err
		}
		if _, err := tw.Write(report); err != nil {
			return []byte{}, err
		}
	}
	if len(missing) > 0 {
		hdr := &tar.Header{
			Name: "missing_reports.txt",
			Mode: 0644,
			Size: int64(len(missing))}
		if err := tw.WriteHeader(hdr); err != nil {
			return []byte{}, err
		}
		if _, err := tw.Write([]byte(missing)); err != nil {
			return []byte{}, err
		}
	}
	if err := tw.Close(); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil	
}
*/