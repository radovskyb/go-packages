package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string that says Hello, Gophers!
	input := "Hello, Gophers!"

	// Store the index of bytes.LastIndex's result
	//
	// bytes.LastIndex returns the index of the last instance
	// of sep in s, or -1 if sep is not present in s.
	index := bytes.LastIndex([]byte(input), []byte("e"))

	// Print out the index if it's found
	//
	// Since bytes.Index will return -1 if it isn't found,
	// we can write an if statement like so
	if index > -1 {
		fmt.Println("Last index found at position:", index)
	} else {
		fmt.Println("Index not found")
	}
}
