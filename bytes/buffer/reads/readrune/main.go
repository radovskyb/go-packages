package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new string s
	s := "Hello, World"

	// Create a new bytes buffer buf with
	// string s as it's initial content
	buf := bytes.NewBuffer([]byte(s))

	// Read one rune from buf into b
	//
	// ReadRune reads and returns the next UTF-8-encoded Unicode
	// code point from the buffer. If no bytes are available, the error
	// returned is io.EOF. If the bytes are an erroneous UTF-8 encoding,
	// it consumes one byte and returns U+FFFD, 1.
	b, size, err := buf.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out b as a string
	fmt.Printf("Read in: %c\nSize in bytes %d\n", b, size)
}
