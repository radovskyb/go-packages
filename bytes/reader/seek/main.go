package main

import (
	"bytes"
	"os"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c', 'd', 'e'})

	// Use r.Seek(1, 1) to place r's position at 1 (byte `b`)
	//
	// Seek implements the io.Seeker interface.
	r.Seek(1, 1)

	// Print r's contents to os.Stdout
	r.WriteTo(os.Stdout) // bcde
}
