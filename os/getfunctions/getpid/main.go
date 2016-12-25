package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the process id of the caller
	//
	// Getpid returns the process id of the caller.
	pid := os.Getpid()

	// Print the process id
	fmt.Println(pid)
}
