package backend

import (
	"database/sql"
	"errors"
)

var (
	errUnknownDefense = errors.New("Unknown defense")
	errDefenseExists = errors.New("The defense already exists")
)

func Defense(db *sql.DB, id string) (string, error) {
	var cnt string
	err := db.QueryRow("select content from defenses where id=$1", id).Scan(&cnt)
	return cnt, err
}

func NewDefense(db *sql.DB, id , cnt string) error {
	return SingleUpdate(db, errDefenseExists, "insert into defenses(id, content) values($1,$2)", id, cnt)
}

func LongDefense(db *sql.DB, id string) (string, error) {
	var cnt string
	err := db.QueryRow("select public from defenses where id=$1", id).Scan(&cnt)
	if err != nil {
		return "", errUnknownDefense
	}
	return cnt, err
}

func SaveDefense(db *sql.DB, id, shortVersion, longVersion string) error {
	return SingleUpdate(db, errUnknownDefense, "update defenses set content=$2,public=$3 where id=$1", id, shortVersion, longVersion)
}
