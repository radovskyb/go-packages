package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Println calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Println.
	logger.Println("Println")
}
