package main

import (
	"fmt"
	"os"
)

func main() {
	// Return all of the environment variables
	//
	// Environ returns a copy of strings representing the environment,
	// in the form "key=value".
	vars := os.Environ()

	// Iterate over all the environment variables and print each of them
	for _, env := range vars {
		fmt.Println(env)
	}
}
