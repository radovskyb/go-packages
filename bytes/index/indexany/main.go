package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string that says Hallo, World!
	input := "Hallo, World!"

	// bytes.IndexAny can also check against utf-8 encoded
	// values, in this case, 97 can be checked as an `a` byte
	// character, therefore it will show up at position 1 here
	//
	// IndexAny interprets s as a sequence of UTF-8-encoded Unicode
	// code points. It returns the byte index of the first occurrence
	// in s of any of the Unicode code points in chars. It returns -1 if
	// chars is empty or if there is no code point in common.
	index := bytes.IndexAny([]byte(input), string(97))

	// Once again print out our results found or not.
	if index > -1 {
		fmt.Println("Index found at position:", index)
	} else {
		fmt.Println("Index not found")
	}
}
