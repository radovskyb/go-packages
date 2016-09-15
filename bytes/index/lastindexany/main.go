package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string that says Hallo, Ma Friend!
	input := "Hallo, Ma Friend!"

	// bytes.LastIndexAny can also check against utf-8 encoded
	// values, in this case, 97 can be checked as an `a` byte
	// character, therefore it will show up at position 1 here
	//
	// LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode
	// code points. It returns the byte index of the last occurrence
	// in s of any of the Unicode code points in chars. It returns -1 if
	// chars is empty or if there is no code point in common.
	index := bytes.LastIndexAny([]byte(input), string(97))

	// Once again print out our results found or not.
	if index > -1 {
		fmt.Println("Last index found at position:", index)
	} else {
		fmt.Println("Index not found")
	}
}
