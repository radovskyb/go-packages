package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if \ is a separator
	//
	// IsPathSeparator reports whether c is a directory separator character.
	issep := os.IsPathSeparator(uint8('\\'))

	// Print where \ is a separator or not
	if issep {
		fmt.Println("\\ is a separator")
	} else {
		fmt.Println("\\ is not a separator")
	}

	// Check again but this time if / is a separator
	issep = os.IsPathSeparator(uint8('/'))

	// Print where / is a separator or not
	if issep {
		fmt.Println("/ is a separator")
	} else {
		fmt.Println("/ is not a separator")
	}
}
