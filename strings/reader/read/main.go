package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!\n")

	// Create a byte slice `buf` with the length of
	// the unread portion of the strings reader `sr`
	buf := make([]byte, sr.Len())

	// Read from sr into buf
	n, err := sr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out buf as a string
	fmt.Printf("Read %d bytes from sr into buf: %s", n, buf) // Hello, World!
}
