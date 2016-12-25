package main

import (
	"log"
	"log/syslog"
)

func main() {
	var wLogger *syslog.Writer

	// Create a new syslog Writer
	//
	// New establishes a new connection to the system log daemon.  Each
	// write to the returned writer sends a log message with the given
	// priority and prefix.
	wLogger, err := syslog.New(syslog.LOG_CRIT, "A critical error occurred! ")
	if err != nil {
		log.Fatalln(err)
	}

	// Close closes a connection to the syslog daemon.
	defer wLogger.Close()

	// Write a message to the system logger through wLogger
	_, err = wLogger.Write([]byte("Hello, wLogger"))
	if err != nil {
		log.Fatalln(err)
	}

	// To see the logged output open /var/log/system.log
}
