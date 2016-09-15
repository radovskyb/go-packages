package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Panic is equivalent to l.Print() followed by a call to panic().
	logger.Panic("Panic")
}
