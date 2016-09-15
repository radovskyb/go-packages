package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a rot13 conversion function to be used with strings.Map
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'A'+13)%26
		}
		return r
	}

	// Create a string `today`
	today := "Today was a very nice day!"

	// Print the rot13 version of the string `today` by calling
	// strings.Map and passing it the rot13 conversion function
	// `rot13` and the string `today`
	//
	// Map returns a copy of the string s with all its characters modified
	// according to the mapping function. If mapping returns a
	// negative value, the character is dropped from the
	// string with no replacement.
	fmt.Println(strings.Map(rot13, today))
}
