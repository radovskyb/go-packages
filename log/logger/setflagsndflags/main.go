package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Print out logger's flags
	//
	// Flags returns the output flags for the logger.
	fmt.Println(logger.Flags())

	logger.Println("Hello, Logger!\n")

	// Set the logger flags to log.Llongfile
	//
	// SetFlags sets the output flags for the logger.
	logger.SetFlags(log.Llongfile)

	// Print out logger's flags
	fmt.Println(logger.Flags())

	logger.Print("Hello, Logger!")
}
