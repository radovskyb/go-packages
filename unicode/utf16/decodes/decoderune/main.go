package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	// DecodeRune returns the UTF-16 decoding of a surrogate pair.
	// If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
	// the Unicode replacement code point U+FFFD.
	valid_pair := utf16.DecodeRune(0xd800, 0xdc00)
	fmt.Printf("%x \n", valid_pair)

	not_valid_pair := utf16.DecodeRune('\u6C34', '水')
	fmt.Printf("%x \n", not_valid_pair)

	not_valid_pair = utf16.DecodeRune(0xd800, '爱')
	fmt.Printf("%x \n", not_valid_pair)
}
