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

	// Read a string from the buffer buf, up until and
	// including the delimiter `,` comma
	//
	// ReadString reads until the first occurrence of delim in the
	// input, returning a string containing the data up to and including
	// the delimiter. If ReadString encounters an error before finding
	// a delimiter, it returns the data read before the error and the error
	// itself (often io.EOF). ReadString returns err != nil if and only
	// if the returned data does not end in delim.
	line, err := buf.ReadString(byte(','))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out line
	fmt.Println(line)
}
