package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Print out the loggers current prefix
	//
	// Prefix returns the output prefix for the logger.
	fmt.Println(logger.Prefix())

	logger.Println("Hello, Logger!\n")

	// Change the loggers prefix
	//
	// SetPrefix sets the output prefix for the logger.
	logger.SetPrefix("Log: ")

	fmt.Println(logger.Prefix())
	logger.Print("Hello, Logger!")
}
