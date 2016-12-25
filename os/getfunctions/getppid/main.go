package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the process id of the caller's parent
	//
	// Getppid returns the process id of the caller's parent.
	ppid := os.Getppid()

	// Print the process id of the callers parent
	fmt.Println(ppid)
}
