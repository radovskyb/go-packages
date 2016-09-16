package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Fprintln formats using the default formats for its operands and writes to w.
	// Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.
	_, err := fmt.Fprintln(os.Stdout, "Hello, World! From Fprintln")
	if err != nil {
		log.Fatalln(err)
	}
}
