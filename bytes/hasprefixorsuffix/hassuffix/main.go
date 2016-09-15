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

	// Check if byte slice b ends with the bytes World!
	if bytes.HasSuffix(b, []byte("World!")) {
		// If it does, then print out the result
		fmt.Println("byte slice b ends with Hello!")
	}
}
