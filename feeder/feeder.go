package feeder

//Package feeder contains what is needed to import the conventions
//and list that indicatesthe students that must perform an internship

import (
	"io"
	"strconv"

	"github.com/fhermeni/wints/schema"
)

type Errors []error

func (e *Errors) Error() string {
	return strconv.Itoa(len([]error(e))) + " error(s)"
}

//Conventions allows to import conventions
type Conventions interface {
	Import() ([]schema.Convention, *Errors)
}

//ConventionReader interfaces the source where to read conventions
type ConventionReader interface {
	Reader(year int, promotion string) (io.Reader, error)
}
