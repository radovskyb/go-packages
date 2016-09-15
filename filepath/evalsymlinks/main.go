package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// Store a symbolic link name in the variable path which links to main.go
	path := "mainsymlink"

	// Return the path name after the evaluation of any symbolic links and print it
	//
	// EvalSymlinks returns the path name after the evaluation of any symbolic
	// links.
	// If path is relative the result will be relative to the current directory,
	// unless one of the components is an absolute symbolic link.
	evaledPath, err := filepath.EvalSymlinks(path)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out evaledPath
	fmt.Println(evaledPath) // main.go
}
