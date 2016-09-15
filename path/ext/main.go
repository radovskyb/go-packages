package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string which is the absolute path to this file
	// including this filename itself (main.go)
	fullPath := "/Users/Benjamin/Workspace/go/src/pkgchallenge/path/base/main.go"

	// Get the file extension of the filename at the end of fullPath
	//
	// Ext returns the file name extension used by path. The extension is the
	// suffix beginning at the final dot in the final slash-separated element
	// of path; it is empty if there is no dot.
	fmt.Println(path.Ext(fullPath)) // .go
}
