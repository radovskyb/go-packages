package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "Hello, World! Wassup?"

	// Run bytes.ToLower on s
	//
	// ToLower returns a copy of the byte slice s
	// with all Unicode letters mapped to their lower case.
	tl := bytes.ToLower([]byte(s))

	// Print the result as a string
	// Prints: hello, world! wassup?
	fmt.Println(string(tl))
}
