package datastore

import (
	"database/sql"
	"log"

	"github.com/fhermeni/wints/internship"
)

var JuriesStmt *sql.Stmt
var JuryUsersStmt *sql.Stmt
var DefStmt *sql.Stmt
var DefsStmt *sql.Stmt
var SessionsStmt *sql.Stmt
var GradeStmt *sql.Stmt

func (srv *Service) DefenseSession(student string) (internship.DefenseSession, error) {
	var err error
	session := internship.DefenseSession{}
	def := internship.Defense{}
	if DefStmt == nil {
		if DefStmt, err = srv.DB.Prepare("select student,date,room,rank,private,remote from defenses where student=$1"); err != nil {
			return session, err
		}
	}
	rows, err := DefStmt.Query(student)
	if err != nil {
		return session, err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&def.Student, &session.Date, &session.Room, &def.Offset, &def.Private, &def.Remote); err != nil {
			return session, nil
		}
		if err = srv.appendGrade(&def); err != nil {
			return session, err
		}
		session.Defenses = []internship.Defense{def}
	}
	err = srv.appendJuries(&session)
	return session, err
}

//DefenseSessions returns all the registered defense sessions
func (srv *Service) DefenseSessions() ([]internship.DefenseSession, error) {
	sessions := make([]internship.DefenseSession, 0, 0)
	var err error
	if SessionsStmt == nil {
		if SessionsStmt, err = srv.DB.Prepare("select date,room from defenseSessions"); err != nil {
			return sessions, err
		}
	}
	rows, err := SessionsStmt.Query()
	if err != nil {
		return sessions, err
	}
	defer rows.Close()
	for rows.Next() {
		s := internship.DefenseSession{}
		if err = rows.Scan(&s.Date, &s.Room); err != nil {
			return sessions, err
		}
		if err = srv.appendJuries(&s); err != nil {
			return sessions, err
		}
		if err = srv.appendDefenses(&s); err != nil {
			return sessions, err
		}
		sessions = append(sessions, s)
	}
	return sessions, err
}

func (srv *Service) appendDefenses(s *internship.DefenseSession) error {
	s.Defenses = make([]internship.Defense, 0, 0)
	var err error
	if DefsStmt == nil {
		if DefsStmt, err = srv.DB.Prepare("select student,rank,private,remote from defenses where date=$1 and room=$2 order by rank"); err != nil {
			return err
		}
	}
	rows, err := DefsStmt.Query(s.Date, s.Room)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		def := internship.Defense{}
		if err = rows.Scan(&def.Student, &def.Offset, &def.Private, &def.Remote); err != nil {
			return err
		}
		if err = srv.appendGrade(&def); err != nil {
			return err
		}
		s.Defenses = append(s.Defenses, def)
	}
	return err
}

func (srv *Service) appendGrade(d *internship.Defense) error {
	var err error
	if GradeStmt == nil {
		if GradeStmt, err = srv.DB.Prepare("select grade from defenseGrades where student=$1"); err != nil {
			return err
		}
	}
	rows, err := GradeStmt.Query(d.Student)
	if err != nil {
		return err
	}
	if rows.Next() {
		var g sql.NullInt64
		rows.Scan(&g)
		if g.Valid {
			d.Grade = int(g.Int64)
		}
	}
	return nil
}
func (srv *Service) appendJuries(s *internship.DefenseSession) error {
	if JuriesStmt == nil {
		var err error
		JuriesStmt, err = srv.DB.Prepare("select firstname,lastname,tel,email  from users,defenseJuries where date=$1 and room=$2 and users.email = defenseJuries.jury")
		if err != nil {
			return err
		}
	}

	rows, err := JuriesStmt.Query(s.Date, s.Room)
	if err != nil {
		return err
	}
	defer rows.Close()
	s.Juries = make([]internship.User, 0, 0)
	for rows.Next() {
		u := internship.User{}
		if err = rows.Scan(&u.Firstname, &u.Lastname, &u.Tel, &u.Email); err != nil {
			return err
		}
		s.Juries = append(s.Juries, u)
	}
	return nil
}

//SetDefenseSessions saves all the defense sessions
func (srv *Service) SetDefenseSessions(sessions []internship.DefenseSession) error {
	tx, err := srv.DB.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.Exec("delete from defenseSessions"); err != nil {
		return rollback(err, tx)
	}
	for _, s := range sessions {
		if _, err = tx.Exec("insert into defenseSessions(date,room) values($1,$2)", s.Date, s.Room); err != nil {
			return rollback(err, tx)
		}
		for _, j := range s.Juries {
			log.Println("J: " + j.Email)
			if _, err = tx.Exec("insert into defenseJuries(date,room,jury) values($1,$2,$3)", s.Date, s.Room, j.Email); err != nil {
				return rollback(err, tx)
			}
		}
		for _, def := range s.Defenses {
			if _, err = tx.Exec("insert into defenses(date,room,student,remote,private,rank) values($1,$2,$3,$4,$5,$6)", s.Date, s.Room, def.Student, def.Remote, def.Private, def.Offset); err != nil {
				return rollback(err, tx)
			}
		}
	}
	return tx.Commit()
}

func (srv *Service) SetDefenseGrade(stu string, g int) error {
	sql := "update defenseGrades set grade=$2 where student=$1"
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	res, err := srv.DB.Exec(sql, stu, g)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb == 1 {
		return nil
	}
	sql = "insert into defenseGrades(student, grade) values($1,$2)"
	_, err = srv.DB.Exec(sql, stu, g)
	return err
}
