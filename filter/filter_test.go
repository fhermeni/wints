package filter

import (
	"testing"

	"github.com/fhermeni/wints/internship"
	"github.com/stretchr/testify/assert"
)

func TestSetUserPassword(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto"}
	f, _ := NewService(&mock, u)
	assert.NoError(t, f.SetUserPassword("toto", []byte{}, []byte{}))
	assert.Equal(t, ErrPermission, f.SetUserPassword("titi", []byte{}, []byte{}))
}

func TestSetUserPassword(t *testing.T) {
	var mock MockBackend
	u := internship.User{Email: "toto"}
	f, _ := NewService(&mock, u)
	assert.NoError(t, f.SetUserPassword("toto", []byte{}, []byte{}))
	assert.Equal(t, ErrPermission, f.SetUserPassword("titi", []byte{}, []byte{}))
}
