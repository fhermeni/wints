package internship

import (
	"database/sql"
	"time"

	"github.com/fhermeni/wints/db"
)

type DbService struct {
	db *sql.DB
}

/*func (srv *DbService) New(stu string, sup string, tut Tutor, from, to time.Time, title string) (Internship, error) {

}*/

func (srv *DbService) SetTutor(stu string, t Tutor) error {
	sql := "update internship set tutorFn=$1, tutorLn=$2, tutorTel=$3, tutorEmail=$4 where student=$5"
	return db.SingleUpdate(srv.db, ErrUnknown, sql, t.Firstname, t.Lastname, t.Tel, t.Email, stu)
}

func (srv *DbService) SetCompany(stu string, c Company) error {
	sql := "update internship set companyWWW=$1, companyName=$2 where student=$3"
	return db.SingleUpdate(srv.db, ErrUnknown, sql, c.WWW, c.Name, stu)
}

func (srv *DbService) Get(stu string) (Internship, error) {
	sql := "select companyName, companyWWW, title, start, end, supervisor, tutorFn, tutorLn, tutorEmail, tutorTel from internship where student=$1"
	var companyName, companyWWW string
	var title string
	var start, end time.Time
	var supervisor string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	err := srv.db.QueryRow(sql, &companyName, &companyWWW, &title, &start, &end, &supervisor, &tutorFn, &tutorLn, &tutorEmail, &tutorTel)
	if err != nil {
		return Internship{}, ErrUnknown
	}
	c := Company{Name: companyName, WWW: companyWWW}
	t := Tutor{Firstname: tutorFn, Lastname: tutorLn, Tel: tutorTel, Email: tutorEmail}

	return Internship{
		Stu:   stu,
		Sup:   supervisor,
		Tut:   t,
		Cpy:   c,
		Begin: start,
		End:   end,
		Title: title,
	}, nil
}

func (srv *DbService) List() ([]Internship, error) {
	sql := "select student, companyName, companyWWW, title, start, end, supervisor, tutorFn, tutorLn, tutorEmail, tutorTel from internship"
	internships := make([]Internship, 0, 0)
	var companyName, companyWWW string
	var title string
	var start, end time.Time
	var student, supervisor string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	rows, err := srv.db.Query(sql)
	if err != nil {
		return internships, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&student, &companyName, &companyWWW, &title, &start, &end, &supervisor, &tutorFn, &tutorLn, &tutorEmail, &tutorTel)
		c := Company{Name: companyName, WWW: companyWWW}
		t := Tutor{Firstname: tutorFn, Lastname: tutorLn, Tel: tutorTel, Email: tutorEmail}
		i := Internship{Stu: student, Sup: supervisor, Tut: t, Cpy: c, Begin: start, End: end, Title: title}
		internships = append(internships, i)
	}
	return internships, nil
}
