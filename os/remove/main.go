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

	// Now create a file called tmpfile.txt
	_, err = os.Create("tmpfile.txt")

	// Check if there were any errors creating the tmp file and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print that tmpfile.txt was created
	fmt.Println("Created tmpfile.txt\n")

	// Remove the tmp directory using os.Remove
	//
	// Remove removes the named file or directory. If there is an
	// error, it will be of type *PathError.
	err = os.Remove("tmp")

	// Check if there was any errors and log any if there was
	if err != nil {
		log.Fatalln(err)
	}

	// Print to stdout that the directory was removed
	fmt.Println("Removed tmp directory")

	// Now remove the file tmpfile.txt using os.Remove
	err = os.Remove("tmpfile.txt")

	// Check again if there was any errors and log any if there was
	if err != nil {
		log.Fatalln(err)
	}

	// Print to stdout that the file was removed
	fmt.Println("Removed tmpfile.txt")
}
