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

	// Read one byte from buf into b
	//
	// ReadByte reads and returns the next byte from the
	// buffer. If no byte is available, it returns error io.EOF.
	b, err := buf.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out b as a string
	fmt.Println(string(b))
}
