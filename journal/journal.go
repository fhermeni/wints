//Package journal allows to store events related to the internship management.
package journal

import (
	"io"

	"github.com/fhermeni/wints/schema"
)

//Journal is the interface that allows to store an event into the journal
type Journal interface {

	//UserLog stores an event emitted by a logged user. The event is described in a message
	//while err denotes the event status
	UserLog(u schema.User, msg string, err error)

	//Log stores an event emitted by a non-logged user
	Log(em, msg string, err error)

	//Wipe indicates the database has been reseted
	Wipe()

	//Access notify an access to a Rest endPoint. The latency is in milliseconds
	Access(method, url string, statusCode, latency int)

	//StreamLog stream a log content
	StreamLog(kind string) (io.ReadCloser, error)

	//Logs returns all the available log files
	Logs() ([]string, error)
}
