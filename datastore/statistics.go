package datastore

import "github.com/fhermeni/wints/internship"

func (srv *Service) Statistics() ([]internship.Stat, error) {

	interns, err := srv.Internships()
	if err != nil {
		return []internship.Stat{}, err
	}
	stats := make([]internship.Stat, len(interns), len(interns))
	i := 0
	for i < len(interns) {
		stats[i] = convert(interns[i])
		i++
	}
	conv, err := srv.Conventions()
	if err != nil {
		return stats, err
	}
	for _, c := range conv {
		if c.Skip {
			stats = append(stats, convertSkipped(c))
		}
	}
	return stats, nil
}

func convertSkipped(c internship.Convention) internship.Stat {
	return internship.Stat{
		Male:           c.Male,
		Creation:       c.Creation,
		Begin:          c.Begin,
		End:            c.End,
		Promotion:      c.Promotion,
		ForeignCountry: c.ForeignCountry,
		Lab:            c.Lab,
		Gratification:  c.Gratification,
		Reports:        make(map[string]int),
		Surveys:        make(map[string]map[string]string),
		Cpy:            c.Cpy,
	}
}

func convert(i internship.Internship) internship.Stat {
	s := internship.Stat{
		Male:           i.Male,
		Creation:       i.Creation,
		Begin:          i.Begin,
		End:            i.End,
		Promotion:      i.Promotion,
		Major:          i.Major,
		Future:         i.Future,
		ForeignCountry: i.ForeignCountry,
		Lab:            i.Lab,
		Gratification:  i.Gratification,
		Reports:        make(map[string]int),
		Surveys:        make(map[string]map[string]string),
		Cpy:            i.Cpy,
	}
	for _, r := range i.Reports {
		if r.ToGrade {
			s.Reports[r.Kind] = r.Grade
		}
	}
	for _, survey := range i.Surveys {
		s.Surveys[survey.Kind] = survey.Answers
	}
	return s
}
