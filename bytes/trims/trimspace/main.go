package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "   Hello, World    "

	// Print string s before bytes.TrimSpace
	// Prints: `   Hello, World    `
	fmt.Println(s)

	// Trim all spaces from eiher side of s
	//
	// TrimSpace returns a subslice of s by slicing off all
	// leading and trailing white space, as defined by Unicode.
	st := bytes.TrimSpace([]byte(s))

	// Print the results
	// Prints: Hello, World
	fmt.Println(string(st))
}
