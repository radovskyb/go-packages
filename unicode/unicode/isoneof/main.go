package main

import (
	"fmt"
	"unicode"
)

func main() {
	// IsOneOf reports whether the rune is a member of one of the ranges.
	// The function "In" provides a nicer signature and should be used
	// in preference to IsOneOf.
	memberShip := unicode.IsOneOf([]*unicode.RangeTable{
		unicode.Han, unicode.Hebrew,
	}, '界')

	fmt.Println(memberShip)
}
