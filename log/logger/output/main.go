package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged - ", log.Lshortfile)

	// Output writes the output for a logging event. The string s contains
	// the text to print after the prefix specified by the flags of the
	// Logger. A newline is appended if the last character of s is not
	// already a newline. Calldepth is used to recover the PC and is
	// provided for generality, although at the moment on all pre-defined
	// paths it will be 2.
	logger.Output(1, "Hello, Logger!")
}
