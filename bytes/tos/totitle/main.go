package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "Hello, World! Wassup?"

	// Run bytes.ToTitle on s, which will titleise all
	// unicode letters in s
	//
	// ToTitle returns a copy of the byte slice s
	// with all Unicode letters mapped to their title case.
	tl := bytes.ToTitle([]byte(s))

	// Print the result as a string
	// Prints: HELLO, WORLD! WASSUP?
	fmt.Println(string(tl))
}
