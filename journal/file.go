package journal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fhermeni/wints/internship"
)

var TIME_FMT = "02/01/06 15:04:05"

var ACCESS_LOG = "access.log"
var EVENT_LOG = "event.log"

//File is a structure to store a journal into a file.
//The file is never zeroed.
type File struct {
	path  string
	mutex sync.Mutex
}

//FileBackend make a new file backend from the path provided as argument
func FileBacked(p string) (*File, error) {
	err := os.MkdirAll(p, 0770)
	if err != nil && !os.IsExist(err) {
		return &File{}, err
	}
	return &File{path: p, mutex: sync.Mutex{}}, nil
}

//Log the event into the file.
func (f *File) UserLog(u internship.User, msg string, err error) {
	go func() {
		f.log(EVENT_LOG, "[%s] %s (%s) - %s: %s\n", time.Now().Format(TIME_FMT), u.Email, u.Role.String(), msg, status(err))
	}()
}

func (f *File) Log(em, msg string, err error) {
	go func() {
		f.log(EVENT_LOG, "[%s] %s - %s: %s\n", time.Now().Format(TIME_FMT), em, msg, status(err))
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
		f.log(EVENT_LOG, "[%s] *** WIPE ***\n", time.Now().Format(TIME_FMT))
	}()
}

func (f *File) Access(method, url string, statusCode, latency int) {
	go func() {
		f.log(ACCESS_LOG, "[%s] \"%s %s\" %d %d\n", time.Now().Format(TIME_FMT), method, url, statusCode, latency)
	}()
}
