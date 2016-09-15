package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "!!Hello, World!!"

	// Trim all `!`'s on either side of string s
	//
	// Trim returns a subslice of s by slicing off all leading and
	// trailing UTF-8-encoded Unicode code points contained in cutset.
	st := bytes.Trim([]byte(s), "!")

	// Print the results
	// Prints: Hello, World
	fmt.Println(string(st))
}
