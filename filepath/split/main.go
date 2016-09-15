package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Create a new path string
	path := "filepath/split/main.go"

	// Split the path immediately following the final separator, spliting
	// it into directory and file name
	//
	// Split splits path immediately following the final Separator,
	// separating it into a directory and file name component.
	// If there is no Separator in path, Split returns an empty dir
	// and file set to path.
	// The returned values have the property that path = dir+file.
	dir, file := filepath.Split(path)

	// Print out the directory name
	fmt.Println(dir) // filepath/split/

	// Print out the file name
	fmt.Println(file) // main.go
}
