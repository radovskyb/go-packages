package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new bytes buffer with 3 bytes `a`, `b` and `c`
	// as its contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c'})

	// Run buf.Reset() which resets the buffer so that it has
	// no content.
	//
	// Reset resets the buffer so it has no content. b.Reset()
	// is the same as b.Truncate(0).
	buf.Reset()

	// Print out buf to show that it will not print anything
	fmt.Println(buf)
}
