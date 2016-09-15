package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
	logger.Fatalf("Fatal%s", "f")
}
