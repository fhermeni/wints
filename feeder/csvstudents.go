package feeder

import (
	"encoding/csv"
	"io"

	"github.com/fhermeni/wints/internship"
)

//CsvStudents embed the parsing of students from a CSV file
type CsvStudents struct {
	reader StudentReader
}

//NewCsvStudents creates a new parser from a given source
func NewCsvStudents(r StudentReader) *CsvStudents {
	return &CsvStudents{reader: r}
}

//Import inserts all the students provided in the given CSV file.
//Expected format: firstname;lastname;email;tel;
//The other fields are set to the default values
func (c *CsvStudents) Import(s internship.Adder) error {
	r, err := c.reader.Reader()
	if err != nil {
		return err
	}
	in := csv.NewReader(r)
	in.Comma = ';'
	//Get rid of the header
	in.Read()
	for {
		record, e2 := in.Read()
		if e2 == io.EOF {
			break
		}
		ln := record[0]
		fn := record[1]
		email := record[2]
		p := record[3]
		major := record[4]

		if err := s.NewUser(fn, ln, "not available", email); err == nil {
			err = s.ToStudent(email, major, p, true) //male by default.Sexist isn't it ? but no choice at this stage
		}
	}
	return nil
}
