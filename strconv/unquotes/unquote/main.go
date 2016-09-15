package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := func(s string) {
		// Unquote interprets s as a single-quoted, double-quoted,
		// or backquoted Go string literal, returning the string value
		// that s quotes.  (If s is single-quoted, it would be a Go
		// character literal; Unquote returns the corresponding
		// one-character string.)
		t, err := strconv.Unquote(s)
		if err != nil {
			fmt.Printf("Unquote(%v): %v\n", s, err)
		} else {
			fmt.Printf("Unquote(%v) = %v\n", s, t)
		}
	}

	s := `\"I'm inside of quotes\t\u263a\"\"`
	// If the string doesn't have quotes, it can't be unquoted.
	test(s) // invalid syntax
	test("`" + s + "`")
	test(`"` + s + `"`)
	test(`'\u263a'`)
}
