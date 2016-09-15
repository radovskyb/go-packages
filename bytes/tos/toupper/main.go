package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "Hello, World! Wassup?"

	// Run bytes.ToUpper on s, which will uppercase all
	// unicode letters in s
	//
	// ToUpper returns a copy of the byte slice s with all
	// Unicode letters mapped to their upper case.
	tl := bytes.ToUpper([]byte(s))

	// Print the result as a string
	// Prints: HELLO, WORLD! WASSUP?
	fmt.Println(string(tl))
}
