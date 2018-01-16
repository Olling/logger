// Package slog implements a simpler version of the golang log package. It predifines a lot of different log levels, so its easy to output data, which is easy to filter.
// Only newline prints are used in this package.

package slog

import (
	"os"
	"io"
	"fmt"
	"log"
)


var (
	logLevel = All
	FatalLogger = log.New(os.Stderr, "FATAL: ", log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Lshortfile)
	WarnLogger = log.New(os.Stdout, "WARN: ", log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Lshortfile)
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Lshortfile)
	TraceLogger = log.New(os.Stdout, "TRACE: ", log.Lshortfile)
)

// The different log levels to be used with SetLogLevel(int)
const (
	Off = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace
	All
)


// PrintFatal calls l.Output to print to the fatal logger.
// Arguments are handled in the manner of fmt.Println.
func PrintFatal (v ...interface{}) {
	if logLevel >= Fatal {
		FatalLogger.Output(2,fmt.Sprintln(v...))
	}
}


// PrintError calls l.Output to print to the error logger.
// Arguments are handled in the manner of fmt.Println.
func PrintError (v ...interface{}) {
	if logLevel >= Error {
		ErrorLogger.Output(2,fmt.Sprintln(v...))
	}
}


// PrintWarn calls l.Output to print to the warn logger.
// Arguments are handled in the manner of fmt.Println.
func PrintWarn (v ...interface{}) {
	if logLevel >= Warn {
		WarnLogger.Output(2,fmt.Sprintln(v...))
	}
}


// PrintInfo calls l.Output to print to the info logger.
// Arguments are handled in the manner of fmt.Println.
func PrintInfo (v ...interface{}) {
	if logLevel >= Info {
		InfoLogger.Output(2,fmt.Sprintln(v...))
	}
}


// PrintDebug calls l.Output to print to the debug logger.
// Arguments are handled in the manner of fmt.Println.
func PrintDebug (v ...interface{}) {
	if logLevel >= Debug {
		DebugLogger.Output(2,fmt.Sprintln(v...))
	}
}


// PrintTrace calls l.Output to print to the trace logger.
// Arguments are handled in the manner of fmt.Println.
func PrintTrace (v ...interface{}) {
	if logLevel >= Trace {
		TraceLogger.Output(2,fmt.Sprintln(v...))
	}
}


// SetLogLevel set the LogLevel for the Print functions.
func SetLogLevel (level int) {
	logLevel = level
}

// SetFlags sets the flag on all the loggers
func SetFlags(flag int) {
      FatalLogger.SetFlags(flag)
      ErrorLogger.SetFlags(flag)
      WarnLogger.SetFlags(flag)
      InfoLogger.SetFlags(flag)
      DebugLogger.SetFlags(flag)
      TraceLogger.SetFlags(flag)
}


// SetOutput sets the output io.writer on all the loggers
func SetOutput(w io.Writer) {
	FatalLogger.SetOutput(w)
	ErrorLogger.SetOutput(w)
	WarnLogger.SetOutput(w)
	InfoLogger.SetOutput(w)
	DebugLogger.SetOutput(w)
	TraceLogger.SetOutput(w)
}
