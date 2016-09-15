package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Panicln is equivalent to l.Println() followed by a call to panic().
	logger.Panicln("Panicln")
}
