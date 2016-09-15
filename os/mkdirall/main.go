package main

import (
	"log"
	"os"
)

func main() {
	// Create the path tmp/tmp1/ as folders using os.MkdirAll
	//
	// MkdirAll creates a directory named path, along with any necessary
	// parents, and returns nil, or else returns an error. The permission
	// bits perm are used for all directories that MkdirAll creates. If path
	// is already a directory, MkdirAll does nothing and returns nil.
	err := os.MkdirAll("tmp/tmp1", 0755)

	// Log any errors
	if err != nil {
		log.Fatalln(err)
	}
}
