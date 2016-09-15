package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Return the actual destination name of the symbolic link mainsymlink
	//
	// Readlink returns the destination of the named symbolic link. If
	// there is an error, it will be of type *PathError.
	name, err := os.Readlink("mainsymlink")

	// Check if os.Readlink returned and error and if it did, log it
	if err != nil {
		log.Fatalln(err)
	}

	// Print the destination name returned from os.Readlink
	fmt.Println(name) // Prints main.go since mainsymlink is linked to main.go
}
