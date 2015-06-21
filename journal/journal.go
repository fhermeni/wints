//Package journal allows to store events related to the internship management.
package journal

import "github.com/fhermeni/wints/internship"

//Journal is the interface that allows to store an event into the journal
type Journal interface {

	//UserLog stores an event emitted by a logged user. The event is described in a message
	//while err denotes the event status
	UserLog(u internship.User, msg string, err error)

	//Log stores an event emitted by a non-logged user
	Log(em, msg string, err error)

	//Wipe indicates the database has been reseted
	Wipe()
}
