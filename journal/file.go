package journal

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fhermeni/wints/schema"
)

var (
	//TimeFmt indicates the output format for timestamps
	TimeFmt = "02/01/06 15:04:05"
	//FileTimeFmt specifies the date format for the log files
	FileTimeFmt = "06-01-02"
	//AccessLog is the filename where the rest call are stored
	AccessLog = "access"
	//EventLog is the file containing wints events
	EventLog = "event"
)

func fileName(key string, t time.Time) string {
	return fmt.Sprintf("%s-%s.log", key, t.Format(FileTimeFmt))
}

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
func (f *File) UserLog(u schema.User, msg string, err error) {
	now := time.Now()
	go func() {
		f.log(fileName(EventLog, now), "[%s] %s (%s) - %s: %s\n", now.Format(TimeFmt), u.Person.Email, u.Role, msg, status(err))
	}()
}

//Log a message with a possible error
func (f *File) Log(em, msg string, err error) {
	now := time.Now()
	go func() {
		f.log(fileName(EventLog, now), "[%s] %s - %s: %s\n", now.Format(TimeFmt), em, msg, status(err))
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

//Wipe signals the application has been restarted
func (f *File) Wipe() {
	now := time.Now()
	go func() {
		f.log(fileName(EventLog, now), "[%s] *** WIPE ***\n", now.Format(TimeFmt))
	}()
}

//Access logs an access to a file through HTTP
func (f *File) Access(method, url string, statusCode, latency int) {
	now := time.Now()
	go func() {
		f.log(fileName(AccessLog, now), "[%s] \"%s %s\" %d %d\n", now.Format(TimeFmt), method, url, statusCode, latency)
	}()
}

func (f *File) StreamLog(kind string) (io.ReadCloser, error) {
	var path = filepath.Join(f.path, kind) + ".log"
	return os.OpenFile(path, os.O_RDONLY, 0660)
}

func (f *File) Logs() ([]string, error) {
	files, err := ioutil.ReadDir(f.path)
	if err != nil {
		return []string{}, err
	}
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = strings.TrimSuffix(f.Name(), ".log")
	}
	return names, err
}
