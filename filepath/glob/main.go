package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// Create a pattern to use with filepath.Glob which will search from up
	// one directory, and search any type of folder and return anything that
	// has the word main inside of it
	pattern := "../*/*main*"

	// Now get any matches from the pattern
	//
	// Glob returns the names of all files matching pattern or nil
	// if there is no matching file. The syntax of patterns is the same
	// as in Match. The pattern may describe hierarchical names such as
	// /usr/*/bin/ed (assuming the Separator is '/').
	//
	// Glob ignores file system errors such as I/O errors reading directories.
	// The only possible returned error is ErrBadPattern, when pattern
	// is malformed.
	matches, err := filepath.Glob(pattern)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the matches returned from filepath.Glob
	fmt.Println(matches)
}
