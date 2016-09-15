package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))

	// RuneCountInString is like RuneCount but its input is a string.
	fmt.Println("runes =", utf8.RuneCountInString(str))
}
