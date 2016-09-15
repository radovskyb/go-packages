package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new section reader sr that reads from a new string
	// reader that reads from string s, that starts at the offset
	// of 6 bytes and can read upto 7 bytes
	// At position at space before World: -> ` World!`
	sr := io.NewSectionReader(strings.NewReader(s), 6, 7)

	// Print out sr's size
	//
	// Size returns the size of the section in bytes.
	fmt.Println(sr.Size()) // 7

	// Now seek the position 1 place after the current offset
	sr.Seek(1, 1) // Now at the position of byte `W

	// Create a byte slice buf to read into
	buf := make([]byte, 7)

	// Read from sr into buf
	sr.Read(buf)

	// Print out buf as a string
	fmt.Println(string(buf)) // Prints: `World!`
}
