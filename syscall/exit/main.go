package main

import (
	"syscall"
)

func main() {
	// Exit causes the current program to exit with the given status code.
	// Conventionally, code zero indicates success, non-zero an error.
	// The program terminates immediately; deferred functions are not run.
	//
	// Exit the program with a status code of 1 representing an error.
	syscall.Exit(1)
}
