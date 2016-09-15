package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Create a FieldsFunc function type that splits when the rune
	// is neither a unicode letter or a unicode number
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	// FieldsFunc splits the string s at each run of Unicode code
	// points c satisfying f(c) and returns an array of slices of s.
	// If all code points in s satisfy f(c) or the string is empty,
	// an empty slice is returned.
	// FieldsFunc makes no guarantees about the order in which it
	// calls f(c).
	// If f does not return consistent results for a given c,
	// FieldsFunc may crash.
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}
