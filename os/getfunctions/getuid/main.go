package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the numberic user id of the caller
	//
	// Getuid returns the numeric user id of the caller.
	uid := os.Getuid()

	// Print the user id
	fmt.Println(uid)
}
