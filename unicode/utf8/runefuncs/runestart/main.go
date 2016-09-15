package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte("a界")

	// RuneStart reports whether the byte could be the first byte of an encoded,
	// possibly invalid rune. Second and subsequent bytes always have the top two
	// bits set to 10.
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
}
