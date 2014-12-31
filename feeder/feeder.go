package feeder

import "github.com/fhermeni/wints/internship"

type Feeder interface {
	InjectConventions(s internship.Service) (int, error)
}
