package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
	logger.Fatalln("Fatalln")
}
