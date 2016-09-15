package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new strings reader to read from s
	sr := strings.NewReader(s)

	// Create a buffer to read into
	buf := make([]byte, len(s))

	// ReadAtLeast reads from r into buf until it has read at least
	// min bytes. It returns the number of bytes copied and an error
	// if fewer bytes were read. The error is EOF only if no bytes
	// were read. If an EOF happens after reading fewer than min bytes,
	// ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the
	// length of buf, ReadAtLeast returns ErrShortBuffer. On return,
	// n >= min if and only if err == nil.
	n, err := io.ReadAtLeast(sr, buf, len(s))

	// Check if io.ReadAtLeast returned any errors
	if err != nil {
		// If it did, log them
		log.Fatalln(err)
	}

	// Print out how many bytes were read
	fmt.Printf("Read %d bytes into buffer buf from string reader sr\n", n)

	// Print out buf's contents, after the ReadAtLeast call, as a string
	fmt.Println(string(buf))
}
