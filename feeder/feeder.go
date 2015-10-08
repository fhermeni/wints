package feeder

//Package feeder contains what is needed to import the conventions
//and list that indicatesthe students that must perform an internship

import (
	"io"

	"github.com/fhermeni/wints/schema"
)

//Importer imports data into the system
type Conventions interface {
	Import() ([]schema.Convention, error)
}

//ConventionReader interfaces the source where to read conventions
type ConventionReader interface {
	Reader(year int, promotion string) (io.Reader, error)
}
