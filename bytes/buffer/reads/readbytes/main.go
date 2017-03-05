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

	// Read any amount of bytes from buf up
	// until the delimiter `,` (comma)
	//
	// ReadBytes reads until the first occurrence of delim in the
	// input, returning a slice containing the data up to and including
	// the delimiter. If ReadBytes encounters an error before finding
	// a delimiter, it returns the data read before the error and the
	// error itself (often io.EOF). ReadBytes returns err != nil if and
	// only if the returned data does not end in delim.
	b, err := buf.ReadBytes(byte(','))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out b as a string
	fmt.Println(string(b))
}
