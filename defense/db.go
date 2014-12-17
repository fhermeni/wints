package defense

import (
	"database/sql"
	"github.com/fhermeni/wints/db"
)

type DbService struct {
	db *sql.DB
}

func (s *DbService) Get(id string) (string, error) {
	var cnt string
	err := s.db.QueryRow("select content from defenses where id=$1", id).Scan(&cnt)
	return cnt, err
}

func (s *DbService) New(id, cnt string) error {
	return db.SingleUpdate(s.db, ErrExists, "insert into defenses(id, content) values($1,$2)", id, cnt)
}

func (s *DbService) Long(id string) (string, error) {
	var cnt string
	err := s.db.QueryRow("select public from defenses where id=$1", id).Scan(&cnt)
	if err != nil {
		return "", ErrUnknown
	}
	return cnt, err
}

func (s *DbService) Save(id, short, long string) error {
	return db.SingleUpdate(s.db, ErrUnknown, "update defenses set content=$2,public=$3 where id=$1", id, short, long)
}
