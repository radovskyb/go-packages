package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new byte slice
	b := []byte{'a', 'b', 'c'}

	// Run the bytes.Repeat function with byte slice
	// b as the byte slice parameter and 3 as the
	// count integer parameter
	b1 := bytes.Repeat(b, 3)

	// Print out the results of bytes.Reapeat (b1)
	fmt.Println(string(b1)) // abcabcabc
}
