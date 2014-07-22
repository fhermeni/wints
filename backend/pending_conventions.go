package backend

import (
	"strings"
	"net/http"
	"encoding/csv"
	"io"
	"fmt"
	"time"
	"strconv"
	"database/sql"
	"log"
)

const (
	ADMIN_URL       = "http://conventions.polytech.unice.fr/admin/admin.cgi"
	stuPromotion    = 1
	stuFn           = 5
	stuLn           = 6
	stuEmail        = 17
	stuTel          = 18
	company         = 22
	companyWWW      = 23
	begin           = 38
	end             = 39
	titleIdx		= 50
	supervisorFn    = 61
	supervisorLn    = 62
	supervisorEmail = 63
	supervisorTel   = 64
	tutorFn         = 67
	tutorLn         = 68
	tutorEmail      = 69
	tutorTel        = 70
)

func cleanUser(fn, ln, email, tel string) User {
	return User{cleanName(fn),cleanName(ln),clean(email),clean(tel),""}
}

func cleanName(str string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(str)))
}

func clean(str string) string {
	return strings.TrimSpace(str)
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

func PullConventions(db *sql.DB, url, login, password string) {
	err := InspectRawConventions2(db,url, login,password,
						[]string{"Master%20IFI", "Master%20IMAFA", "MAM%205", "SI%205"}, time.Now().Year())
	if err != nil {
		log.Printf("Unable to pull the conventions: %s\n", err)
	}
}

func DaemonConventionsPuller(db *sql.DB, url, login, password string, delay time.Duration) {
	go PullConventions(db, url, login, password)
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				PullConventions(db, url, login, password);
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func ScanPendingConvention(rows *sql.Rows) (Convention, error) {
	var stuFn, stuLn, stuTel, stuEmail, promo string
	var tutorFn, tutorLn, tutorTel, tutorEmail string
	var supFn, supLn, supTel, supEmail, company, companyWWW, title string
	var start, end, midDeadline time.Time

	err := rows.Scan(&stuFn, &stuLn, &stuEmail, &stuTel, &promo, &start, &end,
		&tutorFn, &tutorLn, &tutorEmail, &tutorTel, &midDeadline,
		&company, &companyWWW, &supFn, &supLn, &supEmail, &supTel, &title)
	if err != nil {
		return Convention{}, err
	}
	stu := Student{User{stuFn, stuLn, stuEmail, stuTel,""}, promo, ""}
	tutor := User{tutorFn, tutorLn, tutorEmail, tutorTel, ""}
	sup := User{supFn, supLn, supEmail, supTel, ""}
	return Convention{stu, sup, tutor, company, companyWWW, start, end, ReportMetaData{},ReportMetaData{},ReportMetaData{}, title}, nil
}

func GetRawConventions(db *sql.DB) ([]Convention, error) {
	sql := "select users.firstname, users.lastname, users.email, users.tel, students.promotion, startTime, endTime, tutorFn, " +
			"tutorLn, tutorEmail, tutorTel, midTermDeadline," +
			"company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title" +
			" from pending_internships, users, students where students.email = pending_internships.student and users.email = pending_internships.student"

	rows, err := db.Query(sql)
	cc := make([]Convention, 0, 0)
	if err != nil {
		return cc, err
	}
	defer rows.Close()
	for rows.Next() {
		c, err := ScanPendingConvention(rows)
		if err != nil {
			return cc, err
		}
		cc = append(cc, c)
	}
 	return cc, err;
}

func InspectRawConvention(db *sql.DB, c Convention) error {
	_, err := GetStudent(db, c.Stu.P.Email)
	if err == nil {
		//The student is known, ie. there is already either a pending or a committed convention
		return nil
	}
	err = NewStudent(db, c.Stu)
	if err != nil {
 		return err
	}

	tutor := c.Tutor
	_, err = GetUser(db, tutor.Email)
	if err == nil {
		return RegisterInternship(db, c, false)
	}
	return RegisterPendingInternship(db, c)
}

func RescanPending(db *sql.DB, c Convention) error {
	sql := "select student, midTermDeadline," +
			"company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, starttime, endtime, midtermdeadline" +
			" from pending_internships where tutorEmail=$1"
	var student, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title string
	var md, begin, end time.Time

	//Catch pending interviews with a known tutor
	rows, err := db.Query(sql, c.Tutor.Email)

	if err != nil {
		return err
	}
	defer rows.Close();
	for rows.Next() {
		//All
		rows.Scan(&student, &md, &company, &companyWWW, &supervisorFn, &supervisorLn, &supervisorEmail, &supervisorTel, &begin, &end, &md, &title)
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
		err = RegisterInternship(db, c, true)
	}
	return nil
}

func RegisterPendingInternship(db *sql.DB, c Convention) error {
	supervisor := c.Sup
	tutor := c.Tutor
	_, err := db.Exec("insert into pending_internships (student, startTime, endTime, tutorFn, tutorLn, tutorTel, tutorEmail, midtermDeadline, company, companyWWW, supervisorFn, supervisorLn, supervisorEmail, supervisorTel, title) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14, $15)",
		c.Stu.P.Email, c.Begin, c.End, tutor.Firstname, tutor.Lastname, tutor.Tel, tutor.Email, c.Begin.Add(TWO_MONTHS), c.Company, c.CompanyWWW, supervisor.Firstname, supervisor.Lastname, supervisor.Email, supervisor.Tel, c.Title)
	return err
}

func InspectRawConventions2(db *sql.DB, url, login, password string, promotions []string, year int ) error {
	for _, p := range promotions {
		err := getRawConventions2(db, url, login, password, year, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func getRawConventions2(db *sql.DB, url, login, password string, year int, promotion string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url+"?action=down&filiere="+promotion+"&annee="+strconv.Itoa(year), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(login, password)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	in := csv.NewReader(res.Body)
	in.Comma = ';'

	//Get rid of the header
	in.Read()

	nb := 0
	for {
		record, err := in.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		p := cleanUser(record[stuFn], record[stuLn], record[stuEmail], record[stuTel])
		stu := Student{p, clean(record[stuPromotion]), ""}
		supervisor := cleanUser(record[supervisorFn], record[supervisorLn], record[supervisorEmail], record[supervisorTel])
		tutor := cleanUser(record[tutorFn], record[tutorLn], record[tutorEmail], record[tutorTel])
		startTime, err := time.Parse("02/01/2006", clean(record[begin]))
		if err != nil {
			return err
		}
		endTime, err := time.Parse("02/01/2006", clean(record[end]))
		if err != nil {
			return err
		}
		companyWWW := clean(record[companyWWW])
		title := clean(record[titleIdx])
		if len(companyWWW) != 0 && !strings.HasPrefix(companyWWW, "http") {
			companyWWW = "http://" + companyWWW
		}
		c := Convention{stu, supervisor, tutor, clean(record[company]), companyWWW, startTime, endTime, ReportMetaData{},ReportMetaData{}, ReportMetaData{}, title}
		InspectRawConvention(db, c)
		nb++;
	}
	log.Printf("Parsed %d conventions for promotion %s/%d\n", nb, promotion, year);
	return nil
}
