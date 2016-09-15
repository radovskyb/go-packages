package main

import (
	"fmt"
	"unicode"
)

func main() {
	t := unicode.TurkishCase

	const lci = 'i'

	// ToLower maps the rune to lower case giving priority to the special mapping.
	fmt.Printf("%#U\n", t.ToLower(lci))

	// ToTitle maps the rune to title case giving priority to the special mapping.
	fmt.Printf("%#U\n", t.ToTitle(lci))

	// ToUpper maps the rune to upper case giving priority to the special mapping.
	fmt.Printf("%#U\n", t.ToUpper(lci))

	const uci = 'İ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToTitle(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))
}
