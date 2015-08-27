package journal

/*
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fhermeni/wints/internship"
)

var (
	//TimeFmt indicates the output format for timestamps
	TimeFmt = "02/01/06 15:04:05"
	//AccessLog is the filename where the rest call are stored
	AccessLog = "access.log"
	//EventLog is the file containing wints events
	EventLog = "event.log"
)

//File is a structure to store a journal into a file.
//The file is never zeroed.
type File struct {
	path  string
	mutex sync.Mutex
}

//FileBacked makes a new file backend from the path provided as argument
func FileBacked(p string) (*File, error) {
	err := os.MkdirAll(p, 0770)
	if err != nil && !os.IsExist(err) {
		return &File{}, err
	}
	return &File{path: p, mutex: sync.Mutex{}}, nil
}

//UserLog stores the event emitted by a logged user into the file.
func (f *File) UserLog(u internship.User, msg string, err error) {
	go func() {
		f.log(EventLog, "[%s] %s (%s) - %s: %s\n", time.Now().Format(TimeFmt), u.Person.Email, u.Role.String(), msg, status(err))
	}()
}

func (f *File) Log(em, msg string, err error) {
	go func() {
		f.log(EventLog, "[%s] %s - %s: %s\n", time.Now().Format(TimeFmt), em, msg, status(err))
	}()
}

func status(err error) string {
	if err == nil {
		return "OK"
	}
	return err.Error()
}

func (f *File) log(fn, format string, args ...interface{}) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	var path = filepath.Join(f.path, fn)
	fi, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		log.Printf("Unable to open '%s': %s\n", path, err.Error())
		return
	}
	defer fi.Close()
	fmt.Fprintf(fi, format, args...)

}
func (f *File) Wipe() {
	go func() {
		f.log(EventLog, "[%s] *** WIPE ***\n", time.Now().Format(TimeFmt))
	}()
}

func (f *File) Access(method, url string, statusCode, latency int) {
	go func() {
		f.log(AccessLog, "[%s] \"%s %s\" %d %d\n", time.Now().Format(TimeFmt), method, url, statusCode, latency)
	}()
}
*/
