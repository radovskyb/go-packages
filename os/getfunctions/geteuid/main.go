package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the numeric effective user id of the caller
	//
	// Geteuid returns the numeric effective user id of the caller.
	id := os.Geteuid()

	// Print the id
	fmt.Println(id)
}
