package main

import (
	"fmt"
	"strings"
)

func main() {
	// IndexRune returns the index of the first instance
	// of the Unicode code point r, or -1 if rune is not present in s.
	fmt.Println(strings.IndexRune("Hello, World!", rune('o'))) // 4
	fmt.Println(strings.IndexRune("Hello, World!", rune('z'))) // -1
}
