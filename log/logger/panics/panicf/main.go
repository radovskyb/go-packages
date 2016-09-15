package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Panicf is equivalent to l.Printf() followed by a call to panic().
	logger.Panicf("Panic%s", "f")
}
