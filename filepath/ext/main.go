package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Store a path
	path := "filepath/main.go"

	// Print out the path's extension
	//
	// Ext returns the file name extension used by path.
	// The extension is the suffix beginning at the final dot
	// in the final element of path; it is empty if there is
	// no dot.
	fmt.Println(filepath.Ext(path)) // .go
}
