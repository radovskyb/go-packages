package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert an integer to a string
	//
	// Itoa is shorthand for FormatInt(i, 10).
	s := strconv.Itoa(10)

	// Print out s which is the stringified version of integer 10
	fmt.Printf("%T, %v\n", s, s)
}
