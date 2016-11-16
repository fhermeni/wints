//Package logger is a simple package to support rotating log file with multiple files and basic tagging
package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/stathat/go"
)

type defaultLogger struct {
	path  string
	mutex sync.Mutex
	key   string
}

//MultiLogger specifies an interface to log multiple sources than can be streamed
type MultiLogger interface {
	SetRoot(p string) error
	Read(kind string) (io.ReadCloser, error)
	Logs() ([]string, error)
	Log(kind, tag, msg string, err error)
	ReportHit(lbl string)
	ReportValue(lbl string, v int)
	Trace(key string)
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
		key:   "",
	}
}

func (m *defaultLogger) Trace(key string) {
	m.key = key
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

type byDate []os.FileInfo

func (f byDate) Len() int {
	return len(f)
}
func (f byDate) Less(i, j int) bool {
	return f[i].ModTime().Unix() < f[j].ModTime().Unix()
}
func (f byDate) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

//Logs returns all the files in the log folder
func (m *defaultLogger) Logs() ([]string, error) {
	files, err := ioutil.ReadDir(m.path)
	if err != nil {
		return []string{}, err
	}
	sort.Sort(byDate(files))
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

func (m *defaultLogger) ReportHit(label string) {
	if len(m.key) > 0 {
		stathat.PostEZCount(label, m.key, 1)
	} else {
		m.Log("trace", "hit", label, nil)
	}
}

func (m *defaultLogger) ReportValue(label string, v float64) {
	if len(m.key) > 0 {
		stathat.PostEZValue(label, m.key, v)
	} else {
		s := fmt.Sprintf("%s : %f", label, v)
		m.Log("trace", "value", s, nil)
	}
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

func ReportHit(label string) {
	Default.ReportHit(label)
}

func ReportValue(label string, v float64) {
	Default.ReportValue(label, v)
}

func Trace(key string) {
	Default.Trace(key)
}
