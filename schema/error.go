package schema

import "errors"

var (
	//ErrUnknownStudent declares the student is unknown
	ErrUnknownStudent = errors.New("Unknown student")
	//ErrUnknownConvention declares the convention is unknown
	ErrUnknownConvention = errors.New("No convention associated to this student")
	//ErrStudentExists declares the student already exists
	ErrStudentExists = errors.New("Student already exists")
	//ErrReportExists declares the report already exists
	ErrReportExists = errors.New("Report already exists")
	//ErrUnknownReport declares the report does not exist
	ErrUnknownReport = errors.New("Unknown report or student")
	//ErrInvalidGrade declares the grade is not between in 0 and 20
	ErrInvalidGrade = errors.New("The grade must be between 0 and 20 (inclusive)")
	//ErrReportConflict declares a report has not been uploaded
	ErrReportConflict = errors.New("The report has not been uploaded")
	//ErrInternshipExists declares the internship already exists
	ErrInternshipExists = errors.New("Internship already exists")
	//ErrUnknownInternship declares the internship does not exists
	ErrUnknownInternship = errors.New("Unknown internship")
	//ErrUserExists declares the user already exists
	ErrUserExists = errors.New("User already exists")
	//ErrUnknownUser declares the user is unknown
	ErrUnknownUser = errors.New("The email does not match a registered user")
	//ErrUserTutoring declares the user cannot be removed as it is tutoring students
	ErrUserTutoring = errors.New("The user is tutoring students")
	//ErrCredentials declares invalid credentials
	ErrCredentials = errors.New("Incorrect password")
	//ErrPasswordTooShort declares the stated password is too short
	ErrPasswordTooShort = errors.New("Password too short (8 chars. min)")
	//ErrNoPendingRequests declares there is no password renewal request
	ErrNoPendingRequests = errors.New("No pending reset request. You might use a bad or an expired reset token.")
	//ErrInvalidPeriod declared the internship period is incorrect
	ErrInvalidPeriod = errors.New("invalid internship period")
	//ErrConventionExists declares a convention for the student already exists
	ErrConventionExists = errors.New("convention already scanned")
	//ErrInvalidMajor declares the declared major is not supported
	ErrInvalidMajor = errors.New("Unknown major")
	//ErrInvalidPromotion declares the promotion is not supported
	ErrInvalidPromotion = errors.New("Unknown promotion")
	//ErrDeadlinePassed declares the deadline for a report passed
	ErrDeadlinePassed = errors.New("Deadline passed")
	//ErrGradedReport declares the report is already graded
	ErrGradedReport = errors.New("Report already graded")
	//ErrSessionExpired declares an expired session
	ErrSessionExpired = errors.New("Session expired")
	//ErrInvalidToken declares an invalid session token
	ErrInvalidToken = errors.New("Invalid session")
	//ErrUnknownSurvey declares the survey does not exist
	ErrUnknownSurvey = errors.New("Unknown survey or student")
	//ErrSurveyUploaded declares the survey has already been uploaded
	ErrSurveyUploaded = errors.New("Survey already fullfilled")
	//ErrInvalidSurvey declares the answers are invalid
	ErrInvalidSurvey = errors.New("Invalid answers")
	//ErrUnknownAlumni declares there is no alumni information for the students
	ErrUnknownAlumni = errors.New("No informations for future alumni")
	//ErrInvalidAlumniEmail declares the email cannot be used for alumni
	ErrInvalidAlumniEmail = errors.New("Invalid email. It must not be served by polytech' or unice")
	//ErrInvalidEmail declares the email is invalid
	ErrInvalidEmail = errors.New("Invalid email")
	//ErrUnknownDefense declares the defense is unknown
	ErrUnknownDefense = errors.New("Unknown defense")

	//ErrDefenseSessionConflit declares a session that is in conflict with another
	ErrDefenseSessionConflit = errors.New("There is already a session for that slot")
	ErrDefenseExists         = errors.New("The defense is already planned")
	ErrDefenseConflict       = errors.New("A defense is already planned for that slot")
	ErrDefenseJuryConflict   = errors.New("The teacher is already in a jury for that period")
)
