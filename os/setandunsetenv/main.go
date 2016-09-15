package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// To create an environment variable, simply use os.Setenv
	//
	// Setenv sets the value of the environment variable named
	// by the key. It returns an error, if any.
	err := os.Setenv("TMPENV", "values for testenv")

	// Check if there were any errors and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print a success message after calling os.Setenv
	fmt.Println("Successfully set environment variable TMPENV")

	// Now retrieve that environment variables contents using os.Getenv
	tmpenv := os.Getenv("TMPENV")

	// Print its contents
	fmt.Println(tmpenv)

	// Now that we are done with it, delete it using os.Unsetenv
	//
	// Unsetenv unsets a single environment variable.
	err = os.Unsetenv("TMPENV")

	// Check if there were any errors and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print a success message after calling os.Unsetenv
	fmt.Println("Successfully unset environment variable TMPENV")
}
