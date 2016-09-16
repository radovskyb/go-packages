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

	// Create a new buffered reader which reads from
	// a string reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one byte into b from br
	b, err := br.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print that byte
	fmt.Println(string(b)) // H

	// Read one more into b from br
	b, err = br.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print that next byte
	fmt.Println(string(b)) // e

	// Now reset the strings reader which will push the reader
	// back to byte `H` instead of sitting at the next byte `l`
	//
	// Reset discards any buffered data, resets all state, and
	// switches the buffered reader to read from r.
	br.Reset(strings.NewReader(s))

	// Once again read a byte from br
	b, err = br.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Once again print the byte, however
	// this time it will be back at byte `H`
	fmt.Println(string(b)) // H
}
