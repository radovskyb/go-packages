package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Exiting the program before execution naturally finishes using os.Exit")

	// Exit the program
	//
	// Exit causes the current program to exit with the given status code.
	// Conventionally, code zero indicates success, non-zero an error.
	// The program terminates immediately; deferred functions are not run.
	os.Exit(0)
}
