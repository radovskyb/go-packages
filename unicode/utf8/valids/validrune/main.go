package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := 'a'
	invalid := rune(0xfffffff)

	// ValidRune reports whether r can be legally encoded as UTF-8.
	// Code points that are out of range or a surrogate half are illegal.
	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
}
