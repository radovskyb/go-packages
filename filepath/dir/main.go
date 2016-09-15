package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Store a path in the variable path
	path := "~/Workspace/go/src/pkgchallenge/filepath/dir"

	// Get the path's directory and print it
	//
	// Dir returns all but the last element of path, typically the path's directory.
	// After dropping the final element, the path is Cleaned and trailing
	// slashes are removed.
	// If the path is empty, Dir returns ".".
	// If the path consists entirely of separators, Dir returns a single separator.
	// The returned path does not end in a separator unless it is the root directory.
	fmt.Println(filepath.Dir(path)) // ~/Workspace/go/src/pkgchallenge/filepath
}
