package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "!!Hello, World!!"

	// Trim all `!`'s on the left side of string s
	//
	// TrimLeft returns a subslice of s by slicing off all
	// leading UTF-8-encoded Unicode code points contained in cutset.
	st := bytes.TrimLeft([]byte(s), "!")

	// Print the results
	// Prints: Hello, World!!
	fmt.Println(string(st))
}
