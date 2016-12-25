package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string s
	s := "Hello, World! How are you?"

	// Store all subslices of string s separated by a space. The amount
	// of subslices is determined by the integer passed into SplitN
	//
	// SplitN slices s into subslices separated by sep and returns
	// a slice of the subslices between those separators. If sep is
	// empty, SplitN splits after each UTF-8 sequence. The count
	// determines the number of subslices to return:
	subslices := bytes.SplitN([]byte(s), []byte(" "), 3)

	// Iterate over subslices
	for _, slice := range subslices {
		// Print out each slice as a string
		fmt.Println(string(slice))
	}
}
