package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new file called file1.txt
	_, err := os.Create("file1.txt")

	// Check if any errors occurred and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print that file1.txt was created successfully
	fmt.Println("file1.txt was created successfully")

	// Rename the file file1.txt to file.txt
	//
	// Rename renames (moves) a file. OS-specific restrictions might
	// apply. If there is an error, it will be of type *LinkError.
	err = os.Rename("file1.txt", "file.txt")

	// Check if any errors occurred and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print that file1.txt was renamed to file.txt successfully
	fmt.Println("file1.txt was rename to file.txt successfully")

	// Remove the file now that it's no longer needed
	err = os.Remove("file.txt")

	// Check if any errors occurred and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print that file.txt was removed successfully
	fmt.Println("file.txt was removed successfully")
}
