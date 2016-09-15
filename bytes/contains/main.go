package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create 2 byte slices, a and b
	a := []byte{'a', 'b', 'c'}
	b := []byte{'a', 'b', 'c', 'd', 'e', 'f'}

	// Contains reports whether subslice is within b.
	if bytes.Contains(b, a) {
		fmt.Println("a slice is within b slice")
	}
}
