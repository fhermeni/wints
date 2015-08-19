package datastore

/*
func (srv *Service) NewConvention(c internship.Convention) error {
	sql := "insert into conventions(male, creation, studentEmail, studentFn, studentLn, studentTel, promotion, tutorEmail, tutorFn, tutorLn, tutorTel, supervisorEmail, supervisorFn, supervisorLn, supervisorTel, startTime, endTime, title, company, companyWWW, skip, gratification, foreignCountry, lab) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24)"
	_, err := srv.DB.Exec(sql, c.Male, c.Creation, c.Student.Email, c.Student.Firstname, c.Student.Lastname, c.Student.Tel, c.Promotion, c.Tutor.Email, c.Tutor.Firstname, c.Tutor.Lastname, c.Tutor.Tel, c.Supervisor.Email, c.Supervisor.Firstname, c.Supervisor.Lastname, c.Supervisor.Tel, c.Begin, c.End, c.Title, c.Cpy.Name, c.Cpy.WWW, c.Skip, c.Gratification, c.ForeignCountry, c.Lab)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Constraint == "conventions_pkey" {
				return internship.ErrConventionExists
			}
		}
		return err

	}
	return err
}

func (srv *Service) Conventions() ([]internship.Convention, error) {
	sql := "select male, creation, studentEmail, studentFn, studentLn, studentTel, promotion, tutorEmail, tutorFn, tutorLn, tutorTel, supervisorEmail, supervisorFn, supervisorLn, supervisorTel, startTime, endTime, title, company, companyWWW, skip, gratification, foreignCountry, lab from conventions"
	rows, err := srv.DB.Query(sql)
	conventions := make([]internship.Convention, 0, 0)
	if err != nil {
		return conventions, err
	}
	defer rows.Close()
	for rows.Next() {
		var stuEmail, stuFn, stuLn, stuTel, prom string
		var tutEmail, tutFn, tutLn, tutTel string
		var supEmail, supFn, supLn, supTel string
		var cpyName, cpyWWW string
		var title string
		var at, begin, end time.Time
		var male, skip, lab, foreignCountry bool
		var gratification int
		rows.Scan(&male, &at, &stuEmail, &stuFn, &stuLn, &stuTel, &prom, &tutEmail, &tutFn, &tutLn, &tutTel, &supEmail, &supFn, &supLn, &supTel, &begin, &end, &title, &cpyName, &cpyWWW, &skip, &gratification, &foreignCountry, &lab)
		cpy := internship.Company{Name: cpyName, WWW: cpyWWW}
		stu := internship.Person{Firstname: stuFn, Lastname: stuLn, Email: stuEmail, Tel: stuTel}
		tut := internship.Person{Firstname: tutFn, Lastname: tutLn, Email: tutEmail, Tel: tutTel}
		sup := internship.Person{Firstname: supFn, Lastname: supLn, Email: supEmail, Tel: supTel}
		c := internship.Convention{Male: male, Creation: at, Student: stu, Tutor: tut, Supervisor: sup, Promotion: prom, Title: title, Begin: begin, End: end, Cpy: cpy, Skip: skip, ForeignCountry: foreignCountry, Lab: lab, Gratification: gratification}
		conventions = append(conventions, c)
	}
	return conventions, nil
}

func (srv *Service) SkipConvention(student string, skip bool) error {
	sql := "update conventions set skip=$1 where studentEmail=$2"
	return SingleUpdate(srv.DB, internship.ErrUnknownUser, sql, skip, student)
}

func (srv *Service) DeleteConvention(student string) error {
	sql := "delete from conventions where studentEmail=$1"
	return SingleUpdate(srv.DB, internship.ErrUnknownUser, sql, student)
}*/
