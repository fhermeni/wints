package backend

import "database/sql"


func Defense(db *sql.DB, id string) (string, error) {
	var cnt string
	err := db.QueryRow("select content from defenses where id=$1", id).Scan(&cnt)
	return cnt, err
}

func NewDefense(db *sql.DB, id , cnt string) error {
	err := SingleUpdate(db, "insert into defenses(id, content) values($1,$2)", id, cnt)
	return err
}

func LongDefense(db *sql.DB, id string) (string, error) {
	var cnt string
	err := db.QueryRow("select public from defenses where id=$1", id).Scan(&cnt)
	return cnt, err
}

func SaveDefense(db *sql.DB, id, shortVersion, longVersion string) error {
//	log.Printf("%s\n", shortVersion)
	//log.Printf("%s\n", longVersion)
	err := SingleUpdate(db, "update defenses set content=$2,public=$3 where id=$1", id, shortVersion, longVersion)
	return err
	//return nil
}
