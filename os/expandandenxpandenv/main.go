package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the path from the environment variables
	// by using os.Expand with $PATH and passing in the
	// os.Getenv function
	//
	// Expand replaces ${var} or $var in the string based on the mapping
	// function. For example, os.ExpandEnv(s) is equivalent
	// to os.Expand(s, os.Getenv).
	path1 := os.Expand("$PATH", os.Getenv)

	// Print out the path variables
	fmt.Println(path1)

	// Print out a blank line for spacing
	fmt.Println()

	// The same thing as above can be achieved using os.ExpandEnv
	//
	// ExpandEnv replaces ${var} or $var in the string according to the
	// values of the current environment variables. References to
	// undefined variables are replaced by the empty string.
	path2 := os.ExpandEnv("$PATH")

	// Print out the path variables
	fmt.Println(path2)
}
