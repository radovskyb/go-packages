package main

import (
	"log"
	"log/syslog"
)

func main() {
	// Dial establishes a connection to a log daemon by connecting to
	// address raddr on the specified network.  Each write to the returned
	// writer sends a log message with the given facility, severity and
	// tag.
	// If network is empty, Dial will connect to the local syslog server.
	logWriter, err := syslog.Dial("tcp", "localhost:9000", syslog.LOG_ERR, "SysLogger ")
	if err != nil {
		log.Fatal("error")
	}
	defer logWriter.Close()

	// Alert logs a message with severity LOG_ALERT, ignoring the severity
	// passed to New.
	if err := logWriter.Alert("alert"); err != nil {
		log.Fatalln(err)
	}

	// Crit logs a message with severity LOG_CRIT, ignoring the severity
	// passed to New.
	if err := logWriter.Crit("critical"); err != nil {
		log.Fatalln(err)
	}

	// Err logs a message with severity LOG_ERR, ignoring the severity
	// passed to New.
	if err := logWriter.Err("error"); err != nil {
		log.Fatalln(err)
	}

	// Warning logs a message with severity LOG_WARNING, ignoring the
	// severity passed to New.
	if err := logWriter.Warning("warning"); err != nil {
		log.Fatalln(err)
	}

	// Notice logs a message with severity LOG_NOTICE, ignoring the
	// severity passed to New.
	if err := logWriter.Notice("notice"); err != nil {
		log.Fatalln(err)
	}

	// Info logs a message with severity LOG_INFO, ignoring the severity
	// passed to New.
	if err := logWriter.Info("information"); err != nil {
		log.Fatalln(err)
	}

	// Debug logs a message with severity LOG_DEBUG, ignoring the severity
	// passed to New.
	if err := logWriter.Debug("debug"); err != nil {
		log.Fatalln(err)
	}

	// Write sends a log message to the syslog daemon.
	_, err = logWriter.Write([]byte("Hello Logger!"))
	if err != nil {
		log.Fatalln(err)
	}
}
