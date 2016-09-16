package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new buffered reader br which reads from
	// a new string reader that reads from the string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one byte from buffered reader br
	//
	// ReadByte reads and returns a single byte.
	// If no byte is available, returns an error.
	b, err := br.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the byte read from br in string form
	fmt.Println(string(b)) // H
}
