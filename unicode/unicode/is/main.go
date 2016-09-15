package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Is reports whether the rune is in the specified table of ranges.
	//
	// Greek is the set of Unicode characters in script Greek.
	memberShip := unicode.Is(unicode.Greek, 'Ө')
	fmt.Println("Ө(theta) is in the range of unicode.Greek range table ? : ", memberShip)

	memberShip = unicode.Is(unicode.Han, '世')
	fmt.Println("世 is in the range of unicode.Han range table ? : ", memberShip)
}
