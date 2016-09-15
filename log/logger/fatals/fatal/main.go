package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
	logger.Fatal("Fatal")
}
