package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Print calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Print.
	logger.Print("Print")
}
