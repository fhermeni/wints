//Package journal allows to store events related to the internship management.
package journal

import "github.com/fhermeni/wints/internship"

//Journal is the interface that allows to store an event into the journal
type Journal interface {
	//Log store an event emitted by a user. The event is described in a message
	//while err denotes the event status
	UserLog(u internship.User, msg string, err error)

	Log(em, msg string, err error)
	//Wipe indicates the database has been reseted
	Wipe()
}
