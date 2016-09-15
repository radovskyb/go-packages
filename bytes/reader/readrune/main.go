package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Read one rune from r
	b, _, _ := r.ReadRune()

	// Print the rune that was read from r.ReadByte()
	fmt.Println(string(b))
}
