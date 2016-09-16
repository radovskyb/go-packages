package main

import (
	"bytes"
	"fmt"
	"log"
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
	n, err := buf.ReadFrom(sr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read %d bytes from sr\n", n)

	// Write the buffer buf to os.Stdout
	_, err = buf.WriteTo(os.Stdout) // Prints: Hello, World!
	if err != nil {
		log.Fatalln(err)
	}
}
