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

var (
	//FileTimeFmt specifies the date format for the log files
	FileTimeFmt = "06-01-02"
	path        string
	mutex       sync.Mutex
)

//Init initiate the logger that will store the files in the given path
func Init(p string) error {
	err := os.MkdirAll(p, 0770)
	if err != nil && !os.IsExist(err) {
		return err
	}
	path = p
	return nil
}

//Log in the 'kind' file a tagged message
func Log(kind, tag, msg string, err error) {
	mutex.Lock()
	defer mutex.Unlock()
	now := time.Now()

	path := filepath.Join(path, kind+"-"+now.Format(FileTimeFmt))

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
	}
	defer f.Close()

	status := ""
	if err != nil {
		status = ": " + err.Error()
	}
	out := log.New(f, "", log.LstdFlags)
	out.Printf("%s - %s %s\n", tag, msg, status)
}

//Logs returns all the files in the log folder
func Logs() ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name()
	}
	return names, err
}

//StreamLog stream the required log file
func StreamLog(kind string) (io.ReadCloser, error) {
	var path = filepath.Join(path, kind)
	return os.OpenFile(path, os.O_RDONLY, 0660)
}
