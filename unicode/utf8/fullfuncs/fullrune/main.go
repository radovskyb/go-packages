package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte{228, 184, 150} // 世

	// FullRune reports whether the bytes in p begin with a full UTF-8
	// encoding of a rune. An invalid encoding is considered a full Rune
	// since it will convert as a width-1 error rune.
	fmt.Println(utf8.FullRune(buf))
	fmt.Println(utf8.FullRune(buf[:2]))
}
