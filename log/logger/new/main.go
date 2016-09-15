package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer

	// Create a new logger object to log to buf
	//
	// New creates a new Logger. The out variable sets the
	// destination to which log data will be written.
	// The prefix appears at the beginning of each generated log line.
	// The flag argument defines the logging properties.
	logger := log.New(&buf, "Logged: ", log.Lshortfile)

	logger.Print("Hello, Logger!")

	fmt.Print(&buf)
}
