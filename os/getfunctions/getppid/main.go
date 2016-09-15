package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the proccess id of the caller's parent
	//
	// Getppid returns the process id of the caller's parent.
	ppid := os.Getppid()

	// Print the proccess id of the callers parent
	fmt.Println(ppid)
}
