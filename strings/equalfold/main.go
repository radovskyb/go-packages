package main

import (
	"fmt"
	"strings"
)

func main() {
	// Print whether the strings `Go` and `go` are equal under
	// Unicode case-folding, therefore meaning that cases do not
	// matter in the comparison
	//
	// EqualFold reports whether s and t, interpreted as UTF-8 strings,
	// are equal under Unicode case-folding.
	fmt.Println(strings.EqualFold("Go", "go")) // true

	// `GO` and `go`
	fmt.Println(strings.EqualFold("GO", "go")) // true
}
