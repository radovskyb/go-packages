package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(string('\u263a')) // ☺

	// Find out whether the above unicode character is printable
	//
	// IsPrint reports whether the rune is defined as printable
	// by Go, with the same definition as unicode.IsPrint:
	// letters, numbers, punctuation, symbols and ASCII space.
	c := strconv.IsPrint('\u263a')
	fmt.Println(c) // true

	bell := strconv.IsPrint('\007')
	fmt.Println(bell)           // false
	fmt.Println(string('\007')) // Makes the bell sound
}
