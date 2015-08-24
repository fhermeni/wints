package internship

import "time"

//ReportHeader provides the metadata associated to a student report
type ReportHeader struct {
	Kind string
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

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) Graded() bool {
	return m.Reviewed != nil
}

//Graded indicates if a report has been graded by its tutor or not
func (m *ReportHeader) In() bool {
	return m.Delivery != nil
}
