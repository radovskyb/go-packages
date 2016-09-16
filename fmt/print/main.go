package main

import (
	"fmt"
	"log"
)

func main() {
	// Print formats using the default formats for its operands and writes to standard output.
	// Spaces are added between operands when neither is a string.
	// It returns the number of bytes written and any write error encountered.
	n, err := fmt.Print("Hello, World\n")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Wrote %d bytes\n", n)
}
