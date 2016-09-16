package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c', 'd', 'e'})

	// Use r.Seek(1, 1) to place r's position at 1 (byte `b`)
	//
	// Seek implements the io.Seeker interface.
	_, err := r.Seek(1, 1)
	if err != nil {
		log.Fatalln(err)
	}

	// Print r's contents to os.Stdout
	_, err = r.WriteTo(os.Stdout) // bcde
	if err != nil {
		log.Fatalln(err)
	}
}
