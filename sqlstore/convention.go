package sqlstore

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/schema"
)

var (
	updateSupervisor  = "update conventions set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	updateTutor       = "update conventions set tutor=$1 where student=$2"
	updateCompany     = "update conventions set companyWWW=$1, companyName=$2, title=$3 where student=$4"
	insertConvention  = "insert into conventions(student, startTime, endTime, tutor, companyName, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, creation, foreignCountry, lab, gratification) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)"
	selectConventions = "select stup.firstname, stup.lastname, stup.tel, stup.email, stup.lastVisit, " +
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
	insertReport = "insert into reports(student, kind, deadline, private, toGrade) values($1,$2,$3,$4,$5)"
	insertSurvey = "insert into surveys(student, kind, token, invitation, lastInvitation, deadline) values($1,$2,$3,$4,$5,$6)"
)

//SetSupervisor updates the student supervisor
func (s *Store) SetSupervisor(stu string, t schema.Person) error {
	return s.singleUpdate(updateSupervisor, schema.ErrUnknownInternship, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

//SetTutor updates the student tutor
func (s *Store) SetTutor(stu string, t string) error {
	err := s.singleUpdate(updateTutor, schema.ErrUnknownInternship, t, stu)
	if err == schema.ErrUserTutoring {
		//Bad case, here it is because the user does not exists
		return schema.ErrUnknownUser
	}
	return err
}

//SetCompany updates the student company
func (s *Store) SetCompany(stu string, c schema.Company) error {
	return s.singleUpdate(updateCompany, schema.ErrUnknownInternship, c.WWW, c.Name, c.Title, stu)
}

//NewInternship turns a convention into an internship.
//Returns the internship, the student token and an error
func (s *Store) NewInternship(c schema.Convention) (schema.Internship, []byte, error) {
	i := schema.Internship{
		Convention: c,
	}
	if c.Begin.After(c.End) {
		return i, []byte{}, schema.ErrInvalidPeriod
	}

	tx := newTxErr(s.db)
	tx.Exec(insertConvention, c.Student.User.Person.Email, c.Begin.Truncate(time.Minute).UTC(), c.End.Truncate(time.Minute).UTC(), c.Tutor.Person.Email, c.Company.Name, c.Company.WWW, c.Supervisor.Firstname, c.Supervisor.Lastname, c.Supervisor.Email, c.Supervisor.Tel, c.Company.Title, c.Creation.Truncate(time.Minute).UTC(), c.ForeignCountry, c.Lab, c.Gratification)
	token := randomBytes(32)
	//We delete in case we start a reset over an unvalidated account
	tx.Exec(deletePasswordRenewalRequest, c.Student.User.Person.Email)
	tx.Exec(startPasswordRenewal, c.Student.User.Person.Email, token)

	i.Reports = make([]schema.ReportHeader, len(s.config.Reports))
	//Instantiate the reports and the surveys
	for idx, report := range s.config.Reports {
		tx.Exec(insertReport, c.Student.User.Person.Email, report.Kind, report.Delivery.Value(c.Begin).Truncate(time.Minute).UTC(), false, report.Grade)
		i.Reports[idx] = schema.ReportHeader{
			Kind:     report.Kind,
			ToGrade:  report.Grade,
			Grade:    -1,
			Deadline: report.Delivery.Value(c.Begin).Truncate(time.Minute).UTC(),
		}
	}

	i.Surveys = make([]schema.SurveyHeader, len(s.config.Surveys))
	for idx, survey := range s.config.Surveys {
		token16 := randomBytes(16)
		inv := survey.Invitation.Value(c.Begin).Truncate(time.Minute).UTC()
		dead := inv.Add(survey.Deadline.Duration).Truncate(time.Minute).UTC()
		tx.Exec(insertSurvey, c.Student.User.Person.Email, survey.Kind, token16, inv, inv, dead)
		i.Surveys[idx] = schema.SurveyHeader{
			Kind:           survey.Kind,
			Token:          string(token16),
			Deadline:       dead,
			Invitation:     inv,
			LastInvitation: inv,
		}
	}

	//refresh student & tutor informations for a clean version
	//as here, we only consider in c the person emails
	stu, err := s.Student(c.Student.User.Person.Email)
	if err != nil {
		tx.err = err
		return i, []byte{}, tx.err
	}
	tut, err := s.User(c.Tutor.Person.Email)
	if err != nil {
		tx.err = err
		return i, []byte{}, tx.err
	}
	i.Convention.Tutor = tut
	i.Convention.Student = stu
	return i, token, tx.Done()
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
func (s *Store) conventions() ([]schema.Convention, error) {
	var conventions []schema.Convention
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
	conventions, err := s.conventions()
	if err != nil {
		return schema.Internships{}, err
	}
	defs, err := s.defenses()
	if err != nil {
		return schema.Internships{}, err
	}

	surveys, err := s.allSurveys()
	if err != nil {
		return schema.Internships{}, err
	}

	reports, err := s.allReports()
	if err != nil {
		return schema.Internships{}, err
	}

	ints := make([]schema.Internship, 0, 0)
	for _, c := range conventions {
		stu := c.Student.User.Person.Email
		i := schema.Internship{Convention: c}
		d, ok := defs[stu]
		if ok {
			i.Defense = d
		}
		s, ok := surveys[stu]
		if ok {
			i.Surveys = s
		}
		r, ok := reports[stu]
		if ok {
			i.Reports = r
		}
		ints = append(ints, i)
	}
	return ints, err
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
	return schema.Internship{Convention: c, Reports: r, Surveys: surveys}, err
}

func scanConvention(rows *sql.Rows) (schema.Convention, error) {
	var nextContact sql.NullString
	var nextPosition sql.NullString
	var nextFrance, nextPermanent, nextSameCompany sql.NullBool
	var role string
	c := schema.Convention{
		Student: schema.Student{
			User: schema.User{
				Person: schema.Person{},
				Role:   schema.STUDENT,
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
		&role,
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
	c.Begin = c.Begin.UTC()
	c.End = c.End.UTC()
	c.Creation = c.Creation.UTC()

	c.Tutor.Role = schema.Role(role)
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
