package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Move the reader over 1 byte
	//
	// Seek implements the io.Seeker interface.
	_, err := sr.Seek(1, 1)
	if err != nil {
		log.Fatalln(err)
	}

	// Read 1 byte from sr into `b`
	b, err := sr.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out `b` as a string
	fmt.Println(string(b)) // e
}
