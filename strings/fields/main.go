package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Today was a very lovely day.  "
	// Fields splits the string s around each instance of one or
	// more consecutive white space  characters, as defined by
	// unicode.IsSpace, returning an array of substrings of s or an
	// empty list if s contains only white space.
	fields := strings.Fields(str)
	for _, field := range fields {
		fmt.Println(field)
	}
}
