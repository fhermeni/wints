package datastore

import "github.com/fhermeni/wints/internship"

func (s *Service) Survey(kind string) (internship.SurveyHeader, error) {
	s.DB.Query("select path from survey where kind=$1", kind)
}
