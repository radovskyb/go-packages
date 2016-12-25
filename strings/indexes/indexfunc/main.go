package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Create an IndexFunc function that returns true if the
	// character rune is equivalent to a unicode.Han character
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	// IndexFunc returns the index into s of the first Unicode
	// code point satisfying f(c), or -1 if none do.
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
}
