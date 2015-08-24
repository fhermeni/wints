package feeder

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/fhermeni/wints/internship"
)

//InsertStudents inserts all the students provided in the given CSV file.
//Expected format: firstname;lastname;email;tel
//The other fields are set to the default values
func InjectStudents(file string, s internship.Service) error {
	in := csv.NewReader(strings.NewReader(file))
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
