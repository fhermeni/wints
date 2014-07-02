package backend

import (
	"fmt"
	"time"
	"net/http"
	"strconv"
	"encoding/csv"
	"io"
	"database/sql"
	"log"
	"strings"
	"errors"
)

type Convention struct {
	Stu        Student
	Sup        User
	Tutor      User
	Company    string
	CompanyWWW string
	Begin      time.Time
	End        time.Time
	MidtermReport time.Time
}


const (
	ADMIN_URL       = "http://conventions.polytech.unice.fr/admin/admin.cgi"
	LOGIN           = "stage"
	PASSWORD        = "epu2009"
	stuPromotion    = 1
	stuGn           = 4
	stuFn           = 5
	stuLn           = 6
	stuEmail        = 17
	stuTel          = 18
	company         = 22
	companyWWW      = 23
	begin           = 38
	end             = 39
	supervisorGn    = 60
	supervisorFn    = 61
	supervisorLn    = 62
	supervisorEmail = 63
	supervisorTel   = 64
	tutorGn         = 66
	tutorFn         = 67
	tutorLn         = 68
	tutorEmail      = 69
	tutorTel        = 70
	TWO_MONTHS = 2*time.Hour*24*30
)

func InspectRawConvention(db *sql.DB, c Convention) {
	stu := c.Stu
	//check for a committed convention
	_, err := GetConvention(db, stu.P.Email)
	if err == nil {
		return
	}
	//Check for a pending convention
	_, err = GetPendingConvention(db, stu.P.Email)
	if err == nil {
		return
	}
	_, err = GetStudent(db, c.Stu.P.Email)
	if err != nil {
		err := NewStudent(db, c.Stu)
		if err != nil {
			log.Printf("Unable to create student %s: %s\n", stu, err)
			return
		}
	}

	tutor := c.Tutor
	_, err = GetUser(db, tutor.Email)
	if err == nil {
		err = RegisterInternship(db, c, false)
		if err != nil {
			log.Printf("Unable to register the internship of %s: %s\n", stu, err)
			return
		}
	} else {
		err = RegisterPendingInternship(db, c)
		if err != nil {
			log.Printf("Unable to register the pending internship of %s: %s\n", stu, err)
			log.Printf("%s\n", c)
		}
	}
}

func RescanPending(db *sql.DB, c Convention) error {
	log.Printf("Rescaning to get internships tutored by %s\n", c.Tutor)

	sql := "select student, midTermDeadline," +
			"company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, starttime, endtime, midtermdeadline" +
			" from pending_internships where tutorEmail=$1"
	var student, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel string
	var md, begin, end time.Time

	//Catch pending interviews with a known tutor
	rows, err := db.Query(sql, c.Tutor.Email)

	if err != nil {
		return err
	}
	defer rows.Close();
	for rows.Next() {
		//All
		rows.Scan(&student, &md, &company, &companyWWW, &supervisorFn, &supervisorLn, &supervisorEmail, &supervisorTel, &begin, &end, &md)
		//Update the informations and register
		c.Company = company
		c.CompanyWWW = companyWWW
		c.Sup.Firstname = supervisorFn
		c.Sup.Lastname = supervisorLn
		c.Sup.Email = supervisorEmail
		c.Sup.Tel = supervisorTel
		c.Begin = begin
		c.End = end
		c.Stu.P.Email = student
		//log.Printf("Got %s\n", student, c)
		err = RegisterInternship(db, c, true)
	}
	return nil
}

func RegisterInternship(db *sql.DB, c Convention, move bool) error {
	supervisor := c.Sup
	_, err := db.Exec("insert into internships (student, startTime, endTime, tutor, midtermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
		c.Stu.P.Email, c.Begin, c.End, c.Tutor.Email, c.Begin.Add(TWO_MONTHS), c.Company, c.CompanyWWW, supervisor.Firstname, supervisor.Lastname, supervisor.Email, supervisor.Tel)
	if err != nil {
		log.Printf("Error: %s %s\n", c, err)
		return err
	}
	if move {
		res, err := db.Exec("delete from pending_internships where student=$1", c.Stu.P.Email)
		if err != nil {
			return err
		}
		nb, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if nb != 1 {
			return errors.New("No pending internship deleted with email '" + c.Stu.P.Email + "'")
		}
	}
	log.Printf("Convention of %s validated\n", c.Stu)
	return err
}

func RegisterPendingInternship(db *sql.DB, c Convention) error {
	supervisor := c.Sup
	tutor := c.Tutor
	_, err := db.Exec("insert into pending_internships (student, startTime, endTime, tutorFn, tutorLn, tutorTel, tutorEmail, midtermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
		c.Stu.P.Email, c.Begin, c.End, tutor.Firstname, tutor.Lastname, tutor.Tel, tutor.Email, c.Begin.Add(TWO_MONTHS), c.Company, c.CompanyWWW, supervisor.Firstname, supervisor.Lastname, supervisor.Email, supervisor.Tel)
	return err
}

func GetAllRawConventions() ([]Convention, error) {
	year := time.Now().Year()
	Promotions := []string{"Master%20IFI", "Master%20IMAFA", "MAM%205", "SI%205"}
	stats := make([]string, 4, 4)
	conventions := make([]Convention, 0, 0)
	for i, Promotion := range Promotions {
		cc, err := getRawConventions(year, Promotion)
		if err != nil {
			return conventions, err
		}
		stats[i] = fmt.Sprintf("%d %s",len(cc), Promotion)
		for _, c := range cc {
			conventions = append(conventions, c)
		}
	}
	log.Printf("Parsed conventions: %s\n", strings.Join(stats, ", "))
	return conventions, nil
}

func PeekRawConvention(db * sql.DB) (Convention, error) {
	sql := "select users.firstname, users.lastname, users.email, users.tel, students.promotion, startTime, endTime, tutorFn, " +
			"tutorLn, tutorEmail, tutorTel, midTermDeadline," +
			"company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel" +
			" from pending_internships, users, students where students.email = pending_internships.student and users.email = pending_internships.student order by random() limit 1"
	var stuFn, stuLn, stuTel, stuEmail, promo string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW string
	var start, end, midDeadline time.Time

	err := db.QueryRow(sql).Scan(&stuFn, &stuLn, &stuEmail, &stuTel, &promo, &start, &end,
		&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
		&company, &companyWWW, &supFn, &supLn, &supEmail, &supTel)
	if err != nil {
		return Convention{}, err
	}
	stu := Student{User{stuFn, stuLn, stuEmail, stuTel,""}, promo, ""}
	tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel, ""}
	sup := User{supFn, supLn, supEmail, supTel, ""}
	return Convention{stu, sup, tutor, company, companyWWW, start, end, midDeadline}, nil
}

func CountPending(db *sql.DB) (int, error) {
	sql := "select count(*) from pending_internships";
	var nb int
	err := db.QueryRow(sql).Scan(&nb)
	if err != nil {
		return -1, err
	}
	return nb, nil
}

func GetPendingConvention(db *sql.DB, email string) (Convention, error) {
	sql := "select users.firstname, users.lastname, users.tel, students.promotion, startTime, endTime, tutorFn, " +
			"tutorLn, tutorEmail, tutorTel, midTermDeadline," +
			"company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel" +
			" from pending_internships, users, students where students.email = $1 and users.email = $1 and pending_internships.student = $1"
	var stuFn, stuLn, stuTel, promo string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW string
	var start, end, midDeadline time.Time

	err := db.QueryRow(sql, email).Scan(&stuFn, &stuLn, &stuTel, &promo, &start, &end,
		&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
		&company, &companyWWW, &supFn, &supLn,&supEmail, &supTel)
	if err != nil {
		return Convention{}, err
	}
	stu := Student{User{stuFn, stuLn, email, stuTel, ""}, promo, ""}
	tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel, ""}
	sup := User{supFn, supLn, supEmail, supTel, ""}
	return Convention{stu, sup, tutor, company, companyWWW, start, end, midDeadline}, nil
}

func GetConvention(db *sql.DB, email string) (Convention, error) {
	sql := "select stu.firstname, stu.lastname, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel,"+
	"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel "+
	" from internships, users as stu, users as tut, students where students.email = $1 and stu.email = $1 and internships.student = $1 and tut.email = internships.tutor";

	var stuFn, stuLn, stuTel, promo, major string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW string
	var start, end, midDeadline time.Time

	err := db.QueryRow(sql, email).Scan(&stuFn, &stuLn, &stuTel, &promo, &major, &start, &end,
		&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
		&company, &companyWWW, &supFn, &supLn, &supEmail, &supTel)
	if err != nil {
		return Convention{}, err
	}
	stu := Student{User{stuFn, stuLn, email, stuTel, ""}, promo, major}
	tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel,""}
	sup := User{supFn, supLn, supEmail, supTel, ""}
	return Convention{stu, sup, tutor, company, companyWWW, start, end, midDeadline}, nil
}

func GetConventions(db *sql.DB) ([]Convention, error) {
	sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel,"+
			"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel "+
			" from internships, users as stu, users as tut, students where students.email = stu.email and internships.student = stu.email and tut.email = internships.tutor";


	var stuFn, stuLn, stuMail, stuTel, promo, major string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW string
	var start, end, midDeadline time.Time

	rows, err := db.Query(sql)
	conventions := make([]Convention, 0, 0)
	if err != nil {
 		return conventions, err
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&stuFn, &stuLn, &stuMail, &stuTel, &promo, &major, &start, &end,
		&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
		&company, &companyWWW, &supFn, &supLn, &supEmail, &supTel)
		if err != nil {
			return conventions, err
		}
		stu := Student{User{stuFn, stuLn, stuMail, stuTel,""}, promo, major}
		tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel, ""}
		sup := User{supFn, supLn, supEmail, supTel, ""}
		c :=  Convention{stu, sup, tutor, company, companyWWW, start, end, midDeadline}
		conventions = append(conventions, c)
	}
	return conventions, nil
}

func GetConventions2(db *sql.DB, u User) ([]Convention, error) {
	var rows *sql.Rows
	var err error
	if (u.Role == "admin" || u.Role == "root" || u.Role == "major") {
		sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel," +
				"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel " +
				" from internships, users as stu, users as tut, students where students.email = stu.email and internships.student = stu.email and tut.email = internships.tutor";
		rows, err = db.Query(sql)
	} else {
		sql := "select stu.firstname, stu.lastname, stu.email, stu.tel, students.promotion, students.major, startTime, endTime, tut.firstname, tut.lastname, tut.email, tut.tel," +
				"midTermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel " +
				" from internships, users as stu, users as tut, students where students.email = stu.email and internships.student = stu.email and tut.email = $1";
		rows, err = db.Query(sql, u.Email)
	}

	conventions := make([]Convention, 0, 0)
	if err != nil {
		return conventions, err
	}

	var stuFn, stuLn, stuMail, stuTel, promo, major string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW string
	var start, end, midDeadline time.Time


	defer rows.Close()
	for rows.Next() {
		rows.Scan(&stuFn, &stuLn, &stuMail, &stuTel, &promo, &major, &start, &end,
			&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
			&company, &companyWWW, &supFn, &supLn, &supEmail, &supTel)
		if err != nil {
			return conventions, err
		}
		stu := Student{User{stuFn, stuLn, stuMail, stuTel, ""}, promo, major}
		tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel,""}
		sup := User{supFn, supLn, supEmail, supTel,""}
		c :=  Convention{stu, sup, tutor, company, companyWWW, start, end, midDeadline}
		conventions = append(conventions, c)
	}
	return conventions, nil
}

func cleanUser(fn, ln, email, tel string) User {
	return User{
		cleanName(fn),
		cleanName(ln),
		cleanName(email),
		cleanName(tel),
		""}
}

func cleanName(str string) string {
	return strings.ToTitle(strings.ToLower(strings.TrimSpace(str)))
}

func clean(str string) string {
	return strings.TrimSpace(str)
}

func getRawConventions(year int, Promotion string) ([]Convention, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ADMIN_URL+"?action=down&filiere="+Promotion+"&annee="+strconv.Itoa(year), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(LOGIN, PASSWORD)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	in := csv.NewReader(res.Body)
	in.Comma = ';'

	//Get rid of the header
	in.Read()

	RawConventions := make([]Convention, 0, 0)
	for {
		record, err := in.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		p := cleanUser(record[stuFn], record[stuLn], record[stuEmail], record[stuTel])
		stu := Student{p, clean(record[stuPromotion]), ""}
		supervisor := cleanUser(record[supervisorFn], record[supervisorLn], record[supervisorEmail], record[supervisorTel])
		tutor := cleanUser(record[tutorFn], record[tutorLn], record[tutorEmail], record[tutorTel])
		startTime, err := time.Parse("02/01/2006", strings.TrimSpace(record[begin]))
		if err != nil {
			return nil, err
		}
		endTime, err := time.Parse("02/01/2006", strings.TrimSpace(record[end]))
		if err != nil {
			return nil, err
		}
		companyWWW := strings.TrimSpace(record[companyWWW])
		if len(companyWWW) != 0 && !strings.HasPrefix(companyWWW, "http") {
			companyWWW = "http://" + companyWWW
		}
		c := Convention{stu, supervisor, tutor, strings.TrimSpace(record[company]), companyWWW, startTime, endTime, startTime.Add(TWO_MONTHS)}
		RawConventions = append(RawConventions, c)
	}
	return RawConventions, nil
}

func SetMidtermDeadline(db *sql.DB, email string, d time.Time) error {
	return SingleUpdate(db, "update conventions set midtermDeadline=$2 where email=$1", email, d)
}
