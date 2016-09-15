package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Split a list of paths joined by : which is this computer's OS-specific
	// ListSeparator which we can find it places such as PATH or GOPATH
	// environment variables
	//
	// SplitList splits a list of paths joined by the OS-specific ListSeparator,
	// usually found in PATH or GOPATH environment variables.
	// Unlike strings.Split, SplitList returns an empty slice when passed an empty
	// string. SplitList does not replace slash characters in the returned paths.
	split := filepath.SplitList("filepath/splitlist/main.go:anotherfolder/main.go")

	// Range over the split string slice and print both paths
	for _, path := range split {
		fmt.Println(path)
	}
}
