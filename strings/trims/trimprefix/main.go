package main

import (
	"fmt"
	"strings"
)

func main() {
	// Trim the prefix `Hello, ` from the string `Hello, World!`
	//
	// TrimPrefix returns s without the provided leading prefix string.
	// If s doesn't start with prefix, s is returned unchanged.
	fmt.Println(strings.TrimPrefix("Hello, World!", "Hello, ")) // World!

	// Now trim the prefix `!` from the string below
	fmt.Println(strings.TrimPrefix("!Hello, World!", "!")) // Hello, World!
}
