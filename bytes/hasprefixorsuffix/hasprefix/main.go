package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a byte slice that says Hello, World!
	b := []byte("Hello, World!")

	// Print out byte slice b
	fmt.Println(string(b))

	// Check if byte slice b starts with the bytes Hello
	if bytes.HasPrefix(b, []byte("Hello")) {
		// If it does, then print out the result
		fmt.Println("byte slice b starts with Hello")
	}
}
