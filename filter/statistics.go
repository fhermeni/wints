package filter

import "github.com/fhermeni/wints/internship"

func (v *Service) Statistics() ([]internship.Stat, error) {
	return v.srv.Statistics()
}
