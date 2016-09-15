package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "!!Hello, World!!"

	// Trim all `!`'s on the right side of string s
	//
	// TrimRight returns a subslice of s by slicing off all
	// trailing UTF-8-encoded Unicode code points that are contained in cutset.
	st := bytes.TrimRight([]byte(s), "!")

	// Print the results
	// Prints: !!Hello, World
	fmt.Println(string(st))
}
