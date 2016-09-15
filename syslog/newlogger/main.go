package main

import (
	"log"
	"log/syslog"
)

func main() {
	// NewLogger creates a log.Logger whose output is written to
	// the system log service with the specified priority. The logFlag
	// argument is the flag set passed through to log.New to create
	// the Logger.
	sysLogger, err := syslog.NewLogger(syslog.LOG_CRIT, log.Lshortfile)
	if err != nil {
		log.Fatalln(err)
	}

	sysLogger.Print("Hello, sysLogger!")
}
