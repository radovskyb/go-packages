package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	// Create a new byte slice
	b := []byte("hello, world")

	// Use the function bytes.Map and pass in the unicode
	// function unicode.ToUpper, to return the byte slice
	// b, all in capital letters, which is then stored in
	// a byte slice b1
	b1 := bytes.Map(unicode.ToUpper, b)

	// Print the outcome of using the bytes.Map function
	// by printing b1 as a string
	fmt.Println(string(b1))

	// We can also pass in our own functions to map:
	b2 := bytes.Map(PlusOneRune, b)

	// Print the results of using map with the custom
	// function PlusOneRune
	fmt.Println(string(b2))
}

// PlusOneRune simply moves each rune character over by 1.
// Therefore an `a` will become `b`, a `b` will become `c`.
func PlusOneRune(r rune) rune {
	r += 1
	return r
}
