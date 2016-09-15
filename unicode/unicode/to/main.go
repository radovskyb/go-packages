package main

import (
	"fmt"
	"unicode"
)

func main() {
	const lcG = 'g'

	// To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase.
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))

	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG))
}
