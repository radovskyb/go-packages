package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// Create a pattern and also a name string
	//
	// In this case the pattern checks if the name contains the string main,
	// followed by a dot (.) followed by anything else after it
	pattern := "main.*"

	// The name is main.go so therefore it will match below
	name := "main.go"

	// Check whether the name matches the shell file name pattern
	//
	// Match reports whether name matches the shell file name pattern.
	// The pattern syntax is:
	//
	//	pattern:
	//		{ term }
	//	term:
	//		'*'         matches any sequence of non-Separator characters
	//		'?'         matches any single non-Separator character
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
	//
	// On Windows, escaping is disabled. Instead, '\\' is treated as
	// path separator.
	//
	matched, err := filepath.Match(pattern, name)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print whether it matched or not
	if matched {
		fmt.Println("It matched!")
	} else {
		fmt.Println("It didn't match...")
	}
}
