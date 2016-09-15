package main

import (
	"fmt"
	"strings"
)

func main() {
	// Titlise the string below
	//
	// Title returns a copy of the string s with all Unicode
	// letters that begin words mapped to their title case.
	// BUG(rsc): The rule Title uses for word boundaries does
	// not handle Unicode punctuation properly.
	fmt.Println(strings.Title("hello, world!")) // Hello, World!
}
