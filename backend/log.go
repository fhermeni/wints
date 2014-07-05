package backend

import (
	"log"
	"time"
	"os"
	"bufio"
	"fmt"
)

var fd *os.File
var writer *bufio.Writer

func InitLogger(path string) error {
	fd, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0640)
	if err != nil {
		return err
	}
	writer = bufio.NewWriter(fd)
	LogInfo("logging into " + path);
 	return nil
}

func CloseLogger() {
	writer.Flush()
	fd.Close();
}

func LogActionInfo(me string, msg string, args ...interface{}) {
	mylog("INFO (" + me + "): " + msg);
}

func LogActionError(me string, msg string, args ...interface{}) {
	mylog("ERROR (" + me + "): " + msg);
}

func LogInfo(msg string, args ...interface{}) {
	mylog("INFO: " + msg);
}

func LogError(msg string, args ...interface{}) {
	mylog("ERROR: " + msg);
}

func mylog(msg string) {
	now := time.Now();
	buf := fmt.Sprintf("[%s] %s\n",now.Format("2006-01-02 15:04:05"), msg)
	log.Printf(buf);
	_, err := writer.WriteString(buf)
	if err != nil {
		log.Printf("Error while storing the log:%s\n", err)
	}
	err = writer.Flush()
	if err != nil {
		log.Printf("Error while storing the log:%s\n", err)
	}

}
