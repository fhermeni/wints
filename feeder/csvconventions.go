package feeder

import (
	"encoding/csv"
	"io"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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

var (
	//Integer is a regex to express the integer part for a possible real number.
	//We just ignore the floating part
	Integer = regexp.MustCompile(`^(\d+)`)
)

//CsvConventions parses conventions from CSV files
type CsvConventions struct {
	j          journal.Journal
	Reader     ConventionReader
	Year       int
	promotions []string
}

//NewCsvConventions creates an importer from a given reader
//By default, the parsed year is the current year
func NewCsvConventions(r ConventionReader, j journal.Journal, promotions []string) *CsvConventions {
	return &CsvConventions{
		Reader:     r,
		j:          j,
		promotions: promotions,
		Year:       time.Now().Year()}
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

func (f *CsvConventions) log(msg string, err error) {
	f.j.Log("feeder", msg, err)
}

func cleanInt(str string) int {
	s := strings.Replace(str, " ", "", -1)
	m := Integer.FindSubmatch([]byte(s))
	if m == nil || len(m) == 0 {
		return -1
	}
	i, err := strconv.Atoi(string(m[0]))
	if err != nil {
		return -1
	}
	return i
}
func (f *CsvConventions) scan(prom string, s internship.Service) error {
	r, err := f.Reader.Reader(f.Year, prom)
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

//Import imports all the conventions by requesting in parallal the conventions for each registered promotions
func (f *CsvConventions) Import(s internship.Service) error {
	errors := make([]error, len(f.promotions))
	var wg sync.WaitGroup
	for i, prom := range f.promotions {
		wg.Add(1)
		go func(i int, p string) {
			err := f.scan(p, s)
			errors[i] = err
		}(i, prom)
	}
	wg.Wait()
	return nil
}
