package main

import (
	"fmt"
	"strconv"
)

func main() {
	// QuoteRuneToASCII returns a single-quoted Go character
	// literal representing the rune.  The returned string uses
	// Go escape sequences (\t, \n, \xFF, \u0100) for non-ASCII
	// characters and non-printable characters as defined by IsPrint.
	s := strconv.QuoteRuneToASCII('☺')
	fmt.Println(s)
}
