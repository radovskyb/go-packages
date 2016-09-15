package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// Get the current directories absolute path
	//
	// Abs returns an absolute representation of path.
	// If the path is not absolute it will be joined with the current
	// working directory to turn it into an absolute path.  The absolute
	// path name for a given file is not guaranteed to be unique.
	absPath, err := filepath.Abs(".")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out this directories absolute path
	fmt.Println(absPath)
}
