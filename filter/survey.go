package filter

import "github.com/fhermeni/wints/internship"

func (v *Service) SurveyToken(kind string) (string, string, error) {
	return v.srv.SurveyToken(kind)
}

func (v *Service) Survey(student, kind string) (internship.Survey, error) {
	survey := internship.Survey{}
	err := ErrPermission
	if v.my.Role >= internship.MAJOR || v.isTutoring(student) {
		survey, err = v.srv.Survey(student, kind)
	}
	v.UserLog("get '"+kind+"' survey of '"+student+"'", err)
	return survey, err
}

func (v *Service) SetSurveyContent(token string, cnt map[string]string) error {
	return v.SetSurveyContent(token, cnt)
}

func (v *Service) SurveyDefs() []internship.SurveyDef {
	return v.srv.SurveyDefs()
}
