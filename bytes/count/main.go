package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a byte slice of some byte character a through f
	byteSlice := []byte{'a', 'b', 'c', 'a', 'a', 'b', 'd', 'f'}

	// Print how many times the byte slice a occurs inside of byteSlice
	//
	// Count counts the number of non-overlapping instances of
	// sep in s. If sep is an empty slice, Count returns 1 + the
	// number of Unicode code points in s.
	fmt.Println(bytes.Count(byteSlice, []byte{'a'})) // 3
}
