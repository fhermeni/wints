package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/internship"
)

var (
	updateConventionSkip     = "update conventions set skip=$1 where student=$2"
	updateSupervisor         = "update conventions set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	updateTutor              = "update conventions set tutor=$1 where student=$2"
	updateCompany            = "update conventions set companyWWW=$1, company=$2 where student=$3"
	updateTitle              = "update conventions set title=$1 where student=$2"
	insertConvention         = "insert into conventions(student, male, startTime, endTime, tutor, companyName, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, creation, foreignCountry, lab, gratification, skip, valid) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18"
	updateConvention         = "update conventions set startTime=$2, endTime=$3, tutor=$4, companyName=$5, companyWWW=$6, supervisorFn=$7, supervisorLn=$8, supervisorEmail=$9, supervisorTel=$10, title=$11, creation=$12, foreignCountry=$13, lab=$14, gratification=$15, where student=$1"
	selectConventionCreation = "select creation from conventions where student=$1"
	selectConventions        = "select stup.firstname, stup.lastname, stup.tel, stup.email, stup.lastVisit" +
		"students.male, students.promotion, students.major, students.nextPosition, students.nextContact, students.skip," +
		"tutp.firstname, tutp.lastname, tutp.tel, tutp.email, tutp.lastVisit, tutp.role" +
		"startTime, endTime, companyName, companyWWW, title, creation, foreignCountry, lab, gratification, skip, valid," +
		"supervisorFn, supervisorLn, supervisorEmail, supervisorTel " +
		"from conventions " +
		" inner join students on (students.email = conventions.student) " +
		" inner join users as stup on (stup.email = conventions.student)  " +
		" inner join users as tutp on (tutp.email = conventions.tutor)  "
	selectConvention = "select stup.firstname, stup.lastname, stup.tel, stup.email, stup.lastVisit" +
		"students.male, students.promotion, students.major, students.nextPosition, students.nextContact, students.skip," +
		"tutp.firstname, tutp.lastname, tutp.tel, tutp.email, tutp.lastVisit, tutp.role" +
		"startTime, endTime, companyName, companyWWW, title, creation, foreignCountry, lab, gratification, skip, valid," +
		"supervisorFn, supervisorLn, supervisorEmail, supervisorTel " +
		"from conventions " +
		" inner join students on (students.email = conventions.student) " +
		" inner join users as stup on (stup.email = conventions.student)  " +
		" inner join users as tutp on (tutp.email = conventions.tutor)  " +
		" where stup.email=$1"
	insertReport       = "insert into reports(student, kind, deadline, private, toGrade) values($1,$2,$3,$4,$5)"
	insertSurvey       = "insert into surveys(student, kind, token, deadline) values($1,$2,$3,$4)"
	validateConvention = "update conventions set valid=$2 where student=$1"
)

func (s *Store) SetConventionSkippable(student string, skip bool) error {
	return s.singleUpdate(updateConventionSkip, internship.ErrUnknownStudent, skip, student)
}

func (s *Store) SetSupervisor(stu string, t internship.Person) error {
	return s.singleUpdate(updateSupervisor, internship.ErrUnknownInternship, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

func (s *Store) SetTutor(stu string, t string) error {
	return s.singleUpdate(updateTutor, internship.ErrUnknownInternship, t, stu)
}

func (s *Store) SetCompany(stu string, c internship.Company) error {
	return s.singleUpdate(updateCompany, internship.ErrUnknownInternship, c.WWW, c.Name, stu)
}

func (s *Store) SetTitle(stu string, title string) error {
	return s.singleUpdate(updateTutor, internship.ErrUnknownInternship, title, stu)
}

//New convention, the student and the tutor must already be registered
//If a convention already exists for that student but the new data refer to a fresher convention, it is updated
func (s *Store) NewConvention(student string, startTime, endTime time.Time, tutor string, cpy internship.Company, sup internship.Person, title string, creation time.Time, foreignCountry, lab bool, gratification int) (bool, error) {
	err := s.singleUpdate(insertConvention, internship.ErrUnknownStudent, student, startTime, endTime, tutor, cpy.Name, cpy.WWW, sup.Firstname, sup.Lastname, sup.Email, sup.Tel, title, creation, foreignCountry, lab, gratification, false, false)
	if err == internship.ErrConventionExists {
		//Has it been updated ?
		var last time.Time
		st, err := s.stmt(selectConventionCreation)
		if err != nil {
			return false, err
		}
		if err := st.QueryRow(student).Scan(&last); err != nil {
			return false, err
		}
		if creation.After(last) {
			return true, s.singleUpdate(updateConvention, internship.ErrUnknownConvention,
				student,
				startTime,
				endTime,
				tutor,
				cpy.Name,
				cpy.WWW,
				sup.Firstname,
				sup.Lastname,
				sup.Email,
				sup.Tel,
				title,
				creation,
				foreignCountry,
				lab,
				gratification)
		}
	}
	return false, err
}

func (s *Store) Conventions() ([]internship.Convention, error) {
	conventions := make([]internship.Convention, 0, 0)
	st, err := s.stmt(selectConventions)
	if err != nil {
		return conventions, err
	}
	rows, err := st.Query()
	if err != nil {
		return conventions, err
	}
	defer rows.Close()
	for rows.Next() {
		c, err := scanConvention(rows)
		if err != nil {
			return conventions, err
		}
		conventions = append(conventions, c)
	}
	return conventions, nil
}

func prepareReports(tx *TxErr, student string, reports map[string]config.Report) error {
	st, err := tx.tx.Prepare(insertReport)
	if err != nil {
		return err
	}
	for kind, report := range reports {
		_, tx.err = st.Exec(student, kind, report.Deadline, false, report.Grade)
	}
	return tx.Done()
}

func prepareSurveys(tx *TxErr, student string, surveys map[string]config.Survey) error {
	st, err := tx.tx.Prepare(insertSurvey)
	if err != nil {
		return err
	}
	for kind, survey := range surveys {
		token := randomBytes(16)
		_, tx.err = st.Exec(student, kind, token, survey.Deadline)
	}
	return tx.Done()
}

func (s *Store) ValidateConvention(student string, cfg config.Config) error {
	tx := newTxErr(s.db)
	prepareReports(&tx, student, cfg.Reports)
	prepareSurveys(&tx, student, cfg.Surveys)
	//The prepare* makes sure there is a student. No need to check update
	tx.Update(validateConvention, student, true)
	return tx.Done()
}

func (s *Store) Internships() ([]internship.Internship, error) {
	res := make([]internship.Internship, 0, 0)
	conventions, err := s.Conventions()
	if err != nil {
		return res, err
	}
	for _, c := range conventions {

		i, err := s.toInternship(c)
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}
	return res, err
}

func (s *Store) Internship(student string) (internship.Internship, error) {
	i := internship.Internship{}
	st, err := s.stmt(selectConvention)
	if err != nil {
		return i, err
	}
	rows, err := st.Query(student)
	if err != nil {
		return i, err
	}
	defer rows.Close()
	c, err := scanConvention(rows)
	if err != nil {
		return i, err
	}
	return s.toInternship(c)
}

func (s *Store) toInternship(c internship.Convention) (internship.Internship, error) {
	stu := c.Student.User.Person.Email
	r, err := s.Reports(stu)
	if err != nil {
		return internship.Internship{}, err
	}
	surveys, err := s.Surveys(stu)
	if err != nil {
		return internship.Internship{}, err
	}
	d, err := s.Defense(stu)
	if err != nil {
		return internship.Internship{}, err
	}
	return internship.Internship{Convention: c, Reports: r, Surveys: surveys, Defense: d}, err
}

func scanConvention(rows *sql.Rows) (internship.Convention, error) {
	c := internship.Convention{
		Student: internship.Student{
			User: internship.User{
				Person: internship.Person{},
			},
			Alumni: internship.Alumni{},
		},
		Tutor: internship.User{
			Person: internship.Person{},
		},
		Supervisor: internship.Person{},
		Cpy:        internship.Company{},
	}
	err := rows.Scan(
		&c.Student.User.Person.Firstname,
		&c.Student.User.Person.Lastname,
		&c.Student.User.Person.Tel,
		&c.Student.User.Person.Email,
		&c.Student.User.LastVisit,
		&c.Student.Male,
		&c.Student.Promotion,
		&c.Student.Major,
		&c.Student.Alumni.Position,
		&c.Student.Alumni.Contact,
		&c.Student.Skip,
		&c.Tutor.Person.Firstname,
		&c.Tutor.Person.Lastname,
		&c.Tutor.Person.Tel,
		&c.Tutor.Person.Email,
		&c.Tutor.LastVisit,
		&c.Tutor.Role,
		&c.Begin,
		&c.End,
		&c.Cpy.Name,
		&c.Cpy.WWW,
		&c.Title,
		&c.Creation,
		&c.ForeignCountry,
		&c.Lab,
		&c.Gratification,
		&c.Skip,
		&c.Valid,
		&c.Supervisor.Firstname,
		&c.Supervisor.Lastname,
		&c.Supervisor.Email,
		&c.Supervisor.Tel,
	)
	return c, err
}
