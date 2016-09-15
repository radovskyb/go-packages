package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	var rune_array []rune

	// Decode returns the Unicode code point sequence represented
	// by the UTF-16 encoding s.
	rune_array = utf16.Decode([]uint16{0xdfff})
	fmt.Println(rune_array)

	rune_array = utf16.Decode([]uint16{'爱', 0xd800, 'a'})
	fmt.Println(rune_array)
}
