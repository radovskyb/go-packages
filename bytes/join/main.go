package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create 2 byte slices b1 and b2, which will be joined
	// together below using bytes.Join
	b1 := []byte("abc")
	b2 := []byte("def")

	// Join byte slices b1 and b2 together, using a space for a
	// separator and then store the returned byte slice in b3.
	//
	// Join concatenates the elements of s to create a new byte slice.
	// The separator sep is placed between elements in the resulting slice.
	b3 := bytes.Join([][]byte{b1, b2}, []byte(" "))

	// Print out the new concatenated version of b1 and b2, b3.
	fmt.Println(string(b3))
}
