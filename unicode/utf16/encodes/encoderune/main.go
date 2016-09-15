package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	//r1 := '\u6C34'
	//r2 := '水'

	// EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune.
	// If the rune is not a valid Unicode code point or does not need encoding,
	// EncodeRune returns U+FFFD, U+FFFD.
	r1, r2 := utf16.EncodeRune('水')

	fmt.Printf("u+%x u+%x \n", r1, r2)

	// find surrogate for G clef symbol
	r3, r4 := utf16.EncodeRune('\U0001D11E')

	fmt.Printf("G clef surrogate runes are : u+%x u+%x \n", r3, r4)
}
