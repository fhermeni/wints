package sqlstore

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

var (
	UpdateConventionSkip     = "update conventions set skip=$1 where student=$2"
	updateSupervisor         = "update conventions set supervisorFn=$1, supervisorLn=$2, supervisorTel=$3, supervisorEmail=$4 where student=$5"
	updateTutor              = "update conventions set tutor=$1 where student=$2"
	updateCompany            = "update conventions set companyWWW=$1, company=$2 where student=$3"
	updateTitle              = "update conventions set title=$1 where student=$2"
	insertConvention         = "insert into conventions(student, male, startTime, endTime, tutor, companyName, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title, creation, foreignCountry, lab, gratification, skip, valid) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18"
	updateConvention         = "update conventions set male=$2, startTime=$3, endTime=$4, tutor=$5, companyName=$6, companyWWW=$7, supervisorFn=$8, supervisorLn=$9, supervisorEmail=$10, supervisorTel=$11, title=$12, creation=$13, foreignCountry=$14, lab=$15, gratification=$16, where student=$1"
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
)

func (s *Service) SetConventionSkippable(student string, skip bool) error {
	return s.singleUpdate(UpdateConventionSkip, internship.ErrUnknownStudent, skip, student)
}

func (s *Service) SetSupervisor(stu string, t internship.Person) error {
	return s.singleUpdate(updateSupervisor, internship.ErrUnknownInternship, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

func (s *Service) SetTutor(stu string, t string) error {
	return s.singleUpdate(updateTutor, internship.ErrUnknownInternship, t, stu)
}

func (s *Service) SetCompany(stu string, c internship.Company) error {
	return s.singleUpdate(updateCompany, internship.ErrUnknownInternship, c.WWW, c.Name, stu)
}

func (s *Service) SetTitle(stu string, title string) error {
	return s.singleUpdate(updateTutor, internship.ErrUnknownInternship, title, stu)
}

//New convention, the student and the tutor must already be registered
//If a convention already exists for that student but the new data refer to a fresher convention, it is updated
func (s *Service) NewConvention(student string, male bool,
	startTime, endTime time.Time,
	tutor string,
	cpy internship.Company,
	sup internship.Person,
	title string,
	creation time.Time,
	foreignCountry, lab bool,
	gratification int) error {
	err := s.singleUpdate(insertConvention, internship.ErrUnknownStudent,
		student,
		male,
		startTime,
		endTime,
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
		gratification,
		false,
		false)
	if err == internship.ErrConventionExists {
		//Has it been updated ?
		var last time.Time
		st, err := s.stmt(selectConventionCreation)
		if err != nil {
			return err
		}
		if err := st.QueryRow(student).Scan(&last); err != nil {
			return err
		}
		if creation.After(last) {
			return s.singleUpdate(updateConvention, internship.ErrUnknownConvention,
				student,
				male,
				startTime,
				endTime,
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
	return err

}

func (s *Service) Conventions() ([]internship.Convention, error) {
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
		rows.Scan(
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
		conventions = append(conventions, c)
	}
	return conventions, nil
}

/*
func (srv *Service) DeleteConvention(student string) error {
	sql := "delete from conventions where studentEmail=$1"
	return SingleUpdate(srv.DB, internship.ErrUnknownUser, sql, student)
}*/
