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
	defer logWriter.Close()
	if err != nil {
		log.Fatal("error")
	}

	// Alert logs a message with severity LOG_ALERT, ignoring the severity
	// passed to New.
	logWriter.Alert("alert")

	// Crit logs a message with severity LOG_CRIT, ignoring the severity
	// passed to New.
	logWriter.Crit("critical")

	// Err logs a message with severity LOG_ERR, ignoring the severity
	// passed to New.
	logWriter.Err("error")

	// Warning logs a message with severity LOG_WARNING, ignoring the
	// severity passed to New.
	logWriter.Warning("warning")

	// Notice logs a message with severity LOG_NOTICE, ignoring the
	// severity passed to New.
	logWriter.Notice("notice")

	// Info logs a message with severity LOG_INFO, ignoring the severity
	// passed to New.
	logWriter.Info("information")

	// Debug logs a message with severity LOG_DEBUG, ignoring the severity
	// passed to New.
	logWriter.Debug("debug")

	// Write sends a log message to the syslog daemon.
	logWriter.Write([]byte("Hello Logger!"))
}
