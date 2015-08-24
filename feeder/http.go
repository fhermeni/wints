package feeder

import (
	"encoding/csv"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
)

const (
	timestamp       = 0
	stuPromotion    = 1
	foreignCountry  = 2
	lab             = 3
	gender          = 4
	stuFn           = 5
	stuLn           = 6
	stuEmail        = 17
	stuTel          = 18
	company         = 22
	companyWWW      = 23
	begin           = 38
	end             = 39
	gratification   = 47
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
	encoding   string
	j          journal.Journal
}

func NewHTTPFeeder(j journal.Journal, url, login, password string, promotions []string, enc string) *HTTPFeeder {
	return &HTTPFeeder{j: j, url: url, login: login, password: password, promotions: promotions, encoding: enc}
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

func (f *HTTPFeeder) log(msg string, err error) {
	f.j.Log("feeder", msg, err)
}
func cleanInt(str string) int {
	s := strings.Replace(str, " ", "", -1)
	m := regexp.MustCompile(`^(\d+)`).FindSubmatch([]byte(s))
	if m == nil || len(m) == 0 {
		return -1
	}
	i, err := strconv.Atoi(string(m[0]))
	if err != nil {
		return -1
	}
	return i
}
func (f *HTTPFeeder) scan(year int, prom string, s internship.Service) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", f.url+"?action=down&filiere="+prom+"&annee="+strconv.Itoa(year), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(f.login, f.password)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	//Convert to utf8
	r, err := charset.NewReader(f.encoding, res.Body)
	if err != nil {
		return err
	}
	in := csv.NewReader(r)
	in.Comma = ';'

	//Get rid of the header
	in.Read()

	for {
		record, err := in.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		student := cleanPerson(record[stuFn], record[stuLn], record[stuEmail], record[stuTel])
		supervisor := cleanPerson(record[supervisorFn], record[supervisorLn], record[supervisorEmail], record[supervisorTel])
		tutor := cleanPerson(record[tutorFn], record[tutorLn], record[tutorEmail], record[tutorTel])
		cpy := cleanCompany(record[company], record[companyWWW])
		title := clean(record[titleIdx])
		foreign := clean(record[foreignCountry]) != "non"
		inLab := clean(record[lab]) != "non"
		male := clean(record[gender]) == "m."
		gratif := cleanInt(record[gratification])
		ts, err := time.Parse("2006-01-02 15:04", clean(record[timestamp]))
		if err != nil {
			return err
		}
		startTime, err := time.Parse("02/01/2006", clean(record[begin]))
		if err != nil {
			return err
		}
		endTime, err := time.Parse("02/01/2006", clean(record[end]))
		if err != nil {
			return err
		}
		//The to-valid users
		s.NewUser(tutor.Firstname, tutor.Lastname, tutor.Tel, tutor.Email)
		s.NewUser(student.Firstname, student.Lastname, student.Tel, student.Email)
		s.SetMale(student.Email, male)
		s.NewConvention(student.Email,
			startTime,
			endTime,
			tutor.Email,
			cpy,
			supervisor,
			title,
			ts,
			foreign,
			inLab,
			gratif,
		)
	}
	return err
}

func (f *HTTPFeeder) InjectOnePromotion(s internship.Service, y int, p string) error {
	return f.scan(y, p, s)
}

func (f *HTTPFeeder) InjectConventions(s internship.Service) {
	year := time.Now().Year()
	type empty struct{}
	errors := make([]error, len(f.promotions))
	var wg sync.WaitGroup
	for i, prom := range f.promotions {
		wg.Add(1)
		go func(i int, p string) {
			err := f.InjectOnePromotion(s, year, p)
			errors[i] = err
		}(i, prom)
	}
	wg.Wait()
}

func (f *HTTPFeeder) Test() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", f.url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(f.login, f.password)
	_, err = client.Do(req)
	return err
}
