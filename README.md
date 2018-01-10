# slog
Simple Log (slog) is trying to make it ekstremly easy to get started with different logleevels in your golang applications

## Examples
``` 
package main

import (
	"fmt"
	"github.com/olling/slog"
)

func main() {
  // The slog Loggers are ready for use
	slog.PrintInfo("Welcome to the Readme")
	
  // Change the LogLevel for all Loggers (Default LogLevel is 'All')
  slog.SetLogLevel(slog.Trace)
  
  // Test the Loggers
	slog.PrintTrace("Trace")
	slog.PrintDebug("Debug")
	slog.PrintInfo("Info")
	slog.PrintWarn("Warn")
	slog.PrintError("Error")
	slog.PrintFatal("Fatal")

  // Change the flag(s) for all Loggers
	slog.SetFlags(log.Llongfile)

  // Test the new flag(s)
	slog.PrintTrace("Trace")
	slog.PrintDebug("Debug")
	slog.PrintInfo("Info")
	slog.PrintWarn("Warn")
	slog.PrintError("Error")
	slog.PrintFatal("Fatal")

  // Changing the flag for a single Logger
	slog.InfoLogger.SetFlags(log.Lshortfile)

  // Test the new flag(s)
	slog.PrintTrace("Trace")
	slog.PrintDebug("Debug")
	slog.PrintInfo("Info")
	slog.PrintWarn("Warn")
	slog.PrintError("Error")
	slog.PrintFatal("Fatal")
}
```
