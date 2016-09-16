package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	// Create an input string to read into our buffer
	const input = "Hello, World\n"

	// Create a buf variable which is a new bytes string
	// reader which will read from the input string above
	//
	// NewBufferString creates and initializes a new Buffer using string s as its
	// initial contents. It is intended to prepare a buffer to read an existing
	// string.
	//
	// In most cases, new(Buffer) (or just declaring a Buffer variable) is
	// sufficient to initialize a Buffer.
	buf := bytes.NewBufferString(input)

	// Copy the read contents from our buf reader to os.Stdout
	_, err := io.Copy(os.Stdout, buf)
	if err != nil {
		log.Fatalln(err)
	}
}
