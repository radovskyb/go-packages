package main

import (
	"fmt"
	"strings"
)

func main() {
	// Get the position of the first vowel in the string `Hello, World`
	//
	// IndexAny returns the index of the first instance of any
	// Unicode code point from chars in s, or -1 if no Unicode code
	// point from chars is present in s.
	fmt.Println(strings.IndexAny("Hello, World!", "aeiou")) // 1 (e)
	fmt.Println(strings.IndexAny("Hello, World!", "xyz"))   // -1
}
