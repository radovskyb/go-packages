package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the proccess id of the caller
	//
	// Getpid returns the process id of the caller.
	pid := os.Getpid()

	// Print the proccess id
	fmt.Println(pid)
}
