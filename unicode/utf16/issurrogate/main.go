package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	// IsSurrogate reports whether the specified Unicode code point
	// can appear in a surrogate pair.
	answer := utf16.IsSurrogate('水')
	fmt.Println(answer)

	answer = utf16.IsSurrogate('\U0001D11E')
	fmt.Println(answer)

	answer = utf16.IsSurrogate(rune(0xdc00))
	fmt.Println(answer)

	answer = utf16.IsSurrogate('\u6C34')
	fmt.Println(answer)

	answer = utf16.IsSurrogate(rune(0xdfff))
	fmt.Println(answer)
}
