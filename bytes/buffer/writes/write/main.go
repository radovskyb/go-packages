package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new bytes buffer buf with some contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c'})

	// Print buf's contents
	fmt.Println(buf) // abc

	// Append some more contents to buf
	//
	// Write appends the contents of p to the buffer, growing
	// the buffer as needed. The return value n is the length
	// of p; err is always nil. If the buffer becomes too large,
	// Write will panic with ErrTooLarge.
	buf.Write([]byte{'d', 'e', 'f'})

	// Once again print buf's contents
	fmt.Println(buf) // abcdef
}
