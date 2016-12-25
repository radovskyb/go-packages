package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string s
	s := "Hello, World! How are you?"

	// Store all subslices of string s separated after each
	// instance of a space. The amount of slices is determined
	// by the integer passed into the function SplitAfterN
	//
	// SplitAfterN slices s into subslices after each instance
	// of sep and returns a slice of those subslices. If sep
	// is empty, SplitAfterN splits after each UTF-8 sequence.
	// The count determines the number of subslices to return:
	subslices := bytes.SplitAfterN([]byte(s), []byte(" "), 3)

	// Iterate over subslices
	for _, slice := range subslices {
		// Print out each slice as a string
		fmt.Println(string(slice))
	}
}
