package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the PATH environment value
	//
	// Getenv retrieves the value of the environment variable named by the
	// key. It returns the value, which will be empty if the variable
	// is not present.
	path := os.Getenv("PATH")

	// Print the path
	fmt.Println(path)
}
