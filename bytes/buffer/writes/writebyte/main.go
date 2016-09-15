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

	// Append 1 byte to buf
	//
	// WriteByte appends the byte c to the buffer, growing the buffer
	// as needed. The returned error is always nil, but is included to
	// match bufio.Writer's WriteByte. If the buffer becomes too large,
	// WriteByte will panic with ErrTooLarge.
	buf.WriteByte(byte('d'))

	// Once again print buf's contents
	fmt.Println(buf) // abcd
}
