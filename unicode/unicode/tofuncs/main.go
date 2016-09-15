package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'G'

	// ToLower maps the rune to lower case.
	fmt.Printf("%#U\n", unicode.ToLower(ucG))

	const lcG = 'g'

	// ToTitle maps the rune to title case.
	fmt.Printf("%#U\n", unicode.ToTitle(lcG))

	// ToUpper maps the rune to upper case.
	fmt.Printf("%#U\n", unicode.ToUpper(lcG))
}
