package feeder

//Package feeder contains what is needed to import the conventions
//and list that indicatesthe students that must perform an internship

import (
	"fmt"
	"io"

	"github.com/fhermeni/wints/schema"
)

//ImportError to report a critical error and multiple warnings
type ImportError struct {
	Warnings []string
	Fatal    error
}

//NewImportError creates a new error
func NewImportError() *ImportError {
	return &ImportError{
		Fatal:    nil,
		Warnings: make([]string, 0),
	}
}

//Error reports the critical errors and the possible warnings.
func (e *ImportError) Error() string {
	buf := fmt.Sprintf("%d warning(s)", len(e.Warnings))
	if len(e.Warnings) > 0 {
		buf = buf + ":"
		for _, err := range e.Warnings {
			buf = fmt.Sprintf("%s\n%s", buf, err)
		}
	}
	if e.Fatal != nil {
		buf = fmt.Sprintf("%s\nError: %s", buf, e.Fatal.Error())
	}
	return buf
}

//NewWarning signal a new warning to report
func (e *ImportError) NewWarning(err ...string) {
	e.Warnings = append(e.Warnings, err...)
}

//Conventions allows to import conventions
type Conventions interface {
	Import() ([]schema.Convention, *ImportError)
}

//ConventionReader interfaces the source where to read conventions
type ConventionReader interface {
	Reader(year int, promotion string) (io.Reader, error)
}
