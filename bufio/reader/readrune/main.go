package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new buffered reader br which reads from
	// a new string reader that reads from the string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one rune from buffered reader br
	//
	// ReadRune reads a single UTF-8 encoded Unicode character and
	// returns the rune and its size in bytes. If the encoded rune
	// is invalid, it consumes one byte and returns
	// unicode.ReplacementChar (U+FFFD) with a size of 1.
	r, size, _ := br.ReadRune()

	// Print the rune read from br in string form
	fmt.Println("Read", string(r), "which is", size, "byte/s long") // H
}
