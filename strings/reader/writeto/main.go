package main

import (
	"os"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Write from sr to os.Stdout
	sr.WriteTo(os.Stdout) // Prints: Hello, World!
}
