package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string s
	s := "Hello, World!"

	// Store all subslices of string s separated by a space
	//
	// Split slices s into all subslices separated by sep and
	// returns a slice of the subslices between those separators.
	// If sep is empty, Split splits after each UTF-8 sequence.
	// It is equivalent to SplitN with a count of -1.
	subslices := bytes.Split([]byte(s), []byte(" "))

	// Iterate over subslices
	for _, slice := range subslices {
		// Print out each slice as a string
		fmt.Println(string(slice))
	}
}
