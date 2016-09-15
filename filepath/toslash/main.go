package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Create a new path string
	path := "/Users/Benjamin/Workspace"

	// Return the result of replaces each separator character in path
	// with a slash character
	//
	// ToSlash returns the result of replacing each separator character
	// in path with a slash ('/') character. Multiple separators are
	// replaced by multiple slashes.
	fmt.Println(filepath.ToSlash(path))
}
