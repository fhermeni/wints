package sqlstore

import "github.com/fhermeni/wints/internship"

var (
	updateDefensePrivacy  = "update defenses set private=$1 where student=$2"
	updateDefenseLocality = "update defenses set local=$1 where student=$2"
	selectDefense         = "select date, room, grade, private, local from defenses where student=$1"
	updateDefenseGrade    = "update defenses set grade=$2 where student=$1"
)

func (s *Service) SetDefensePrivacy(student string, private bool) error {
	return s.singleUpdate(updateDefensePrivacy, internship.ErrUnknownStudent, private, student)
}

func (s *Service) SetDefenseLocality(student string, local bool) error {
	return s.singleUpdate(updateDefenseLocality, internship.ErrUnknownStudent, local, student)
}

func (s *Service) SetDefenseGrade(student string, g int) error {
	if g < 0 || g > 20 {
		return internship.ErrInvalidGrade
	}
	return s.singleUpdate(updateDefenseGrade, internship.ErrUnknownDefense, student, g)
}

func (s *Service) Defense(student string) (internship.Defense, error) {
	d := internship.Defense{}
	st, err := s.stmt(selectDefense)
	if err != nil {
		return d, err
	}
	err = st.QueryRow(student).Scan(
		&d.Time,
		&d.Room,
		&d.Grade,
		&d.Private,
		&d.Local,
	)
	return d, err
}

func (s *Service) DefenseSessions() ([]internship.DefenseSession, error) {
	ss := make([]internship.DefenseSession, 0, 0)
	return ss, nil
}

/*
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
		if DefStmt, err = srv.db.Prepare("select firstname, lastname, tel, email,major, promotion,company, companyWWW, title, date,room,rank,private,remote,defenseGrades.grade from defenses inner join users on (defenses.student=users.email) inner join internships on (defenses.student=internships.student) left outer join defenseGrades on (defenses.student=defenseGrades.student) where defenses.student=$1"); err != nil {
			return session, err
		}
	}
	stu := internship.User{}
	cpy := internship.Company{}
	rows, err := DefStmt.Query(student)
	if err != nil {
		return session, err
	}
	defer rows.Close()
	var grade sql.NullInt64
	if rows.Next() {
		if err = rows.Scan(&stu.Firstname, &stu.Lastname, &stu.Tel, &stu.Email, &def.Major, &def.Promotion, &cpy.Name, &cpy.WWW, &def.Title, &session.Date, &session.Room, &def.Offset, &def.Private, &def.Remote, &grade); err != nil {
			return session, nil
		}
		def.Student = stu
		def.Cpy = cpy
		if grade.Valid {
			def.Grade = int(grade.Int64)
		} else {
			def.Grade = -1
		}
		def.Surveys, err = srv.surveys(student)
		if err != nil {
			return session, nil
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
		if SessionsStmt, err = srv.db.Prepare("select date,room from defenseSessions"); err != nil {
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
		if DefsStmt, err = srv.db.Prepare("select firstname, lastname, tel, email,major, promotion,company, companyWWW, title, rank,private,remote,defenseGrades.grade, nextPosition from defenses inner join users on (defenses.student=users.email) inner join internships on (defenses.student=internships.student) left outer join defenseGrades on (defenses.student=defenseGrades.student) where date=$1 and room=$2 order by rank"); err != nil {
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
		stu := internship.User{}
		cpy := internship.Company{}
		var grade, nextPosition sql.NullInt64
		if err = rows.Scan(&stu.Firstname, &stu.Lastname, &stu.Tel, &stu.Email, &def.Major, &def.Promotion, &cpy.Name, &cpy.WWW, &def.Title, &def.Offset, &def.Private, &def.Remote, &grade, &nextPosition); err != nil {
			return err
		}
		def.Student = stu
		def.Cpy = cpy
		if grade.Valid {
			def.Grade = int(grade.Int64)
		} else {
			def.Grade = -1
		}
		if nextPosition.Valid {
			def.NextPosition = int(nextPosition.Int64)
		}
		def.Surveys, err = srv.surveys(def.Student.Email)
		if err != nil {
			return err
		}
		s.Defenses = append(s.Defenses, def)
	}
	return err
}

func (srv *Service) appendJuries(s *internship.DefenseSession) error {
	if JuriesStmt == nil {
		var err error
		JuriesStmt, err = srv.db.Prepare("select firstname,lastname,tel,email from users inner join defenseJuries on users.email = defenseJuries.jury where date=$1 and room=$2")
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
	tx, err := srv.db.Begin()
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
			if _, err = tx.Exec("insert into defenses(date,room,student,remote,private,rank) values($1,$2,$3,$4,$5,$6)", s.Date, s.Room, def.Student.Email, def.Remote, def.Private, def.Offset); err != nil {
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
	res, err := srv.db.Exec(sql, stu, g)
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
	_, err = srv.db.Exec(sql, stu, g)
	return err
}
*/
