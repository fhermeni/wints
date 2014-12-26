package filter

import (
	"errors"

	"github.com/fhermeni/wints/internship"
)

var (
	//ErrPermission indicates an operation that is forbidden by the current user
	ErrPermission = errors.New("Permission denied")
)

/*
Service restricts the operation that can be executed by the current user with regards
to its role and or relationships
*/
type Service struct {
	my  internship.User
	srv internship.Service
}
