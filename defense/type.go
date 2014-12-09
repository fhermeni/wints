package defense;


import (
	"errors"
)

var (
	ErrUnknown = errors.New("Unknown defense")
	ErrExists = errors.New("The defense already exists")
)

type DefenseService interface {
	Get(id string) (string, error)
	New(id, cnt string) error
	Long(id string) (string, error)
	Set(id, short, long string) error
}
