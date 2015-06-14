package datastore

import (
	"database/sql"

	"github.com/fhermeni/wints/internship"
)

var JuriesStmt *sql.Stmt
var JuryUsersStmt *sql.Stmt
var DefStmt *sql.Stmt
var DefsStmt *sql.Stmt
var SessionsStmt *sql.Stmt
var pubDefStmt *sql.Stmt

func (srv *Service) DefenseSession(student string) (internship.DefenseSession, error) {
	var err error
	session := internship.DefenseSession{}
	def := internship.Defense{}
	if DefStmt == nil {
		if DefStmt, err = srv.DB.Prepare("select defenses.student,date,room,rank,private,remote,defenseGrades.grade from defenses left outer join defenseGrades on defenses.student=defenseGrades.student where defenses.student=$1"); err != nil {
			return session, err
		}
	}
	rows, err := DefStmt.Query(student)
	if err != nil {
		return session, err
	}
	defer rows.Close()
	var grade sql.NullInt64
	if rows.Next() {
		if err = rows.Scan(&def.Student, &session.Date, &session.Room, &def.Offset, &def.Private, &def.Remote, &grade); err != nil {
			return session, nil
		}
		if grade.Valid {
			def.Grade = int(grade.Int64)
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
		if DefsStmt, err = srv.DB.Prepare("select defenses.student,rank,private,remote,defenseGrades.grade from defenses left outer join defenseGrades on defenses.student=defenseGrades.student where date=$1 and room=$2 order by rank"); err != nil {
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
		var grade sql.NullInt64
		if err = rows.Scan(&def.Student, &def.Offset, &def.Private, &def.Remote, &grade); err != nil {
			return err
		}
		if grade.Valid {
			def.Grade = int(grade.Int64)
		}
		s.Defenses = append(s.Defenses, def)
	}
	return err
}

func (srv *Service) appendJuries(s *internship.DefenseSession) error {
	if JuriesStmt == nil {
		var err error
		JuriesStmt, err = srv.DB.Prepare("select firstname,lastname,tel,email from users inner join defenseJuries on users.email = defenseJuries.jury where date=$1 and room=$2")
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

func (srv *Service) PublicDefenseSessions() ([]internship.PublicDefenseSession, error) {
	sessions, err := srv.DefenseSessions()
	pubs := make([]internship.PublicDefenseSession, 0, 0)
	if err != nil {
		return pubs, err
	}

	if pubDefStmt == nil {
		var err error
		pubDefStmt, err = srv.DB.Prepare("select firstname,lastname,tel,email, title, company, major, promotion from internships,users where internships.student=$1 and users.email=$1")
		if err != nil {
			return pubs, err
		}
	}

	for _, s := range sessions {
		ps := internship.PublicDefenseSession{Room: s.Room, Date: s.Date, Juries: s.Juries, Defenses: make([]internship.PublicDefense, 0, 0)}
		for _, d := range s.Defenses {
			stu := internship.User{}
			var title, cpy, major, promotion string
			row := pubDefStmt.QueryRow(d.Student)
			err = row.Scan(&stu.Firstname, &stu.Lastname, &stu.Tel, &stu.Email, &title, &cpy, &major, &promotion)
			if err != nil {
				return pubs, err
			}
			p := internship.PublicDefense{
				Student:   stu,
				Major:     major,
				Promotion: promotion,
				Private:   d.Private,
				Remote:    d.Remote,
				Company:   cpy,
				Title:     title,
				Offset:    d.Offset,
			}
			ps.Defenses = append(ps.Defenses, p)
		}
		pubs = append(pubs, ps)
	}
	return pubs, err
}
