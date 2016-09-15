package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Join 3 path elements into a single path and print the results
	//
	// Join joins any number of path elements into a single path, adding
	// a Separator if necessary. The result is Cleaned, in particular
	// all empty strings are ignored.
	// On Windows, the result is a UNC path if and only if the first path
	// element is a UNC path.
	//
	// Prints: filepath/join/main.go
	fmt.Println(filepath.Join("filepath", "join", "main.go"))
}
