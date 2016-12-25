package main

import (
	"fmt"
	"log"
)

func main() {
	// Below, fmt.Println prints the given string and adds a space around any separate
	// operands such as "string" in this case as well as adds a \n character at the end
	//
	// Println formats using the default formats for its operands and writes to standard output.
	// Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.
	_, err := fmt.Println("This is a", "string",
		"that will end with a \\n character added by fmt.Println")
	if err != nil {
		log.Fatalln(err)
	}
}
