package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a string s
	s := "Hello, World!"

	// Store all subslices of string s separated after each
	// instance of a space
	//
	// SplitAfter slices s into all subslices after each instance
	// of sep and returns a slice of those subslices. If sep is empty,
	// SplitAfter splits after each UTF-8 sequence. It is
	// equivalent to SplitAfterN with a count of -1.
	subslices := bytes.SplitAfter([]byte(s), []byte(" "))

	// Iterate over subslices
	for _, slice := range subslices {
		// Print out each slice as a string
		fmt.Println(string(slice))
	}
}
