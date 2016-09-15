package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Printf calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	logger.Printf("Print%s", "f")
}
