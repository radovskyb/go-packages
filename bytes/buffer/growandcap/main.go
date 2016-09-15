package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a byte slice
	b := []byte{'a', 'b'}

	// Create a new bytes buffer with b as its initial content
	var buf = bytes.NewBuffer(b)

	// Print the current buffer capacity of buf
	//
	// Cap returns the capacity of the buffer's underlying byte
	// slice, that is, the total space allocated for the buffer's data.
	fmt.Println(buf.Cap()) // 2

	// Grow buf's capacity by 12 bytes
	//
	// Grow grows the buffer's capacity, if necessary, to
	// guarantee space for another n bytes. After Grow(n), at
	// least n bytes can be written to the buffer without another
	// allocation. If n is negative, Grow will panic. If the
	// buffer can't grow it will panic with ErrTooLarge.
	buf.Grow(12)

	// Print the current buffer capacity of buf after buf.Grow(12)
	fmt.Println(buf.Cap()) // 16
}
