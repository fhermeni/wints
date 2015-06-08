package datastore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
)

var JuriesStmt *sql.Stmt
var DefStmt *sql.Stmt

func (srv *Service) scanDefense(rows *sql.Rows) (internship.Defense, error) {
	var room, student string
	var date time.Time
	var remote, private bool
	var grade sql.NullInt64
	err := rows.Scan(&student, &date, &room, &grade, &private, &remote)
	if err != nil {
		return internship.Defense{}, err
	}
	def := internship.Defense{
		Student: student,
		Date:    date,
		Room:    room,
		Private: private,
		Remote:  remote,
	}
	juries, err := srv.jury(student)
	if err != nil {
		return def, err
	}
	def.Juries = juries
	if grade.Valid {
		def.Grade = int(grade.Int64)
	}
	return def, nil
}

//DefenseSessions returns all the registered defense sessions
func (srv *Service) Defenses() ([]internship.Defense, error) {
	//all the defenses
	query := "select student, date, room, grade, private, remote from defenses  order by date"
	defs := make([]internship.Defense, 0, 0)
	rows, err := srv.DB.Query(query)
	if err != nil {
		return defs, err
	}
	defer rows.Close()
	for rows.Next() {
		var room, student string
		var date time.Time
		var remote, private bool
		var grade sql.NullInt64
		err = rows.Scan(&student, &date, &room, &grade, &private, &remote)
		def, err := srv.scanDefense(rows)
		if err != nil {
			return defs, err
		}
		defs = append(defs, def)
	}
	return defs, nil
}

//DefenseSessions returns all the registered defense sessions
func (srv *Service) Defense(student string) (internship.Defense, error) {
	if DefStmt == nil {
		var err error
		DefStmt, err = srv.DB.Prepare("select student, date, room, grade, private, remote from defenses where student=$1")
		if err != nil {
			return internship.Defense{}, err
		}
	}
	//all the defenses
	rows, err := DefStmt.Query(student)
	if err != nil {
		return internship.Defense{}, err
	}
	if rows.Next() {
		defer rows.Close()
		return srv.scanDefense(rows)
	}
	return internship.Defense{}, nil
}

func (srv *Service) jury(student string) ([]internship.User, error) {
	if JuriesStmt == nil {
		var err error
		JuriesStmt, err = srv.DB.Prepare("select firstname, lastname, users.email, tel from users, juries where juries.student=$1 and users.email = juries.jury")
		if err != nil {
			return []internship.User{}, err
		}
	}

	rows, err := JuriesStmt.Query(student)
	if err != nil {
		return []internship.User{}, err
	}
	defer rows.Close()
	juries := make([]internship.User, 0, 0)
	for rows.Next() {
		var fn, ln, email, tel string
		err = rows.Scan(&fn, &ln, &email, &tel)
		if err != nil {
			return []internship.User{}, err
		}
		j := internship.User{Firstname: fn, Lastname: ln, Tel: tel, Email: email}
		juries = append(juries, j)
	}
	return juries, nil
}

func (srv *Service) SetDefenseGrade(stu string, g int) error {
	sql := "update defenses set grade=$2 where student=$1"
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	return SingleUpdate(srv.DB, internship.ErrUnknownUser, sql, stu, g)
}

func (srv *Service) SetDefenses(defs []internship.Defense) error {
	tx, err := srv.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from defenses")
	if err != nil {
		return rollback(err, tx)
	}
	for _, def := range defs {
		_, err = tx.Exec("insert into defenses(student,private,remote,date,room,grade) values($1,$2,$3,$4,$5,$6)", def.Student, def.Private, def.Remote, def.Date, def.Room, def.Grade)
		if err != nil {
			return rollback(err, tx)
		}
		for _, j := range def.Juries {
			_, err = tx.Exec("insert into juries(student,jury) values($1,$2)", def.Student, j.Email)
			if err != nil {
				return rollback(err, tx)
			}
		}
	}
	return tx.Commit()
}
