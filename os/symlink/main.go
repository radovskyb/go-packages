package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new symbolic link for main.go
	//
	// Symlink creates newname as a symbolic link to oldname. If
	// there is an error, it will be of type *LinkError.
	err := os.Symlink("main.go", "mainsymlink")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print a success message after calling os.Symlink
	fmt.Println("Created symlink mainsymlink that links to main.go")

	// Now that we are done with mainsymlink, remove it using os.Remove
	err = os.Remove("mainsymlink")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print a success message after calling os.Remove on mainsymlink
	fmt.Println("Successfully removed mainsymlink symbolic link to main.go")
}
