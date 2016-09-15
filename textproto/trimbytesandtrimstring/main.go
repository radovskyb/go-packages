package main

import (
	"fmt"
	"net/textproto"
)

func main() {
	// Create a new byte slice `b` that has
	// leading and trailing ASCII space characters
	b := []byte("  Hello, World!   ")

	// Print out the byte slice `b` as it is as a string
	fmt.Println(string(b))

	// Trim the leading and trailing ASCII space characters from []byte `b`
	//
	// TrimBytes returns `b` without leading and trailing ASCII space.
	b = textproto.TrimBytes(b)

	// Now after calling textproto.TrimBytes, re-print `b` as a string
	fmt.Println(string(b))

	// Print a blank line between using TrimBytes and TrimString
	fmt.Println()

	// Create a new string `s` that has
	// leading and trailing ASCII space characters
	s := "    I am a string.    "

	// Print out the string `s` as it is
	fmt.Println(s)

	// Trim the leading and trailing ASCII space characters from string `s`
	//
	// TrimString returns `s` without leading and trailing ASCII space.
	s = textproto.TrimString(s)

	// Now after calling textproto.TrimString, re-print `s`
	fmt.Println(s)
}
