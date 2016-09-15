package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a byte slice called b
	b := []byte("cbabc")

	// LastIndexByte returns the index of the last instance
	// of c in s, or -1 if c is not present in s.
	index := bytes.LastIndexByte(b, byte('c'))

	// If the byte `c` is present in the byte slice b,
	// index will return a positive value, otherwise
	// it will return -1. Print out the results here.
	if index > -1 {
		fmt.Println("Last index found:", index)
	} else {
		fmt.Println("Index not found.")
	}
}
