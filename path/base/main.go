package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string which is the absolute path to this file
	fullPath := "/Users/Benjamin/Workspace/go/src/pkgchallenge/path/base"

	// Print the base of the path (the last element in the path)
	//
	// Base returns the last element of path. Trailing slashes are removed
	// before extracting the last element. If the path is empty, Base
	// returns ".". If the path consists entirely of slashes, Base returns "/".
	fmt.Println(path.Base(fullPath)) // base
}
