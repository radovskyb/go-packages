package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Create a function that trims all uppercase
	// unicode characters on either side of a string
	f := func(r rune) bool {
		return unicode.IsUpper(r)
	}

	// Call all three of the strings package's trimfunc functions
	// and pass them all function f from above which will trim any
	// unicode uppercase characters

	// Call TrimFunc which will trim from both sides
	//
	// TrimFunc returns a slice of the string s with all leading
	// and trailing Unicode code points c satisfying f(c) removed.
	fmt.Println(strings.TrimFunc("Hello, WorlD", f)) // ello, Worl

	// Call TrimLeftFunc which will trim only from the left side
	//
	// TrimLeftFunc returns a slice of the string s with all leading
	// Unicode code points c satisfying f(c) removed.
	fmt.Println(strings.TrimLeftFunc("Hello, WorlD", f)) // ello, WorlD

	// Call TrimRightFunc which will trim only from the right side
	//
	// TrimRightFunc returns a slice of the string s with all trailing
	// Unicode code points c satisfying f(c) removed.
	fmt.Println(strings.TrimRightFunc("Hello, WorlD", f)) // Hello, Worl
}
