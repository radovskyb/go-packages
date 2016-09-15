package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// Create a string slice of paths
	paths := []string{
		"/Users/Benjamin/Workspace",
		"/Benjamin/Workspace",
		"/../../../Benjamin/Workspace",
	}

	// Create a path string
	base := "/Users/Benjamin"

	// Return the relative paths for each of the paths in the paths slice
	// relative to string base
	//
	// Rel returns a relative path that is lexically equivalent to targpath when
	// joined to basepath with an intervening separator. That is,
	// Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
	// On success, the returned path will always be relative to basepath,
	// even if basepath and targpath share no elements.
	// An error is returned if targpath can't be made relative to basepath or if
	// knowing the current working directory would be necessary to compute it.
	for _, path := range paths {
		// Get the relative path to the base
		rel, err := filepath.Rel(base, path)

		// Check if there were any errors and if so, log them
		if err != nil {
			log.Fatalln(err)
		}

		// Print out the relative path
		fmt.Println(rel)
	}
}
