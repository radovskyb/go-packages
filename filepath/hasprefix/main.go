package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Check if main.go starts with main and print a message
	//
	// HasPrefix exists for historical compatibility and should not be used.
	if filepath.HasPrefix("main.go", "main") {
		fmt.Println("main.go has the prefix main")
	}
}
