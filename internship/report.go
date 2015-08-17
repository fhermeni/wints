package internship

import (
	"errors"
	"time"
)

//ReportHeader provides the metadata associated to a student report
type ReportHeader struct {
	//The report deadline
	Deadline time.Time
	Delivery *time.Time `,json:"omitempty"`
	Reviewed *time.Time `,json:"omitempty"`
	Grade    int
	//The tutor comment
	Comment string
	Private bool
	ToGrade bool
}

//ReportsDef allows to define internship reports
type ReportDef struct {
	Name     string
	Deadline string
	Value    string
	ToGrade  bool
}

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) Graded() bool {
	return m.Reviewed != nil
}

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) In() bool {
	return m.Delivery != nil
}

//ReportHeader allows to instantiate a report definition to a header
func (def *ReportDef) Instantiate(from time.Time) (ReportHeader, error) {
	switch def.Deadline {
	case "relative":
		shift, err := time.ParseDuration(def.Value)
		if err != nil {
			return ReportHeader{}, err
		}
		return ReportHeader{Deadline: from.Add(shift), ToGrade: def.ToGrade}, nil
	case "absolute":
		at, err := time.Parse(DateLayout, def.Value)
		if err != nil {
			return ReportHeader{}, err
		}
		return ReportHeader{Deadline: at, ToGrade: def.ToGrade}, nil
	}
	return ReportHeader{}, errors.New("Unsupported time definition '" + def.Deadline + "'")
}
