package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// First create a directory named tmp
	err := os.Mkdir("tmp", 0755)

	// Check if there were any errors creating the tmp directory
	// and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print that directory tmp was created
	fmt.Println("Created directory tmp")

	// Now create a file called tmpfile.txt in the tmp directory created above
	_, err = os.Create("tmp/tmpfile.txt")

	// Check if there were any errors creating the tmp file and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print that tmp/tmpfile.txt was created
	fmt.Println("Created tmp/tmpfile.txt\n")

	// Remove the tmp directory and tmp/tmpfile.txt using os.RemoveAll
	//
	// RemoveAll removes path and any children it contains. It removes
	// everything it can but returns the first error it encounters. If
	// the path does not exist, RemoveAll returns nil (no error).
	err = os.RemoveAll("tmp")

	// Check if there was any errors and log any if there was
	if err != nil {
		log.Fatalln(err)
	}

	// Print to stdout that the directory was removed with tmpfile.txt too
	fmt.Println("Removed tmp directory with tmpfile.txt in it")

}
