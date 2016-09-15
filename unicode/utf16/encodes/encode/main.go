package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	rune_array := []rune{'a', 'b', '好', 0xfffd}

	var uint16_array []uint16

	// Encode returns the UTF-16 encoding of the Unicode code point sequence s.
	uint16_array = utf16.Encode(rune_array)

	fmt.Println(uint16_array)

	// 0xfffd should not be encoded.

	for _, item := range uint16_array {
		fmt.Printf("0x%x \n", item)
	}
}
