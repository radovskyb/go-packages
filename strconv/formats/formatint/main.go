package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := int64(-15)

	// Convert the int64 integer v to it's string representation
	// based on the base passed to FormatInt
	//
	// FormatInt returns the string representation of i in the given base,
	// for 2 <= base <= 36. The result uses the lower-case letters
	// 'a' to 'z' for digit values >= 10.
	s10 := strconv.FormatInt(v, 10)

	// Print out s10, the string representation of v in base 10
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v, 16)

	// Print out s16, the string representation of v in base 16
	fmt.Printf("%T, %v\n", s16, s16)
}
