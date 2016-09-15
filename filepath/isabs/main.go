package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Create an absolute path
	path := "/Im/An/Absolute/Path"

	// Check if path is an absolute path
	CheckIsAbs(path)

	// Create a non absolute path
	path = "Im/Not/An/Absolute/Path"

	CheckIsAbs(path)
}

func CheckIsAbs(path string) {
	// IsAbs reports whether the path is absolute.
	if filepath.IsAbs(path) {
		fmt.Println(path, "is an absolute path")
	} else {
		fmt.Println(path, "is not an absolute path")
	}
}
