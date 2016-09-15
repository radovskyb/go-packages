package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Quote returns a double-quoted Go string literal representing s.  The
	// returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
	// control characters and non-printable characters as defined by
	// IsPrint.
	s := strconv.Quote(`"I'm a string inside of quotes"`)
	fmt.Println(s)
}
