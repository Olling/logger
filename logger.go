package logger

import (
	"os"
    	"log"
)


var (
	debug bool
	infolog *log.Logger
	debuglog *log.Logger
	errorlog *log.Logger
)


func Debug (a ...interface{}) {
	if debug {
		debuglog.Println(a...)
	}
}

func Info (a ...interface{}) {
	infolog.Println(a...)
}


func Error (a ...interface{}) {
	errorlog.Println(a...)
}

func SetDebugState (state bool) {
	debug = state
}


func Initialize() {
	debuglog = log.New(os.Stdout, "DEBUG: ", 0)
	debuglog.Println("Debug logging started")

	infolog = log.New(os.Stdout, "INFO: ", 0)
	debuglog.Println("Info logging started")

	errorlog = log.New(os.Stderr, "ERROR: ", 0)
	debuglog.Println("Error logging started")
}
