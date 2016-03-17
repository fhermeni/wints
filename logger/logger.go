//Package logger is a simple package to support rotating log file with multiple files and basic tagging
package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type defaultLogger struct {
	path  string
	mutex sync.Mutex
}

//MultiLogger specifies an interface to log multiple sources than can be streamed
type MultiLogger interface {
	SetRoot(p string) error
	Read(kind string) (io.ReadCloser, error)
	Logs() ([]string, error)
	Log(kind, tag, msg string, err error)
}

var (
	//FileTimeFmt specifies the date format for the log files
	FileTimeFmt = "06-01-02"
	//Default is the default logger
	Default *defaultLogger
)

func init() {
	Default = &defaultLogger{
		path:  "./",
		mutex: sync.Mutex{},
	}
}

//SetRoot set the folder where the log files will be stored
func (m *defaultLogger) SetRoot(p string) error {
	err := os.MkdirAll(p, 0770)
	if err != nil && !os.IsExist(err) {
		return err
	}
	m.path = p
	return nil

}

//Read stream the required log file
func (m *defaultLogger) Read(kind string) (io.ReadCloser, error) {
	var path = filepath.Join(m.path, kind)
	return os.OpenFile(path, os.O_RDONLY, 0660)
}

//Logs returns all the files in the log folder
func (m *defaultLogger) Logs() ([]string, error) {
	files, err := ioutil.ReadDir(m.path)
	if err != nil {
		return []string{}, err
	}
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name()
	}
	return names, err
}

//Log in the 'kind' file a tagged message
func (m *defaultLogger) Log(kind, tag, msg string, e error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	now := time.Now()

	path := filepath.Join(m.path, now.Format(FileTimeFmt)+"-"+kind)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
	}
	defer f.Close()

	status := ""
	if e != nil {
		status = ": " + e.Error()
	}
	out := log.New(f, "", log.LstdFlags)
	out.Printf("%s - %s %s\n", tag, msg, status)
}

//SetRoot set the folder where the log files will be stored
func SetRoot(p string) error {
	return Default.SetRoot(p)
}

//Log in the 'kind' file a tagged message
func Log(kind, tag, msg string, err error) {
	Default.Log(kind, tag, msg, err)
}

//Logs returns all the files in the log folder
func Logs() ([]string, error) {
	return Default.Logs()
}

//Read stream the required log file
func Read(kind string) (io.ReadCloser, error) {
	return Default.Read(kind)
}
