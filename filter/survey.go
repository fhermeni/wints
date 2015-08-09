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
	err := v.SetSurveyContent(token, cnt)
	v.log.Log(token, "uploaded the survey", err)
	if err == nil {
		stu, kind, err2 := v.srv.SurveyToken(token)
		if err2 != nil {
			v.log.Log(token, "Unable to mail about the uploaded survey", err2)
			return err
		}
		u, err2 := v.srv.User(stu)
		if err2 != nil {
			v.log.Log(token, "Unable to mail about the uploaded survey", err2)
			return err
		}
		v.mailer.SendSurveyUploaded(v.my, u, kind)
	}
	return err
}

func (v *Service) SurveyDefs() []internship.SurveyDef {
	return v.srv.SurveyDefs()
}

func (s *Service) RequestSurvey(stu, kind string) error {
	if s.my.Role != internship.ROOT {
		return ErrPermission
	}
	return s.srv.RequestSurvey(stu, kind)
}