package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!\n")

	// Write from sr to os.Stdout
	n, err := sr.WriteTo(os.Stdout) // Prints: Hello, World!
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Wrote %d bytes to os.Stdout\n", n)
}
