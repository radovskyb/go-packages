package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Create a path
	path := "pkgchallenge/filepath/fromslash/main.go"

	// Now use filepath.FromSlash to replace any slash characters in path
	// with a separator character and print the result
	//
	// FromSlash returns the result of replacing each slash ('/') character
	// in path with a separator character. Multiple slashes are replaced
	// by multiple separators.
	fmt.Println(filepath.FromSlash(path))
}
