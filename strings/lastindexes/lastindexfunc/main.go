package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Create a function `f` to be used with strings.LastIndexFunc
	f := func(r rune) bool {
		return unicode.IsLetter(r)
	}

	// Get the last index of a unicode letter using the func f
	// with strings.LastIndexFunc. It will get the index of `d`
	// from Hello, Worl`d` <- and ignore the `!!!` afterwards since
	// they do not satisfy the function f above
	//
	// LastIndexFunc returns the index into s of the last
	// Unicode code point satisfying f(c), or -1 if none do.
	fmt.Println(strings.LastIndexFunc("Hello, World!!!", f)) // 11

	// This one returns -1 since there are no unicode
	// letters in the string `!!!`
	fmt.Println(strings.LastIndexFunc("!!!", f)) // -1
}
