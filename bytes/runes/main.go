package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string s
	s := "Hello, World!"

	// Return s as a slice of runes into runes
	//
	// Runes returns a slice of runes (Unicode
	// code points) equivalent to s.
	runes := bytes.Runes([]byte(s))

	// Iterate over runes range
	for _, r := range runes {
		// Print out each rune followed by a space
		fmt.Print(string(r), " ")
	}
}
