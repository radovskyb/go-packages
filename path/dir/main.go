package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string which is the absolute path to this file
	fullPath := "/Users/Benjamin/Workspace/go/src/pkgchallenge/path/base"

	// Print out the fullPath's directory that it is located in
	//
	// Dir returns all but the last element of path, typically the path's directory.
	// After dropping the final element using Split, the path is Cleaned and trailing
	// slashes are removed.
	// If the path is empty, Dir returns ".".
	// If the path consists entirely of slashes followed by non-slash bytes, Dir
	// returns a single slash. In any other case, the returned path does not end in a
	// slash.
	fmt.Println(path.Dir(fullPath)) // /Users/Benjamin/Workspace/go/src/pkgchallenge/path
}
