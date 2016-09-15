package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// RuneLen returns the number of bytes required to encode the rune.
	// It returns -1 if the rune is not a valid value to encode in UTF-8.
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
}
