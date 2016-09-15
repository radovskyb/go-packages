package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Create a byte slice `buf` with the length of
	// the unread portion of the strings reader `sr`
	buf := make([]byte, sr.Len())

	// Read from sr into buf
	sr.Read(buf)

	// Print out buf as a string
	fmt.Println(string(buf)) // Hello, World!
}
