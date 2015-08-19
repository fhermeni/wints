package datastore

/*
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
	st :=
		internship.Stat{
			Male:           c.Male,
			Creation:       c.Creation,
			Begin:          c.Begin,
			End:            c.End,
			Promotion:      c.Promotion,
			ForeignCountry: c.ForeignCountry,
			Lab:            c.Lab,
			Gratification:  c.Gratification,
			Reports:        make([]internship.ReportHeader, 0, 0),
			Surveys:        make([]internship.Survey, 0, 0),
			Cpy:            c.Cpy,
		}
	return st
}

func anonReport(r internship.ReportHeader) internship.ReportHeader {
	//hide the comments
	return internship.ReportHeader{
		Private:  r.Private,
		Grade:    r.Grade,
		Kind:     r.Kind,
		Deadline: r.Deadline,
		Delivery: r.Delivery,
		Reviewed: r.Reviewed,
		ToGrade:  r.ToGrade,
	}
}

func anonSurvey(s internship.Survey) internship.Survey {
	//hide the comments
	return internship.Survey{
		Kind:      s.Kind,
		Deadline:  s.Deadline,
		Timestamp: s.Timestamp,
		Token:     "",
		Answers:   s.Answers,
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
		Reports:        make([]internship.ReportHeader, 0, 0),
		Surveys:        make([]internship.Survey, 0, 0),
		Cpy:            i.Cpy,
	}
	for _, r := range i.Reports {
		s.Reports = append(s.Reports, anonReport(r))
	}
	for _, survey := range i.Surveys {
		s.Surveys = append(s.Surveys, anonSurvey(survey))
	}
	return s
}
*/
