package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string which is the absolute path leading upto this file
	fullPath := "/Users/Benjamin/Workspace/go/src/pkgchallenge/path/base"

	// Join the name main.go onto fullPath to get the correct full path of this
	// file, including the filename itself too
	//
	// Join joins any number of path elements into a single path, adding a
	// separating slash if necessary. The result is Cleaned; in particular,
	// all empty strings are ignored.
	//
	// Prints: /Users/Benjamin/Workspace/go/src/pkgchallenge/path/base/main.go
	fmt.Println(path.Join(fullPath, "main.go"))
}
