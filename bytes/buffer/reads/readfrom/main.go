package main

import (
	"bytes"
	"os"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new bytes buffer buf
	var buf bytes.Buffer

	// Create a new strings reader to read from s
	sr := strings.NewReader(s)

	// Buffer buf now reads from the string reader sr
	//
	// ReadFrom reads data from r until EOF and appends it to
	// the buffer, growing the buffer as needed. The return value n
	// is the number of bytes read. Any error except io.EOF encountered
	// during the read is also returned. If the buffer becomes too large,
	// ReadFrom will panic with ErrTooLarge.
	buf.ReadFrom(sr)

	// Write the buffer buf to os.Stdout
	buf.WriteTo(os.Stdout) // Prints: Hello, World!
}
