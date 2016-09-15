package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a new strings replacer `r`
	//
	// NewReplacer returns a new Replacer from a list of
	// old, new string pairs.
	// Replacements are performed in order, without overlapping matches.
	rep := strings.NewReplacer("<", "&lt;", ">", "&gt;")

	// Call rep.Replace on the string below
	//
	// Replace returns a copy of s with all replacements performed.
	fmt.Println(rep.Replace("This is <b>HTML</b>!"))

	// After replacing the characters correctly based on
	// the replacer's old new string pairs, the above
	// rep.Replace prints: `This is &lt;b&gt;HTML&lt;/b&gt;!`
}
