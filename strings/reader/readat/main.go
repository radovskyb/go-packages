package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Create a byte slice `buf` with the length of
	// the unread portion of the strings reader `sr`
	buf := make([]byte, sr.Len())

	// Read from sr into buf starting from position 7
	_, err := sr.ReadAt(buf, 7)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	// Print out buf as a string
	fmt.Println(string(buf)) // World!
}
