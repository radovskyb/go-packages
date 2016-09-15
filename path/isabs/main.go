package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string which is the absolute path to this file
	fullPath := "/Users/Benjamin/Workspace/go/src/pkgchallenge/path/base"

	// Print out if fullPath is an absolute path or not
	//
	// IsAbs reports whether the path is absolute.
	fmt.Println(path.IsAbs(fullPath)) // true
}
