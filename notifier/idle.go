package notifier

import "github.com/fhermeni/wints/schema"

//IdleNotifier defines how to report a user that never connected
type IdleNotifier interface {
	//ReportIdleAccount report to a user that he never connected
	ReportIdleAccount(u schema.User)
}
