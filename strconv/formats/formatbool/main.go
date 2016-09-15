package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a bool v - true
	v := true

	// Convert v - bool - to a string "true"
	//
	// FormatBool returns "true" or "false" according to the value of b
	s := strconv.FormatBool(v)

	// Print out s which was the string returned from FormatBool
	fmt.Printf("%T, %v\n", s, s)
}
