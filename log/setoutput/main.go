package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create a new logfile use for logging to
	logfile, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer logfile.Close()

	// Set the output destination for the standard logger to logfile
	//
	// SetOutput sets the output destination for the standard logger.
	log.SetOutput(logfile)

	// Log some text to logfile
	for i := 1; i <= 5; i++ {
		log.Print("Logged to logfile: ", i)
	}

	// Get log.txt's contents
	logfileContents, err := ioutil.ReadFile("log.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out log.txt's contents as a string
	fmt.Printf("%s", logfileContents)

	// Remove log.txt now that we are done with it
	err = os.Remove("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
