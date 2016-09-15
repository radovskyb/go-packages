package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r := '世'
	buf := make([]byte, 3)

	// EncodeRune writes into p (which must be large enough) the UTF-8
	// encoding of the rune. It returns the number of bytes written.
	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
}
