package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Environ returns a copy of strings representing the environment,
	// in the form "key=value".
	//
	// Print out the environment variables.
	for _, val := range syscall.Environ() {
		fmt.Println(val)
	}
}
