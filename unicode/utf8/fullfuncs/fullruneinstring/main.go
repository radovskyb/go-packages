package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "世"

	// FullRuneInString is like FullRune but its input is a string.
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))
}
