package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the numeric effective group id of the caller
	//
	// Getgid returns the numeric group id of the caller.
	id := os.Getgid()

	// Print the id
	fmt.Println(id)
}
