package main

import (
	"fmt"
	"unicode"
)

// RangeTable defines a set of Unicode code points by listing the ranges of
// code points within the set. The ranges are listed in two slices
// to save space: a slice of 16-bit ranges and a slice of 32-bit ranges.
// The two slices must be in sorted order and non-overlapping.
// Also, R32 should contain only values >= 0x10000 (1<<16).
var tab unicode.RangeTable

func main() {
	// In reports whether the rune is a member of one of the ranges.
	memberShip := unicode.In('A', &tab)
	fmt.Println("A is a member of tab ? : ", memberShip)

	// Han is the set of Unicode characters in script Han.
	memberShip = unicode.In('界', unicode.Han)
	fmt.Println("界 is a member of unicode.Han ? : ", memberShip)
}
