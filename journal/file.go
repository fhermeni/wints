package journal

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fhermeni/wints/internship"
)

var TIME_FMT = "02/01/06 15:04:05"

//File is a structure to store a journal into a file
type File struct {
	path  string
	mutex sync.Mutex
}

//FileBackend make a new file backend from the path provided as argument
func FileBacked(p string) *File {
	return &File{path: p, mutex: sync.Mutex{}}
}

//Log the event into the file.
func (f *File) UserLog(u internship.User, msg string, err error) {
	go func() {
		st := "OK"
		if err != nil {
			st = err.Error()
		}
		f.log("[%s] %s (%s) - %s: %s\n", time.Now().Format(TIME_FMT), u.Email, u.Role.String(), msg, st)
	}()
}

func (f *File) Log(em, msg string, err error) {
	go func() {
		st := "OK"
		if err != nil {
			st = err.Error()
		}
		f.log("[%s] %s - %s: %s\n", time.Now().Format(TIME_FMT), em, msg, st)
	}()
}

func (f *File) log(format string, args ...interface{}) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	fi, err := os.OpenFile(f.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		log.Printf("Unable to open journal '%s': %s\n", f.path, err.Error())
		return
	}
	defer fi.Close()
	fmt.Fprintf(fi, format, args...)

}
func (f *File) Wipe() {
	go func() {
		f.log("[%s] *** WIPE ***\n", time.Now().Format(TIME_FMT))
	}()
}
