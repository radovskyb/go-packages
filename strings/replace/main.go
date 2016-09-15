package main

import (
	"fmt"
	"strings"
)

func main() {
	// Pass -1 so the replace function runs through the whole string
	//
	// Replace returns a copy of the string s with the first n
	// non-overlapping instances of old replaced by new.
	// If old is empty, it matches at the beginning of the string
	// and after each UTF-8 sequence, yielding up to k+1 replacements
	// for a k-rune string.
	// If n < 0, there is no limit on the number of replacements.
	fmt.Println(strings.Replace("Hello, World!", "l", "1", -1))

	// Now pass 1 so the replace function only replaces the first
	// instance of old with new in the string below
	fmt.Println(strings.Replace("blah blah blah", "bla", "blah", 1))
}
