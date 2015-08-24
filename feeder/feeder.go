package feeder

//Package feeder contains what is needed to import the conventions
//and list that indicatesthe students that must perform an internship

import (
	"io"

	"github.com/fhermeni/wints/internship"
)

//Importer imports data into the system
type Importer interface {
	Import(s internship.Adder) error
}

//ConventionReader interfaces the source where to read conventions
type ConventionReader interface {
	Reader(year int, promotion string) (io.Reader, error)
}

//StudentReader interfaces the source where to extract students
type StudentReader interface {
	Reader() (io.Reader, error)
}
