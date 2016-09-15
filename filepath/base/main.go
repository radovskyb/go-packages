package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get the current working directories path
	dir, err := os.Getwd()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Return the last element in the current working directory's path
	//
	// Base returns the last element of path.
	// Trailing path separators are removed before extracting the last element.
	// If the path is empty, Base returns ".".
	// If the path consists entirely of separators, Base returns a single separator.
	base := filepath.Base(dir)

	// Print out base's contents
	fmt.Println(base) // base
}
