package datastore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/internship"
)

func (srv *Service) DefenseSessions() ([]internship.DefenseSession, error) {
	sql := "select pause, date, room from defenseSessions"
	rows, err := srv.DB.Query(sql)
	defs := make([]internship.DefenseSession, 0, 0)
	if err != nil {
		return defs, err
	}
	defer rows.Close()
	for rows.Next() {
		var pause int
		var room string
		var date time.Time
		err = rows.Scan(&pause, &date, &room)
		df := internship.DefenseSession{Pause: pause, Date: date, Room: room, Juries: make([]string, 0, 0), Defenses: make([]internship.Defense, 0, 0)}
		//pick the jury member
		juries, err := srv.jury(date, room)
		if err != nil {
			return defs, err
		}
		df.Juries = juries
		//The defenses
		dd, err := srv.Defenses(date, room)
		if err != nil {
			return defs, err
		}
		df.Defenses = dd
		defs = append(defs, df)
	}
	return defs, nil
}

func (srv *Service) defenseSession(room string, date time.Time) (internship.DefenseSession, error) {
	sql := "select pause from defenseSessions where room=$1 and date=$2"
	row := srv.DB.QueryRow(sql, room, date)

	var pause int
	err := row.Scan(&pause)
	if err != nil {
		return internship.DefenseSession{}, err
	}
	df := internship.DefenseSession{Pause: pause, Date: date, Room: room, Juries: make([]string, 0, 0), Defenses: make([]internship.Defense, 0, 0)}
	//pick the jury member
	juries, err := srv.jury(date, room)
	if err != nil {
		return internship.DefenseSession{}, err
	}
	df.Juries = juries
	//The defenses
	dd, err := srv.Defenses(date, room)
	if err != nil {
		return internship.DefenseSession{}, err
	}
	df.Defenses = dd
	return df, nil
}

func (srv *Service) jury(date time.Time, room string) ([]string, error) {
	sql := "select jury from defenseJuries where date=$1 and room=$2"
	rows, err := srv.DB.Query(sql, date, room)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()
	juries := make([]string, 0, 0)
	for rows.Next() {
		var j string
		rows.Scan(&j)
		juries = append(juries, j)
	}
	return juries, nil
}

//JuryDefenseSessions returns all the registered defense sessions for a given jury member
func (srv *Service) JuryDefenseSessions(jury string) ([]internship.DefenseSession, error) {
	sql := "select distinct room, date from defenseJuries where jury=$1"
	defs := make([]internship.DefenseSession, 0, 0)
	rows, err := srv.DB.Query(sql, jury)
	if err != nil {
		return defs, err
	}
	defer rows.Close()
	for rows.Next() {
		var room string
		var date time.Time
		rows.Scan(&room, &date)
		d, err := srv.defenseSession(room, date)
		if err != nil {
			return defs, err
		}
		defs = append(defs, d)
	}
	return defs, err
}

func (srv *Service) Defenses(date time.Time, room string) ([]internship.Defense, error) {
	req := "select student, private, remote, grade from defenses where date=$1 and room=$2"
	defs := make([]internship.Defense, 0, 0)
	rows, err := srv.DB.Query(req, date, room)
	if err != nil {
		return defs, err
	}
	defer rows.Close()
	for rows.Next() {
		var grade sql.NullInt64
		var remote, private bool
		var student string
		rows.Scan(&student, &private, &remote, &grade)
		d := internship.Defense{Student: student, Private: private, Remote: remote, Grade: -1}
		if grade.Valid {
			d.Grade = int(grade.Int64)
		}
		defs = append(defs, d)
	}

	if err != nil {
		return defs, err
	}
	return defs, nil
}

func (srv *Service) Defense(student string) (internship.StudentDefense, error) {
	req := "select private, remote, date, room, grade from defenses where student=$1"
	var remote, private bool
	var room string
	var date time.Time
	var grade sql.NullInt64
	rows, err := srv.DB.Query(req, student)
	if err != nil {
		return internship.StudentDefense{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return internship.StudentDefense{Grade: -1}, nil
	}
	err = rows.Scan(&private, &remote, &date, &room, &grade)
	if err != nil {
		return internship.StudentDefense{}, err
	}
	d := internship.StudentDefense{Private: private, Remote: remote, Room: room, Date: date, Juries: make([]internship.User, 0, 0), Grade: -1}
	if grade.Valid {
		d.Grade = int(grade.Int64)
	}
	rows, err = srv.DB.Query("select firstname, lastname, email, tel from defenseJuries,users where date=$1 and room=$2 and defenseJuries.jury=users.Email", date, room)
	if err != nil {
		return d, err
	}
	defer rows.Close()
	for rows.Next() {
		var fn, ln, tel, email string
		rows.Scan(&fn, &ln, &email, &tel)
		u := internship.User{Firstname: fn, Lastname: ln, Tel: tel, Email: email}
		d.Juries = append(d.Juries, u)
	}
	return d, nil
}

func (srv *Service) SetDefenseGrade(stu string, g int) error {
	sql := "update defenses set grade=$2 where student=$1"
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	return SingleUpdate(srv.DB, internship.ErrUnknownUser, sql, stu, g)
}

func (srv *Service) SetDefenseSessions(defs []internship.DefenseSession) error {
	tx, err := srv.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from defenseSessions")
	if err != nil {
		return rollback(err, tx)
	}
	for _, def := range defs {
		_, err = tx.Exec("insert into defenseSessions(date, room, pause) values($1,$2,$3)", def.Date, def.Room, def.Pause)
		if err != nil {
			return rollback(err, tx)
		}
		for _, j := range def.Juries {
			_, err = tx.Exec("insert into defenseJuries(date,room,jury) values($1,$2,$3)", def.Date, def.Room, j)
			if err != nil {
				return rollback(err, tx)
			}
		}
		for _, s := range def.Defenses {
			_, err = tx.Exec("insert into defenses(student,private,remote,date,room) values($1,$2,$3,$4,$5)", s.Student, s.Private, s.Remote, def.Date, def.Room)
			if err != nil {
				return rollback(err, tx)
			}
		}
	}
	return tx.Commit()
}
