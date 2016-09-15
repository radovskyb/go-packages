package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string that says Hello, World!
	input := "Hello, World!"

	// Store the index of bytes.Index's result
	//
	// Index returns the index of the first instance of
	// sep in s, or -1 if sep is not present in s.
	index := bytes.Index([]byte(input), []byte("W"))

	// Print out the index if it's found
	//
	// Since bytes.Index will return -1 if it isn't found,
	// we can write an if statement like so
	if index > -1 {
		fmt.Println("Index found at position:", index)
	} else {
		fmt.Println("Index not found")
	}
}
