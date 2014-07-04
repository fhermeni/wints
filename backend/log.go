package backend

import (
	"log"
	"time"
)

const (
	LOGFILE = "log.txt"
)
func Log(me string, msg string, args ...interface{}) {
	now := time.Now();
	str := "[" + now.String() + "] " + me + ": " + msg
	log.Println(str);
}
