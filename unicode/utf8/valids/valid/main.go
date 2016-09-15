package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}

	// Valid reports whether p consists entirely of valid UTF-8-encoded runes.
	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
}
