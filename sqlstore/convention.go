package sqlstore

import (
	"database/sql"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/schema"
)

var (
	updateConventionSkip = "update conventions set skip=$1 where student=$2"
	updateSupervisor     = "update conventions set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	updateTutor          = "update conventions set tutor=$1 where student=$2"
	updateCompany        = "update conventions set companyWWW=$1, companyName=$2 where student=$3"
	updateTitle          = "update conventions set title=$1 where student=$2"
	insertConvention     = "insert into conventions(student, startTime, endTime, tutor, companyName, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, creation, foreignCountry, lab, gratification) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)"
	updateConvention     = "update conventions set startTime=$2, endTime=$3, tutor=$4, companyName=$5, companyWWW=$6, supervisorFn=$7, supervisorLn=$8, supervisorEmail=$9, supervisorTel=$10, title=$11, creation=$12, foreignCountry=$13, lab=$14, gratification=$15, where student=$1 and creation < $12"
	selectConventions    = "select stup.firstname, stup.lastname, stup.tel, stup.email, stup.lastVisit, " +
		"students.male, students.promotion, students.major, students.nextPosition, students.nextFrance, students.nextPermanent, students.nextSameCompany, students.nextContact, students.skip," +
		"tutp.firstname, tutp.lastname, tutp.tel, tutp.email, tutp.lastVisit, tutp.role, " +
		"startTime, endTime, companyName, companyWWW, title, creation, foreignCountry, lab, gratification, " +
		"supervisorFn, supervisorLn, supervisorEmail, supervisorTel " +
		"from conventions " +
		" inner join students on (students.email = conventions.student) " +
		" inner join users as stup on (stup.email = conventions.student)  " +
		" inner join users as tutp on (tutp.email = conventions.tutor)  "
	selectConvention = "select stup.firstname, stup.lastname, stup.tel, stup.email, stup.lastVisit, " +
		"students.male, students.promotion, students.major, students.nextPosition, students.nextFrance, students.nextPermanent, students.nextSameCompany, students.nextContact, students.skip, " +
		"tutp.firstname, tutp.lastname, tutp.tel, tutp.email, tutp.lastVisit, tutp.role, " +
		"startTime, endTime, companyName, companyWWW, title, creation, foreignCountry, lab, gratification, " +
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

//SetConventionSkippable stores the skipable status for a student convention
/*func (s *Store) SetConventionSkippable(student string, skip bool) error {
	return s.singleUpdate(updateConventionSkip, schema.ErrUnknownStudent, skip, student)
}*/

//SetSupervisor updates the student supervisor
func (s *Store) SetSupervisor(stu string, t schema.Person) error {
	return s.singleUpdate(updateSupervisor, schema.ErrUnknownInternship, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

//SetTutor updates the student tutor
func (s *Store) SetTutor(stu string, t string) error {
	return s.singleUpdate(updateTutor, schema.ErrUnknownInternship, t, stu)
}

//SetCompany updates the student company
func (s *Store) SetCompany(stu string, c schema.Company) error {
	return s.singleUpdate(updateCompany, schema.ErrUnknownInternship, c.WWW, c.Name, stu)
}

func (s *Store) NewInternship(c schema.Convention, cfg config.Internships) (schema.Internship, error) {
	i := schema.Internship{
		Convention: c,
	}
	if c.Begin.After(c.End) {
		return i, schema.ErrInvalidPeriod
	}

	tx := newTxErr(s.db)
	tx.Exec(insertConvention, c.Student.User.Person.Email, c.Begin, c.End, c.Tutor.Person.Email, c.Company.Name, c.Company.WWW, c.Supervisor.Firstname, c.Supervisor.Lastname, c.Supervisor.Email, c.Supervisor.Tel, c.Company.Title, c.Creation, c.ForeignCountry, c.Lab, c.Gratification)
	for kind, report := range cfg.Reports {
		tx.Exec(insertReport, c.Student.User.Person.Email, kind, report.Delivery.Value(c.Begin), false, report.Grade)
	}
	for kind, survey := range cfg.Surveys {
		token := randomBytes(16)
		tx.Exec(insertSurvey, c.Student.User.Person.Email, kind, token, survey.Deadline.Value(c.Begin))
	}

	//refresh student & tutor informations for a clean version
	stu, err := s.Student(c.Student.User.Person.Email)
	if err != nil {
		tx.err = err
		return i, tx.Done()
	}
	tut, err := s.User(c.Tutor.Person.Email)
	if err != nil {
		tx.err = err
		return i, tx.Done()
	}
	i.Convention.Tutor = tut
	i.Convention.Student = stu
	return i, tx.Done()
}

//Convention returns the convention of a given student
func (s *Store) Convention(student string) (schema.Convention, error) {
	st := s.stmt(selectConvention)
	rows, err := st.Query(student)
	if err != nil {
		return schema.Convention{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return schema.Convention{}, schema.ErrUnknownConvention
	}
	return scanConvention(rows)
}

//Conventions lists all the registered conventions
func (s *Store) Conventions() ([]schema.Convention, error) {
	conventions := make([]schema.Convention, 0, 0)
	st := s.stmt(selectConventions)
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

//Internships returns all the internships. Not necessarily validated
func (s *Store) Internships() (schema.Internships, error) {
	res := make([]schema.Internship, 0, 0)
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

//Internship returns the internship for a given student
func (s *Store) Internship(student string) (schema.Internship, error) {
	i := schema.Internship{}
	st := s.stmt(selectConvention)
	rows, err := st.Query(student)
	if err != nil {
		return i, err
	}
	defer rows.Close()
	if !rows.Next() {
		return i, schema.ErrUnknownInternship
	}
	c, err := scanConvention(rows)
	if err != nil {
		return i, err
	}
	return s.toInternship(c)
}

func (s *Store) toInternship(c schema.Convention) (schema.Internship, error) {
	stu := c.Student.User.Person.Email
	r, err := s.Reports(stu)
	if err != nil {
		return schema.Internship{}, err
	}
	surveys, err := s.Surveys(stu)
	if err != nil {
		return schema.Internship{}, err
	}
	d, err := s.Defense(stu)
	if err != nil {
		return schema.Internship{}, err
	}
	return schema.Internship{Convention: c, Reports: r, Surveys: surveys, Defense: d}, err
}

func scanConvention(rows *sql.Rows) (schema.Convention, error) {
	var nextContact sql.NullString
	var nextPosition sql.NullString
	var nextFrance, nextPermanent, nextSameCompany sql.NullBool

	c := schema.Convention{
		Student: schema.Student{
			User: schema.User{
				Person: schema.Person{},
			},
			Alumni: &schema.Alumni{},
		},
		Tutor: schema.User{
			Person: schema.Person{},
		},
		Supervisor: schema.Person{},
		Company:    schema.Company{},
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
		&nextPosition,
		&nextFrance,
		&nextPermanent,
		&nextSameCompany,
		&nextContact,
		&c.Student.Skip,
		&c.Tutor.Person.Firstname,
		&c.Tutor.Person.Lastname,
		&c.Tutor.Person.Tel,
		&c.Tutor.Person.Email,
		&c.Tutor.LastVisit,
		&c.Tutor.Role,
		&c.Begin,
		&c.End,
		&c.Company.Name,
		&c.Company.WWW,
		&c.Company.Title,
		&c.Creation,
		&c.ForeignCountry,
		&c.Lab,
		&c.Gratification,
		&c.Supervisor.Firstname,
		&c.Supervisor.Lastname,
		&c.Supervisor.Email,
		&c.Supervisor.Tel,
	)
	if !nextPosition.Valid {
		c.Student.Alumni = nil
	} else {
		c.Student.Alumni.Position = nullableString(nextPosition)
		c.Student.Alumni.France = nullableBool(nextFrance, false)
		c.Student.Alumni.Permanent = nullableBool(nextPermanent, false)
		c.Student.Alumni.SameCompany = nullableBool(nextSameCompany, false)
		c.Student.Alumni.Contact = nullableString(nextContact)

	}
	return c, err
}