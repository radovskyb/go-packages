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

	// ReadFull reads exactly len(buf) bytes from r into buf. It
	// returns the number of bytes copied and an error if fewer bytes
	// were read. The error is EOF only if no bytes were read. If an
	// EOF happens after reading some but not all the bytes, ReadFull
	// returns ErrUnexpectedEOF. On return, n == len(buf) if
	// and only if err == nil.
	n, err := io.ReadFull(sr, buf)

	// Check if io.ReadFull returned any errors
	if err != nil {
		// If it did, log them
		log.Fatalln(err)
	}

	// Print out how many bytes were read
	fmt.Printf("Read %d bytes into buffer buf from string reader sr\n", n)

	// Print out buf's contents, after the ReadFull call, as a string
	fmt.Println(string(buf))
}
