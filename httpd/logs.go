package httpd

import (
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/session"
)

func streamLog(ex Exchange) error {
	kind := ex.V("k")
	if ex.s.Me().Role.Level() < schema.AdminLevel {
		return session.ErrPermission
	}
	in, err := ex.not.Log.StreamLog(kind)
	return ex.out("text/plain", in, err)
}

func logs(ex Exchange) error {
	if ex.s.Me().Role.Level() < schema.AdminLevel {
		return session.ErrPermission
	}
	in, err := ex.not.Log.Logs()
	return ex.outJSON(in, err)
}
