package main

import "log"

func main() {
	log.Println("Logging without any prefix")

	// SetPrefix sets the output prefix for the standard logger.
	log.SetPrefix("Log output: ")

	// Prefix returns the output prefix for the standard logger.
	log.Printf("Logging with the prefix `%s`", log.Prefix())
}
