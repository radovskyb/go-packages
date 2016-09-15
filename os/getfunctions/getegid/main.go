package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the numeric effective group id of the caller
	//
	// Getegid returns the numeric effective group id of the caller.
	id := os.Getegid()

	// Print the id
	fmt.Println(id)
}
