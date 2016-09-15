package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string that holds a path
	fullPath := "~/Workspace/go/src/goproj/gomain/main.go"

	// Split the fullPath string into directory and file
	//
	// Split splits path immediately following the final slash,
	// separating it into a directory and file name component.
	// If there is no slash path, Split returns an empty dir and
	// file set to path.
	// The returned values have the property that path = dir+file.
	dir, file := path.Split(fullPath)

	// Print the full path
	fmt.Println("Full path:", fullPath)

	// Print the directory returned from path.Split(fullPath)
	fmt.Println("Directory:", dir) // ~/Workspace/go/src/goproj/gomain/

	// Print the file returned from path.Split(fullPath)
	fmt.Println("File:", file) // main.go
}
