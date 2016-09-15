package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new bytes buffer with some contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c', 'd', 'e', 'f'})

	// Remove all of the unread contents in buf
	// except for the first 3 bytes
	//
	// Truncate discards all but the first n unread bytes
	// from the buffer. It panics if n is negative or greater
	// than the length of the buffer.
	buf.Truncate(3)

	// Print out bufs remaining contents
	fmt.Println(buf) // abc
}
