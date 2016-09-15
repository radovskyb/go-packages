package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a default directory name to use for creating a temporary directory
	//
	// TempDir returns the default directory to use for temporary files.
	tmpDir := os.TempDir()

	// Print the name that would be used to create a temporary directory with
	fmt.Println(tmpDir)
}
