package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := uint64(15)

	// Convert the uint64 integer v into it's string
	// representation in base 10
	//
	// FormatUint returns the string representation of i in the given base,
	// for 2 <= base <= 36. The result uses the lower-case letters
	// 'a' to 'z' for digit values >= 10.
	s10 := strconv.FormatUint(v, 10)

	// Print out s10
	fmt.Printf("%T, %v\n", s10, s10)

	// Convert the uint64 integer v into it's string
	// representation in base 16
	s16 := strconv.FormatUint(v, 16)

	// Print out s16
	fmt.Printf("%T, %v\n", s16, s16)
}
