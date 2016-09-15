package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))

	// RuneCount returns the number of runes in p.  Erroneous and short
	// encodings are treated as single runes of width 1 byte.
	fmt.Println("runes =", utf8.RuneCount(buf))
}
