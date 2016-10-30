package main

import "log"

func main() {
	// Flags returns the output flags for the standard logger.
	log.Println(log.Flags()) // 3

	// Set flags to change the log output.
	//
	// SetFlags sets the output flags for the standard logger.
	log.SetFlags(log.Ldate)

	log.Println(log.Flags()) // 1

	// Set flags to change the log output again.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println(log.Flags()) // 19
}
