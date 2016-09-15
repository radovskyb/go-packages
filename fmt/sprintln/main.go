package main

import "fmt"

func main() {
	// Sprintln formats using the default formats for its operands
	// and returns the resulting string.
	// Spaces are always added between operands and a newline is appended.
	numStr := fmt.Sprintln("Number:", 5)
	fmt.Print(numStr)
}
