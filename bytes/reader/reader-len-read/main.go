package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Print the lenth of the number of bytes in r
	//
	// Len returns the number of bytes of the unread portion of the slice.
	fmt.Println(r.Len())

	// Create a new byte slice of 3 bytes in size
	b := make([]byte, 3)
	// Now read from reader r into b
	r.Read(b)
	// Print the contents of b that were read in from r
	fmt.Println(string(b))
}
