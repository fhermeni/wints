package spy

import "github.com/fhermeni/wints/sqlstore"

type Spy interface {
	Spy(s *sqlstore.Store)
}
