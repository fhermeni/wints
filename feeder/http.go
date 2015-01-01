package feeder

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fhermeni/wints/internship"
)

const (
	stuPromotion    = 1
	stuFn           = 5
	stuLn           = 6
	stuEmail        = 17
	stuTel          = 18
	company         = 22
	companyWWW      = 23
	begin           = 38
	end             = 39
	titleIdx        = 50
	supervisorFn    = 61
	supervisorLn    = 62
	supervisorEmail = 63
	supervisorTel   = 64
	tutorFn         = 67
	tutorLn         = 68
	tutorEmail      = 69
	tutorTel        = 70
)

type HTTPFeeder struct {
	url        string
	login      string
	password   string
	promotions []string
}

func NewHTTPFeeder(url, login, password string, promotions []string) *HTTPFeeder {
	return &HTTPFeeder{url: url, login: login, password: password, promotions: promotions}
}

func cleanPerson(fn, ln, email, tel string) internship.Person {
	return internship.Person{
		Firstname: clean(fn),
		Lastname:  clean(ln),
		Email:     clean(email),
		Tel:       clean(tel),
	}
}

func cleanCompany(name, WWW string) internship.Company {
	if len(WWW) != 0 && !strings.HasPrefix(WWW, "http") {
		WWW = "http://" + WWW
	}
	return internship.Company{
		Name: clean(name),
		WWW:  clean(WWW),
	}
}

func clean(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func (f *HTTPFeeder) scan(year int, prom string) ([]internship.Convention, error) {
	client := &http.Client{}
	conventions := make([]internship.Convention, 0, 0)
	req, err := http.NewRequest("GET", f.url+"?action=down&filiere="+prom+"&annee="+strconv.Itoa(year), nil)
	if err != nil {
		return conventions, err
	}
	req.SetBasicAuth(f.login, f.password)

	res, err := client.Do(req)
	if err != nil {
		return conventions, err
	}
	in := csv.NewReader(res.Body)
	in.Comma = ';'

	//Get rid of the header
	in.Read()

	for {
		record, err := in.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			return conventions, err
		}
		student := cleanPerson(record[stuFn], record[stuLn], record[stuEmail], record[stuTel])
		prom := clean(record[stuPromotion])
		supervisor := cleanPerson(record[supervisorFn], record[supervisorLn], record[supervisorEmail], record[supervisorTel])
		tutor := cleanPerson(record[tutorFn], record[tutorLn], record[tutorEmail], record[tutorTel])
		cpy := cleanCompany(record[company], record[companyWWW])
		title := clean(record[titleIdx])
		startTime, err := time.Parse("02/01/2006", clean(record[begin]))
		if err != nil {
			return conventions, err
		}
		endTime, err := time.Parse("02/01/2006", clean(record[end]))
		if err != nil {
			return conventions, err
		}

		c := internship.Convention{
			Student:    student,
			Supervisor: supervisor,
			Tutor:      tutor,
			Promotion:  prom,
			Cpy:        cpy,
			Begin:      startTime,
			End:        endTime,
			Title:      title,
		}
		conventions = append(conventions, c)
	}
	return conventions, err
}

func (f *HTTPFeeder) injectOnePromotion(s internship.Service, y int, p string) (int, error) {
	nb := 0
	conventions, err := f.scan(y, p)
	if err != nil {
		log.Println("Unable to scan promotion" + p + ": " + err.Error())
		return nb, err
	}
	for _, c := range conventions {
		err = s.NewConvention(c)
		if err == nil {
			nb++
		} else if err != internship.ErrConventionExists {
			log.Println("Unable insert a convention: " + err.Error())
			return nb, err
		}
	}
	return nb, err
}

func (f *HTTPFeeder) InjectConventions(s internship.Service) {
	year := time.Now().Year()
	type empty struct{}
	res := make([]int, len(f.promotions))
	errors := make([]error, len(f.promotions))
	sem := make(chan empty, len(f.promotions)) // semaphore pattern

	for i, prom := range f.promotions {
		go func(i int, p string) {
			nb, err := f.injectOnePromotion(s, year, p)
			res[i] = nb
			errors[i] = err
			sem <- empty{}
		}(i, prom)
	}
	// wait for goroutines to finish
	for i := 0; i < len(f.promotions); i++ {
		<-sem
	}
}
