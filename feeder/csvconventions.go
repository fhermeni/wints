package feeder

import (
	"encoding/csv"
	"io"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/schema"
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
	Reader     ConventionReader
	Year       int
	promotions []string
}

//NewCsvConventions creates an importer from a given reader
//By default, the parsed year is the current year
func NewCsvConventions(r ConventionReader, promotions []string) *CsvConventions {
	return &CsvConventions{
		Reader:     r,
		promotions: promotions,
		Year:       time.Now().Year()}
}

func cleanPerson(fn, ln, email, tel string) schema.Person {
	return schema.Person{
		Firstname: clean(fn),
		Lastname:  clean(ln),
		Email:     clean(email),
		Tel:       clean(tel),
	}
}

func cleanCompany(name, WWW, title string) schema.Company {
	if len(WWW) != 0 && !strings.HasPrefix(WWW, "http") {
		WWW = "http://" + WWW
	}
	return schema.Company{
		Name:  clean(name),
		WWW:   clean(WWW),
		Title: clean(title),
	}
}

func clean(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
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
func parseTime(fmt, buf string, student schema.Student) (time.Time, error) {
	return time.Parse(fmt, clean(buf))
}
func (f *CsvConventions) scan(prom string) ([]schema.Convention, *ImportError) {
	ierr := NewImportError()
	var conventions []schema.Convention
	lasts := make(map[string]schema.Convention)
	r, err := f.Reader.Reader(f.Year, prom)
	if err != nil {
		ierr.NewWarning("(" + prom + ") : " + err.Error())
		return conventions, ierr
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
			ierr.NewWarning("(" + prom + ") : " + err.Error())
			return conventions, ierr
		}
		student := cleanPerson(record[stuFn], record[stuLn], record[stuEmail], record[stuTel])
		supervisor := cleanPerson(record[supervisorFn], record[supervisorLn], record[supervisorEmail], record[supervisorTel])
		tutor := cleanPerson(record[tutorFn], record[tutorLn], record[tutorEmail], record[tutorTel])
		cpy := cleanCompany(record[company], record[companyWWW], record[titleIdx])
		foreign := clean(record[foreignCountry]) != "non"
		inLab := clean(record[lab]) != "non"
		male := clean(record[gender]) == "m."
		gratif := cleanInt(record[gratification])

		stu := schema.Student{
			User: schema.User{
				Person: student,
				Role:   schema.STUDENT,
			},
			Promotion: prom,
			Skip:      false,
			Male:      male,
		}
		ts, err := parseTime("2006-01-02 15:04", record[timestamp], stu)
		if err != nil {
			ierr.NewWarning("(" + prom + ") " + stu.User.Fullname() + ": " + err.Error())
			continue
		}
		startTime, err := parseTime("02/01/2006", record[begin], stu)
		if err != nil {
			ierr.NewWarning("(" + prom + ") " + stu.User.Fullname() + ": " + err.Error())
			continue
		}
		endTime, err := parseTime("02/01/2006", record[end], stu)
		if err != nil {
			ierr.NewWarning("(" + prom + ") " + stu.User.Fullname() + ": " + err.Error())
			continue
		}
		//The to-valid users
		c := schema.Convention{
			Creation:       ts,
			Student:        stu,
			Gratification:  gratif,
			Lab:            inLab,
			ForeignCountry: foreign,
			Begin:          startTime,
			End:            endTime,
			Supervisor:     supervisor,
			Company:        cpy,
			Tutor: schema.User{
				Person: tutor,
				Role:   schema.TUTOR,
			},
		}
		lasts[c.Student.User.Person.Email] = c
	}
	for _, c := range lasts {
		conventions = append(conventions, c)
	}
	return conventions, ierr
}

//Import imports all the conventions by requesting in parallel the conventions for each registered promotions
func (f *CsvConventions) Import() ([]schema.Convention, *ImportError) {
	lock := &sync.Mutex{}
	var all []schema.Convention
	var wg sync.WaitGroup
	ierr := NewImportError()
	for i, prom := range f.promotions {
		wg.Add(1)
		go func(i int, p string) {
			convs, err := f.scan(p)
			if err != nil {
				ierr.NewWarning(err.Warnings...)
			}
			lock.Lock()
			defer lock.Unlock()
			all = append(all, convs...)
			wg.Done()
		}(i, prom)
	}
	wg.Wait()
	logger.Log("event", "feeder", "Scanning conventions", ierr)
	return all, ierr
}
