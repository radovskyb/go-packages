package main

import (
	"fmt"
	"log"
	"path"
)

func main() {
	// Create a new string that holds a path
	fullPath := "~/Workspace/go/src"

	// Return whether fullPath above matches the shell file name pattern below
	//
	// Match reports whether name matches the shell file name pattern.
	// The pattern syntax is:
	//
	//	pattern:
	//		{ term }
	//	term:
	//		'*'         matches any sequence of non-/ characters
	//		'?'         matches any single non-/ character
	//		'[' [ '^' ] { character-range } ']'
	//		            character class (must be non-empty)
	//		c           matches character c (c != '*', '?', '\\', '[')
	//		'\\' c      matches character c
	//
	//	character-range:
	//		c           matches character c (c != '\\', '-', ']')
	//		'\\' c      matches character c
	//		lo '-' hi   matches character c for lo <= c <= hi
	//
	// Match requires pattern to match all of name, not just a substring.
	// The only possible returned error is ErrBadPattern, when pattern
	// is malformed.
	matched, err := path.Match("~/*/*/*", fullPath)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print whether there was a match or not
	fmt.Println(matched) // true
}
