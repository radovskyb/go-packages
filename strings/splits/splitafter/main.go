package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split the string after each instance of a comma `,`
	//
	// SplitAfter slices s into all substrings after each
	// instance of sep and returns a slice of those substrings.
	// If sep is empty, SplitAfter splits after each UTF-8 sequence.
	// It is equivalent to SplitAfterN with a count of -1.
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
}
