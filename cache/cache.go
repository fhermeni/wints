package cache

import (
	"time"

	"github.com/fhermeni/wints/internship"
)

type Cache struct {
	hot         AtomicBool
	backend     internship.Service
	internships []internship.Internship
}

func NewCache(backend internship.Service) (*Cache, error) {
	b := AtomicBool{}
	b.Set(false)
	return &Cache{hot: b, backend: backend}, nil
}

//Things we cache
func (cache *Cache) Internships() ([]internship.Internship, error) {
	var err error
	err = nil
	if !cache.hot.Get() {
		cache.internships, err = cache.backend.Internships()
		if err != nil {
			cache.hot.Set(false)
		} else {
			cache.hot.Set(true)
		}
	}
	return cache.internships, err
}

//Get the internship associated to a given student
func (cache *Cache) Internship(stu string) (internship.Internship, error) {
	var err error
	err = nil
	if !cache.hot.Get() {
		return cache.backend.Internship(stu)
	}
	for _, i := range cache.internships {
		if i.Student.Email == stu {
			return i, err
		}
	}
	return internship.Internship{}, err
}

//Things we don't cache
func (cache *Cache) NewInternship(c internship.Convention) ([]byte, error) {
	cache.hot.Set(false)
	return cache.backend.NewInternship(c)
}

//Set the supervisor for the internship of a given student
func (cache *Cache) SetSupervisor(stu string, sup internship.Person) error {
	cache.hot.Set(false)
	return cache.backend.SetSupervisor(stu, sup)
}

//Set the tutor for the internship of a given student
func (cache *Cache) SetTutor(stu string, t string) error {
	cache.hot.Set(false)
	return cache.backend.SetTutor(stu, t)
}

//Set the company for the internship of a given student
func (cache *Cache) SetCompany(stu string, c internship.Company) error {
	cache.hot.Set(false)
	return cache.backend.SetCompany(stu, c)
}

//Set the internship title
func (cache *Cache) SetTitle(stu string, title string) error {
	cache.hot.Set(false)
	return cache.backend.SetTitle(stu, title)
}

//Set the student major
func (cache *Cache) SetMajor(stu string, m string) error {
	cache.hot.Set(false)
	return cache.backend.SetMajor(stu, m)
}

//Get the possible majors
func (cache *Cache) Majors() []string {
	return cache.backend.Majors()
}

//Set the student promotion
func (cache *Cache) SetPromotion(stu string, p string) error {
	cache.hot.Set(false)
	return cache.backend.SetPromotion(stu, p)
}

//Test if the credentials match a user, return a session token
func (cache *Cache) Registered(email string, password []byte) ([]byte, error) {
	return cache.backend.Registered(email, password)
}

//Check if a session is opened for a given user and token
func (cache *Cache) OpenedSession(email, token string) error {
	return cache.backend.OpenedSession(email, token)
}

//Destroy a session
func (cache *Cache) Logout(email, token string) error {
	return cache.backend.Logout(email, token)
}

func (cache *Cache) Sessions() (map[string]time.Time, error) {
	cache.hot.Set(false)
	return cache.backend.Sessions()
}

//Create a new user
//Returns the resulting password
func (cache *Cache) NewTutor(p internship.User) ([]byte, error) {
	cache.hot.Set(false)
	return cache.backend.NewTutor(p)
}

//Delete the user if it is not tutoring anyone
func (cache *Cache) RmUser(email string) error {
	cache.hot.Set(false)
	return cache.backend.RmUser(email)
}

//Get the user
func (cache *Cache) User(email string) (internship.User, error) {
	return cache.backend.User(email)
}

//List the users
func (cache *Cache) Users() ([]internship.User, error) {
	return cache.backend.Users()
}

//Change the user password
func (cache *Cache) SetUserPassword(email string, oldP, newP []byte) error {
	cache.hot.Set(false)
	return cache.backend.SetUserPassword(email, oldP, newP)
}

//Change user profile
func (cache *Cache) SetUserProfile(email, fn, ln, tel string) error {
	cache.hot.Set(false)
	return cache.backend.SetUserProfile(email, fn, ln, tel)
}

//Change user role
func (cache *Cache) SetUserRole(email string, priv internship.Privilege) error {
	cache.hot.Set(false)
	return cache.backend.SetUserRole(email, priv)
}

//Ask for a password reset.
//Return a token used to declare the new password (see NewPassword)
func (cache *Cache) ResetPassword(email string) ([]byte, error) {
	cache.hot.Set(false)
	return cache.backend.ResetPassword(email)
}

//Declare the new password using an authentication token
func (cache *Cache) NewPassword(token, newP []byte) (string, error) {
	cache.hot.Set(false)
	return cache.backend.NewPassword(token, newP)
}

func (cache *Cache) PlanReport(student string, r internship.ReportHeader) error {
	cache.hot.Set(false)
	return cache.backend.PlanReport(student, r)
}
func (cache *Cache) ReportDefs() []internship.ReportDef {
	return cache.backend.ReportDefs()
}
func (cache *Cache) Report(kind, email string) (internship.ReportHeader, error) {
	return cache.backend.Report(kind, email)
}
func (cache *Cache) ReportContent(kind, email string) ([]byte, error) {
	return cache.backend.ReportContent(kind, email)
}
func (cache *Cache) SetReportContent(kind, email string, cnt []byte) error {
	cache.hot.Set(false)
	return cache.backend.SetReportContent(kind, email, cnt)
}
func (cache *Cache) SetReportGrade(kind, email string, r int, comment string) error {
	cache.hot.Set(false)
	return cache.backend.SetReportGrade(kind, email, r, comment)
}
func (cache *Cache) SetReportDeadline(kind, email string, t time.Time) error {
	cache.hot.Set(false)
	return cache.backend.SetReportDeadline(kind, email, t)
}
func (cache *Cache) SetReportPrivate(kind, email string, p bool) error {
	cache.hot.Set(false)
	return cache.backend.SetReportPrivate(kind, email, p)
}

//Survey management
func (cache *Cache) SurveyToken(kind string) (string, string, error) {
	return cache.backend.SurveyToken(kind)
}
func (cache *Cache) Survey(student, kind string) (internship.Survey, error) {
	return cache.backend.Survey(student, kind)
}
func (cache *Cache) SetSurveyContent(token string, cnt map[string]string) error {
	cache.hot.Set(false)
	return cache.backend.SetSurveyContent(token, cnt)
}
func (cache *Cache) SurveyDefs() []internship.SurveyDef {
	return cache.backend.SurveyDefs()
}

func (cache *Cache) NewConvention(c internship.Convention) error {
	cache.hot.Set(false)
	return cache.backend.NewConvention(c)
}
func (cache *Cache) Conventions() ([]internship.Convention, error) {
	return cache.backend.Conventions()
}
func (cache *Cache) SkipConvention(student string, skip bool) error {
	cache.hot.Set(false)
	return cache.backend.SkipConvention(student, skip)
}
func (cache *Cache) DeleteConvention(student string) error {
	cache.hot.Set(false)
	return cache.backend.DeleteConvention(student)
}
func (cache *Cache) SetAlumni(student string, a internship.Alumni) error {
	cache.hot.Set(false)
	return cache.backend.SetAlumni(student, a)
}

//public statistics
func (cache *Cache) Statistics() ([]internship.Stat, error) {
	return cache.backend.Statistics()
}

//Students that have to make an internship
func (cache *Cache) Students() ([]internship.Student, error) {
	return cache.backend.Students()

}
func (cache *Cache) InsertStudents(csv string) error {
	cache.hot.Set(false)
	return cache.backend.InsertStudents(csv)

}
func (cache *Cache) AddStudent(s internship.Student) error {
	cache.hot.Set(false)
	return cache.backend.AddStudent(s)
}
func (cache *Cache) AlignWithInternship(student string, internship string) error {
	cache.hot.Set(false)
	return cache.backend.AlignWithInternship(student, internship)
}
func (cache *Cache) HideStudent(em string, st bool) error {
	cache.hot.Set(false)
	return cache.backend.HideStudent(em, st)
}

func (cache *Cache) DefenseSessions() ([]internship.DefenseSession, error) {
	return cache.backend.DefenseSessions()
}

func (cache *Cache) PublicDefenseSessions() ([]internship.PublicDefenseSession, error) {
	return cache.backend.PublicDefenseSessions()
}

func (cache *Cache) DefenseSession(student string) (internship.DefenseSession, error) {
	return cache.backend.DefenseSession(student)
}

func (cache *Cache) SetDefenseGrade(student string, g int) error {
	cache.hot.Set(false)
	return cache.backend.SetDefenseGrade(student, g)
}

func (cache *Cache) SetDefenseSessions(defs []internship.DefenseSession) error {
	cache.hot.Set(false)
	return cache.backend.SetDefenseSessions(defs)
}
