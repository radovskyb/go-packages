package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a byte slice of words split by spaces that
	// we will want to tokenize into their own fields
	bslice := []byte("Hello  how are ya    ??")

	// Get all fields from our byte slice by using bytes.Fields
	//
	// Fields splits the slice s around each instance of one
	// or more consecutive white space characters, returning a
	// slice of subslices of s or an empty list if s contains
	// only white space.
	fields := bytes.Fields(bslice)

	// Now range over all of the fields and print out each
	// field on it's own line.
	for _, field := range fields {
		fmt.Println(string(field))
	}
}
