package schema

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

//StudentReports aggregate the reports of a given student
type StudentReports struct {
	Student Student
	Reports []ReportHeader
}

//Graded indicates if a report has been graded
func (m *ReportHeader) Graded() bool {
	return m.Reviewed != nil
}

//In indicates if a report has been uploaded
func (m *ReportHeader) In() bool {
	return m.Delivery != nil
}

//Anonymise the report comment
func (m *ReportHeader) Anonymise() {
	m.Comment = ""
}
