package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Read one byte from r
	b, _ := r.ReadByte()

	// Print the byte that was read from r.ReadByte()
	fmt.Println(string(b))
}
