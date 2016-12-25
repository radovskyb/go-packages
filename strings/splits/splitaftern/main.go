package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split after each comma in the string below, returning
	// at most 2 separated strings by calling 2 as it's n integer
	//
	// SplitAfterN slices s into substrings after each instance of sep and
	// returns a slice of those substrings.
	// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
	// The count determines the number of substrings to return:
	//   n > 0: at most n substrings; the last substring will
	//			be the unsplit remainder.
	//   n == 0: the result is nil (zero substrings)
	//   n < 0: all substrings
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))

	// Call the same thing again but this time with 3 as it's n int
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 3))

	// Calling -1 as it's n integer will split on all occurrences
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", -1))

	// Calling 0 as it's n integer will return 0 substrings
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 0))
}
