package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a string v that contains the string "true"
	v := "true"

	// Store the boolean value representation of the string v ("true")
	// in the variable s
	//
	// ParseBool returns the boolean value represented by the string.
	// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	// Any other value returns an error.
	if s, err := strconv.ParseBool(v); err == nil {
		// Print out s and it's type
		fmt.Printf("%T, %v\n", s, s) // bool, true
	}
}
