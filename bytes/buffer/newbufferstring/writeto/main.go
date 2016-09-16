package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new bytes buffer that's initial contents is string s
	//
	// NewBufferString creates and initializes a new Buffer
	// using string s as its initial contents. It is intended
	// to prepare a buffer to read an existing string.
	//
	// In most cases, new(Buffer) (or just declaring a Buffer
	// variable) is sufficient to initialize a Buffer.
	b := bytes.NewBufferString(s)

	// Write the bytes buffer b to os.Stdout
	_, err := b.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
