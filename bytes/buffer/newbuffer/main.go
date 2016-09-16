package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	// Create a byte slice with some values
	s := []byte{'a', 'b', 'c'}

	// Create a new bytes buffer that's initial contents is s
	//
	// NewBuffer creates and initializes a new Buffer using buf
	// as its initial contents. It is intended to prepare a Buffer
	// to read existing data. It can also be used to size the
	// internal buffer for writing. To do that, buf should have
	// the desired capacity but a length of zero.
	//
	// In most cases, new(Buffer) (or just declaring a Buffer variable)
	// is sufficient to initialize a Buffer.
	b := bytes.NewBuffer(s)

	// Write the bytes buffer b to os.Stdout
	_, err := b.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
