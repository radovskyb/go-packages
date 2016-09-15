package main

import (
	"fmt"
	"strings"
)

func main() {
	// ContainsRune reports whether the Unicode code point r is within s.
	fmt.Println(strings.ContainsRune("abc", 'a')) // true
	fmt.Println(strings.ContainsRune("abc", 'z')) // false
}
