package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})

	// ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
	fmt.Println(utf8.ValidString(valid))
	fmt.Println(utf8.ValidString(invalid))
}
